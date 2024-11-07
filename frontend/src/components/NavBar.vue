<template>
  <div class="nav">
    <router-link to="/" class="nav-item" :class="{ active: activeItem === 'home' }" @click="setActiveItem('home')">
      Home
    </router-link>
    <router-link :to="tree_id" class="nav-item" :class="{ active: activeItem === 'mytree' }"
      @click="setActiveItem('mytree')">
      MyTree
    </router-link>
    <router-link to="/edit" class="nav-item" :class="{ active: activeItem === 'edit' }" @click="setActiveItem('edit')">
      Edit
    </router-link>
    <button class="nav-item" :class="{ active: activeItem === 'logout' }"
      @click="() => { handlelogout(); setActiveItem('logout'); }">
      Logout
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import router from '@/router';

const tree_id = ref(null);
const activeItem = ref(null);

onMounted(() => {
  console.log(localStorage.getItem("tree_id"));
  tree_id.value = "/tree/" + localStorage.getItem("tree_id");
});

const setActiveItem = (item) => {
  activeItem.value = item;
};

const handlelogout = () => {
  if (localStorage.getItem("token") == null) {
    alert("You are not logged in");
    return;
  }

  const logout = confirm("Are you sure you want to logout?");
  if (!logout) {
    return;
  }
  localStorage.removeItem("token");
  localStorage.removeItem("tree_id");
  router.push('/');
};
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

.nav-item.active {
  font-weight: bold;
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

button.active {
  font-weight: bold;
}

button:hover {
  font-weight: bold;
}

@media only screen and (max-width: 800px) {
  .nav {
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
}
</style>
