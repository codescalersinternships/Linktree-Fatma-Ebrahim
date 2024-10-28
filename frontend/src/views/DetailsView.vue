<template>
    <div class="details">
        <h1>Add your Info and Links</h1>
        <p>Enter your fullname, bio and links</p>
        <form>
            <label for="fullname">Fullname</label>
            <input type="text" name="fullname">
            <label for="bio">Bio</label>
            <input type="text" name="bio">
            <label for="link">Link</label>
            <input type="text" name="link">
            <label for="link">Link</label>
            <input type="text" name="link">
            <label for="link">Link</label>
            <input type="text" name="link">
            <router-link to="/tree"> <input type="submit" value="Submit"></router-link>

        </form>
        <router-link to="/tree"> tree</router-link>
    </div>
</template>
<style scoped>
.details {
    display: flex;
    flex-direction: column;
    align-items: center;
}

</style>