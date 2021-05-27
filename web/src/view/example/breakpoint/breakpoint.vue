<template>
  <div class="hello">
     <el-divider content-position="left">Big file upload</el-divider>
    <form id="fromCont" method="post" >
      <div class="fileUpload" @click="inputChange">
        Select a document
        <input v-show="false"  @change="choseFile" id="file" multiple="multiple" type="file" ref="Input"  />
      </div>
    </form>
     <el-button @click="getFile" :disabled="limitFileSize" type="primary" size="medium" class="uploadBtn">upload files</el-button>
    <div class="el-upload__tip">Please upload no more than 5MB files</div>
    <div class="list">
      <transition  name="list" tag="p">
        <div class="list-item" v-if="file" >
          <i class="el-icon-document"></i>
          <span>{{ file.name }}</span>
          <span class="percentage" >{{percentage}}%</span>
          <el-progress  :show-text='false' :text-inside="false" :stroke-width="2" :percentage="percentage"></el-progress>
        </div> 
      </transition>
   </div>
     <!-- <span
      v-if="this.file"
    >{{Math.floor(((this.formDataList.length-this.waitNum)/this.formDataList.length)*100)}}%</span> -->
    <div class="tips">This version is the first experience function beta, style beautification, and performance optimization is in progress, upload sliced files and synthetic full files separately QMPLUSSERVER directory BreakPointDir folders and FileDir folders</div>
  </div>
</template>
<script>
import SparkMD5 from 'spark-md5'
import axios from 'axios'
import {
  findFile,
  breakpointContinueFinish,
  removeChunk
} from '@/api/breakpoint'
export default {
  name: 'HelloWorld',
  data() {
    return {
      file: null,
      fileMd5: '',
      formDataList: [],
      waitUpLoad: [],
      waitNum: 0,
      limitFileSize: false,
      percentage:0,
      percentageFlage: true,
      customColor: '#409eff'
    }
  },
  created(){
   
  },
  methods: {
    // Select the function of the file
    async choseFile(e) {
      const fileR = new FileReader() // 创建一个reader用来读取文件流
      const file = e.target.files[0] // 获取当前文件
      const maxSize = 5*1024*1024
      this.file = file // file 丢全局方便后面用 可以改进为func传参形式
      this.percentage = 0
    if(file.size<maxSize){
      fileR.readAsArrayBuffer(file) // 把文件读成ArrayBuffer  主要为了保持跟后端的流一致
      fileR.onload = async e => {
        // 读成arrayBuffer的回调 e 为方法自带参数 相当于 dom的e 流存在e.target.result 中
        const blob = e.target.result
        let spark = new SparkMD5.ArrayBuffer() // 创建md5制造工具 （md5用于检测文件一致性 这里不懂就打电话问我）
        spark.append(blob) // 文件流丢进工具
        this.fileMd5 = spark.end() // 工具结束 产生一个a 总文件的md5
        const FileSliceCap = 1 * 1024 * 1024 // 分片字节数
        let start = 0 // Define fragmentation to start cutting
        let end = 0 // Each piece is cut into place A
        let i = 0 // Second
        this.formDataList = [] // Piece of fragmentation is lost
        while (end < file.size) {
          // 当结尾数字大于文件总size的时候 结束切片
          start = i * FileSliceCap // 计算每片开始位置
          end = (i + 1) * FileSliceCap // 计算每片结束位置
          var fileSlice = this.file.slice(start, end) // 开始切  file.slice 为 h5方法 对文件切片 参数为 起止字节数
          const formData = new window.FormData() // 创建FormData用于存储传给后端的信息
          formData.append('fileMd5', this.fileMd5) // 存储总文件的Md5 让后端知道自己是谁的切片
          formData.append('file', fileSlice) //当前的切片
          formData.append('chunkNumber', i) // 当前是第几片
          formData.append('fileName', this.file.name) //当前文件的文件名 用于后端文件切片的命名  formData.appen 为 formData对象添加参数的方法
          this.formDataList.push({ key: i, formData }) // 把当前切片信息 自己是第几片 存入我们方才准备好的池子
          i++
        }
        const params = {
          fileName: this.file.name,
          fileMd5: this.fileMd5,
          chunkTotal: this.formDataList.length
        }
        const res = await findFile(params)
        // 全部切完以后 发一个请求给后端 拉当前文件后台存储的切片信息 用于检测有多少上传成功的切片
        const finishList = res.data.file.ExaFileChunk // 上传成功的切片
        const IsFinish = res.data.file.IsFinish // 是否是同文件不同命 （文件md5相同 文件名不同 则默认是同一个文件但是不同文件名 此时后台数据库只需要拷贝一下数据库文件即可 不需要上传文件 即秒传功能）
        if (!IsFinish) {
          // 当是断点续传时候
          this.waitUpLoad = this.formDataList.filter(all => {
            return !(
              finishList &&
              finishList.some(fi => fi.FileChunkNumber === all.key)
            ) // 找出需要上传的切片
          })
        } else {
          this.waitUpLoad = [] // 秒传则没有需要上传的切片
        }
        this.waitNum = this.waitUpLoad.length // 记录长度用于百分比展示
      }
      } else {
         this.limitFileSize = true
         this.$message('Please upload less than 5m files')
      }
    },
    getFile() {
      // Confirm button
      if (this.file == null) {
        this.$message('Please upload the file first')
        return
      }
      this.percentage = Math.floor(((this.formDataList.length-this.waitNum)/this.formDataList.length)*100)
      if(this.percentage == 100){
        this.percentageFlage = false
      }
      this.sliceFile() // Upload slice
    },
    sliceFile() {
      this.waitUpLoad &&
        this.waitUpLoad.map(item => {
          //Need to upload the slice
          item.formData.append('chunkTotal', this.formDataList.length) // The total number of slits is carried to the background.
          const fileR = new FileReader() // Functionally
          const file = item.formData.get('file')
          fileR.readAsArrayBuffer(file)
          fileR.onload = e => {
            let spark = new SparkMD5.ArrayBuffer()
            spark.append(e.target.result)
            item.formData.append('chunkMd5', spark.end()) // Get the current slice MD5 rear end for verifying the finish
            this.upLoadFileSlice(item)
          }
        })
    },
    async upLoadFileSlice(item) {
      // Sliced upload
      await axios.post(process.env.VUE_APP_BASE_API+"/fileUploadAndDownload/breakpointContinue",item.formData)
      this.waitNum-- // Percentage
      if (this.waitNum == 0) {
        // Sliced files after passing
        const params = {
          fileName: this.file.name,
          fileMd5: this.fileMd5
        }
        const res = await breakpointContinueFinish(params)
        if (res.success) {
          // After the synthesis file, remove the cache slice
          const params = {
            fileName: this.file.name,
            fileMd5: this.fileMd5,
            filePath: res.data.filePath
          }
          await removeChunk(params)
        }
      }
    },
    inputChange(){
      this.$refs.Input.dispatchEvent(new MouseEvent('click'))
    }
  }
}
</script>

<style lang='scss' scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#fromCont{
  display: inline-block;
}
.fileUpload{
    padding: 4px 10px;
    height: 20px;
    line-height: 20px;
    position: relative;
    cursor: pointer;
    color: #000;
    border: 1px solid #c1c1c1;
    border-radius: 4px;
    overflow: hidden;
    display: inline-block;
    input{
      position: absolute;
      font-size: 100px;
      right: 0;
      top: 0;
      opacity: 0;
      cursor: pointer;
    }
}
 .fileName{
    display: inline-block;
    vertical-align: top;
    margin: 6px 15px 0 15px;
  }
  .uploadBtn{
    position: relative;
    top: -10px;
    margin-left: 15px;
  }
  .tips{
    margin-top: 30px;
    font-size: 14px;
    font-weight: 400;
    color: #606266;
  }
  .el-divider{
    margin: 0 0 30px 0;
  }
 
 .list{
   margin-top:15px;
 }
 .list-item {
  display: block;
  margin-right: 10px;
  color: #606266;
  line-height: 25px;
  margin-bottom: 5px;
  width: 40%;
   .percentage{
          float: right;
        }
}
.list-enter-active, .list-leave-active {
  transition: all 1s;
}
.list-enter, .list-leave-to
/* .list-leave-active for below version 2.1.8 */ {
  opacity: 0;
  transform: translateY(-30px);
}
</style>