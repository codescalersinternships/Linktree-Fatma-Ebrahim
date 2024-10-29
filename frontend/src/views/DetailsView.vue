<template>
    <div class="details">

        <img src="../assets/tree.svg" width="100" alt="logo">
        <h1>Add your Info</h1>
        <p>Enter your fullname, bio and links</p>
        <form class="form">
            <input type="text" name="fullname" placeholder="Enter Fullname">
            <input type="text" name="bio" placeholder="Enter Bio">
            <input type="text" name="link" placeholder="Enter Link">
            <input type="text" name="link" placeholder="Enter Link">
            <input type="text" name="link" placeholder="Enter Link">
            <router-link to="/tree"> <input class="btn" type="submit" value="Submit"></router-link>

        </form>
    </div>
</template>

<script>

export default {
  name: 'DetailsView',
}
</script>
<style scoped>
.details {
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
    color: #C8826B;
}
p {
    color: #F2F3EB;
    font-size: 17px;
    width: 90%;
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