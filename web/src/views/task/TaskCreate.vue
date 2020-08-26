<template>
  <div id="task-create">
    <TaskInfo :disabled="false" />
  </div>
</template>

<script>
// import Sticky from '@/components/sticky/Sticky'

import TaskInfo from '@/views/task/componets/TaskInfo'

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
  components: { TaskInfo },
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
