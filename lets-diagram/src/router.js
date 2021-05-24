import Vue from "vue";
import Router from "vue-router"; // 引入vue-router
import Login from './components/user/Login.vue'
import Register from "./components/user/Register.vue";
import Index from './components/Index.vue'
import Canvas from './components/Canvas.vue'
import Cooperate from './components/Cooperate.vue'

Vue.use(Router); //使用vue-router

export default new Router({
  routes: [
    {
      path: "/",
      redirect: "/login",
    },

    {
      path: "/login",
      name: "login",
      component: Login,
    },
    {
      path: "/register",
      name: "register",
      component: Register,
    },
    {
      path: "/index",
      name: "index",
      component: Index,
    },
    {
      path: "/newCanvas/:id",
      name: "newCanvas",
      component: Canvas,
    },
    {
      path: "/cooperate/:code",
      name: "cooperate",
      component: Cooperate,
    },
  ],
});
