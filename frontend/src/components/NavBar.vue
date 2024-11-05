<template>
  <div class="nav">
    <router-link to="/" class="nav-item">Home</router-link>
    <router-link :to="tree_id" class="nav-item">MyTree</router-link>
    <router-link to="/edit" class="nav-item">Edit</router-link>
    <button class="nav-item" @click="handlelogout">Logout</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import router from '@/router';

const tree_id = ref(null);

onMounted(() => {
  tree_id.value = "/tree/"+localStorage.getItem("tree_id"); 
});

const handlelogout = () => {
  const logout=confirm("Are you sure You want to logout?");
  if (!logout) {
    return
  }
  localStorage.removeItem("token");
  localStorage.removeItem("tree_id");
  router.push('/');
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
  width: 100%;
  height: 50px;
}


.nav-item {
  margin: 20px;
  text-decoration: none;
  color: #F2F3EB;
  font-size: 15px;

}
.nav-item:hover {
  cursor: pointer;
  font-weight: bold;
}
button {
  margin: 20px;
  padding: 0;
  background: none;
  border: none;
  text-decoration: none;
  color: #F2F3EB;
  font-size: 15px;
  cursor: pointer;
}

button:hover {
  font-weight: bold;
}


@media  only screen and (max-width: 800px) {
  .nav {
    flex-direction: row;
    align-items: center;
    justify-content: center;
 
  }
  
}
</style>
