<template>
<div id="profile">
    <div id="head">
        <div class="user-profile-icon">
            <div class="user-profile-icon-inner">
                <img :src="user_icon_url">
            </div>
        </div>
        <div class="user-profile-details">
            <div class="user-profile-name">
                {{ user_name }}
            </div>
            <div class="user-profile-liked-num">
                {{ liked_num }}イイネ
            </div>
        </div>
    </div>
    <div class="user-profile-posts">
        <user-post-list class="user-profile-post-list" :user_posts="user_posts"/>
    </div>
</div>
    
</template>

<script>
import UserPostList from "~/components/UserPostList"
export default {
    components: { UserPostList },
    async asyncData ({ $axios, params }) {
        var options = {
            withCredentials: true
            }
        const res = await $axios.get(`http://localhost:8080/users/${params.id}`, options);
        if (res.status == 200) {
            var liked_num = 0;
            for (var i in res.data.posts) {
                liked_num += res.data.posts[i].like_users.length
            }
            return {
                user_icon_url: res.data.icon,
                user_name: res.data.name,
                user_posts: res.data.posts,
                liked_num: liked_num,
            }
        }
    }
}
</script>>

<style>
#profile {
    padding: 10% 10%;
}
#head {
    display: flex;
    padding: 0 5%;
}
.user-profile-icon {
    width: 20%;
}
.user-profile-icon-inner {
    padding-top: 100%;
    position: relative;
}
.user-profile-icon-inner img {
    position: absolute;
    top: 0;
    left: 0;
    display: block;
    width: 100%;
    height: 100%;
    object-fit: cover;
    border: solid 1px #f04262;
}
.user-profile-details {
    width: 80%;
    padding-left: 7%;
    padding-top: 2%;
}
.user-profile-name {
    font-size: 2rem;
}
.user-profile-liked-num {
    padding-left: 1rem;
    font-size: 0.8rem;
}
.user-profile-posts {
    margin-top: 5%;
}
@media (max-width: 768px) {
    .user-profile-name {
        font-size: 1.5rem;
    }
    .user-profile-liked-num {
        padding-left: 0.4rem;
        font-size: 0.6rem;
    }
}
</style>