<template>
    <div>
       <el-col :span="14">
        <el-image src="/letdiagram.png"></el-image>
        <div class="intro">
          <p>
            <i class="el-icon-magic-stick"></i> Let Diagram 是一款完全基于 Web
            的绘图工具，您无需下载任何软件就可以使用。
          </p>
          <p>
            <i class="el-icon-pie-chart"></i> Let Diagram
            支持程序流程图，状态转换图，活动图，时序图等多种图表类型，满足您的大部分需求。
          </p>
          <p>
            <i class="el-icon-data-line"></i> Let Diagram
            支持多人同时绘图，分享协作码，与你的团队一起头脑风暴吧！
          </p>
          <p>
            <i class="el-icon-coordinate"></i> Let Diagram
            拥有完善的权限管理机制，协作过程中，您可以动态调整每个人的权限。
          </p>
        </div>
      </el-col>
    </div>
</template>

<script>
export default {
    name: "Left"
}
</script>

<style scoped>
@import "./user.css"
</style>