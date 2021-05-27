<template>
  <uploader
    :options="options"
    :file-status-text="statusText"
    :autoStart="false"
    @file-added="fileAdded"
    @file-progress="onFileProgress"
    @file-success="onFileSuccess"
    @file-error="onFileError"
    class="uploader-example"
  >
    <uploader-unsupport></uploader-unsupport>
    <uploader-drop>
      <p>Drag and drop files or click</p>
      <uploader-btn>Select a document</uploader-btn>
    </uploader-drop>
    <uploader-list></uploader-list>
  </uploader>
</template>

<script>
var notUploadedChunks = []; // 已经上传过的文件chunkNumber数组
var isUploaded = false; // The file has been uploaded successfully.
import { mapGetters } from "vuex";
import { checkFileMd5,mergeFileMd5 } from "@/api/simpleUploader";
import SparkMD5 from "spark-md5";
const path = process.env.VUE_APP_BASE_API;
export default {
  name: "simpleUploader",
  data(){
    return{
      md5:""
    }
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"]),
    statusText() {
      return {
        success: "Successful",
        error: "Error",
        uploading: "Uploading",
        paused: "Paused",
        waiting: "Waiting"
      };
    },
    options() {
      return {
        target: path + "/simpleUploader/upload",
        testChunks: false,
        simultaneousUploads: 5,
        chunkSize: 2 * 1024 * 1024,
        headers: {
          "x-token": this.token,
          "x-user-id": this.userInfo.ID
        },
        checkChunkUploadedByResponse(chunk) {
          if (isUploaded) {
            return true; // return true Will ignore the current file, will not be sent to the background
          } else {
              // Ignore according to the slice that has been uploaded.
              return (
                notUploadedChunks &&
                notUploadedChunks.some(
                  item => item.chunkNumber == chunk.offset + 1
                )
              );
          }
        }
      };
    }
  },
  methods: {
    // Upload a single file
    fileAdded(file) {
      this.computeMD5(file); // Generate MD5
    },
    // Calculate MD5 value
    computeMD5(file) {
      var that = this;
      isUploaded = false; // Whether this file has been uploaded
      notUploadedChunks = []; // Unsuccessful Chunknumber
      var fileReader = new FileReader();
      var md5 = "";

      file.pause();

      fileReader.readAsArrayBuffer(file.file);
      fileReader.onload = async function(e) {
        if (file.size != e.target.result.byteLength) {
          this.error(
            "Browser reported success but could not read the file until the end."
          );
          return false;
        }
        md5 = SparkMD5.ArrayBuffer.hash(e.target.result, false);

        file.uniqueIdentifier = md5;
        if (md5 != "") {
          const res = await checkFileMd5({ md5: md5 });
          if (res.code == 0) {
            if (res.data.isDone) {
              // Upload success
              isUploaded = true;
              that.$message({
                message: "The file has been successfully uploaded, and the second is successful.",
                type: "success"
              });

              file.cancel();
            } else {
              isUploaded = false;
              notUploadedChunks = res.data.chunks;
              if(notUploadedChunks.length){
                file.resume();
              }
            }
          }
        }

        
      };
      fileReader.onerror = function() {
        this.error(
          "When Generate MD5, FileReader reads files.，FileReader onerror was triggered, maybe the browser aborted due to high memory usage."
        );
        return false;
      };
    },
    // Upload progress
    onFileProgress() {},
    // Upload success
    async onFileSuccess(rootFile, file) {
      await mergeFileMd5({md5:file.uniqueIdentifier,fileName:file.name})
    },
    onFileError(rootFile, file, response) {
      this.$message({
        message: response,
        type: "error"
      });
    }
  }
};
</script>

<style>
.uploader-example {
  width: 880px;
  padding: 15px;
  margin: 115px 15px 20px;
  font-size: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
}
.uploader-example .uploader-btn {
  margin-right: 4px;
}
.uploader-example .uploader-list {
  max-height: 440px;
  overflow: auto;
  overflow-x: hidden;
  overflow-y: auto;
}
</style>