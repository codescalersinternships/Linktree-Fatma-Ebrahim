<template>
    <div class="login">
        <img src="../assets/tree.svg" width="100" alt="logo">
        <h1>Login</h1>
        <form class="form">
            <input type="text" name="username" placeholder="Enter Username">
            <input type="password" name="password" placeholder="Enter Password"> 
            <router-link  to="/tree" ><input class="btn" type="submit" value="Login"> </router-link>
        </form>
        <router-link  to="/tree" >login</router-link>
    </div>
</template>
<style  scoped>
.login {
    padding: 0px;
    margin: 0px;
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;


}


img:hover{  
    transform:rotate(-15deg);
    transition: 0.3s ease-in-out;
}

.form {
    width: 500px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

}

h1 {
    margin: 15px;
    padding: 0px;
    font-size: 40px;
    color: #90a2a4;
}


input {
    width: 70%;
    height: 30px;
    margin: 10px;
    padding: 10px;
    border: none;
    border-radius: 5px;
    font-size: 15px;
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



.btn {
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

.btn:hover {
    cursor: pointer;
    background: #F2F3EB;
    color: #C8826B;
    transform: scale(1.05);
    transition: 0.3s ease-in-out;
}
</style>
