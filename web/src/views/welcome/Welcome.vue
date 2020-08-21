<template>
  <div id="welcome">
    <span class="welcome-tips">{{ name }}，欢迎使用Linking分发服务</span>
    <br>
    <span class="welcome-ip">当前登录ip：{{ name }}</span>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Welcome',
  computed: {
    ...mapGetters([
      'name'
    ])
  }
}
</script>

<style lang="scss" scoped>
  #welcome {
    margin: 30px;

    .welcome-tips {
      font-size: 30px;
      line-height: 46px;
    }

    .welcome-ip {
      font-size: 20px;
      line-height: 46px;
    }
  }
</style>
