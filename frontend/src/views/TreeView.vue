<template>
    <div class="tree">
        <img class="logo" src="../assets/tree.svg" width="100" alt="Tree Logo">
        <h1>{{ tree.fullname }}</h1>
        <p>{{ tree.bio }}</p>
        <ul>
            <li class="link-item" v-for="(link, index) in tree.links" :key="index" @click="handlelinkclick(link)">
                <img src="../assets/link.svg" width="20" alt="Link Icon" class="link-icon" />
                <span>{{ link.Name }}</span>
                <div class="views">
                    <span> {{ link.Visits }}</span> <img src="../assets/views.svg" width="20" alt="Link Icon"
                        class="link-icon" />
                </div>


            </li>
        </ul>

        <button class="btn" type="button" @click="handlelinktoedit">Edit</button>

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

const handlelinktoedit = () => {
    router.push(`/edit/${tree.ID}`)
}
onMounted(async () => {
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


        } else if (response.status === 404) {
            alert('Tree not found');
        }
    } catch (err) {
        console.log(err);
    }
})

const handlelinkclick = async (link) => {
    if (localStorage.getItem("token") == null) {
        try {
            const response = await fetch(`/linktree/${tree.ID}/addvisit`, {
                method: 'PUT',
                headers: {
                    'link_id': link.ID,
                },
            })
            const data = await response.json()
            window.open(link.Link, '_blank');
            link.Visits+=1

        }
        catch (err) {
            console.log(err)
        }
    } else {
        window.open(link.Link, '_blank');
    }


}

</script>


<style scoped>
.tree {
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

h1 {
    color: #8d565f;
}

p {
    color: #C8826B;
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
    margin: 0;
}

li {
    margin: 10px;
    width: 40%;
}

li:hover {
    cursor: pointer;
}

.link-item {
    background: #90a2a4;
    border-radius: 5px;
    padding: 15px;
    display: flex;
    justify-content: space-between; 
    text-decoration: none;
    color: #F2F3EB;
    transition: transform 0.3s ease-in-out;
}

.link-item:hover {
    transform: scale(1.01);
}

.link-icon {
    margin-right: 5px;
    margin-left:5px ;
}
.views {
    display: flex;
    align-items: center;
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