<template>
  <div id="app">
     <router-view></router-view>
    <!-- <TopBanner /> -->
    <!-- <HelloWorld /> -->
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
// import TopBanner from './components/banner/TopBanner.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld,
    // TopBanner
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
