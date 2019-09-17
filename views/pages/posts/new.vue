<template>
    <div id="post-new">
        <form id="form" @submit.prevent="createPost" enctype="multipart/form-data">
            <div id="image">
                <input id="upload" @change="selectedFile" name="image" type="file"/>
            </div>
            <div id="caption-form">
                <textarea id="caption" placeholder="キャプション" v-model="caption" name="caption"></textarea>
            </div>
            <div id="post">
                <button id="post_button">Post</button>
            </div>
        </form>
    </div>
</template>

<script>
import axios from "axios"

export default {
    data: function () {
        return {
            caption: ''
        }
    },
    methods: {
        selectedFile: function(e) {
            e.preventDefault();
            let files = e.target.files;
            this.uploadFile = files[0];
        },
        createPost: function() {
            let formData = new FormData();
            formData.append('user_id', this.$store.state.login_user_id)
            formData.append('caption', this.caption)
            formData.append('image', this.uploadFile);
            const res = axios.post('http://localhost:8080/posts/new', formData)
                .then(res => {
                    if(res.status == 200){
                        this.$router.push("/posts")
                    } 
                })
                .catch ( err => {
                    console.log(err)
                    this.$router.push("/posts")
                })
        }

    }
}
</script>

<style>
#post-new {
    padding: 10% 25%;
    height: 100vh;
}
#image {
    text-align: center;
    height: 15%;
}
#form {
    height: 100%;
}
#caption-form{
    text-align: center;
    height: 40%;
}
#post{
    height: 55%;
    text-align: center;
}
#caption{
    width: 100%;
    height: 80%;
    border: solid 1px #f04262;
}
#post_button {
    display: inline-block;
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

#post_button:hover{
    opacity: 0.7;
}

</style>