<template>
  <section>
    <h2>物件更新</h2>
    <form @submit="e => (submit(), false)">
      <section>
        <dl>
          <dt><label for="article_name">物件名</label></dt>
          <dd><input id="article_name" type="text" v-model="article.data.name" /></dd>
        </dl>
        <section>
          <button type="button" @click="addRoom()">追加</button>
        </section>
        <ul>
           <li v-for="(room, index) in rooms" :key="room.key">
             <dl>
               <dt><label :for="`room_name_${index}`">部屋名</label></dt>
               <dd><input type="text" :id="`room_name_${index}`" v-model="room.data.name" /></dd>
             </dl>
             <button type="button" @click="removeRoom(index)">削除</button>
           </li>
        </ul>
        <button type="submit">更新する</button>
      </section>
    </form>
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
          console.log(roomUpdateQueue)
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

