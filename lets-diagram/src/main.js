import Vue from 'vue'
import App from './App.vue'
import topology from "topology-vue";
import router from './router.js'
import ElementUI from "element-ui";
import { Message } from "element-ui";
import vueAxios from "vue-axios";
import axios from "axios";
import cookie from "vue-cookie";
import VueLazyload from "vue-lazyload";
import "topology-vue/topology-vue.css";
import "element-ui/lib/theme-chalk/index.css";

Vue.use(topology);
Vue.use(ElementUI);
Vue.use(vueAxios, axios)
Vue.use(VueLazyload, {
  preLoad: 1.3,
  error: "/letdiagram.png",
  loading: "/loading.gif",
  attempt: 3,
});

Vue.config.productionTip = false
Vue.prototype.$axio = axios
Vue.prototype.$message = Message
Vue.prototype.$cookie = cookie;
axios.defaults.withCredentials = true;
Vue.config.productionTip = false;
axios.defaults.baseURL = "http://139.9.117.134:8889";
Vue.prototype.$base_api = "139.9.117.134:8889";
// axios.defaults.baseURL = "http://localhost:8888";
// Vue.prototype.$base_api = "localhost:8888";

// 响应拦截器
axios.interceptors.response.use(
  function(response) {
    if (response.data["header"]["code"] != 200) {
      Message.error(response.data["header"]["msg"]);
    }
    return response;
  },
  function(error) { 
    if (error.response.status == 401) {
      Message.error("您还没有登录呢！");
      router.push("/login");
    } else if (error.response.status == 403) {
      Message.error("您貌似没有这项权限！");
      router.push("/index");
    } else if (error.response.status == 419) {
      Message.error("请求太快了，休息一下吧！")
    } else {
      Message.error("请求服务失败，请检查您的网络连接！");
    }
    return Promise.reject(error);
  }
);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
