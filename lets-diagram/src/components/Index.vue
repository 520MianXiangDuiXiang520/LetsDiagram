<template>
  <!-- 登录之后的首页，展示用户所有画布 -->
  <div class="index">
    <!-- 顶部导航栏 -->
    <div class="top-banner">
      <el-menu
        :default-active="activeIndex"
        class="el-menu-demo top-banner"
        mode="horizontal"
        @select="handleSelect"
      >
        <div class="menu-left" @click="enterNameFlag = true">
          <el-col :span="16">
            <div>
              <el-card shadow="hover" class="my-button">
                新建文件
              </el-card>
            </div>
          </el-col>
        </div>
        <div class="menu-right">
          <el-col :span="16">
            <div @click="openCooperate">
              <el-card shadow="hover" class="my-button">
                加入协作
              </el-card>
            </div>
          </el-col>
        </div>
      </el-menu>
    </div>
    <!-- 协作模态框 -->
    <el-dialog title="加入协作" :visible.sync="dialogTableVisible">
      <el-input
        placeholder="请输入协作码"
        v-model="cooperateCode"
        clearable
      ></el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogTableVisible = false">取 消</el-button>
        <el-button type="primary" @click="checkCode">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="画布重命名" :visible.sync="rename.visible">
      <el-form :model="form">
        <el-form-item label="新名称" :label-width="formLabelWidth">
          <el-input v-model="rename.name" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="rename.visible = false">取 消</el-button>
        <el-button type="primary" @click="doCanvasRename">确 定</el-button>
      </div>
    </el-dialog>

    <!-- canvas 列表 -->
    <div class="canvas-list">
      <el-row :gutter="20">
        <!-- canvas 列表 -->
        <el-col
          :span="4"
          v-for="canvas in canvasList"
          :key="canvas.id"
          class="index-list-out"
          ><el-card shadow="hover" class="canvas-cover index-list">
            <el-row :gutter="20">
              <el-col :span="18" v-if="canvas.name != ''">
                <p
                  class="canvasName"
                  @click="openRenameForm(canvas.id, canvas.name)"
                >
                  {{ canvas.name.slice(0, 20) }}
                </p></el-col
              >
              <el-col :span="18" v-else>
                <p
                  class="canvasName"
                  @click="openRenameForm(canvas.id, '未命名文件')"
                >
                  未命名文件
                </p>
              </el-col>
              <el-col
                :span="6"
                class="delete-race"
                @click="openDelectSure(canvas.id)"
                ><i
                  @click="openDelectSure(canvas.id)"
                  class="el-icon-delete"
                ></i
              ></el-col>
            </el-row>
            <div class="cover" @click="openCanvas(canvas.id)">
              <img :src="getCover(canvas.id)" :id="'cover' + canvas.id" />
              <!-- <lazy-component>
                <img :src="getCover(canvas.id)" :id="'cover'+canvas.id"/>
              </lazy-component> -->
              <!-- <el-image
                :src="covers[canvas.id]"
                :id="'cover' + canvas.id" layz
              ></el-image> -->
            </div>
          </el-card></el-col
        >
      </el-row>
    </div>

    <el-dialog title="请输入文件名" :visible.sync="enterNameFlag">
      <el-input
        placeholder="请输入文件名"
        v-model="newCanvasName"
        maxlength="7"
        clearable
      ></el-input>
      <div slot="footer" class="dialog-footer">
        <el-button @click="enterNameFlag = false">取 消</el-button>
        <el-button type="primary" @click="newCanvers">确 定</el-button>
      </div>
    </el-dialog>

    <el-pagination
      @current-change="changePage"
      background
      layout="prev, pager, next"
      :total="total"
      :page-size="size"
      :pager-count="10"
      :hide-on-single-page="true"
    >
    </el-pagination>
  </div>
</template>

<script>
import * as storage from "../utils/coverStorage.js";
export default {
  name: "Index",
  data() {
    return {
      canvasList: [],
      page: 0,
      size: 10,
      total: 0,
      dialogTableVisible: false,
      cooperateCode: "",
      deleteSure: false,
      enterNameFlag: false,
      newCanvasName: "",
      covers: {},
      rename: {
        id: 0,
        name: "",
        visible: false,
      },
    };
  },
  components: {},
  mounted() {
    this.getAll();
  },
  methods: {
    changePage: function(cur) {
      console.log(cur);
      this.page = cur;
      this.getAll();
    },
    openCooperate: function() {
      this.dialogTableVisible = true;
    },
    checkCode: function() {
      if (this.cooperateCode == "") {
        this.$message.error("请输入协作码！");
        this.dialogTableVisible = false;
        return;
      }
      let self = this;
      this.axios({
        method: "post",
        url: "canvas/check_cooperate/",
        data: {
          code: self.cooperateCode,
        },
      }).then(function(response) {
        if (response.data["header"]["code"] === 200) {
          if (response.data["result"]) {
            self.dialogTableVisible = false;
            self.$router.push("/cooperate/" + self.cooperateCode);
          } else {
            self.dialogTableVisible = false;
            self.$message.error("请输入正确的协作码！");
          }
        }
      });
    },
    newCanvers: function() {
      let self = this;
      this.axios({
        method: "post",
        url: "canvas/new/",
        data: {
          name: this.newCanvasName,
        },
      }).then(function(response) {
        if (response.data["header"]["code"] == 200) {
          self.$router.push("/newCanvas/" + response.data["canvas_id"]);
        }
      });
      this.enterNameFlag = false;
      // this.$router.push("/newCanvas");
    },
    deleteCanvas: function(id) {
      let self = this;
      this.axios({
        method: "post",
        url: "canvas/delete/",
        data: {
          canvas_id: id,
        },
      }).then(function(response) {
        if (response.data["header"]["code"] === 200) {
          self.$message.success("删除成功");
          self.getAll();
        }
      });
    },
    openDelectSure: function(id) {
      let self = this;
      this.$confirm("此操作将永久删除该文件, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        self.deleteCanvas(id);
      });
    },
    openRenameForm: function(id, oldName) {
      this.rename.visible = true;
      this.rename.name = oldName;
      this.rename.id = id;
    },
    closeRenameForm: function() {
      this.rename.visible = false;
      this.rename.name = "";
      this.rename.id = 0;
    },
    doCanvasRename: function() {
      let self = this;
      self
        .axios({
          method: "post",
          url: "canvas/rename/",
          data: {
            canvas_id: this.rename.id,
            name: this.rename.name,
          },
        })
        .then(function(response) {
          if (response.data["header"]["code"] === 200) {
            self.$message.success("重命名成功");
            self.getAll();
          }
        })
        .catch(function(err) {
          console.log(err);
          self.$message.error("重命名失败");
        });
      this.closeRenameForm();
    },
    getCovers: async function(id) {
      let coverFromLocal = storage.get(id);
      if (coverFromLocal != null && coverFromLocal.length > 0) {
        this.covers[id] = coverFromLocal;
      } else {
        let self = this;
      await self
        .axios({
          method: "post",
          url: "canvas/cover/get/",
          data: {
            canvas_id: id,
          },
        })
        .then(function(response) {
          if (response.data["header"]["code"] === 200) {
            self.covers[id] = response.data["cover"];
            storage.set(id, response.data["cover"]);
          }
        })
        .catch(function(err) {
          console.log(err);
        });
      }
      
    },
    getCover: function(id) {
      let self = this;
        this.getCovers(id).then(function() {
          let doc = document.getElementById("cover" + id);
          doc.setAttribute("src", self.covers[id]);
        });
    },
    getAll: function() {
      let self = this;
      this.axios({
        method: "post",
        url: "canvas/all/",
        data: {
          page: self.page,
          size: self.size,
        },
      }).then(function(response) {
        if (response.data["header"]["code"] === 200) {
          self.canvasList = response.data["list"];
          self.total = response.data["total"];
        }
      });
    },
    openCanvas: function(id) {
      storage.remove(id)
      this.$router.push("/newCanvas/" + id);
    },
  },
};
</script>

<style scoped>
.canvasName {
  font-size: 14px;
  color: #303133;
}
.delete-race {
  border-radius: 6px;
  font-size: 15px;
}
.delete-race:hover {
  background: #ae1100;
  cursor: pointer;
}
.cover {
  margin-top: 20px;
}
.canvas-cover {
  height: 250px;
  border-block-width: 0px;
  border-left-width: 0px;
  border-right-width: 0px;
}
.canvas-list {
  margin: 20px auto auto 130px;
}
.index-list {
  border-radius: 22px;
  background: linear-gradient(145deg, #cacaca, #f0f0f0);
  box-shadow: 20px 20px 60px #bebebe, -20px -20px 60px #ffffff;
}
.index-list-out {
  margin-right: 25px;
  margin-bottom: 35px;
}
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
