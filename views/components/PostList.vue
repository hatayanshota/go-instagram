<template>
    <div class="post-list">
        <post
        class="posts"
        v-for="post in this.$store.state.posts"
        :key="post.id"
        :post_id="post.id"
        :user_id="post.user.id"
        :user_name="post.user.name"
        :user_icon_path="post.user.icon"
        :post_image_path="post.image"
        :post_caption="post.caption"
        :like_users="post.like_users"
        :login_user_id="login_user_id"
        v-on:postsChanged="postsChanged"
        />
        <div class="post-page-nation">
            <a v-if="this.$store.state.page_num > 1" @click="goPrev" class="post-page-nation-prev"><span>Previous</span></a>
            <a v-if="this.$store.state.page_num < this.$store.state.max_page" @click="goNext" class="post-page-nation-next"><span>Next</span></a>
        </div>
    </div>
</template>

<script>
import Post from "~/components/Post"
export default {
    components: { Post },
    props: {
        login_user_id: {
            type: Number,
            required: true,
        },
    },
    methods: {
        postsChanged () {
            this.$store.dispatch("updatePostList",this.$store.state.page_num)
        },
        goPrev () {
            this.$store.dispatch("updatePostList",this.$store.state.page_num-1)
        },
        goNext () {
            this.$store.dispatch("updatePostList",this.$store.state.page_num+1)
        },
    },
}
</script>

<style>
.post-list {
    text-align: center;
    padding: 5px 0;
}
.posts {
    display: inline-block;
}
.post-page-nation {
    padding-top: 1rem;
    padding-bottom: 5rem;
    display: inline-block;
}
.post-page-nation span {
    padding: 0.5rem 2rem;
    text-decoration: none;
    color: #FFF;
    background: #f04262;
    border-radius: 4px;
    box-shadow: inset 0 1px 0 #f16b4f, inset 0 -2px 0 rgba(0, 0, 0, 0.05);
    font-weight: bold;
    font-size: 1rem;
    border: solid 1px #f16b4f;
    vertical-align: bottom;
    transition: all 0.4s ease;
    line-height: 1.2rem;
}
.post-page-nation-prev {
    display: inline-block;
    margin-right: 1rem;
}
.post-page-nation-next {
    display: inline-block;
    margin-left: 1rem;
}
</style>