<template>
    <div class="signup">
        <img src="../assets/tree.svg" width="100" alt="logo">
        <h1>Sign up</h1>
        <form class="form" @submit.prevent="handlesubmit">
            <input type="text" name="username" placeholder="Enter Username" v-model="user.username">
            <input type="email" name="email" placeholder="Enter Email" v-model="user.email">
            <input type="password" name="password" placeholder="Enter Password" v-model="user.password">
            <input class="btn" type="submit" value="Signup">
        </form>
    </div>
</template>
<script setup>
import router from '@/router';
import { reactive } from 'vue';
const user = reactive({
    username: "",
    email: "",
    password: ""
})
const handlesubmit = async () => {
    try {
        const response = await fetch("/linktree/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(user)
        })
        if (response.status == 201) {
            const data = await response.json()
            localStorage.setItem("token", data.token)
            router.push("/details")
        } else if (response.status == 400) {
            alert("User already exists, please login or try again")
        }
        console.log(response.status)
    }
    catch (err) {
        console.log(err)
    }

}

</script>
<style scoped>
.signup {
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
    color: #C8826B;
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
    background: #90a2a4;
    border: none;
    border-radius: 5px;
    color: #F2F3EB;
    font-size: 17px;
    font-weight: bold;
}

.btn:hover {
    cursor: pointer;
    background: #F2F3EB;
    color: #90a2a4;
    transform: scale(1.05);
    transition: 0.3s ease-in-out;
}
</style>