const OSS = require('ali-oss')

// LTAI4GLCwofVysdJqYgFF5sQ
// XP8nyjdTL55DIwMv7dB7IdMaY3I0Bo
// const client = oss({
//   accessKeyId: 'LTAI4GLCwofVysdJqYgFF5sQ',
//   accessKeySecret: 'XP8nyjdTL55DIwMv7dB7IdMaY3I0Bo',
//   bucket: 'sdkfile',
//   region: 'oss-cn-beijing'
// })

const client = new OSS({
  accessKeyId: 'LTAI4GLCwofVysdJqYgFF5sQ',
  accessKeySecret: 'XP8nyjdTL55DIwMv7dB7IdMaY3I0Bo',
  // bucket: 'sdkfile',
  region: 'oss-cn-beijing'
})

export async function testOSS() {
  try {
    const result = await client.listBuckets()
    console.log(result.bucket.Location)
  } catch (err) {
    console.log(err)
  }
}

