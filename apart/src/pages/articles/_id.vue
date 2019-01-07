<template>
  <section>
    <b-breadcrumb>
      <b-breadcrumb-item :to="{ name: 'articles' }">物件一覧</b-breadcrumb-item>
      <b-breadcrumb-item :active="true">物件詳細</b-breadcrumb-item>
    </b-breadcrumb>
    <b-navbar>
      <h2>物件詳細</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'articles' }">一覧へ</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-form @submit.prevent="submit()">
      <b-form-group label="建物名">
        <b-input type="text" v-model="article.data.name" required />
      </b-form-group>
      <b-form-group label="管理者">
        <b-input type="text" v-model="article.data.administrator" required />
      </b-form-group>
      <b-form-group label="共益費">
        <b-input-group append="円 / 月">
          <b-form-input
            type="number"
            v-model="article.data.commonAreaCharge"
            min="0"
            required
            />
        </b-input-group>
      </b-form-group>
      <b-form-group label="駐車場料金">
        <b-input-group append="円 / 月">
          <b-form-input
            type="number"
            v-model="article.data.parkingFee"
            min="0"
            required
            />
        </b-input-group>
      </b-form-group>
      <b-button-group class="d-block pb-3">
        <b-button type="button" @click="addRoom()">部屋を追加する</b-button>
      </b-button-group>
      <b-list-group class="pb-3" v-if="rooms.length !== 0">
        <b-list-group-item
          v-for="(room, index) in rooms"
          :key="room.key"
          >
          <b-form-group label="部屋名">
            <b-form-input
              type="text"
              v-model="room.data.name"
              required
              />
          </b-form-group>
          <b-form-group label="家賃">
            <b-input-group append="円 / 月">
              <b-form-input
                type="number"
                v-model="room.data.rent"
                required
                min="0"
                />
            </b-input-group>
          </b-form-group>
          <b-button-group>
            <b-button type="button" variant="danger" @click="removeRoom(index)">削除</b-button>
          </b-button-group>
        </b-list-group-item>
      </b-list-group>
      <b-button-group>
        <b-button variant="primary" type="submit">更新する</b-button>
      </b-button-group>
    </b-form>
  </section>
</template>

<script>
import shortid from 'shortid';
import Vue from 'vue';
import { normalizeArticle, normalizeRoom } from '@/util/normalize';

export default Vue.extend({
  async asyncData({ params, error, $firestore }) {
    const articlesRef = $firestore.collection('articles');
    const roomsRef = $firestore.collection('rooms');

    const articleId = params.id;
    const articleRef = articlesRef.doc(articleId);

    let article;
    const rooms = [];
    try {
      const articleDoc = await articleRef.get();

      // 記事が存在しなければ表示できないので404にする
      if (!articleDoc.exists) {
        return error({ statusCode: 404, message: '指定された物件が存在しません' });
      }

      article = normalizeArticle(articleDoc.id, articleDoc.data());

      const roomsDoc = await roomsRef
        .where('articleId', '==', articleDoc.id)
        .orderBy('index')
        .get();
      roomsDoc.forEach(room => rooms.push({
        key: room.id,
        ...normalizeRoom(room.id, room.data()),
      }));

    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }

    return {
      article,
      rooms
    };
  },
  methods: {
    // 部屋を追加する
    addRoom() {
      const lastIndex = this.rooms.length;
      this.rooms.push({
        key: shortid.generate(),
        ...normalizeRoom(null, {
          articleId: this.article.id,
          index: lastIndex,
        }),
      });
    },
    // 部屋を削除する
    // @param index 削除するインデックス
    removeRoom(index) {
      this.rooms.splice(index, 1);
      this.computeRoomIndex();
    },
    // 順序を計算し直す
    computeRoomIndex() {
      this.rooms.forEach((room, index) => {
        room.data.index = index;
      });
    },
    // 送信する
    async submit() {
      const articleId = this.article.id;
      const articleRef = this.$firestore.collection('articles').doc(articleId);
      const roomsRef = this.$firestore.collection('rooms');
      this.$firestore.runTransaction(async t => {
        try {
          await articleRef.update(this.article.data);
          // 自分の部屋一覧
          const rooms = await roomsRef.where('articleId', '==', articleId).get();
          // 未更新の部屋
          const roomUpdateQueue = [...this.rooms];
          // 部屋更新処理
          rooms.forEach(room => {
            const roomId = room.id;
            const index = roomUpdateQueue.findIndex(queue => queue.id === roomId);

            const storeRoom = roomsRef.doc(roomId)
            // 編集後存在しなくなった部屋は削除
            if (index === -1) {
              return;
            }
            // 存在している既存の物件に対して更新
            const roomData = roomUpdateQueue[index];
            storeRoom.update(roomData.data);
            // 追加したのでバッファーから削除
            roomUpdateQueue.splice(index, 1);
          })
          // 新規追加
          roomUpdateQueue.forEach(room => roomsRef.add(room.data));
        } catch(e) {
          return console.log(e);
        }
      })
    }
  }
})
</script>

