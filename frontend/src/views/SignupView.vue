<template>
    <div class="signup">
        <h1>Signup</h1>
        <form class="submit-form">
            <label for="username">Username</label>
            <input type="text" name="username" placeholder="Enter Username">
            <label for="email">Email</label>
            <input type="email" name="email" placeholder="Enter Email">
            <label for="password">Password</label>
            <input type="password" name="password" placeholder="Enter Password"> 
            <router-link  to="/details" ><input type="submit" value="Signup"> </router-link> 
        </form>
        <router-link  to="/details" >signup</router-link>
    </div>
</template>