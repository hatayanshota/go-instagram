<template>
<div class="user-profile-post">
    <div class="user-profile-post-inner">
        <img class="user-profile-post-image" :src="post_image_path">
    </div>
</div>
</template>

<script>
export default {
    props: {
        post_id: {
            type: Number,
            required: true,
        },
        post_image_path: {
            type: String,
            required: true,
        },
        post_caption: {
            type: String,
            required: true,
        },
    }
}
</script>

<style>
.user-profile-post {
    width: 25%;
    padding: 1px;
}
.user-profile-post-inner {
    padding-top: 100%;
    position: relative;
}
.user-profile-post-image {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
    display: block;
}
@media (max-width: 768px) {
    .user-profile-post {
        width: calc(100%/3);
    }
}
@media (max-width: 576px) {
    .user-profile-post {
        width: 50%;
    }
}
</style>