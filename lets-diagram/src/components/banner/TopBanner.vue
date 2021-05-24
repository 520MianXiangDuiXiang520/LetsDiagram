<template>
  <div class="top-banner">
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo top-banner"
      mode="horizontal"
      @select="handleSelect"
    >
      <div class="menu-left">
        <el-col :span="16" >
          <el-card shadow="hover" class="my-button">
            新建文件
          </el-card>
        </el-col>
      </div>
      <div class="menu-right">
        <el-col :span="16" >
          <el-card shadow="hover" class="my-button">
            加入协作
          </el-card>
        </el-col>
      </div>
    </el-menu>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeIndex: "1",
      activeIndex2: "1",
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
  },
};
</script>

<style scoped>
.menu-right {
  float: right;
  margin-right: 30px;
  margin-top: 10px;
  margin-bottom: 10px;
}
.menu-left {
  float: left;
  margin-left: 30px;
  margin-top: 10px;
  margin-bottom: 10px;
}
.my-button {
  border-radius: 50px;
  background: linear-gradient(145deg, #cacaca, #f0f0f0);
  box-shadow: 20px 20px 60px #bebebe, -20px -20px 60px #ffffff;
  border-left-width: 0px;
  border-right-width: 0px;
  border-bottom-width: 0px;
  border-top-width: 0px;
  width: 140px;
  height: 60px;
}
.top-banner {
  background: #e0e0e0;
  border-bottom-width: 0px;
}
</style>
