import $axios from "axios"

export const state = () => {
  return {
    login_user_id: 0,
    posts: [],
    max_page: 1,
    page_num: 1,
    login: false,
  }
}

export const mutations = {
  setUserID(state, user_id) {
    state.login_user_id = user_id
  },
  setPosts(state, posts) {
    state.posts = posts
  },
  setMaxPage(state, max_page) {
    state.max_page = max_page
  },
  setPageNum(state, page_num) {
    state.page_num = page_num
  },
  setLogin(state) {
    state.login = true
  },
  setLogout(state) {
    state.login = false
  }
}

export const actions = {
  async fetchUser(ctx) {
    var options = {
      withCredentials: true
    }
    const res = await $axios.get('http://localhost:8080/posts', options)
    ctx.commit("setUserID", res.data.login_user_id)
    ctx.commit("setPosts", res.data.posts)
    ctx.commit("setMaxPage", res.data.max_page)
  },
  async updatePostList(ctx, page_num) {
    var options = {
      withCredentials: true
    }
    const res = await $axios.get(`http://localhost:8080/posts?page_num=${page_num}`, options)
    if (res.status == 200) {
      ctx.commit("setPosts", res.data.posts)
      ctx.commit("setMaxPage", res.data.max_page)
      ctx.commit("setPageNum", page_num)
    }
  },
  async authUserLogin(ctx) {
    var options = {
      withCredentials: true
    }
    const res = await $axios.get("http://localhost:8080/login", options)
    console.log(res)
    if (res.status == 200) {
      ctx.commit("setLogin")
    } else {
      ctx.commit("setUserID", 0)
      ctx.commit("setLogout")
    }
  }
}