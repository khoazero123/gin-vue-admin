<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Request method">
          <el-input placeholder="search condition" v-model="searchInfo.method"></el-input>
        </el-form-item>
        <el-form-item label="Request path">
          <el-input placeholder="search condition" v-model="searchInfo.path"></el-input>
        </el-form-item>
        <el-form-item label="Result status code">
          <el-input placeholder="search condition" v-model="searchInfo.status"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">Inquire</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>You sure you want to delete it?</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text">cancel</el-button>
              <el-button @click="onDelete" size="mini" type="primary">determine</el-button>
            </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">batch deletion</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="Operator" width="140">
        <template slot-scope="scope">
          <div>{{scope.row.user.userName}}({{scope.row.user.nickName}})</div>
        </template>
      </el-table-column>
      <el-table-column label="date" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="status code" prop="status" width="120">
        <template slot-scope="scope">
          <div>
            <el-tag type="success">{{ scope.row.status }}</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Request IP" prop="ip" width="120"></el-table-column>
      <el-table-column label="Request method" prop="method" width="120"></el-table-column>
      <el-table-column label="Request path" prop="path" width="240"></el-table-column>
      <el-table-column label="request" prop="path" width="80">
        <template slot-scope="scope">
          <div>
            <el-popover placement="top-start" trigger="hover" v-if="scope.row.body">
              <div class="popover-box">
                <pre>{{fmtBody(scope.row.body)}}</pre>
              </div>
              <i class="el-icon-view" slot="reference"></i>
            </el-popover>

            <span v-else>no</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="response" prop="path" width="80">
        <template slot-scope="scope">
          <div>
            <el-popover placement="top-start" trigger="hover" v-if="scope.row.resp">
              <div class="popover-box">
                <pre>{{fmtBody(scope.row.resp)}}</pre>
              </div>
              <i class="el-icon-view" slot="reference"></i>
            </el-popover>
            <span v-else>no</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-popover placement="top" v-model="scope.row.visible" width="160">
            <p>You sure you want to delete it?</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="scope.row.visible = false" size="mini" type="text">cancel</el-button>
              <el-button @click="deleteSysOperationRecord(scope.row)" size="mini" type="primary">determine</el-button>
            </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">delete</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
import {
  deleteSysOperationRecord,
  getSysOperationRecordList,
  deleteSysOperationRecordByIds
} from "@/api/sysOperationRecord"; //  Please replace the address here
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "SysOperationRecord",
  mixins: [infoList],
  data() {
    return {
      listApi: getSysOperationRecordList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        ip: null,
        method: null,
        path: null,
        status: null,
        latency: null,
        agent: null,
        error_message: null,
        user_id: null
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "Yes" : "no";
      } else {
        return "";
      }
    }
  },
  methods: {
    //Condition Search Front End View this method
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    async onDelete() {
      const ids = [];
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID);
        });
      const res = await deleteSysOperationRecordByIds({ ids });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "successfully deleted"
        });
        if (this.tableData.length == ids.length && this.page > 1) {
          this.page--;
        }
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    async deleteSysOperationRecord(row) {
      row.visible = false;
      const res = await deleteSysOperationRecord({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "successfully deleted"
        });
        if (this.tableData.length == 1 && this.page > 1 ) {
          this.page--;
        }
        this.getTableData();
      }
    },
    fmtBody(value) {
      try {
        return JSON.parse(value);
      } catch (err) {
        return value;
      }
    }
  },
  created() {
    this.getTableData();
  }
};
</script>

<style lang="scss">
.table-expand {
  padding-left: 60px;
  font-size: 0;
  label {
    width: 90px;
    color: #99a9bf;
    .el-form-item {
      margin-right: 0;
      margin-bottom: 0;
      width: 50%;
    }
  }
}
.popover-box {
  background: #112435;
  color: #f08047;
  height: 600px;
  width: 420px;
  overflow: auto;
}
.popover-box::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>