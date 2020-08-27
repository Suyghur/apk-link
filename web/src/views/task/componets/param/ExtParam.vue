<template>
  <div>
    <params-input />
    <el-button>增加</el-button>
    <el-button>删除</el-button>
  </div>
</template>

<script>
import ParamsInput from '@/views/task/componets/param/ParamsInput'

export default {
  name: 'ExtParam',
  components: { ParamsInput }
}
</script>

<style scoped>

</style>
