const OSS = require('ali-oss')

const client = new OSS({
  accessKeyId: 'LTAI4GLCwofVysdJqYgFF5sQ',
  accessKeySecret: 'XP8nyjdTL55DIwMv7dB7IdMaY3I0Bo',
  bucket: 'sdkfile',
  region: 'oss-cn-beijing',
  endpoint: 'sdkfile.hihoulang.com',
  cname: true
})

const linkingDir = '/linking/'

export async function listBuckets() {
  try {
    // console.log(client)
    const result = await client.list()
    console.log(result)
  } catch (err) {
    console.log(err)
  }
}

export async function put2oss(file) {
  try {
    // object-key可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。
    const filePath = linkingDir + file.uid + '_' + file.name
    return await client.put(filePath, file)
  } catch (e) {
    console.log(e)
  }
}
