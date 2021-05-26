<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addMenu('0')" type="primary">New root menu</el-button>
    </div>

    <!-- Since the menu is followed by a list of the left side, it is not necessary to paid PageSize default 999 -->
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
      <el-table-column label="Display name" min-width="120" prop="authorityName">
        <template slot-scope="scope">
          <span>{{scope.row.meta.title}}</span>
        </template>
      </el-table-column>
      <el-table-column label="icon" min-width="140" prop="authorityName">
        <template slot-scope="scope">
          <i :class="`el-icon-${scope.row.meta.icon}`"></i>
          <span>{{scope.row.meta.icon}}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="operating" width="300">
        <template slot-scope="scope">
          <el-button
            @click="addMenu(scope.row.ID)"
            size="small"
            type="primary"
            icon="el-icon-edit"
          >Add submenu</el-button>
          <el-button
            @click="editMenu(scope.row.ID)"
            size="small"
            type="primary"
            icon="el-icon-edit"
          >edit</el-button>
          <el-button
            @click="deleteMenu(scope.row.ID)"
            size="small"
            type="danger"
            icon="el-icon-delete"
          >delete</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :before-close="handleClose" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
        label-width="85px"
        ref="menuForm"
      >
        <el-form-item label="Route Name" prop="path" style="width:30%">
          <el-input
            @change="changeName"
            autocomplete="off"
            placeholder="Single English string"
            v-model="form.name"
          ></el-input>
        </el-form-item>
        <el-form-item prop="path" style="width:30%">
          <div style="display:inline-block" slot="label">
            Routing Path
            <el-checkbox style="float:right;margin-left:20px;" v-model="checkFlag">Add parameters</el-checkbox>
          </div>
          <el-input
            :disabled="!checkFlag"
            autocomplete="off"
            placeholder="It is recommended to splicing parameters in the rear"
            v-model="form.path"
          ></el-input>
        </el-form-item>
        <el-form-item label="Whether hidden" style="width:30%">
          <el-select placeholder="Is it hidden in the list?" v-model="form.hidden">
            <el-option :value="false" label="no"></el-option>
            <el-option :value="true" label="Yes"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Parent node ID" style="width:30%">
          <el-cascader
            :disabled="!this.isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
            v-model="form.parentId"
          ></el-cascader>
        </el-form-item>
        <el-form-item label="file path" prop="component" style="width:60%">
          <el-input autocomplete="off" v-model="form.component"></el-input>
          <span style="font-size:12px;margin-right:12px;">If the menu contains a submenu, create a Router-View secondary routing page or</span><el-button size="mini" @click="form.component = 'view/routerHolder.vue'">Click me to set up</el-button>
        </el-form-item>
        <el-form-item label="Display name" prop="meta.title" style="width:30%">
          <el-input autocomplete="off" v-model="form.meta.title"></el-input>
        </el-form-item>
        <el-form-item label="icon" prop="meta.icon" style="width:30%">
          <icon :meta="form.meta">
            <template slot="prepend">el-icon-</template>
          </icon>
        </el-form-item>
        <el-form-item label="Sort tag" prop="sort" style="width:30%">
          <el-input autocomplete="off" v-model.number="form.sort"></el-input>
        </el-form-item>
        <el-form-item label="keepAlive" prop="meta.keepAlive" style="width:30%">
          <el-select placeholder="Whether KeePalive Cache page" v-model="form.meta.keepAlive">
            <el-option :value="false" label="no"></el-option>
            <el-option :value="true" label="Yes"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="closeTab" prop="meta.closeTab" style="width:30%">
          <el-select placeholder="Whether automatically close TAB" v-model="form.meta.closeTab">
            <el-option :value="false" label="no"></el-option>
            <el-option :value="true" label="Yes"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div class="warning">The new menu needs to configure permissions within role management.</div>
      <div>
        <el-button
          size="small"
          type="primary"
          icon="el-icon-edit"
          @click="addParameter(form)"
        >Add menu parameters</el-button>
        <el-table :data="form.parameters" stripe style="width: 100%">
          <el-table-column prop="type" label="Parameter Type" width="180">
            <template slot-scope="scope">
              <el-select v-model="scope.row.type" placeholder="please choose">
                <el-option key="query" value="query" label="query"></el-option>
                <el-option key="params" value="params" label="params"></el-option>
              </el-select>
            </template>
          </el-table-column>
          <el-table-column prop="key" label="Parameters Key" width="180">
            <template slot-scope="scope">
              <div>
                <el-input v-model="scope.row.key"></el-input>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="value" label="Parameter value">
            <template slot-scope="scope">
              <div>
                <el-input v-model="scope.row.value"></el-input>
              </div>
            </template>
          </el-table-column>
          <el-table-column>
            <template slot-scope="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="el-icon-delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >delete</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">Take</el-button>
        <el-button @click="enterDialog" type="primary">Confirm</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
// Get list content package In Mixins internal GetTableData method initialized packaged

import {
  updateBaseMenu,
  getMenuList,
  addBaseMenu,
  deleteBaseMenu,
  getBaseMenuById
} from "@/api/menu";
import infoList from "@/mixins/infoList";
import icon from "@/view/superAdmin/menu/icon";
export default {
  name: "Menus",
  mixins: [infoList],
  data() {
    return {
      checkFlag: false,
      listApi: getMenuList,
      dialogFormVisible: false,
      dialogTitle: "Add menu",
      menuOption: [
        {
          ID: "0",
          title: "Root menu"
        }
      ],
      form: {
        ID: 0,
        path: "",
        name: "",
        hidden: "",
        parentId: "",
        component: "",
        meta: {
          title: "",
          icon: "",
          defaultMenu: false,
          closeTab: false,
          keepAlive: false
        },
        parameters: []
      },
      rules: {
        path: [{ required: true, message: "Please enter a menu name", trigger: "blur" }],
        component: [
          { required: true, message: "Please enter the file path", trigger: "blur" }
        ],
        "meta.title": [
          { required: true, message: "Please enter the menu display name", trigger: "blur" }
        ]
      },
      isEdit: false,
      test: ""
    };
  },
  components: {
    icon
  },
  methods: {
    addParameter(form) {
      if (!form.parameters) {
        this.$set(form, "parameters", []);
      }
      form.parameters.push({
        type: "query",
        key: "",
        value: ""
      });
    },
    deleteParameter(parameters, index) {
      parameters.splice(index, 1);
    },
    changeName() {
      this.form.path = this.form.name;
    },
    setOptions() {
      this.menuOption = [
        {
          ID: "0",
          title: "Root directory"
        }
      ];
      this.setMenuOptions(this.tableData, this.menuOption, false);
    },
    setMenuOptions(menuData, optionsData, disabled) {
      menuData &&
        menuData.map(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID == this.form.ID,
              children: []
            };
            this.setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID == this.form.ID
            );
            optionsData.push(option);
          } else {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID == this.form.ID
            };
            optionsData.push(option);
          }
        });
    },
    handleClose(done) {
      this.initForm();
      done();
    },
    // Lazy load menu
    load(tree, treeNode, resolve) {
      resolve([
        {
          id: 31,
          date: "2016-05-01",
          name: "Wang Xiaohu",
          address: "1519 Jinshajiang Road, Putuo District, Shanghai"
        },
        {
          id: 32,
          date: "2016-05-01",
          name: "Wang Xiaohu",
          address: "1519 Jinshajiang Road, Putuo District, Shanghai"
        }
      ]);
    },
    // Delete menu
    deleteMenu(ID) {
      this.$confirm("This will be permanently deleted under all roles, do you continue?", "prompt", {
        confirmButtonText: "determine",
        cancelButtonText: "cancel",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteBaseMenu({ ID });
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
    // Initialize the table method
    initForm() {
      this.checkFlag = false;
      this.$refs.menuForm.resetFields();
      this.form = {
        ID: 0,
        path: "",
        name: "",
        hidden: "",
        parentId: "",
        component: "",
        meta: {
          title: "",
          icon: "",
          defaultMenu: false,
          keepAlive: ""
        }
      };
    },
    // Close the pop-up window
    closeDialog() {
      this.initForm();
      this.dialogFormVisible = false;
    },
    // Add menu
    async enterDialog() {
      this.$refs.menuForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateBaseMenu(this.form);
          } else {
            res = await addBaseMenu(this.form);
          }
          if (res.code == 0) {
            this.$message({
              type: "success",
              message: this.isEdit ? "Editor success" : "Added successfully!"
            });
            this.getTableData();
          }
          this.initForm();
          this.dialogFormVisible = false;
        }
      });
    },
    // Add a menu method, ID is 0, add the root menu
    addMenu(id) {
      this.dialogTitle = "Add menu";
      this.form.parentId = String(id);
      this.isEdit = false;
      this.setOptions();
      this.dialogFormVisible = true;
    },
    // Modify the meter method
    async editMenu(id) {
      this.dialogTitle = "Edit menu";
      const res = await getBaseMenuById({ id });
      this.form = res.data.menu;
      this.isEdit = true;
      this.setOptions();
      this.dialogFormVisible = true;
    }
  },
  async created() {
    this.pageSize = 999;
    await this.getTableData();
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
.warning {
  color: #dc143c;
}
</style>
