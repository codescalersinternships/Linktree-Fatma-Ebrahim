<template>
    <div class="login">
        <img src="../assets/tree.svg" width="100" alt="logo">
        <h1>Login</h1>
        <form class="form" @submit.prevent="handlesubmit">
            <input type="text" name="username" placeholder="Enter Username" v-model="user.username">
            <input type="password" name="password" placeholder="Enter Password" v-model="user.password">
            <input class="btn" type="submit" value="Login">
        </form>
    </div>
</template>
<script setup>
import router from '@/router';
import { reactive } from 'vue';
const user = reactive({
    username: "",
    password: ""
})
const handlesubmit = async () => {
    if (user.username == "" || user.password == "") {
        alert("Empty fields, please enter username and password")
    } else {
        try {
            const response = await fetch("/linktree/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(user)
            })
            if (response.status == 200) {
                const data = await response.json()
                localStorage.setItem("token", data.Token)
                router.push("/tree/" + data.LinkTreeID)
            } else if (response.status == 400) {
                alert("Unauthorized user, please sign up or try again")
            }
        }
        catch (err) {
            console.log(err)
        }
    }

}

</script>

<style scoped>
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


img:hover {
    transform: rotate(-15deg);
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

@media only screen and (max-width: 800px) {
  .login {
    scale: 0.8;
  }

}
</style>
