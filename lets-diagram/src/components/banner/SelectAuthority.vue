<template>
  <div>
    <el-dialog
      :title="'选择权限 '+userID"
      :visible.sync="dialogFormVisible"
      :modal="false"
      :before-close="close"
    >
      <el-form>
        <el-form-item label="权限">
          <el-select v-model="selectAuthority" placeholder="请选择权限">
            <el-option
              label="管理员"
              :value="c.CanvasJurisdictionMarkManager"
              v-if="authority == c.CanvasJurisdictionMarkRoot"
            ></el-option>
            <el-option
              label="可读可写"
              :value="c.CanvasJurisdictionMarkWrite"
            ></el-option>
            <el-option
              label="只读"
              :value="c.CanvasJurisdictionMarkReadOnly"
            ></el-option>
            <el-option
              label="踢出房间"
              :value="c.CanvasJurisdictionMarkNotAnyAuthorized"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="close">取 消</el-button>
        <el-button type="primary" @click="changeAuthority"
          >确 定</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
import * as c from "../../utils/const.js";
import * as core from "../../utils/connection.js"
export default {
  name: "SelectAuthority",
  props: {
    canvas_id: Number,
    authority: Number,
    dialogFormVisible: Boolean,
    userID: Number
  },
  data() {
    return {
      c: c,
      selectAuthority: "请选择权限"
    };
  },
  methods: {
    close: function() {
      this.$emit('setDialogFormVisible', false);
    },
    changeAuthority: function() {
      console.log(this.userID);
      console.log(this.selectAuthority);
      this.$emit('setDialogFormVisible', false);
      core.sendSetPermission(this.userID, this.selectAuthority)
      this.$message.success("修改成功");
    },
  },
};
</script>
