<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="upload">
      <el-row>
        <el-col :span="12">
          <el-upload
            :action="`${path}/fileUploadAndDownload/upload`"
            :before-upload="checkFile"
            :headers="{ 'x-token': token }"
            :on-error="uploadError"
            :on-success="uploadSuccess"
            :show-file-list="false"
          >
            <el-button size="small" type="primary">Click to upload</el-button>
            <div class="el-upload__tip" slot="tip">Only upload JPG / PNG files，Not more than 500KB</div>
          </el-upload>
        </el-col>
        <el-col :span="12">
          Uploaded with compression (512 (k) is compressed)
          <upload-image v-model="imageUrl" :fileSize="512" :maxWH="1080" />
          Upload file {{ imageUrl }}
        </el-col>
      </el-row>

      <el-table :data="tableData" border stripe>
        <el-table-column label="Preview" width="100">
          <template slot-scope="scope">
            <CustomPic picType="file" :picSrc="scope.row.url" />
          </template>
        </el-table-column>
        <el-table-column label="date" prop="UpdatedAt" width="180">
          <template slot-scope="scope">
            <div>{{ scope.row.UpdatedAt | formatDate }}</div>
          </template>
        </el-table-column>
        <el-table-column label="file name" prop="name" width="180"></el-table-column>
        <el-table-column label="link" prop="url" min-width="300"></el-table-column>
        <el-table-column label="label" prop="tag" width="100">
          <template slot-scope="scope">
            <el-tag
              :type="scope.row.tag === 'jpg' ? 'primary' : 'success'"
              disable-transitions
            >{{ scope.row.tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="operating" width="160">
          <template slot-scope="scope">
            <el-button @click="downloadFile(scope.row)" size="small" type="text">download</el-button>
            <el-button @click="deleteFile(scope.row)" size="small" type="text">delete</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{ float: 'right', padding: '20px' }"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
      ></el-pagination>
    </div>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import { mapGetters } from "vuex";
import infoList from "@/mixins/infoList";
import { getFileList, deleteFile } from "@/api/fileUploadAndDownload";
import { downloadImage } from "@/utils/downloadImg";
import { formatTimeToStr } from "@/utils/date";
import CustomPic from "@/components/customPic";
import UploadImage from "@/components/upload/image.vue";
export default {
  name: "Upload",
  mixins: [infoList],
  components: {
    CustomPic,
    UploadImage
  },
  data() {
    return {
      fullscreenLoading: false,
      listApi: getFileList,
      path: path,
      tableData: [],
      imageUrl: ""
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"])
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    }
  },
  methods: {
    async deleteFile(row) {
      this.$confirm("This will be permanently filed, Whether to continue?", "prompt", {
        confirmButtonText: "determine",
        cancelButtonText: "cancel",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteFile(row);
          if (res.code == 0) {
            this.$message({
              type: "success",
              message: "successfully deleted!"
            });
            if (this.tableData.length == 1 && this.page > 1 ) {
              this.page--;
            }
            this.getTableData();
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "Delete"
          });
        });
    },
    checkFile(file) {
      this.fullscreenLoading = true;
      const isJPG = file.type === "image/jpeg";
      const isPng = file.type === "image/png";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG && !isPng) {
        this.$message.error("Upload avatar pictures can only be JPG or PNG format!");
        this.fullscreenLoading = false;
      }
      if (!isLt2M) {
        this.$message.error("Upload avatar image size can not exceed 2MB!");
        this.fullscreenLoading = false;
      }
      return (isPng || isJPG) && isLt2M;
    },
    uploadSuccess(res) {
      this.fullscreenLoading = false;
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "Upload success"
        });
        if (res.code == 0) {
          this.getTableData();
        }
      } else {
        this.$message({
          type: "warning",
          message: res.msg
        });
      }
    },
    uploadError() {
      this.$message({
        type: "error",
        message: "upload failed"
      });
      this.fullscreenLoading = false;
    },
    downloadFile(row) {
      downloadImage(row.url, row.name);
    }
  },
  created() {
    this.getTableData();
  }
};
</script>
