<template>
  <div class="param-input">
    <!--    <el-row>-->
    <!--      <el-col :span="10">-->
    <!--        <el-input-->
    <!--          v-model="inputKey"-->
    <!--          class="param-key"-->
    <!--          placeholder="请输入参数名"-->
    <!--          clearable-->
    <!--        />-->
    <!--      </el-col>-->
    <!--      <el-col :span="10">-->
    <!--        <el-input-->
    <!--          v-model="inputValue"-->
    <!--          class="param-value"-->
    <!--          placeholder="请输入参数值"-->
    <!--          clearable-->
    <!--        />-->
    <!--      </el-col>-->
    <!--    </el-row>-->
    <el-input
      v-model="inputKey"
      class="param-key"
      placeholder="请输入参数名"
      clearable
    />
    <el-input
      v-model="inputValue"
      class="param-value"
      placeholder="请输入参数值"
      clearable
    />
  </div>
</template>

<script>
const defaultParam = {
  paramKey: '',
  paramValue: ''
}
export default {
  name: 'ParamsInput',
  model: {
    prop: 'param',
    event: 'inputParam'
  },
  props: {
    param: {
      type: Object,
      default: defaultParam
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      tempParam: Object.assign({}, defaultParam)
    }
  },
  watch: {
    tempParam(value) {
      this.$emit('inputParam', value)
    }
  }
}
</script>

<style lang="scss" scoped>
  /*.param-input {*/
  /*  width: 100%;*/
  /*  text-align: center;*/

  .param-key,
  .param-value {
    /*display: block;*/
    /*position: relative;*/
    width: 350px;
  }

  .param-value {
    margin-left: 20px;
  }

</style>
