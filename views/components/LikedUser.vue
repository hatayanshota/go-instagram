<template>
<a class="liked-user" @click="$router.push(`/users/${user_id}`)">
    <div class="liked-user-icon"><div class="liked-user-icon-inner"><img :src="user_icon_path"></div></div>
    <div class="liked-user-name">{{ user_name }}</div>
</a>
</template>

<script>
export default {
    props: {
        user_id: {
            type: Number,
            required: true,
        },
        user_name: {
            type: String,
            required: true,
        },
        user_icon_path: {
            type: String,
            required: true,
        }
    }
}
</script>

<style>
.liked-user {
    background-color: white;
    border: solid 1px #f04262;
    margin-top: -1px;
}
.liked-user-icon {
    width: 20%;
    display: inline-block;
    vertical-align: middle;
}
.liked-user-icon-inner {
    padding-top: 100%;
    position: relative;
}
.liked-user-icon-inner img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
}
.liked-user-name {
    display: inline-block;
    vertical-align: middle;
    padding-left: 10%;
}
.liked-user img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    display: block;
}
</style>