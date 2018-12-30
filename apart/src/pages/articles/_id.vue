<template>
  <section>
    <b-navbar class="pt-3">
      <h2>物件詳細</h2>
      <b-button-toolbar>
        <b-button-group>
          <b-button :to="{ name: 'articles-create' }">一覧へ</b-button>
        </b-button-group>
      </b-button-toolbar>
    </b-navbar>
    <b-form @submit="e => (submit(), false)">
      <b-form-group label="建物名">
        <b-input type="text" v-model="article.data.name" />
      </b-form-group>
      <b-button-group class="d-block pb-3">
        <b-button type="button" @click="addRoom()">追加</b-button>
      </b-button-group>
      <b-list-group class="pb-3" v-if="rooms.length !== 0">
        <b-list-group-item
          v-for="(room, index) in rooms"
          :key="room.key"
          >
          <b-form-group label="部屋名">
            <b-input type="text" :id="`room_name_${index}`" v-model="room.data.name" />
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
import shortid from 'shortid'
import Vue from 'vue'

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
        return error({ statusCode: 404, message: '記事が存在しません' });
      }

      article = {
        id: articleId,
        data: articleDoc.data()
      };

      const roomsDoc = await roomsRef.where('articleId', '==', articleDoc.id).get();
      roomsDoc.forEach(room => rooms.push({
        key: room.id,
        id: room.id,
        data: room.data(),
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
        id: null,
        data: {
          articleId: this.article.id,
          name: '',
          index: lastIndex,
        },
      })
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
              console.log(roomId)
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

