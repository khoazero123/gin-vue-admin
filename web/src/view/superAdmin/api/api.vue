<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="path">
          <el-input placeholder="path" v-model="searchInfo.path"></el-input>
        </el-form-item>
        <el-form-item label="description">
          <el-input placeholder="description" v-model="searchInfo.description"></el-input>
        </el-form-item>
        <el-form-item label="API">
          <el-input placeholder="API" v-model="searchInfo.apiGroup"></el-input>
        </el-form-item>
        <el-form-item label="request">
          <el-select clearable placeholder="please choose" v-model="searchInfo.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">Inquire</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('addApi')" type="primary">New API</el-button>
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
    <el-table :data="tableData" @sort-change="sortChange" border stripe @selection-change="handleSelectionChange">
       <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column label="id" min-width="60" prop="ID" sortable="custom"></el-table-column>
      <el-table-column label="API path" min-width="150" prop="path" sortable="custom"></el-table-column>
      <el-table-column label="API group" min-width="150" prop="apiGroup" sortable="custom"></el-table-column>
      <el-table-column label="Introduction to API" min-width="150" prop="description" sortable="custom"></el-table-column>
      <el-table-column label="request" min-width="150" prop="method" sortable="custom">
        <template slot-scope="scope">
          <div>
            {{scope.row.method}}
            <el-tag
              :key="scope.row.methodFiletr"
              :type="scope.row.method|tagTypeFiletr"
              effect="dark"
              size="mini"
            >{{scope.row.method|methodFiletr}}</el-tag>
            <!-- {{scope.row.method|methodFiletr}} -->
          </div>
        </template>
      </el-table-column>

      <el-table-column fixed="right" label="operating" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="primary" icon="el-icon-edit">edit</el-button>
          <el-button
            @click="deleteApi(scope.row)"
            size="small"
            type="danger"
            icon="el-icon-delete"
          >delete</el-button>
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

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="path" prop="path">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="request" prop="method">
          <el-select placeholder="please choose" v-model="form.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="API group" prop="apiGroup">
          <el-input autocomplete="off" v-model="form.apiGroup"></el-input>
        </el-form-item>
        <el-form-item label="Introduction to API" prop="description">
          <el-input autocomplete="off" v-model="form.description"></el-input>
        </el-form-item>
      </el-form>
      <div class="warning">New API needs to configure permissions within role management.</div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">Take</el-button>
        <el-button @click="enterDialog" type="primary">Confirm</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// Get List Content Package In Mixins Internal GetTableData Method Initialization Package Completed Condition Search, place the conditional sector in this.SearchInfo when you are searched.

import {
  getApiById,
  getApiList,
  createApi,
  updateApi,
  deleteApi,
  deleteApisByIds
} from "@/api/api";
import infoList from "@/mixins/infoList";
import { toSQLLine } from "@/utils/stringFun";
const methodOptions = [
  {
    value: "POST",
    label: "create",
    type: "success"
  },
  {
    value: "GET",
    label: "Check",
    type: ""
  },
  {
    value: "PUT",
    label: "Update",
    type: "warning"
  },
  {
    value: "DELETE",
    label: "delete",
    type: "danger"
  }
];

export default {
  name: "Api",
  mixins: [infoList],
  data() {
    return {
      deleteVisible:false,
      listApi: getApiList,
      dialogFormVisible: false,
      dialogTitle: "New API",
      apis:[],
      form: {
        path: "",
        apiGroup: "",
        method: "",
        description: ""
      },
      methodOptions: methodOptions,
      type: "",
      rules: {
        path: [{ required: true, message: "Please enter an API path", trigger: "blur" }],
        apiGroup: [
          { required: true, message: "Please enter the group name", trigger: "blur" }
        ],
        method: [
          { required: true, message: "Please select a request method", trigger: "blur" }
        ],
        description: [
          { required: true, message: "Please enter the API introduction", trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    //  Select API
      handleSelectionChange(val) {
        this.apis = val;
      },
      async onDelete(){
        const ids = this.apis.map(item=>item.ID)
        const res = await deleteApisByIds({ids})
        if(res.code==0){
          this.$message({
            type:"success",
            message:res.msg
          })
         if (this.tableData.length == ids.length && this.page > 1) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    // Sort
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop);
        this.searchInfo.desc = order == "descending";
      }
      this.getTableData();
    },
    //Condition Search Front End View this method
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    initForm() {
      this.$refs.apiForm.resetFields();
      this.form = {
        path: "",
        apiGroup: "",
        method: "",
        description: ""
      };
    },
    closeDialog() {
      this.initForm();
      this.dialogFormVisible = false;
    },
    openDialog(type) {
      switch (type) {
        case "addApi":
          this.dialogTitle = "New API";
          break;
        case "edit":
          this.dialogTitle = "Edit API";
          break;
        default:
          break;
      }
      this.type = type;
      this.dialogFormVisible = true;
    },
    async editApi(row) {
      const res = await getApiById({ id: row.ID });
      this.form = res.data.api;
      this.openDialog("edit");
    },
    async deleteApi(row) {
      this.$confirm("This action will permanently delete the API under all roles, is it?", "Tips", {
        confirmButtonText: "determine",
        cancelButtonText: "cancel",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteApi(row);
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
    async enterDialog() {
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case "addApi":
              {
                const res = await createApi(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "Added successfully",
                    showClose: true
                  });
                }
                this.getTableData();
                this.closeDialog();
              }

              break;
            case "edit":
              {
                const res = await updateApi(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "Editor success",
                    showClose: true
                  });
                }
                this.getTableData();
                this.closeDialog();
              }
              break;
            default:
              {
                this.$message({
                  type: "error",
                  message: "Unknown operation",
                  showClose: true
                });
              }
              break;
          }
        }
      });
    }
  },
  filters: {
    methodFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0];
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`;
    },
    tagTypeFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0];
      return target && `${target.type}`;
    }
  },
  created() {
    this.getTableData();
  }
};
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>