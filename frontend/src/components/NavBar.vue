<template>
    <div class="nav">
    <router-link to="/" class="nav-item">Home</router-link>
    <router-link to="/tree" class="nav-item">Tree</router-link>
    <router-link to="/details" class="nav-item">Details</router-link>
    <router-link to="/edit" class="nav-item">Edit</router-link>
  </div>
</template>

<script>
export default {
  name: 'NavBar',
}
</script>

<style scoped>


.nav {
  padding: 30px;
  text-decoration: none;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: end;
  background: #364748;
  color: #F2F3EB;
  width: 100vw;
}


.nav-item {
  margin: 10px;
  text-decoration: none;
  color: #F2F3EB;
  font-size: 17px;

}
</style>
