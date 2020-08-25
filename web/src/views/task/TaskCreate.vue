<template>
  <div id="task-create">
    <el-form ref="postForm" :model="postForm" :rules="rules" class="task-form">
      <!--      <Sticky :z-index="10" :class-name="'sub-navbar '+postForm.status">-->
      <Sticky :z-index="10" :class-name="'sub-navbar '+postForm.status">

        <!--        <CommentDropdown v-model="postForm.comment_disabled" />-->
        <!--        <PlatformDropdown v-model="postForm.platforms" />-->
        <!--        <SourceUrlDropdown v-model="postForm.source_uri" />-->
        <el-button v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm">
          Publish
        </el-button>
        <el-button v-loading="loading" type="warning" @click="draftForm">
          Draft
        </el-button>
        <div>aaaa</div>
      </Sticky>
    </el-form>
  </div>
</template>

<script>
import Sticky from '@/components/sticky/Sticky'

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
  components: { Sticky },
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
  @import "~@/styles/mixin.scss";

</style>
