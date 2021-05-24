<template>
  <div class="hello">
    <!-- 顶部导航栏 -->
    <un-display-menu
      :canvas_id="parseInt(this.$route.params.id)"
      :authority="authority"
    />

    <!-- 请求权限时的模态框 -->
    <el-dialog title="提示" :visible.sync="dialogVisible" width="30%">
      <span>{{ permissionMessage }}</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="allowed">同 意</el-button>
        <el-button type="primary" @click="denied">拒 绝</el-button>
      </span>
    </el-dialog>

    <writer :writer_name="writer_name" />
    <div class="topology">

    </div>
    <topology :configs="topologyConfigs"/>
  </div>
</template>

<script>
// 导入topology-vue组件
import * as core from "../utils/connection.js";
import { insertVersion } from "../utils/versions.js";
import UnDisplayMenu from "./banner/UnDisplayMenu.vue";
import { deepClone } from "../utils/deepcopy.js"
import Writer from "./banner/Writer.vue"
import {start_save_image_task, destory_save_image_timer} from "../utils/saveImage.js"

export default {
  name: "Canvas",
  components: {
    UnDisplayMenu,
    Writer,
  },
  data: function() {
    return {
      topologyConfigs: {},
      oldTopologyData: {},
      timer: "",
      webSocket: "",
      version: 1,
      authority: 0,
      dialogVisible: false,
      permissionMessage: "",
      permissionPingUserID: 0,
      permissionStyle: "pointer-events: none",
      cooperateCode: "",
      writer_name: "",
      last_push_cover_version: 0,
      openLoading: ""
    };
  },
  methods: {
    getData: async function() {
      const openLoading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        });
      let self = this;
      await this.axios({
        method: "post",
        url: "canvas/open/",
        data: {
          canvas_id: parseInt(self.$route.params.id),
        },
      }).then(function(response) {
        if (response.data["header"]["code"] === 200) {
          let serverVersion = response.data["version"];
          let sd = JSON.parse(response.data["data"]);
          let opSD = deepClone(sd);
          self.oldTopologyData = sd;
          self.version = serverVersion;
          self.authority = response.data["authority"];
          if (self.authority === 3) {
            self.permissionStyle = "";
          }
          if (opSD["pens"]) {
            insertVersion(serverVersion, sd);
            window.topology.open(opSD);
          } else {
            let sp = {};
            insertVersion(serverVersion, sp);
          }
          openLoading.close();
          console.log("准备建立 websocket 连接");
          // 建立 ws 连接
          let onMessage = core.disponse(self);
          let onOpen = core.withOpen(self);
          self.webSocket = core.newWSConn(
            "ws://"+ self.$base_api +"/canvas/paint/?token=" +
              self.$cookie.get("SESSIONID") +
              "&canvas_id=" +
              self.$route.params.id +
              "&version=" +
              self.version,
            onOpen,
            onMessage
          );
          start_save_image_task(self);
        }
      }).catch(function(){
        openLoading.close();
      });
      
    },

    deleteDocument() {
      let rightDocument = document.querySelector(
        "#app > div > div.le5le-topology > div.menus > div.flex > div.pr16"
      );
      let und = document.querySelector("#un-display > div");
      rightDocument.replaceChildren(...und.children);
    },

    asyncMounted: async function() {
      await this.getData();
      this.deleteDocument();
    },
    allowed: function() {
      core.sendPermissionAllowed(this.permissionPingUserID);
      this.dialogVisible = false;
    },
    denied: function() {
      core.sendPermissionDenied(this.permissionPingUserID);
      this.dialogVisible = false;
    }
  },
  mounted() {
    this.asyncMounted();
  },
  beforeDestroy() {
    console.log("DESTORY");
    core.closeWebsocket(this)
    clearInterval(this.timer);
    destory_save_image_timer();
  },
};
</script>

<style scoped>

</style>
