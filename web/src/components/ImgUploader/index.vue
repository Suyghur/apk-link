<template>
  <div class="uploader">
    <el-upload
      class="avatar-uploader"
      action="#"
      :show-file-list="false"
      :on-success="handleSuccess"
      :before-upload="beforeUpload"
      :http-request="handleUpload"
      :disabled="disabled"
    >
      <img v-if="url" :src="url" class="img-uploader">
      <i v-else class="el-icon-plus img-uploader-icon" />
    </el-upload>
  </div>
</template>
<script>
import { put2oss } from '@/utils/aliyun-oss'

export default {
  name: 'ImgUploader',
  model: {
    prop: 'resultUrl',
    event: 'uploadSuccess'
  },
  props: {
    resultUrl: {
      type: String,
      default: ''
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      url: ''
    }
  },
  watch: {
    url(value) {
      this.$emit('uploadSuccess', value)
    }
  },
  methods: {

    handleSuccess(res, file) {
      this.url = URL.createObjectURL(file.raw)
    },
    beforeUpload(file) {
      const isJPG = file.type === 'image/jpeg'
      const isPNG = file.type === 'image/png'
      const isLt5M = file.size / 1024 / 1024 < 5
      if (!isLt5M) {
        this.$message.error('上传图片大小不能超过5MB!')
      }
      if (!isJPG && !isPNG) {
        this.$message.error('图片只能是JPG或PNG格式!')
      }
      console.log((isJPG || isPNG) && isLt5M)
      return (isJPG || isPNG) && isLt5M
    },
    handleUpload(option) {
      console.log(option)
      put2oss(option.file).then(res => {
        this.url = res.url
      })
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

  .img-uploader-icon {
    font-size: 28px;
    color: #03a9f4;
    width: 150px;
    height: 150px;
    line-height: 150px;
    text-align: center;
  }

  .img-uploader {
    width: 150px;
    height: 150px;
    display: block;
  }
</style>
