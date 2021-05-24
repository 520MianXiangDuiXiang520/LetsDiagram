<template id="login-body">
  <div class="datas user-page">
    <el-row :gutter="20">
      <left />
      <el-col :span="10"
        ><form>
          <div class="segment">
            <p class="title">注 册</p>
            <button class="little" type="button" @click="tologin">
            去 登 录
          </button>
          </div>
          <label>
            <input
              type="text"
              v-model="username"
              placeholder="请输入注册用户名"
            />
          </label>
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
          <label>
            <input
              type="password"
              v-model="againPsw"
              placeholder="请输入密码"
            />
          </label>
          <button class="red" type="button" @click="register">
            注 册
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
  name: "Register",
  components: {
    Left,
  },
  // inject: ["reload", "serverIP"],
  data() {
    return {
      this: null,
      username: "",
      email: "",
      password: "",
      againPsw: "",
    };
  },
  methods: {
    tologin(){
      this.$router.push("/login")
    },
    register() {
      let self = this;
      if (this.username == "") {
        this.$message.error("用户名不能为空！");
      } else if (this.email == "") {
        this.$message.error("邮箱不能为空！");
      } else if (this.password == "") {
        this.$message.error("密码不能为空！");
      } else if (this.againPsw != this.password) {
        this.$message.error("两次密码不一致！");
      } else {
        this.axios({
          method: "post",
          url: "user/register",
          data: {
            username: this.username,
            email: this.email,
            password: this.password,
          },
        }).then(function(response) {
          console.log(self);
          if (response.data["header"]["code"] != 200) {
            self.$message.error("参数错误");
          } else {
            self.$router.push("/login");
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
