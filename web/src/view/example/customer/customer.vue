<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">New customer</el-button>
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
      <el-table-column label="Access date" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="Name" prop="customerName" width="120"></el-table-column>
      <el-table-column label="Phone" prop="customerPhoneData" width="120"></el-table-column>
      <el-table-column label="Access person ID" prop="sysUserId" width="120"></el-table-column>
      <el-table-column label="Button group" min-width="160">
        <template slot-scope="scope">
          <el-button @click="updateCustomer(scope.row)" size="small" type="text">change</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>You sure you want to delete it?</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">cancel</el-button>
              <el-button type="primary" size="mini" @click="deleteCustomer(scope.row)">determine</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="client">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="Customer name">
          <el-input autocomplete="off" v-model="form.customerName"></el-input>
        </el-form-item>
        <el-form-item label="Client phone">
          <el-input autocomplete="off" v-model="form.customerPhoneData"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">Take</el-button>
        <el-button @click="enterDialog" type="primary">Confirm</el-button>
      </div>
    </el-dialog>
    <div class="tips">In the resource permission, the resource permission of this role is emptied or does not include the display of this customer resource.</div>
  </div>
</template>

<script>
import {
  createExaCustomer,
  updateExaCustomer,
  deleteExaCustomer,
  getExaCustomer,
  getExaCustomerList
} from "@/api/customer";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "Customer",
  mixins: [infoList],
  data() {
    return {
      listApi: getExaCustomerList,
      dialogFormVisible: false,
      type: "",
      form: {
        customerName: "",
        customerPhoneData: ""
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
    }
  },
  methods: {
    async updateCustomer(row) {
      const res = await getExaCustomer({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.form = res.data.customer;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.form = {
        customerName: "",
        customerPhoneData: ""
      };
    },
    async deleteCustomer(row) {
      row.visible = false;
      const res = await deleteExaCustomer({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: ""
        });
        if (this.tableData.length == 1 && this.page > 1 ) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createExaCustomer(this.form);
          break;
        case "update":
          res = await updateExaCustomer(this.form);
          break;
        default:
          res = await createExaCustomer(this.form);
          break;
      }

      if (res.code == 0) {
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.getTableData();
  }
};
</script>

<style>
</style>