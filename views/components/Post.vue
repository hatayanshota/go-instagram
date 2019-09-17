<template>
    <div class="post">
        <div class="post-inner-wrapper">
            <div class="post-inner">
                <div class="post-header">
                    <a class="post-user" @click="$router.push(`users/${user_id}`)">
                        <div class="post-user-icon-wrapper">
                            <div class="post-user-icon">
                                <img :src="user_icon_path">
                            </div>
                        </div>
                        <div class="post-user-name">{{user_name}}</div>
                    </a>
                    <div class="post-delete" v-if="user_id==login_user_id"><a @click="deleteConfirmModal" class="post-delete-link"><img src="~/assets/images/delete_white.png"></a></div>
                </div>
                <delete-confirm-modal :confirm_image_path="confirm_image_path" :confirm_caption="confirm_caption" v-if="show_delete_modal" @close="closeDeleteModal" @ok="deletePost"></delete-confirm-modal>
                <img class="post-image" :src="post_image_path">
                <div class="post-footer">
                    <div class="post-caption">{{post_caption}}</div>
                    <a class="people-icon" @click="$router.push(`posts/${post_id}/likes`)"><div><img src="~/assets/images/liked_user.png"></div><div class="like_num">{{like_num}}</div></a>
                    <div class="liked" v-if="is_liked"><a @click="deleteLike"><div><img src="~/assets/images/liked.png"></div></a></div>
                    <div class="before-liked" v-else><a @click="createLike"><div><img src="~/assets/images/not_liked.png"></div></a></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
// axios moduleだとpost出来なかったのでこちらを採用
import axios from "axios";
import deleteConfirmModal from '~/components/DeleteConfirmModal'
export default {
    components: {
        deleteConfirmModal
    },
    props: {
        post_id: {
            type: Number,
            required: true,
        },
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
        },
        post_image_path: {
            type: String,
            required: true,
        },
        post_caption: {
            type: String,
            required: true,
        },
        like_users: {
            type: Array,
            required: true,
        },
        login_user_id: {
            type: Number,
            required: true,
        }
    },
    data({ like_users, login_user_id }) {
        var like_user_ids = []
        for (var i in like_users) {
            like_user_ids.push(like_users[i].id)
        }
        return {
            is_liked: like_user_ids.includes(login_user_id),
            like_num: like_user_ids.length,
            show_delete_modal: false,
        }
    },
    methods: {
        deleteConfirmModal() {
            this.confirm_image_path = this.post_image_path;
            this.confirm_caption = this.post_caption;
            this.show_delete_modal = true;
        },
        closeDeleteModal() {
            this.show_delete_modal = false;
        },
        deletePost () {
            var options = {
                withCredentials: true,
            }
            const res = axios.post(`http://localhost:8080/posts/${this.post_id}/delete`)
                .then( res => {
                    console.log(res)
                    // 親要素を変更するために親の関数を呼び出し
                    this.$emit('postsChanged')
                })
                .catch ( err => {
                    console.log(err)
                    // エラー時の動作を入れる
                })
        },
        createLike () {
            // フォームパラメータとして送信
            var params = new URLSearchParams();
            params.append('user_id', this.login_user_id);
            params.append('post_id', this.post_id);
            
            var options = {
                withCredentials: true,
            }
            const res = axios.post('http://localhost:8080/posts/like', params, options)
                .then ( res => {
                    console.log(res)
                    this.is_liked = true
                    this.like_num += 1
                })
                .catch ( err => {
                    console.log(err)
                    // エラー時の動作を入れる
                })
        },
        deleteLike() {
            var params = new URLSearchParams();
            params.append('user_id', this.login_user_id);
            params.append('post_id', this.post_id);

            var options = {
                withCredentials: true,
            }

            const res = axios.post('http://localhost:8080/posts/delete/like', params, options)
                .then ( res => {
                    console.log(res)
                    this.is_liked = false
                    this.like_num -= 1
                })
                .catch ( err => {
                    console.log(err)
                    // エラー時の動作を入れる
                })
        }
    }
}
</script>

<style>
.post {
    width: 40%;
    margin: 0 30%;
    margin-top: -5px;
    border: solid 1px #FF0005;
    background-color: white;
}
.post-inner-wrapper {
    padding-top: 100%;
    position: relative;
}
.post-inner {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}

.post-header {
    height: 8%;
    padding: 0 0.5rem;
    text-align: left;
    display: table;
    width: 100%;
}
.post-footer {
    height: 20%;
    display: flex;
    padding-top: 5px;
    padding-bottom: 1rem;
}
img.post-image {
    width: 100%;
    height: 72%;
    object-fit: cover;
    display: block;
}
.post-user {
    display: table-cell;
    vertical-align: middle;
}
.post-user-icon-wrapper {
    width: 8%;
    display: inline-block;
    vertical-align: middle;
}
.post-user-icon {
    padding-top: 100%;
    position: relative;
    display: block;
}
.post-user img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
}
.post-user .post-user-name {
    display: inline-block;
    vertical-align: middle;
}
.post-delete {
    display: table-cell;
    text-align: center;
    vertical-align: middle;
    width: 5%;
}
.post-delete-link img {
    width: 0.8rem;
}
.post-delete-link {
    padding: 0.2rem 0.5rem;
    text-decoration: none;
    color: #FFF;
    background: #FF0010;
    border-radius: 4px;/*角の丸み*/
    box-shadow: inset 0 1px 0 #f16b4f, inset 0 -2px 0 rgba(0, 0, 0, 0.05);
    font-weight: bold;
    border: solid 1px #f16b4f;
    transition: all 0.4s ease;
}
.post-delete-link:hover {
    opacity: 0.7
}
.post-caption {
    width: 80%;
    text-align: left;
    padding-left: 0.5rem;
    height: 100%;
    overflow: scroll;
}
.people-icon,
.liked,
.before-liked {
    width: 10%;
}
.like_num {
    color: #FF0010;
    font-size: 0.5rem;
    margin-top: -0.5rem;
}
.post-footer img {
    width: 70%;
}
@media (max-width: 768px) {
    .post {
        width: 50%;
        margin: 0 25%;
        margin-top: -5px;
    }
    .post-delete-link img {
        width: 0.7rem;
    }
}
@media (max-width: 576px) {
    .post {
        width: 80%;
        margin: 0 10%;
        margin-top: -5px;
    }
}
</style>