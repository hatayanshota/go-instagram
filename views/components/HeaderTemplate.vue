<template>
    <div class="header">
        <h1 class="header-logo-title">
            <a @click="goHome">
                <img class="header-logo" src="~/assets/images/logo.jpg"/>
                <span class="header-title">
                    <div>Go-Instagram</div>
                    <div style="font-size: 0.5rem;">teamLabインターン</div>
                </span>
            </a> 
        </h1>
        <nav class="go-insta-navigation">
            <ul class="menu">
                <li class="header-item header-login" v-if="$store.state.login"><a class="logout_link" href="http://localhost:8080/logout">Logout</a></li>
                <li class="header-item header-login" v-else><a class="login_link" href="/sign_up">Login</a></li>
                <li class="header-item header-home"><a class="home_link" @click="goHome">Home</a></li>
                <li class="header-item header-post"><a @click="goPostsNew">Post</a></li>
            </ul>
        </nav>
    </div>
</template>

<script>
export default {
    methods: {
        goHome () {
            this.$store.dispatch("authUserLogin")
            this.$router.push("/posts")
        },
        goPostsNew () {
            this.$store.dispatch("authUserLogin")
            if (this.$store.state.login) {
                this.$router.push("/posts/new")
            } else {
                this.$router.push("/sign_up")
            }
        },
    }
}
</script>

<style>
.header {
    position: sticky;
    left: 0;
    top: 0;
    z-index: 1;
    background-color: white;
}
.header-logo-title {
    display: inline-block;
    margin: 2rem 0 0.5rem 7%;
}
.header-logo {
    height: 3rem;
}
.header-title {
    display: inline-block;
    margin-left: 0.5rem;
    color: #f04262;
}
.go-insta-navigation {
    float: right;
    margin: 3.5rem 7% 0.5rem 0;
}
.go-insta-navigation li {
    display: inline-block;
    vertical-align: top;
    font-size: 1rem;
    margin-left: 3rem;
    color: #f04262;
}
.go-insta-navigation a {
    color: #f04262;
    padding-bottom: 0.2rem;
    border-bottom: 2px solid transparent;
}
.go-insta-navigation a:hover {
    border-color: #f04262;
    opacity: 1;
}
.go-insta-navigation li:first-child {
    margin-left: 0;
}
.menu {
    list-style: none;
    padding: 0;
    margin: 0;
}

@media (max-width: 768px) {
    .header {
        text-align: center;
    }
    .header-logo-title {
        margin-left: 0;
    }
    .go-insta-navigation {
        display: none;
    }
}
</style>