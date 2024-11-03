<template>
  <div class="homepage">
    <div class="right">
      <img src="../assets/tree.svg" width="240" alt="logo">
    </div>
    <div class="left">
      <h1>Welcome To Linktree App</h1>
      <p>Save All your links in one place!</p>
      <div class="btns">
        <div class="btn-group">
          New to Linktree? <router-link to="/signup"><button class="btn-left">Signup</button></router-link>

        </div>
        <div class="btn-group">
          Have an account? <router-link to="/login"><button class="btn-right">Login</button></router-link>
        </div>
      </div>

    </div>
  </div>
</template>

<script>

export default {
  name: 'HomeView',

}
</script>

<style scoped>
.homepage {
  padding: 0px;
  margin: 0px;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;


}

@media only screen and (max-width: 800px) {
  .homepage {
    flex-direction: column;
  }

}

.left {
  widows: 80%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

}

h1 {
  margin: 10px;
  padding: 0px;
  font-size: 40px;
  color:#C8826B;
}

p {
  margin-bottom: 20px;
  padding: 0px;
  font-size: 20px;
  color:#F2F3EB;
}

.btns {
  width: 30%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-evenly;
}

.btn-group {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.btn-right {
  margin: 10px;
  padding: 10px;
  width: 120px;
  height: 40px;
  background: #C8826B;
  border: none;
  border-radius: 5px;
  color: #F2F3EB;
  font-size: 17px;
  font-weight: bold;
}

.btn-right:hover {
  cursor: pointer;
  background: #F2F3EB;
  color: #C8826B;
  transform: scale(1.05);
  transition: 0.3s ease-in-out;
}

.btn-left {
  margin: 10px;
  padding: 10px;
  width: 120px;
  height: 40px;
  background: #90a2a4;
  border: none;
  border-radius: 5px;
  color: #F2F3EB;
  font-size: 17px;
  font-weight: bold;
}

.btn-left:hover {
  cursor: pointer;
  background: #F2F3EB;
  color: #90a2a4;
  transform: scale(1.05);
  transition: 0.3s ease-in-out;
}
</style>