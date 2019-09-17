<template>
    <div class="liked-users-list">
        <h2 class="liked-users-list-header">Liked Users</h2>
        <liked-user
        class="liked-users"
        v-for="user in users"
        :key="user.id"
        :user_id="user.id"
        :user_name="user.name"
        :user_icon_path="user.icon"
        />
    </div>
</template>

<script>
import LikedUser from "~/components/LikedUser"
export default {
    components: { LikedUser },
    props: {
        users: {
            type: Array,
            required: true,
        }
    }
}
</script>

<style>
.liked-users-list-header {
    padding: 1rem 1rem;
    color: #f04262;
}
</style>