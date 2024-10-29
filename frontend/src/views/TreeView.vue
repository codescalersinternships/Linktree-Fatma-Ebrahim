<template>
    <div class="tree">
        <img src="../assets/tree.svg" width="100" alt="logo">
        <h1>fullname</h1>
        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Numquam at est eveniet libero.</p>
        <ul>
            <li>link</li>
            <li>link</li>
            <li>link</li>
        </ul>
        <router-link to="/edit"><button class="btn">Edit</button></router-link>
    </div>
</template>

<style scoped>
.tree {
    padding-top: 100px;
    margin: 0px;
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: start;

}

img:hover {
    transform: rotate(-15deg);
    transition: 0.3s ease-in-out;
}

h1 {
    color: #8d565f;
}

p {
    color: #F2F3EB;
    font-size: 17px;
    width: 90%;
}

ul {
    color: #F2F3EB;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    list-style: none;
    padding: 10px;
    margin: 0px;

}

li {

    background: #90a2a4;
    border-radius: 5px;
    padding: 15px;
    margin: 10px;
    width: 60%;

}
li:hover{
cursor: pointer;

transform: scale(1.01);
transition: 0.3s ease-in-out
}
.btn {
   
    padding: 10px;
    width: 150px;
    height: 40px;
    background: #8d565f;
    border: none;
    border-radius: 5px;
    color: #F2F3EB;
    font-size: 17px;
    font-weight: bold;
}

.btn:hover {
    cursor: pointer;
    background: #F2F3EB;
    color: #8d565f;
    transform: scale(1.05);
    transition: 0.3s ease-in-out;
}
</style>