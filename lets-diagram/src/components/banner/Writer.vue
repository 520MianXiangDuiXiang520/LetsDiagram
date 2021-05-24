<template>
    <div class="writer" v-if="writer_name != ''">
        <el-tag>{{writer_name}} 正在画</el-tag>
    </div>
</template>

<script>
export default {
    name: "Writer",
    props: {
    writer_name: String
  },
}
</script>

<style scoped>
.writer {
    position: fixed;
    top: 55px;
    left: 1100px;
    z-index: 999;
}
</style>