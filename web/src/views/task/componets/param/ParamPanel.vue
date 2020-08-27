<template>
  <div class="param-panel">
    <el-divider content-position="left">渠道参数</el-divider>
    <el-row>
      <el-form-item class="task-info-item" label-width="200px" label="参数1：">
        <!--        <params-input />-->
        <div class="channel-params">
          <span class="params-key">appid</span>
          <span class="params-value" style="margin-left: 40px">12345678</span>
        </div>
      </el-form-item>
    </el-row>
    <el-divider content-position="left">插件参数</el-divider>
    <el-row>
      <el-form-item class="task-info-item" label-width="200px" label="参数1：">
        <div class="channel-params">
          <span class="params-key">appid</span>
          <span class="params-value" style="margin-left: 40px">12345678</span>
        </div>
      </el-form-item>
    </el-row>
    <!--    <el-divider content-position="left">扩展参数</el-divider>-->
    <!--    <el-row>-->
    <!--      <el-form-item class="task-info-item" label-width="200px" label="参数1：">-->
    <!--        <ext-param />-->
    <!--      </el-form-item>-->
    <!--    </el-row>-->
  </div>
</template>

<script>

export default {
  name: 'ParamPanel',
  props: {
    gameName: {
      type: String,
      default: ''
    },
    channelName: {
      type: String,
      default: ''
    },
    pluginName: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss" scoped>
  .channel-params {
    font-size: 20px;
  }
</style>
