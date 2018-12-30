<template>
  <div>
    <b-navbar toggleable="md" type="dark" variant="info">
      <b-navbar-brand :to="{ name: 'index'}">物件管理システム</b-navbar-brand>
    </b-navbar>
    <main>
      <nuxt />
    </main>
    <footer />
  </div>
</template>
