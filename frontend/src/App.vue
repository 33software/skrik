<template>
  <div id="app">
    <router-view />  <!-- Сюда будет подгружаться компонент в зависимости от маршрута -->
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
