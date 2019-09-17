<template>
    <div id="posts-index">
        <post-list :posts="posts" :login_user_id="login_user_id" :max_page="max_page"/>
    </div>
</template>

<script>
import PostList from "~/components/PostList"

export default {
    components: {PostList},
    computed: {
        login_user_id(){
            return this.$store.state.login_user_id
        },
        posts() {
            return this.$store.state.posts
        },
        max_page() {
            return this.$store.state.max_page
        }
    },
    async mounted(){
        this.$store.dispatch("fetchUser")
        this.$store.dispatch("authUserLogin")
    }
}
</script>

<style>
</style>