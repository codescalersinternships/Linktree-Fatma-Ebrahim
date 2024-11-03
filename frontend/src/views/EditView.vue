<template>
    <div class="edit">
        <img class="logo" src="../assets/tree.svg" width="100" alt="Tree Logo">
        <div class="fname-group">
            <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Link Icon" />
            <input class="fullname" type="text" name="fullname" value={{tree.fullname}} placeholder="Enter Fullname"
                v-model="tree.fullname" @change="handlefullnamechange" />
        </div>
        <div class="bio-group">
            <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Edit Icon" />

            <input class="bio" type="text" name="bio" value={{tree.bio}} placeholder="Enter Bio" v-model="tree.bio"
                @change="handlebiochange" />
        </div>
        <ul>
            <li class="link-item" v-for="(link, index) in tree.links" :key="index" @change="handlelinkchange(index)">
                <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Edit Icon" />

                <input class="name" type="text" :placeholder="'Enter Link Name ' + (index + 1)" value={{link.Name}}
                    v-model="tree.links[index].Name" />
                <img class="edit-icon" src="../assets/edit.svg" width="15" alt="Edit Icon"/>

                <input class="link" type="text" :placeholder="'Enter Link URL ' + (index + 1)" value={{link.Link}}
                    v-model="tree.links[index].Link" />
                    <img class="delete-icon" src="../assets/delete.svg" width="20" alt="Delete Icon" @click="deletelink(index)" />

            </li>
            <div class="btns">
                <button class="btn-left" type="button" @click="addLinkField">Add link</button>
                <button class="btn-right" type="button" @click="addLinkField">Save</button>

            </div>

        </ul>

    </div>
</template>

<script setup>
import router from '@/router';
import { onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const tree = reactive({
    ID: route.params.id,
    fullname: '',
    bio: '',
    links: [{
        Link: "",
        Name: "",
        Visits: "",
        ID: ""
    }],
});
const addLinkField = () => {
    tree.links.push({
        link: "",
        name: ""
    });
};


const handlelinkchange = (index) => {
    console.log(tree.links[index].ID, tree.links[index].Name, tree.links[index].Link)
}
const handlefullnamechange = () => {
    console.log(tree.fullname)
}
const handlebiochange = () => {
    console.log(tree.bio)
}

const deletelink = (index) => {

    console.log(tree.links[index].ID)
}

onMounted(async () => {
    console.log(tree.ID)
    try {
        const tree_id = route.params.id
        const token = localStorage.getItem("token");
        const response = await fetch('/linktree/' + tree_id, {
            method: 'GET',
            headers: {
                token: token,
            },
        });

        if (response.status === 200) {
            const data = await response.json()
            tree.fullname = data.Fullname
            tree.bio = data.Bio
            tree.links = data.Links
            console.log(tree)


        } else if (response.status === 404) {
            alert('Tree not found');
        }
    } catch (err) {
        console.log(err);
    }
})


</script>


<style scoped>
.edit {
    padding-top: 100px;
    margin: 0;
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: start;
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
    width: 30%;
    border-radius: 5px;
    background: #F2F3EB;
    display: flex;
    flex-direction: row;

}

.fullname {
    color: #8d565f;
    border: none;
    font-size: 20px;
    margin-left: 20px;
    width: 80%;
    height: 100%;
    border-radius: 5px;
    background: #F2F3EB;
}

.edit-icon {
    margin-left: 10px;
    margin-right: 10px;
}
.delete-icon {
    margin-left: 20px;
    margin-right: 10px;
}
.delete-icon:hover{
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
    width: 50%;
    border-radius: 5px;
    background: #F2F3EB;
    display: flex;
    flex-direction: row;

}

.bio {
    color: #8d565f;
    font-size: 17px;
    font-size: 20px;
    margin-left: 20px;
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

li .name {
    margin-left: 15px;
    font-size: 15px;
    color: #8d565f;
    background: #F2F3EB;
    width: 30%;
    height: 30px;
    border: none;
    border-radius: 5px;
}

li .link {
    color: #8d565f;
    margin-left: 15px;
    background: #F2F3EB;
    width: 50%;
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
</style>