<template>
    <div class="liked-users-index">
        <liked-user-list :users="users"/>
    </div>
</template>

<script>
import LikedUserList from "~/components/LikedUserList"
import liked_users from "~/assets/json/liked-users.json"
export default {
    components: { LikedUserList },
    data() {
        return {
            users: liked_users,
        }
    },
    async asyncData ({$axios, params}) {
        var options = {
            withCredentials: true,
        }
        const res = await $axios.get(`http://localhost:8080/posts/${params.id}/likes`, options);
        if (res.status == 200) {
            return {
                users: res.data
            }
        }
    }
}
</script>

<style>
.liked-users-index {
    padding: 10% 30%;
}
@media (max-width: 756px) {
    .liked-users-index {
        padding: 10% 20%;
    }
}
</style>