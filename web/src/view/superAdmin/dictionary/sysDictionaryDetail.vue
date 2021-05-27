<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Display value">
          <el-input placeholder="search condition" v-model="searchInfo.label"></el-input>
        </el-form-item>
        <el-form-item label="Dictionary value">
          <el-input placeholder="search condition" v-model="searchInfo.value"></el-input>
        </el-form-item>
        <el-form-item label="Enable state" prop="status">
          <el-select v-model="searchInfo.status" placeholder="please choose">
            <el-option key="true" label="Yes" value="true"></el-option>
            <el-option key="false" label="no" value="false"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">Inquire</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">New word type</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="date" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="Display value" prop="label" width="120"></el-table-column>

      <el-table-column label="Dictionary value" prop="value" width="120"></el-table-column>

      <el-table-column label="Enable state" prop="status" width="120">
        <template slot-scope="scope">{{scope.row.status|formatBoolean}}</template>
      </el-table-column>

      <el-table-column label="Sort tag" prop="sort" width="120"></el-table-column>

      <el-table-column label="Button group">
        <template slot-scope="scope">
          <el-button @click="updateSysDictionaryDetail(scope.row)" size="small" type="primary">change</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>You sure you want to delete it?</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">cancel</el-button>
              <el-button type="primary" size="mini" @click="deleteSysDictionaryDetail(scope.row)">determine</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">delete</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="Pop-up operation">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="110px">
        <el-form-item label="Display value" prop="label">
          <el-input
            v-model="formData.label"
            placeholder="Please enter the display value"
            clearable
            :style="{width: '100%'}"
          ></el-input>
        </el-form-item>
        <el-form-item label="Dictionary value" prop="value">
          <el-input-number
            v-model.number="formData.value"
            step-strictly
            :step="1"
            placeholder="Please enter a dictionary value"
            clearable
            :style="{width: '100%'}"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="Enable state" prop="status" required>
          <el-switch v-model="formData.status" active-text="Open" inactive-text="Deactivate"></el-switch>
        </el-form-item>
        <el-form-item label="Sort tag" prop="sort">
          <el-input-number v-model.number="formData.sort" placeholder="Sort tag"></el-input-number>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">Take</el-button>
        <el-button @click="enterDialog" type="primary">Confirm</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createSysDictionaryDetail,
  deleteSysDictionaryDetail,
  updateSysDictionaryDetail,
  findSysDictionaryDetail,
  getSysDictionaryDetailList
} from "@/api/sysDictionaryDetail"; //  Please replace the address here
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "SysDictionaryDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getSysDictionaryDetailList,
      dialogFormVisible: false,
      type: "",
      formData: {
        label: null,
        value: null,
        status: true,
        sort: null
      },
      rules: {
        label: [
          {
            required: true,
            message: "Please enter the display value",
            trigger: "blur"
          }
        ],
        value: [
          {
            required: true,
            message: "Please enter a dictionary value",
            trigger: "blur"
          }
        ],
        sort: [
          {
            required: true,
            message: "Sort tag",
            trigger: "blur"
          }
        ]
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
      if (this.searchInfo.status == "") {
        this.searchInfo.status = null;
      }
      this.getTableData();
    },
    async updateSysDictionaryDetail(row) {
      const res = await findSysDictionaryDetail({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.resysDictionaryDetail;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        label: null,
        value: null,
        status: true,
        sort: null,
        sysDictionaryID: ""
      };
    },
    async deleteSysDictionaryDetail(row) {
      row.visible = false;
      const res = await deleteSysDictionaryDetail({ ID: row.ID });
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
    async enterDialog() {
      this.formData.sysDictionaryID = Number(this.$route.params.id);
      this.$refs["elForm"].validate(async valid => {
        if (!valid) return;
        let res;
        switch (this.type) {
          case "create":
            res = await createSysDictionaryDetail(this.formData);
            break;
          case "update":
            res = await updateSysDictionaryDetail(this.formData);
            break;
          default:
            res = await createSysDictionaryDetail(this.formData);
            break;
        }
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "Create / change success"
          });
          this.closeDialog();
          this.getTableData();
        }
      });
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.searchInfo.sysDictionaryID = Number(this.$route.params.id);
    this.getTableData();
  }
};
</script>

<style>
</style>