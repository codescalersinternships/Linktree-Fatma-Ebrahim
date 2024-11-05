<template>
  <div class="app" >
<NavBar />
  <router-view></router-view>
  </div>
</template>

<script setup>
import NavBar from '@/components/NavBar.vue'

</script>


<style>
body {
  padding: 0px;
  margin: 0px;
  background: #364748;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding: 0px;
  margin: 0px;
  background: #364748;
  color: #F2F3EB;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.app{
  width: 100%;
  padding: 0px;
  margin: 0px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

}

</style>
