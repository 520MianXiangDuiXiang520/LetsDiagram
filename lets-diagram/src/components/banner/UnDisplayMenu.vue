<template>
  <div>
    <div id="un-display">
      <div>
        <div v-if="authority >= c.CanvasJurisdictionMarkManager">
          <a class="menu" @click="share">
            <div class="icon myicon">
              <i class="el-icon-share"></i>
            </div>
            <div>分享</div>
          </a>
        </div>
        <div>
          <a class="menu" @click="openOperation">
            <div class="icon myicon">
              <i class="el-icon-s-operation"></i>
            </div>
            <div>操作</div>
          </a>
        </div>
        <div>
          <a class="menu">
            <div class="icon myicon">
              <i class="el-icon-user-solid"></i>
            </div>
            <div>个人</div>
          </a>
        </div>
      </div>
    </div>

    <!-- 分享协作码时的模态框 -->
    <el-dialog
      title="提示"
      :visible.sync="centerDialogVisible"
      width="30%"
      center
    >
      <span>Let Diagram: 您的协作码是 【 <b>{{ cooperateCode }}</b> 】</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="centerDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="copyCooperateCode"
          >复制到剪贴板</el-button
        >
        
      </span>
    </el-dialog>

    <!-- 操作权限的抽屉 -->
    <el-drawer title="权限管理面板" :visible.sync="drawer" direction="rtl">
      <div class="collaborators-list">
        <el-collapse v-model="activeNames">
          <el-collapse-item title="Rooter" name="1">
            <el-row v-for="root in rooters" :key="root.id">
              <el-col :span="18">
                <div class="grid-content">
                  <el-row>
                    <el-col :span="4">
                      <div class="grid-content">
                        <div>
                          <el-avatar class="big-font">
                            {{ root.username[0] }}
                          </el-avatar>
                        </div>
                      </div>
                    </el-col>
                    <el-col :span="20">
                      <div class="grid-content">
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo big-font">{{
                                root.username
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo smill-font">{{
                                root.email
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
              <el-col :span="6"
                ><div class="grid-content">
                  <el-button type="primary" round disabled>修改权限</el-button>
                </div></el-col
              >
            </el-row>
          </el-collapse-item>

          <el-collapse-item title="管理员" name="2">
            <el-row v-for="root in managers" :key="root.id" class="list">
              <el-col :span="18">
                <div class="grid-content">
                  <el-row>
                    <el-col :span="4">
                      <div class="grid-content">
                        <div>
                          <el-avatar class="big-font">
                            {{ root.username[0] }}
                          </el-avatar>
                        </div>
                      </div>
                    </el-col>
                    <el-col :span="20">
                      <div class="grid-content">
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo big-font">{{
                                root.username
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo smill-font">{{
                                root.email
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
              <el-col :span="6"
                ><div class="grid-content">
                  <el-button
                    @click="changePermission(root.id)"
                    type="primary"
                    round
                    :disabled="authority != c.CanvasJurisdictionMarkRoot"
                    >修改权限</el-button
                  >
                </div></el-col
              >
            </el-row>
          </el-collapse-item>

          <el-collapse-item title="创作者" name="3">
            <el-row v-for="root in writers" :key="root.id" class="list">
              <el-col :span="18">
                <div class="grid-content">
                  <el-row>
                    <el-col :span="4">
                      <div class="grid-content">
                        <div>
                          <el-avatar class="big-font">
                            {{ root.username[0] }}
                          </el-avatar>
                        </div>
                      </div>
                    </el-col>
                    <el-col :span="20">
                      <div class="grid-content">
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo big-font">{{
                                root.username
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo smill-font">{{
                                root.email
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
              <el-col :span="6"
                ><div class="grid-content">
                  <el-button
                    type="primary"
                    @click="changePermission(root.id)"
                    :disabled="authority < c.CanvasJurisdictionMarkManager"
                    round
                    >修改权限</el-button
                  >
                </div></el-col
              >
            </el-row>
          </el-collapse-item>

          <el-collapse-item title="观众" name="4">
            <el-row v-for="root in readers" :key="root.id" class="list">
              <el-col :span="18">
                <div class="grid-content">
                  <el-row>
                    <el-col :span="4">
                      <div class="grid-content">
                        <div>
                          <el-avatar class="big-font">
                            {{ root.username[0] }}
                          </el-avatar>
                        </div>
                      </div>
                    </el-col>
                    <el-col :span="20">
                      <div class="grid-content">
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo big-font">{{
                                root.username
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                        <el-row>
                          <el-col :span="24">
                            <div class="grid-content bg-purple-dark">
                              <span class="uname userinfo smill-font">{{
                                root.email
                              }}</span>
                            </div>
                          </el-col>
                        </el-row>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
              <el-col :span="6"
                ><div class="grid-content">
                  <el-button
                    @click="changePermission(root.id)"
                    type="primary"
                    round
                    :disabled="authority < c.CanvasJurisdictionMarkManager"
                    >修改权限</el-button
                  >
                </div></el-col
              >
            </el-row>
          </el-collapse-item>
        </el-collapse>
      </div>
    </el-drawer>

    <select-authority
      :userID="selectUserID"
      :canvas_id="canvas_id"
      :authority="authority"
      :dialogFormVisible="dialogFormVisible"
      @setDialogFormVisible="setDialogFormVisible"
    />
  </div>
</template>

<script>
import * as c from "../../utils/const.js";
import SelectAuthority from "./SelectAuthority.vue";

export default {
  name: "UnDisplayMenu",
  components: {
    SelectAuthority,
  },
  props: {
    canvas_id: Number,
    authority: Number,
  },
  data: function() {
    return {
      cooperateCode: "",
      centerDialogVisible: false,
      drawer: false,
      dialogFormVisible: false,
      activeNames: ["1"],
      rooters: [],
      managers: [],
      writers: [],
      readers: [],
      c: c,
      selectAuthority: 0,
      selectUserID: 0,
    };
  },
  methods: {
    share: function() {
      let self = this;
      this.axios({
        method: "post",
        url: "canvas/cooperate/",
        data: {
          canvas_id: self.canvas_id,
        },
      }).then(function(response) {
        self.cooperateCode = response.data["cooperate_code"];
      });
      this.centerDialogVisible = true;
    },
    copyCooperateCode: async function() {
      try {
        if (this.cooperateCode != "") {
          await navigator.clipboard.writeText(this.cooperateCode);
          this.$message.success("已复制，快去分享给小伙伴吧！")
        }
      } catch (err) {
        this.$message.error("糟糕，您的浏览器似乎不支持操作剪贴板！")
      }
      this.centerDialogVisible = false
    },
    openOperation: function() {
      this.getCollaborators();
      this.drawer = true;
    },
    getCollaborators: function() {
      let self = this;
      this.axios({
        method: "post",
        url: "canvas/collaborators/",
        data: {
          canvas_id: self.canvas_id,
        },
      }).then(function(response) {
        self.readers = response.data["reader"];
        self.writers = response.data["writer"];
        self.rooters = response.data["rooter"];
        self.managers = response.data["manger"];
      });
    },
    deleteDocument: function() {
      let rightDocumrnt = document.querySelector(
        "#app > div > div.le5le-topology > div.menus > div.flex > div.pr16"
      );
      let und = document.querySelector("#un-display > div");
      rightDocumrnt.replaceChildren(...und.children);
    },
    changePermission: function(user) {
      this.dialogFormVisible = true;
      this.selectUserID = user;
    },
    setDialogFormVisible: function(v) {
      this.dialogFormVisible = v;
      this.drawer = false;
    },
  },
  mounted() {},
};
</script>

<style scoped>
#un-display {
  display: none;
}
.myicon {
  font-size: 21px;
}
.collaborators-list {
  margin: 15px;
}
.userinfo {
  float: left;
  margin-left: 7px;
}
.uname {
  color: #303133;
}
.email {
  color: #606266;
}
.big-font {
  font-size: 15px;
}
.smill-font {
  font-size: 13px;
}
.list {
  margin-bottom: 7px;
}
</style>
