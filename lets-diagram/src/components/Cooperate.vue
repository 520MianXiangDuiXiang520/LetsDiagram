<template>
  <div class="hello">
    <!-- 顶部导航栏 -->
    <un-display-menu :canvas_id="canvasID" :authority="authority" />

    <!-- 请求权限时的模态框 -->
    <el-dialog title="提示" :visible.sync="dialogVisible" width="30%">
      <span>{{ permissionMessage }}</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="allowed">同 意</el-button>
        <el-button type="primary" @click="denied">拒 绝</el-button>
      </span>
    </el-dialog>

    <div v-if="authority == c.CanvasJurisdictionMarkReadOnly && !stw">
      <el-button
        type="primary"
        round
        class="raiseHead"
        style="z-index:17"
        @click="openApplication"
        ><svg
          t="1618715311404"
          class="icon"
          viewBox="0 0 1024 1024"
          version="1.1"
          xmlns="http://www.w3.org/2000/svg"
          p-id="10415"
          width="32"
          height="32"
        >
          <path
            d="M759.702069 421.534897a18.643862 18.643862 0 0 1-13.241379-5.543725l-98.127449-98.127448a18.749793 18.749793 0 0 1 0-26.482758l51.517793-51.482483a353.809655 353.809655 0 0 0-187.46262-53.21269c-24.364138 0-48.657655 2.471724-72.527448 7.379862 0.070621 24.505379 0.247172 53.495172 0.706206 82.944 0.282483 19.244138 0.635586 36.369655 1.059311 51.659035 13.135448-6.355862 44.243862-17.867034 69.773241 0.105931 22.068966 15.536552 20.162207 44.032 20.162207 44.032l5.649655 162.286345a154.412138 154.412138 0 0 1 40.077242 91.806896c9.392552 98.445241-35.769379 131.813517-40.96 135.309241a16.843034 16.843034 0 0 0-7.379862 15.148138l8.756965 120.76138a356.422621 356.422621 0 0 0 284.460138-531.244138l-49.18731 49.18731c-3.531034 3.531034-8.297931 5.473103-13.27669 5.473104z m-420.334345 337.92a317.334069 317.334069 0 0 0-1.906758-13.771035c-20.268138-138.063448-30.578759-264.827586-30.614069-376.690759 0-38.770759 0.176552-79.130483 1.412413-118.077793-17.019586 11.970207-32.979862 25.423448-47.668965 40.112552a356.140138 356.140138 0 0 0 55.260689 548.89931l22.845794-74.151724a16.666483 16.666483 0 0 0 0.706206-6.355862z"
            fill="#C5EAFF"
            p-id="10416"
          ></path>
          <path
            d="M536.293517 762.244414c5.15531-3.495724 50.317241-36.864 40.96-135.309242a154.447448 154.447448 0 0 0-40.077241-91.806896l-5.649655-162.286345s1.906759-28.495448-20.197518-44.032c-25.529379-18.008276-56.602483-6.461793-69.737931-0.141241a5458.555586 5458.555586 0 0 1-1.765517-134.567724c-0.105931-38.912 0.070621-66.56 0.070621-67.019035V126.411034c0-0.282483 1.412414-28.212966-5.296552-56.602482C426.548966 35.80469 411.012414 17.513931 388.413793 15.536552a28.354207 28.354207 0 0 0-2.577655-0.105931c-18.114207 0-42.301793 15.042207-55.931586 57.273379-13.771034 42.760828-19.42069 108.155586-21.680552 178.21131a3637.812966 3637.812966 0 0 0-1.412414 118.042483c0.070621 111.863172 10.381241 238.62731 30.64938 376.761379 1.306483 8.827586 1.765517 12.005517 1.871448 13.700414a16.666483 16.666483 0 0 1-0.670897 6.320552l-22.845793 74.222345-41.666207 135.485793h269.135449l-5.614345-77.329655-8.721655-120.76138a16.843034 16.843034 0 0 1 7.344551-15.112827z"
            fill="#FFD4C7"
            p-id="10417"
          ></path>
          <path
            d="M871.106207 291.416276l-98.162759-98.092138a18.749793 18.749793 0 0 0-26.482758 0l-46.609656 46.609655-51.482482 51.482483a18.749793 18.749793 0 0 0 0 26.482758l98.127448 98.162759a18.749793 18.749793 0 0 0 26.482759 0l49.18731-49.222621 48.904828-48.940138a18.749793 18.749793 0 0 0 0-26.482758z"
            fill="#F9DB91"
            p-id="10418"
          ></path>
          <path
            d="M881.240276 328.103724a33.085793 33.085793 0 0 0 0-46.821517l-98.127448-98.127448a33.085793 33.085793 0 0 0-46.821518 0l-38.735448 38.770758a371.005793 371.005793 0 0 0-243.394207-44.950069c0-29.907862 0.141241-49.787586 0.141242-49.787586S460.729379 7.415172 389.684966 1.306483a43.431724 43.431724 0 0 0-3.813518-0.176552c-23.375448 0-53.671724 17.937655-69.561379 67.230897-13.735724 42.513655-19.667862 104.977655-22.104276 175.13931a370.052414 370.052414 0 0 0-152.293517 299.431724 370.052414 370.052414 0 0 0 157.131034 302.892138l-39.476965 128.141241a12.25269 12.25269 0 0 0 11.723034 15.889656h274.184828a12.25269 12.25269 0 0 0 12.25269-13.170759l-4.766897-65.536c185.520552-20.232828 329.904552-177.328552 329.904552-368.216276a368.745931 368.745931 0 0 0-43.043311-173.373793l41.419035-41.419035zM756.559448 203.387586a4.413793 4.413793 0 0 1 6.249931 0l98.092138 98.092138a4.413793 4.413793 0 0 1 0 6.249931l-98.127448 98.127448a4.413793 4.413793 0 0 1-6.214621 0l-98.092138-98.021517a4.413793 4.413793 0 0 1 0-6.249931l98.092138-98.198069zM270.689103 784.560552a341.451034 341.451034 0 0 1 22.45738-503.913931c-0.529655 28.813241-0.706207 58.544552-0.670897 88.346482 0.105931 195.336828 31.77931 381.351724 32.591448 391.627035 0.03531 0.317793 0 0.635586-0.105931 0.918069l-17.019586 55.296a344.452414 344.452414 0 0 1-37.252414-32.273655z m257.588966-34.215724a31.179034 31.179034 0 0 0-13.700414 28.107034l13.241379 182.660414H293.499586l58.862345-191.099586a30.967172 30.967172 0 0 0 1.271172-11.758345c-0.176552-2.189241-0.600276-5.014069-1.977379-14.618483a3022.565517 3022.565517 0 0 1-17.796414-141.947586c-8.403862-83.261793-12.711724-161.509517-12.711724-232.695173-0.070621-107.343448 1.624276-227.29269 22.422069-291.84 5.579034-17.302069 13.629793-30.825931 23.304828-39.053241 8.474483-7.20331 15.536552-8.297931 18.996965-8.297931l1.341793 0.03531c5.790897 0.494345 23.446069 2.048 33.262345 42.372414 6.426483 26.341517 5.15531 53.248 5.155311 53.459862v1.306483c0 0.706207-0.564966 73.939862 0.600275 150.280828 2.365793 159.214345 9.604414 176.551724 13.064828 184.849655 7.309241 17.584552 23.728552 29.872552 42.725517 44.137931 33.262345 24.929103 74.681379 55.931586 80.966621 122.067862 3.672276 38.347034-1.05931 71.468138-13.665104 95.832276-9.357241 18.008276-19.879724 25.423448-21.044965 26.200276z m-6.638345-249.679449c-25.56469-21.115586-49.646345-34.710069-55.860965-49.575724-5.014069-12.040828-7.909517-58.721103-9.533793-113.063724 12.464552-4.590345 31.426207-8.333241 46.856827 2.507035 15.006897 10.593103 14.124138 31.249655 14.124138 31.355586v0.635586l4.413793 128.141241z m305.681655-90.818207a341.662897 341.662897 0 0 1-276.409379 472.664276l-7.768276-106.178207a2.471724 2.471724 0 0 1 1.129931-2.189241c8.262621-5.543724 57.30869-43.184552 47.245242-148.585931-4.025379-42.548966-20.44469-72.986483-40.183173-96.150069l-5.473103-156.530759c0.211862-5.896828-0.070621-37.393655-26.270897-55.790344-21.786483-15.32469-46.009379-13.771034-64.123586-8.827587-0.776828-35.204414-1.129931-71.468138-1.235862-102.223448a343.569655 343.569655 0 0 1 58.156138-4.943448 339.756138 339.756138 0 0 1 164.122483 41.948689l-38.311725 38.311725a33.085793 33.085793 0 0 0 0 46.786206l98.092138 98.162759a33.121103 33.121103 0 0 0 46.821518 0l35.451586-35.451586a388.413793 388.413793 0 0 1 8.756965 18.996965z"
            fill="#512C56"
            p-id="10419"
          ></path></svg
      ></el-button>
      <div
         class="divMask" style="position: absolute; pointer-events: none;
        width: 100%; height: 100%; left: 0px; top: 0px; background: #8f8f8f;
        opacity: 0.15; filter: alpha(opacity=0);z-index:5;" >
      </div>
    </div>

    <!-- 请求权限时的选择框 -->
    <el-dialog title="您要向谁申请？" :visible.sync="selectManagerDialog">
      <el-select v-model="selectManager" placeholder="请选择">
        <el-option
          v-for="item in managers"
          :key="item.id"
          :label="item.username"
          :value="item.id"
        >
          <span style="float: left">{{ item.username }}</span>
          <span style="float: right; color: #8492a6; font-size: 13px">{{
            item.email
          }}</span>
        </el-option>
      </el-select>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="application">确 定</el-button>
      </div>
    </el-dialog>

    <!-- stw 时的提示框 -->
    <el-dialog title="提示" :visible.sync="stwDialog" width="30%" :close-on-click-modal="false" :close-on-press-escape="false" :show-close="false" center>
      <div>
        <div class="icountdown">
          <el-avatar class="seconds"> {{ seconds }} </el-avatar>
          <p class="hint">主持人已离开协作，将在 {{ seconds }} 秒后返回！</p>
        </div>
      </div>
    </el-dialog>

    <writer :writer_name="writer_name" />

    <topology :configs="topologyConfigs" :style="permissionStyle" />
  </div>
</template>

<script>
// 导入topology-vue组件
import * as core from "../utils/connection.js";
import UnDisplayMenu from "./banner/UnDisplayMenu.vue";
import * as c from "../utils/const.js";
import { insertVersion } from "../utils/versions.js";
import { deepClone } from "../utils/deepcopy.js";
import Writer from "./banner/Writer.vue";

export default {
  name: "Canvas",
  components: {
    UnDisplayMenu,
    Writer,
  },
  data: function() {
    return {
      c: c,
      topologyConfigs: {},
      oldTopologyData: {},
      timer: "",
      webSocket: "",
      version: 1,
      authority: 0,
      dialogVisible: false,
      selectManagerDialog: false,
      permissionMessage: "",
      permissionPingUserID: 0,
      permissionStyle: "pointer-events: none",
      core: null,
      cooperateCode: "",
      canvasID: 0,
      rooters: [],
      managers: [],
      selectManager: "",
      writer_name: "",
      seconds: 60,
      stw: false,
      stwDialog: false,
    };
  },
  methods: {
    getData: async function() {
      let self = this;
      await this.axios({
        method: "post",
        url: "canvas/join/",
        data: {
          code: self.$route.params.code,
        },
      }).then(function(response) {
        if (response.data["header"]["code"] === 200) {
          let serverVersion = response.data["version"];
          let sd = JSON.parse(response.data["data"]);
          let opSD = deepClone(sd);
          self.oldTopologyData = sd;
          self.version = serverVersion;
          self.canvasID = response.data["canvas_id"];
          self.authority = response.data["authority"];
          if (self.authority == 3) {
            self.permissionStyle = "";
          }
          if (opSD["pens"]) {
            insertVersion(serverVersion, sd);
            window.topology.open(opSD);
          } else {
            let sp = {};
            insertVersion(serverVersion, sp);
          }

          // 建立 ws 连接
          let onMessage = core.disponse(self);
          let onOpen = core.withOpen(self);
          self.webSocket = core.newWSConn(
            "ws://"+ self.$base_api +"/canvas/paint/?token=" +
              self.$cookie.get("SESSIONID") +
              "&canvas_id=" +
              self.canvasID +
              "&version=" +
              self.version,
            onOpen,
            onMessage
          );
        }
      });
    },

    deleteDocument() {
      let rightDocumrnt = document.querySelector(
        "#app > div > div.le5le-topology > div.menus > div.flex > div.pr16"
      );
      let und = document.querySelector("#un-display > div");
      rightDocumrnt.replaceChildren(...und.children);
    },

    getManagers: function() {
      let self = this;
      this.axios({
        method: "post",
        url: "canvas/collaborators/",
        data: {
          canvas_id: self.canvasID,
        },
      }).then(function(response) {
        self.rooters = response.data["rooter"];
        self.managers = response.data["manger"];
        self.managers.push(...self.rooters);
      });
    },

    openApplication: function() {
      this.getManagers();
      this.selectManagerDialog = true;
    },

    allowed: function() {
      core.sendPermissionAllowed(this.permissionPingUserID);
      this.dialogVisible = false;
    },

    denied: function() {
      core.sendPermissionDenied(this.permissionPingUserID);
      this.dialogVisible = false;
    },

    application: function() {
      console.log(this.selectManager);
      this.selectManagerDialog = false;
      core.sendWriteApplication(this, this.selectManager);
    },
  },
  mounted() {
    this.getData();
    this.deleteDocument();
  },
  beforeDestroy() {
    console.log("DESTORY");
    core.closeWebsocket(this)
    clearInterval(this.timer);
  },
};
</script>

<style scoped>
.countdown {
  position: absolute;
  pointer-events: none;
  width: 100%;
  height: 100%;
  left: 0px;
  top: 0px;
  background: #8f8f8f;
  opacity: 0.15;
  filter: alpha(opacity=0);
  z-index: 19;
}
.raiseHead {
  position: absolute;
  right: 260px;
  bottom: 60px;
}

.icountdown {
  text-align: center;
  color: #ae1100;
  font-size: 25px;
}
.seconds {
  font-size: 25px;
}
</style>
