<template id="login-body">
  <div class="datas user-page">
    <el-row :gutter="20">
      <left />
      <el-col :span="10"
        ><form>
          <div class="segment">
            <p class="title">登 录</p>
            <button class="little" type="button" @click="toRegister">
            去 注 册
          </button>
          </div>
          <label>
            <input type="text" v-model="email" placeholder="请输入注册邮箱" />
          </label>
          <label>
            <input
              type="password"
              v-model="password"
              placeholder="请输入密码"
            />
          </label>
          <button class="red" type="button" @click="login">
            登 录
          </button>
        </form></el-col
      >
    </el-row>
    <el-footer class="footer">
      <hr />
      <el-row :gutter="20">
        <el-col :span="8"> </el-col>
        <el-col :span="8">
          <p>©2020-2021 Let Diagram</p>
        </el-col>
        <el-col :span="8"> </el-col>
      </el-row>
    </el-footer>
  </div>
</template>

<script>
import Left from "./Left.vue";
export default {
  name: "Login",
  components: {
    Left,
  },
  // inject: ["reload", "serverIP"],
  data() {
    return {
      this: null,
      email: "",
      password: "",
      token: "",
    };
  },
  methods: {
    toRegister(){
      this.$router.push("/register")
    },
    login() {
      let self = this;
      if (this.email == "" || this.password == "") {
        this.$message.error("you mast input password and username");
      } else {
        this.axios({
          method: "post",
          url: "user/login",
          data: {
            email: this.email,
            password: this.password,
          },
        }).then(function(response) {
          console.log(self);
          if (response.data["header"]["code"] != 200) {
            self.$message.error("用户名或密码错误！");
          } else {
            self.token = response.data["token"];
            self.$cookie.set("SESSIONID", response.data["token"], 1);
            self.$router.push("/index");
            // self.reload();
          }
        });
      }
    },
  },
  mounted() {},
};
</script>

<style scoped>
@import "./user.css";
</style>
