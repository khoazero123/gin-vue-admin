<template>
  <div class="upload">
    <el-row>
      <el-col :span="2">
        <el-upload
          :action="`${path}/excel/importExcel`"
          :headers="{'x-token':token}"
          :on-success="loadExcel"
          :show-file-list="false"
        >
          <el-button size="small" type="primary" icon="el-icon-upload2">Import</el-button>
        </el-upload>
      </el-col>
      <el-col :span="2">
        <el-button size="small" type="primary" icon="el-icon-download" @click="handleExcelExport('ExcelExport.xlsx')">Export</el-button>
      </el-col>
      <el-col :span="2">
        <el-button size="small" type="success" icon="el-icon-download" @click="downloadExcelTemplate()">Download template</el-button>
      </el-col>
    </el-row>
    <el-table :data="tableData" border row-key="ID" stripe>
      <el-table-column label="ID" min-width="100" prop="ID"></el-table-column>
      <el-table-column label="Route Name" min-width="160" prop="name"></el-table-column>
      <el-table-column label="Routing Path" min-width="160" prop="path"></el-table-column>
      <el-table-column label="Whether hidden" min-width="100" prop="hidden">
        <template slot-scope="scope">
          <span>{{scope.row.hidden?"hide":"display"}}</span>
        </template>
      </el-table-column>
      <el-table-column label="Parent" min-width="90" prop="parentId"></el-table-column>
      <el-table-column label="Sort" min-width="70" prop="sort"></el-table-column>
      <el-table-column label="file path" min-width="360" prop="component"></el-table-column>
    </el-table>
  </div>
</template>
<script>
const path = process.env.VUE_APP_BASE_API;
import { mapGetters } from 'vuex';
import infoList from "@/mixins/infoList";
import { exportExcel, loadExcelData, downloadTemplate } from "@/api/excel";
import { getMenuList } from "@/api/menu";
export default {
  name: 'Excel',
  mixins: [infoList],
  data() {
    return {
      listApi: getMenuList,
      path: path
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods: {
    handleExcelExport(fileName) {
      if (!fileName || typeof fileName !== "string") {
        fileName = "ExcelExport.xlsx";
      }
      exportExcel(this.tableData, fileName);
    },
    loadExcel() {
      this.listApi = loadExcelData;
      this.getTableData();
    },
    downloadExcelTemplate() {
      downloadTemplate('ExcelTemplate.xlsx')
    }
  },
  created() {
    this.pageSize = 999;
    this.getTableData();
  }
}
</script>