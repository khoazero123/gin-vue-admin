<template>
  <div>
    <!-- Get fields directly from the database -->
    <el-collapse v-model="activeNames">
      <el-collapse-item name="1">
        <template slot="title">
          <div :style="{fontSize:'16px',paddingLeft:'20px'}">
            Click here to create code from existing database
            <i class="header-icon el-icon-thumb"></i>
          </div>
        </template>
        <el-form ref="getTableForm" :inline="true" :model="dbform" label-width="120px">
          <el-form-item label="data storage name" prop="structName">
            <el-select @change="getTable" v-model="dbform.dbName" filterable placeholder="Please select a database">
              <el-option
                v-for="item in dbOptions"
                :key="item.database"
                :label="item.database"
                :value="item.database"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Table Name" prop="structName">
            <el-select
              v-model="dbform.tableName"
              :disabled="!dbform.dbName"
              filterable
              placeholder="Please select a table"
            >
              <el-option
                v-for="item in tableOptions"
                :key="item.tableName"
                :label="item.tableName"
                :value="item.tableName"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button @click="getColumn" type="primary">Create this table</el-button>
          </el-form-item>
        </el-form>
      </el-collapse-item>
    </el-collapse>

    <el-divider></el-divider>
    <!-- Initial version Automation Code Tool -->
    <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="120px" :inline="true">
      <el-form-item label="ClassName" prop="structName">
        <el-input v-model="form.structName" placeholder="First letter automatic conversion capital"></el-input>
      </el-form-item>
      <el-form-item label="Table name" prop="tableName">
        <el-input v-model="form.tableName" placeholder="Specify a table name (non-required)"></el-input>
      </el-form-item>
      <el-form-item label="Route name" prop="abbreviation">
        <el-input v-model="form.abbreviation" placeholder="Abbreviation will be used as an entrance to object name and routing group"></el-input>
      </el-form-item>
      <el-form-item label="Description" prop="description">
        <el-input v-model="form.description" placeholder="Chinese description as an automatic API description"></el-input>
      </el-form-item>
      <el-form-item label="Package name" prop="packageName">
        <el-input v-model="form.packageName" placeholder="The default name for generating files"></el-input>
      </el-form-item>
      <el-form-item label="Automatically create an API">
        <el-checkbox v-model="form.autoCreateApiToSql"></el-checkbox>
      </el-form-item>
      <el-form-item label="Automatic move file">
        <el-checkbox v-model="form.autoMoveFile"></el-checkbox>
      </el-form-item>
    </el-form>
    <!-- Component list -->
    <div class="button-box clearflex">
      <el-button @click="editAndAddField()" type="primary">New Field</el-button>
    </div>
    <el-table :data="form.fields" border stripe>
      <el-table-column type="index" label="sequence" width="100"></el-table-column>
      <el-table-column prop="fieldName" label="FIELD name"></el-table-column>
      <el-table-column prop="fieldDesc" label="Field description"></el-table-column>
      <el-table-column prop="fieldJson" label="FieldJson"></el-table-column>
      <el-table-column prop="fieldType" label="Field data type" width="130"></el-table-column>
      <el-table-column prop="dataType" label="Database field type" width="130"></el-table-column>
      <el-table-column prop="dataTypeLong" label="Database field length" width="130"></el-table-column>
      <el-table-column prop="columnName" label="Database field" width="130"></el-table-column>
      <el-table-column prop="comment" label="Database field description" width="130"></el-table-column>
      <el-table-column prop="fieldSearchType" label="search condition" width="130"></el-table-column>
      <el-table-column prop="dictType" label="dictionary" width="130"></el-table-column>
      <el-table-column label="operating" width="300">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-edit"
            @click="editAndAddField(scope.row)"
          >edit</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="scope.$index == 0"
            @click="moveUpField(scope.$index)"
          >Move up</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="(scope.$index + 1) == form.fields.length"
            @click="moveDownField(scope.$index)"
          >Move down</el-button>
          <el-popover placement="top" v-model="scope.row.visible">
            <p>confirm to delete?</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">cancel</el-button>
              <el-button type="primary" size="mini" @click="deleteField(scope.$index)">determine</el-button>
            </div>
            <el-button size="mini" type="danger" icon="el-icon-delete" slot="reference">delete</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-tag type="danger">id , created_at , updated_at , deleted_at Will automatically generate, do not create repeatedly</el-tag>
    <!-- Component list -->
    <div class="button-box clearflex">
      <el-button @click="enterForm(true)" type="primary">Preview code</el-button>
      <el-button @click="enterForm(false)" type="primary">Code</el-button>
    </div>
    <!-- Component pop-up window -->
    <el-dialog title="Component content" :visible.sync="dialogFlag">
      <FieldDialog v-if="dialogFlag" :dialogMiddle="dialogMiddle" ref="fieldDialog" />
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">Take</el-button>
        <el-button type="primary" @click="enterDialog">Confirm</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="previewFlag">
      <PreviewCodeDialg v-if="previewFlag" :previewCode="preViewCode"></PreviewCodeDialg>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="previewFlag = false">Confirm</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
const fieldTemplate = {
  fieldName: "",
  fieldDesc: "",
  fieldType: "",
  dataType: "",
  fieldJson: "",
  columnName: "",
  dataTypeLong: "",
  comment: "",
  fieldSearchType: "",
  dictType: ""
};

import FieldDialog from "@/view/systemTools/autoCode/component/fieldDialog.vue";
import PreviewCodeDialg from "@/view/systemTools/autoCode/component/previewCodeDialg.vue";
import { toUpperCase, toHump } from "@/utils/stringFun.js";
import { createTemp, getDB, getTable, getColumn, preview } from "@/api/autoCode.js";
import { getDict } from "@/utils/dictionary";

export default {
  name: "autoCode",
  data() {
    return {
      activeNames: [""],
      preViewCode:{},
      dbform: {
        dbName: "",
        tableName: ""
      },
      dbOptions: [],
      tableOptions: [],
      addFlag: "",
      fdMap: {},
      form: {
        structName: "",
        tableName: "",
        packageName: "",
        abbreviation: "",
        description: "",
        autoCreateApiToSql: false,
        autoMoveFile: false,
        fields: []
      },
      rules: {
        structName: [
          { required: true, message: "Please enter the structure name", trigger: "blur" }
        ],
        abbreviation: [
          { required: true, message: "Please enter the structure [a-zA-Z]", trigger: "blur", pattern: /[a-zA-Z]/ }
        ],
        description: [
          { required: true, message: "Please enter a structure description", trigger: "blur" }
        ],
        packageName: [
          {
            required: true,
            message: "file name：sys_xxxx_xxxx",
            trigger: "blur"
          }
        ]
      },
      dialogMiddle: {},
      bk: {},
      dialogFlag: false,
      previewFlag:false
    };
  },
  components: {
    FieldDialog,
    PreviewCodeDialg
  },
  methods: {
    editAndAddField(item) {
      this.dialogFlag = true;
      if (item) {
        this.addFlag = "edit";
        this.bk = JSON.parse(JSON.stringify(item));
        this.dialogMiddle = item;
      } else {
        this.addFlag = "add";
        this.dialogMiddle = JSON.parse(JSON.stringify(fieldTemplate));
      }
    },
    moveUpField(index) {
      if (index == 0) {
        return;
      }
      const oldUpField = this.form.fields[index - 1];
      this.form.fields.splice(index - 1, 1);
      this.form.fields.splice(index, 0, oldUpField);
    },
    moveDownField(index) {
      const fCount = this.form.fields.length;
      if (index == fCount - 1) {
        return;
      }
      const oldDownField = this.form.fields[index + 1];
      this.form.fields.splice(index + 1, 1);
      this.form.fields.splice(index, 0, oldDownField);
    },
    enterDialog() {
      this.$refs.fieldDialog.$refs.fieldDialogFrom.validate(valid => {
        if (valid) {
          this.dialogMiddle.fieldName = toUpperCase(
            this.dialogMiddle.fieldName
          );
          if (this.addFlag == "add") {
            this.form.fields.push(this.dialogMiddle);
          }
          this.dialogFlag = false;
        } else {
          return false;
        }
      });
    },
    closeDialog() {
      if (this.addFlag == "edit") {
        this.dialogMiddle = this.bk;
      }
      this.dialogFlag = false;
    },
    deleteField(index) {
      this.form.fields.splice(index, 1);
    },
    async enterForm(isPreview) {
      if (this.form.fields.length <= 0) {
        this.$message({
          type: "error",
          message: "Please fill in at least one field"
        });
        return false;
      }
      if (
        this.form.fields.some(item => item.fieldName == this.form.structName)
      ) {
        this.$message({
          type: "error",
          message: "There is a field with the same name with the structure"
        });
        return false;
      }
      this.$refs.autoCodeForm.validate(async valid => {
        if (valid) {
          this.form.structName = toUpperCase(this.form.structName).replace(/\s/g, '');
          if (this.form.structName == this.form.abbreviation) {
            this.$message({
              type: "error",
              message: "StructName and Struct are abbrevite"
            });
            return false;
          }
          if(isPreview){
            const data = await preview(this.form);
            console.log(data.code == 0)
            this.preViewCode = data.data.autoCode
            this.previewFlag = true
          }else{
            const data = await createTemp(this.form);
            if (data.headers?.success == "false") {
              return;
            } else {
              this.$message({
                type: "success",
                message: "Automation code creates success, downloading"
              });
            }
            const blob = new Blob([data]);
            const fileName = "ginvueadmin.zip";
            if ("download" in document.createElement("a")) {
              // Not IE browser
              let url = window.URL.createObjectURL(blob);
              let link = document.createElement("a");
              link.style.display = "none";
              link.href = url;
              link.setAttribute("download", fileName);
              document.body.appendChild(link);
              link.click();
              document.body.removeChild(link); // Download completion removal element
              window.URL.revokeObjectURL(url); // Release Blob object
            } else {
              // IE 10+
              window.navigator.msSaveBlob(blob, fileName);
            }
          }
        } else {
          return false;
        }
      });
    },
    async getDb() {
      const res = await getDB();
      if (res.code == 0) {
        this.dbOptions = res.data.dbs;
      }
    },
    async getTable() {
      const res = await getTable({ dbName: this.dbform.dbName });
      if (res.code == 0) {
        this.tableOptions = res.data.tables;
      }
      this.dbform.tableName = "";
    },
    async getColumn() {
      const gormModelList = ["id", "created_at", "updated_at", "deleted_at"];
      const res = await getColumn(this.dbform);
      if (res.code == 0) {
        const tbHump = toHump(this.dbform.tableName);
        this.form.structName = toUpperCase(tbHump);
        this.form.tableName = this.dbform.tableName;
        this.form.packageName = tbHump;
        this.form.abbreviation = tbHump;
        this.form.description = tbHump + "表";
        this.form.autoCreateApiToSql = true;
        this.form.fields = [];
        res.data.columns &&
          res.data.columns.map(item => {
            if (!gormModelList.some(gormfd => gormfd == item.columnName)) {
              const fbHump = toHump(item.columnName);
              this.form.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + "Field",
                fieldType: this.fdMap[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong,
                columnName: item.columnName,
                comment: item.columnComment,
                fieldSearchType: "",
                dictType: ""
              });
            }
          });
      }
    },
    async setFdMap() {
      const fdTypes = ["string", "int", "bool", "float64", "time.Time"];
      fdTypes.map(async fdtype => {
        const res = await getDict(fdtype);
        res&&res.map(item => {
          this.fdMap[item.label] = fdtype;
        });
      });
    }
  },
  created() {
    this.getDb();
    this.setFdMap();
  }
};
</script>
<style scope lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    margin-right: 20px;
    float: right;
  }
}
</style>
