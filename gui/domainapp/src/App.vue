<template>
  <div id="app">
    <DomainApp msg="Welcome to my Domain Api"/>
  </div>
</template>

<script>
import DomainApp from './components/DomainApp.vue'

export default {
  name: 'app',
  components: {
    DomainApp
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
