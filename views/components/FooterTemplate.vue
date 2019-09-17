<template>
    <div class="footer">
        <nav class="go-insta-navigation-footer">
            <ul class="menu">
                <li class="footer-item" v-if="$route.path === '/sign_up'"><a class="login_link" href="http://localhost:8080/signup"><img src="~/assets/images/login_white.png"></a></li>
                <li class="footer-item" v-else><a class="logout_link" href="http://localhost:8080/logout"><img src="~/assets/images/logout_white.png"></a></li>
                <li class="footer-item"><a class="home_link" @click="goHome"><img src="~/assets/images/home_white.png"></a></li>
                <li class="footer-item"><a @click="goPostsNew"><img src="~/assets/images/post_white.png"></a></li>
            </ul>
        </nav>
    </div>
</template>

<script>
export default {
    methods: {
        goHome () {
            var options = {
                withCredentials: true
            }
            const res = this.$axios.get("http://localhost:8080/auth", options)
                .then( res => {
                    this.$router.push("/posts")
                })
                .catch( err => {
                    this.$router.push("/sign_up")
                })
        },
        goPostsNew () {
            var options = {
                withCredentials: true
            }
            const res = this.$axios.get("http://localhost:8080/auth", options)
                .then( res => {
                    this.$router.push("/posts/new")
                })
                .catch( err => {
                    this.$router.push("/sign_up")
                })
        }
    }
}
</script>

<style>
@media (min-width: 769px) {
    .footer {
        display: none;
    }
}
.footer {
    width: 100%;
    position: fixed;
    bottom: 0;
}
.footer-item img {
    width: 1.5rem;
    height: 1.5rem;
}
.go-insta-navigation-footer {
    display: block;
    background-color: white;
}
.go-insta-navigation-footer ul {
    display: flex;
}
.go-insta-navigation-footer li {
    display: inline-block;
    width: calc(100%/3);
    text-align: center;
    background-color: #f04261;
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
}
.go-insta-navigation-footer li:hover {
    opacity: 0.7;
    transition: all 0.4s ease;
}
</style>