<template>
    <div class="details">
        <img src="../assets/tree.svg" width="100" alt="logo" />
        <h1>Add your Info</h1>
        <p>Enter your fullname, bio and links</p>
        <form class="form" @submit.prevent="handlesubmit">
            <input type="text" name="fullname" placeholder="Enter Fullname" v-model="tree.fullname" />
            <input type="text" name="bio" placeholder="Enter Bio" v-model="tree.bio" />
            <div>
                <div class="link" v-for="(link, index) in tree.links" :key="index">
                    <input type="text" :placeholder="'Enter Link Name ' + (index + 1)"
                        v-model="tree.links[index].name" />
                    <input type="text" :placeholder="'Enter Link URL ' + (index + 1)"
                        v-model="tree.links[index].link" />
                </div>
                <button class="addlink-btn" type="button" @click="addLinkField">Add link</button>

            </div>
            <input class="btn" type="submit" value="Submit" />
        </form>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();
const tree = reactive({
    fullname: '',
    bio: '',
    links: [{
        link: "",
        name: ""
    }],
});


const addLinkField = () => {
    tree.links.push({
        link: "",
        name: ""
    });
};

const handlesubmit = async () => {
    if (tree.fullname == "" || tree.bio == "") {
        alert("Please add your fullname and bio");
    } else {
        try {

            const token = localStorage.getItem("token");
            const response = await fetch('/linktree', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    token: token,
                },
                body: JSON.stringify(tree),
            });

            if (response.status === 201) {
                const data = await response.json()
                router.push('/tree/' + data.LinkTreeID);

            } else if (response.status === 400) {
                alert('Unauthorized user, please sign up');
            }
        } catch (err) {
            console.log(err);
        }
    }
};
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

img:hover {
    transform: rotate(-15deg);
    transition: 0.3s ease-in-out;
}

.form {
    width: 60%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

h1 {
    margin: 15px;
    padding: 0px;
    font-size: 40px;
    color: #c8826b;
}

p {
    color: #f2f3eb;
    font-size: 17px;
    width: 90%;
}

input {
    width: 50%;
    height: 30px;
    margin: 10px;
    padding: 10px;
    border: none;
    border-radius: 5px;
    font-size: 15px;
}

.btn {
    margin: 10px;
    padding: 10px;
    width: 150px;
    height: 40px;
    background: #8d565f;
    border: none;
    border-radius: 5px;
    color: #f2f3eb;
    font-size: 17px;
    font-weight: bold;
}


.addlink-btn {
    margin: 10px;
    padding: 10px;
    width: 150px;
    height: 40px;
    background: #C8826B;
    border: none;
    border-radius: 5px;
    color: #f2f3eb;
    font-size: 17px;
    font-weight: bold;
}

.addlink-btn:hover {
    cursor: pointer;
    background: #f2f3eb;
    color: #C8826B;
    transform: scale(1.05);
    transition: 0.3s ease-in-out;
}

.btn:hover {
    cursor: pointer;
    background: #f2f3eb;
    color: #8d565f;
    transform: scale(1.05);
    transition: 0.3s ease-in-out;
}

.link {
    width: 100%;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;

}

@media only screen and (max-width: 800px) {
    .link {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

}
</style>