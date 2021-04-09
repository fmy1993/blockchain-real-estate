<template>
  <div id="crop-form">
    <!-- <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="活动名称1">
        <el-input v-model="form.name1"></el-input>
      </el-form-item>
      <el-form-item label="活动名2">
        <el-input v-model="form.name2"></el-input>
      </el-form-item>
    </el-form> -->

    <el-alert type="success">
      <p>信息</p>
    </el-alert>
  </div>
</template>

<script>
// import { mapGetters } from 'vuex'
import { createCrop } from "@/api/crop";

export default {
  name: "AddCrop",
  data() {},
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        createCrop({
          accountId: this.accountId,
          proprietor: this.ruleForm.proprietor,
          totalArea: this.ruleForm.totalArea,
          livingSpace: this.ruleForm.livingSpace,
        }).then((response) => {
          this.loading = false;
          if (response !== null) {
            this.$message({
              type: "success",
              message: "创建成功!",
            });
          } else {
            this.$message({
              type: "error",
              message: "创建失败!",
            });
          }
        });
      });
    },
  },
};
</script>

<style scoped>
</style>
