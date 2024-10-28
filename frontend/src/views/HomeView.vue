<template>
  <div class="homepage">
    <h1>Welcome To Linktree App</h1>
    <h2>Save All your links in one place !</h2>
    <div class="btns">
     New to Linktree? <router-link to="/signup"><button>Sign Up</button> <br></router-link> 
     Have an account?  <router-link to="/login" ><button>Login</button></router-link> 
    </div>
   
  </div>
</template>

<script>

export default {
  name: 'HomeView',

}
</script>
