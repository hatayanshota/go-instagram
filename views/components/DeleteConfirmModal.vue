<template>
    <transition appear name="delete-confirm-modal" mode="out-in">
        <div class="delete-confirm-modal-mask" @click="$emit('close')">
            <div class="delete-confirm-modal-wrapper">
                <div class="delete-confirm-modal-container">
                    <a @click="$emit('close')" class="cancel">Cancel</a>
                    <img class="delete-confirm-modal-image" :src="confirm_image_path">
                    <div class="delete-confirm-modal-footer">
                        <div class="delete-confirm-modal-caption">{{confirm_caption}}</div>
                    </div>
                </div>
                <div class="delete-confirm-modal-description-container">
                    <div class="delete-confirm-modal-description">
                        <h4 >
                            Are you sure you want to <span style="color: red">delete</span> this post?
                        </h4>
                    </div>
                    <div class="delete-confirm-modal-link">
                        <a @click="$emit('ok')">Delete</a>
                    </div>
                </div>
            </div>
        </div>
    </transition>
</template>

<script>
export default {
    props: {
        confirm_image_path: {
            type: String,
            required: true,
        },
        confirm_caption: {
            type: String,
            required: true,
        },
    },
}
</script>

<style scoped>
.delete-confirm-modal-enter-active, .delete-confirm-modal-leave-active {
    transition: all 1s;
}
.delete-confirm-modal-enter, .delete-confirm-modal-leave-to {
    opacity: 0;
}
.post-inner {
    position: relative;
}
.delete-confirm-modal-mask {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    padding: 1px;
}
.delete-confirm-modal-wrapper {
    width: 100%;
    height: 100%;
}
.delete-confirm-modal-container {
    background-color: #fef8f5;
    height: 80%;
    position: relative;
    padding: 1%;
}
.delete-confirm-modal-container .cancel {
    position: absolute;
    top: 0.5rem;
    right:0.5rem;
    height: 1.3rem;
    color: #4a87ee;
    border-bottom: 1px solid transparent;
}
.delete-confirm-modal-container .cancel:hover {
    border-color: #4a87ee;
    opacity: 1;
}
.delete-confirm-modal-container img {
    height: 80%;
    width: 100%;
    object-fit: contain;
    display: block;
    padding-top: 10%;
}
.delete-confirm-modal-footer {
    height: 20%;
    position: absolute;
    bottom: 0;
    width: 100%;
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
}
.delete-confirm-modal-caption {
    height: 100%;
    overflow: scroll;
}
.delete-confirm-modal-description-container {
    background-color: white;
    height: 20%;
    display: flex;
}
.delete-confirm-modal-description {
    display: inline-block;
    width: 80%;
    position: relative;
}
.delete-confirm-modal-description h4 {
    position: absolute;
    margin:  auto;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 1.3rem;
}
.delete-confirm-modal-link {
    display: inline-block;
    width: 20%;
    position: relative;
}
.delete-confirm-modal-link a {
    position: absolute;
    margin: auto;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    width: 55%;
    height: 1.3rem;
    color: red;
    font-weight: bold;
    border-bottom: 2px solid transparent;
}
.delete-confirm-modal-link a:hover {
    border-color: red;
    opacity: 1;
}
</style>