<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Dictionary name (middle)">
          <el-input placeholder="search condition" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="Dictionary name (English)">
          <el-input placeholder="search condition" v-model="searchInfo.type"></el-input>
        </el-form-item>
        <el-form-item label="status" prop="status">
          <el-select v-model="searchInfo.status" clear placeholder="please choose">
            <el-option key="true" label="Yes" value="true"></el-option>
            <el-option key="false" label="no" value="false"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="description">
          <el-input placeholder="search condition" v-model="searchInfo.desc"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">Inquire</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">New dictionary</el-button>
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

      <el-table-column label="Dictionary name (middle)" prop="name" width="120"></el-table-column>

      <el-table-column label="Dictionary name (English)" prop="type" width="120"></el-table-column>

      <el-table-column label="status" prop="status" width="120">
        <template slot-scope="scope">{{scope.row.status|formatBoolean}}</template>
      </el-table-column>

      <el-table-column label="description" prop="desc" width="280"></el-table-column>

      <el-table-column label="Button group">
        <template slot-scope="scope">
          <el-button @click="toDetile(scope.row)" size="small" type="success">Detail</el-button>
          <el-button @click="updateSysDictionary(scope.row)" size="small" type="primary">change</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>You sure you want to delete it?</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">cancel</el-button>
              <el-button type="primary" size="mini" @click="deleteSysDictionary(scope.row)">determine</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference" style="margin-left:10px">delete</el-button>
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
        <el-form-item label="Dictionary name (middle)" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="Please enter a dictionary name (medium)"
            clearable
            :style="{width: '100%'}"
          ></el-input>
        </el-form-item>
        <el-form-item label="Dictionary name (English)" prop="type">
          <el-input
            v-model="formData.type"
            placeholder="Please enter a dictionary (English)"
            clearable
            :style="{width: '100%'}"
          ></el-input>
        </el-form-item>
        <el-form-item label="status" prop="status" required>
          <el-switch v-model="formData.status" active-text="Open" inactive-text="Deactivate"></el-switch>
        </el-form-item>
        <el-form-item label="description" prop="desc">
          <el-input v-model="formData.desc" placeholder="Please enter a description" clearable :style="{width: '100%'}"></el-input>
        </el-form-item>
      </el-form>

      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">Take</el-button>
        <el-button @click="enterDialog" type="primary">Confirm</el-button>
      </div>
    </el-dialog>

    <div style="margin-top:40px;color:red">Get a dictionary and the cache method is already on the front end utils/dictionary Has been packaged, do not have to write usage, see files</div>
  </div>
</template>

<script>
import {
  createSysDictionary,
  deleteSysDictionary,
  updateSysDictionary,
  findSysDictionary,
  getSysDictionaryList
} from "@/api/sysDictionary"; //  Please replace the address here
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "SysDictionary",
  mixins: [infoList],
  data() {
    return {
      listApi: getSysDictionaryList,
      dialogFormVisible: false,
      type: "",
      formData: {
        name: null,
        type: null,
        status: true,
        desc: null
      },
      rules: {
        name: [
          {
            required: true,
            message: "Please enter a dictionary name (medium)",
            trigger: "blur"
          }
        ],
        type: [
          {
            required: true,
            message: "Please enter a dictionary (English)",
            trigger: "blur"
          }
        ],
        desc: [
          {
            required: true,
            message: "Please enter a description",
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
    toDetile(row) {
      this.$router.push({
        name: "dictionaryDetail",
        params: {
          id: row.ID
        }
      });
    },
    //Condition Search Front End View this method
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      if (this.searchInfo.status == "") {
        this.searchInfo.status = null;
      }
      this.getTableData();
    },
    async updateSysDictionary(row) {
      const res = await findSysDictionary({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.resysDictionary;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        name: null,
        type: null,
        status: true,
        desc: null
      };
    },
    async deleteSysDictionary(row) {
      row.visible = false;
      const res = await deleteSysDictionary({ ID: row.ID });
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
      this.$refs["elForm"].validate(async valid => {
        if (!valid) return;
        let res;
        switch (this.type) {
          case "create":
            res = await createSysDictionary(this.formData);
            break;
          case "update":
            res = await updateSysDictionary(this.formData);
            break;
          default:
            res = await createSysDictionary(this.formData);
            break;
        }
        if (res.code == 0) {
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
  async created() {
    this.getTableData();
  }
};
</script>

<style>
</style>