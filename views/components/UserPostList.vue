<template>
    <div class="user-profile-post-list">
        <user-post
        class="user-profile-post"
        v-for="user_post in user_posts"
        :key="user_post.id"
        :post_id="user_post.id"
        :post_image_path="user_post.image"
        :post_caption="user_post.caption"
        />
    </div>
</template>

<script>
import UserPost from "~/components/UserPost"
export default {
    components: { UserPost },
    props: {
        user_posts: {
            type: Array,
            required: true,
        }
    }
}
</script>

<style>
.user-profile-post-list {
    display: flex;
    flex-wrap:wrap;
}
</style>