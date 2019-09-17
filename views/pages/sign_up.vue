<template>
    <div id="sign-up">
        <div class="github-icon"><img src="~/assets/images/github-icon.png"></div>
        <div class="signup_link"><a href="http://localhost:8080/signup"><button class="github-sign-up">GitHub<br>Login</button></a></div>
    </div>
</template>

<script>
export default {
}
</script>

<style>
#sign-up {
    margin-top: 15rem;
}
.github-icon {
    text-align: center;
}
.signup_link {
    text-align: center;
    margin-top: 1rem;
}
#sign-up a {
    display: inline-block;
}
.github-sign-up {
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
.github-sign-up:hover {
    opacity: 0.7;
}
</style>