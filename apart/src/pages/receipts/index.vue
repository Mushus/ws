<template>
  <section class="container">
    <h2>領収書印刷</h2>
    <ul>
      <li v-for="article in articles" :key="article.id">
        <h3>{{ article.data.name }}</h3>
        <table>
          <thead>
            <tr>
              <th>部屋名</th><th>入居者</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="room in rooms.filter(room => room.data.articleId === article.id)"
              :key="room.id">
              <td>
                <nuxt-link :to="{ name: 'receipts-id', params: { id: room.id } }">
                  {{ room.data.name }}
                </nuxt-link>
              </td>
            </tr>
          </tbody>
        </table>
      </li>
    </ul>
  </section>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  async asyncData({ $firestore }) {
    const articlesRef = $firestore.collection('articles');
    const roomsRef = $firestore.collection('rooms');
    const articles = [];
    const rooms = [];
    try {
      const articlesDoc = await articlesRef.get();
      articlesDoc.forEach(article => articles.push({
        id: article.id,
        data: article.data(),
      }));

      const roomsDoc = await roomsRef.get();
      roomsDoc.forEach(room => rooms.push({
        id: room.id,
        data: room.data(),
      }));
    } catch(e) {
      console.log(e);
      return error({ statusCode: 500, message: 'データ取得失敗' });
    }
    return {
      articles,
      rooms,
    };
  }
})
</script>
