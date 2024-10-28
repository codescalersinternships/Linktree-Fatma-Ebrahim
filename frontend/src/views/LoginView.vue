<template>
    <div class="login">
        <h1>Login</h1>
        <form class="login-form">
            <label for="username">Username</label>
            <input type="text" name="username" placeholder="Enter Username">
            <label for="password">Password</label>
            <input type="password" name="password" placeholder="Enter Password"> 
            <router-link  to="/tree" ><input type="submit" value="Login"> </router-link>
        </form>
        <router-link  to="/tree" >login</router-link>
    </div>
</template>