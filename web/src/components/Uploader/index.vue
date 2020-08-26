<template>
  <div class="uploader">
    <el-upload
      class="avatar-uploader"
      action="https://jsonplaceholder.typicode.com/posts/"
      :show-file-list="false"
      :on-success="handleAvatarSuccess"
      :before-upload="beforeAvatarUpload"
    >
      <img v-if="url" :src="url" class="avatar">
      <i v-else class="el-icon-plus avatar-uploader-icon" />
    </el-upload>
  </div>
</template>
<script>
export default {
  name: 'Uploader',
  // props: {
  //   url: {
  //     type: String,
  //     default: ''
  //   }
  //   // visible: {
  //   //   type: Boolean,
  //   //   default: false
  //   // },
  //   // disabled: {
  //   //   type: Boolean,
  //   //   default: false
  //   // }
  //
  // },
  data() {
    return {
      url: ''
    }
  },
  methods: {
    handleAvatarSuccess(res, file) {
      this.url = URL.createObjectURL(file.raw)
      console.log(this.url)
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 2

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return isJPG && isLt2M
    }
  }
}

</script>

<style lang="scss" scoped>
  .uploader {
    border: 1px dashed #a8a8a8;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }

  .uploader:hover {
    border-color: orange;
  }

  .avatar-uploader-icon {
    font-size: 28px;
    color: #03a9f4;
    width: 150px;
    height: 150px;
    line-height: 150px;
    text-align: center;
  }

  .avatar {
    width: 150px;
    height: 150px;
    display: block;
  }
</style>
