<template>
    <div class="details">
        <form @submit.prevent="handlesubmit">
            <img class="logo" src="../assets/tree.svg" width="100" alt="Tree Logo">
            <h1 data-test="details-header">Add all your Details and Links</h1>
            <div class="fname-group">
                <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Link Icon" />
                <input class="fullname" type="text" name="fullname" placeholder="Enter Fullname"
                    v-model="tree.fullname" />
            </div>
            <div class="bio-group">
                <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Edit Icon" />

                <input class="bio" type="text" name="bio" placeholder="Enter Bio" v-model="tree.bio" />
            </div>
            <ul>
                <li class="link-item" v-for="(link, index) in tree.links" :key="index">
                    <div class="link-group">
                        <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Edit Icon" />

                        <input class="name" type="text" :placeholder="'Enter Link Name ' + (index + 1)"
                            v-model="tree.links[index].Name" />
                    </div>
                    <div class="link-group">
                        <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Edit Icon" />

                        <input class="link" type="text" :placeholder="'Enter Link URL ' + (index + 1)"
                            v-model="tree.links[index].Link" />
                    </div>
                    <img class="delete-icon" src="../assets/delete.svg" width="20" alt="Delete Icon"
                        @click="handledeltelink(index)" />

                </li>
            </ul>
            <div class="btns">
                <button class="btn-left" type="button" @click="addLinkField">Add link</button>
                <input class="btn-right" type="button" @click="handlesubmit" value="Save">

            </div>


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
        Link: "",
        Name: ""
    }],
});


const addLinkField = () => {
    tree.links.push({
        Link: "",
        Name: ""
    });
};

const handledeltelink = (index) => {
    tree.links.splice(index, 1);
}

const handlesubmit = async () => {

    if (tree.fullname == "" || tree.bio == "") {
        alert("Please add your fullname and bio");
        return
    }
    tree.links = tree.links.filter(link => {
        if (link.Name === "" && link.Link === "") {
            return false;
        } else if (link.Name === "" || link.Link === "") {
            alert("Please fill all links fields");
            return true;
        }
        return true;
    });
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
            localStorage.setItem("tree_id", data.LinkTreeID)
            router.push('/tree/' + data.LinkTreeID);


        } else if (response.status === 400) {
            alert('Unauthorized user, please sign up');
        }
    } catch (err) {
        console.log(err);
    }

};
</script>


<style scoped>
.details {
    width: 90%;
    height: 100%;
}

form {
    padding-top: 50px;
    margin: 0;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: start;
}
h1 {
    color: #C8826B;
}
.logo:hover {
    transform: rotate(-15deg);
    transition: 0.3s ease-in-out;
}

.fname-group {
    margin: 15px;
    padding: 5px;
    color: #8d565f;
    border: none;
    font-size: 20px;
    text-align: center;
    height: 30px;
    width: 40%;
    border-radius: 5px;
    background: #F2F3EB;
    display: flex;
    flex-direction: row;

}

.fullname {
    color: #8d565f;
    border: none;
    font-size: 17px;
    margin-left: 10px;
    width: 80%;
    height: 100%;
    border-radius: 5px;
    background: #F2F3EB;
}

.edit-icon {
    margin-left: 10px;
    margin-right: 5px;
}

.delete-icon {
    margin-left: 20px;
    margin-right: 10px;
    margin: 10px;
}

.delete-icon:hover {
    cursor: pointer;
    transform: scale(1.05);
    transition: transform 0.3s ease-in-out;
}

.bio-group {
    margin: 5px;
    padding: 5px;
    color: #8d565f;
    border: none;
    font-size: 20px;
    text-align: center;
    height: 30px;
    width: 60%;
    border-radius: 5px;
    background: #F2F3EB;
    display: flex;
    flex-direction: row;

}

.link-group {
    margin: 5px;
    padding: 5px;
    color: #8d565f;
    border: none;
    font-size: 20px;
    text-align: center;
    height: 30px;
    width: 50%;
    border-radius: 5px;
    background: #F2F3EB;
    display: flex;
    flex-direction: row;

}


.bio {
    color: #8d565f;
    font-size: 15px;
    margin-left: 10px;
    width: 80%;
    height: 100%;
    border-radius: 5px;
    border: none;
    background: #F2F3EB;
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
    margin: 0;
}

li input {
    width: 100%;
}

li .name {
    margin-left: 15px;
    font-size: 15px;
    color: #8d565f;
    background: #F2F3EB;
    width: 100%;
    height: 30px;
    border: none;
    border-radius: 5px;
}

li .link {
    color: #8d565f;
    margin-left: 15px;
    background: #F2F3EB;
    width: 100%;
    height: 40px;
    border: none;
    border-radius: 5px;
    font-size: 15px;
}


input:focus {
    color: #8d565f;
    outline: none;
    background: #F2F3EB;
    transform: scale(1.05);
}


.link-item {
    width: 70%;
    background: #F2F3EB;
    border-radius: 5px;
    padding: 5px;
    margin: 5px;
    display: flex;
    align-items: center;
    text-decoration: none;
    color: #8d565f;
    transition: transform 0.3s ease-in-out;
}



.link-icon {
    margin-right: 10px;

}

.btns {
    width: 50%;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
}

.btn-left {
    margin: 10px;
    padding: 10px;
    width: 150px;
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

.btn-right {
    margin: 10px;
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

.btn-right:hover {
    cursor: pointer;
    background: #F2F3EB;
    color: #8d565f;
    transform: scale(1.05);
    transition: 0.3s ease-in-out;
}

@media only screen and (max-width: 800px) {

    form {
        scale: 0.8;
    }

    .link-item {
        width: 90%;
        background: #F2F3EB;
        border-radius: 5px;
        display: flex;
        flex-direction: column;
        align-items: center;
        text-decoration: none;
        color: #8d565f;
        transition: transform 0.3s ease-in-out;

    }

    .link-group {
        margin: 5px;
        padding: 5px;
        color: #8d565f;
        border: none;
        font-size: 20px;
        text-align: center;
        height: 30px;
        width: 100%;
        border-radius: 5px;
        background: #F2F3EB;
        display: flex;
        flex-direction: row;

    }

    li .name {
        margin-left: 15px;
        font-size: 15px;
        color: #8d565f;
        background: #F2F3EB;
        width: 80%;
        height: 30px;
        border: none;
        border-radius: 5px;
    }

    li .link {
        color: #8d565f;
        margin-left: 15px;
        background: #F2F3EB;
        width: 80%;
        height: 40px;
        border: none;
        border-radius: 5px;
        font-size: 15px;
    }

    .btns {
        width: 90%;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
    }




}
</style>