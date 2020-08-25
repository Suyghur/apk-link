<template>
  <div id="task-create">
    <el-form ref="postForm" :model="postForm" :rules="rules" class="task-form">
      <el-form-item label="活动名称" prop="name">
        <el-input v-model="postForm.id" />
      </el-form-item>
      <el-form-item label="活动区域" prop="region">
        <el-select v-model="postForm.id" placeholder="请选择活动区域">
          <el-option label="区域一" value="shanghai" />
          <el-option label="区域二" value="beijing" />
        </el-select>
      </el-form-item>
      <el-form-item label="活动时间" required>
        <el-col :span="11">
          <el-form-item prop="date1">
            <el-date-picker
              v-model="postForm.id"
              type="date"
              placeholder="选择日期"
              style="width: 100%;"
            />
          </el-form-item>
        </el-col>
        <el-col class="line" :span="2">-</el-col>
        <el-col :span="11">
          <el-form-item prop="date2">
            <el-time-picker v-model="postForm.id" placeholder="选择时间" style="width: 100%;" />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="即时配送" prop="delivery">
        <el-switch v-model="postForm.id" />
      </el-form-item>
      <el-form-item label="活动性质" prop="type">
        <el-checkbox-group v-model="postForm.id">
          <el-checkbox label="美食/餐厅线上活动" name="type" />
          <el-checkbox label="地推活动" name="type" />
          <el-checkbox label="线下主题活动" name="type" />
          <el-checkbox label="单纯品牌曝光" name="type" />
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="特殊资源" prop="resource">
        <el-radio-group v-model="postForm.id">
          <el-radio label="线上品牌商赞助" />
          <el-radio label="线下场地免费" />
        </el-radio-group>
      </el-form-item>
      <el-form-item label="活动形式" prop="desc">
        <el-input v-model="postForm.id" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button
          v-loading="loading"
          size="medium"
          circle
          type="success"
          icon="el-icon-check"
          @click="submitForm"
        />
        <el-button
          v-loading="loading"
          size="medium"
          circle
          type="warning"
          icon="el-icon-close"
          @click="draftForm"
        />
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
// import Sticky from '@/components/sticky/Sticky'

const defaultForm = {
  status: 'draft',
  title: '', // 文章题目
  content: '', // 文章内容
  content_short: '', // 文章摘要
  source_uri: '', // 文章外链
  image_uri: '', // 文章图片
  display_time: undefined, // 前台展示时间
  id: undefined,
  platforms: ['a-platform'],
  comment_disabled: false,
  importance: 0
}

export default {
  name: 'TaskCreate',
  // components: { Sticky },
  data() {
    return {
      postForm: Object.assign({}, defaultForm),
      rules: {},
      loading: false

    }
  },
  methods: {
    draftForm() {
    },
    submitForm() {

    }
  }
}
</script>

<style lang="scss" scoped>
  /*@import "~@/styles/mixin.scss";*/
  .task-form {
    text-align: center;
  }

</style>
