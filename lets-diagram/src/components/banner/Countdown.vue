<template>
  <div>
    <!-- 画布创作者退出后的倒计时遮罩层 -->

    
  </div>
</template>

<script>
export default {
  name: "CountDown",
  data: function() {
    return {
      seconds: 60,
    };
  },
};
</script>

<style scoped>
.icountdown {
  text-align: center;
  color: #ae1100;
  font-size: 25px;
  margin-top: 100px;
  z-index: 19;
}
.seconds {
  font-size: 25px;
}
</style>
