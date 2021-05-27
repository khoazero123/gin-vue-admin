<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
           {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
            <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clear placeholder="please choose">
                <el-option
                    key="true"
                    label="Yes"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="no"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
                  {{- else }}
        <el-form-item label="{{.FieldDesc}}">
          <el-input placeholder="search condition" v-model="searchInfo.{{.FieldJson}}"></el-input>
        </el-form-item> {{ end }} {{ end }}  {{ end }}
        <el-form-item>
          <el-button @click="onSubmit" type="primary">Inquire</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">Add new {{.Description}}</el-button>
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
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{ "{{scope.row.CreatedAt|formatDate}}" }}</template>
    </el-table-column>
    {{range .Fields}}
    {{- if .DictType}}
      <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
        <template slot-scope="scope">
          {{"{{"}}filterDict(scope.row.{{.FieldJson}},"{{.DictType}}"){{"}}"}}
        </template>
      </el-table-column>
    {{- else if eq .FieldType "bool" }}
    <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
         <template slot-scope="scope">{{ "{{scope.row."}}{{.FieldJson}}{{"|formatBoolean}}" }}</template>
    </el-table-column> {{- else }}
    <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"></el-table-column> {{ end }}
    {{ end }}
      <el-table-column label="Button group">
        <template slot-scope="scope">
          <el-button class="table-button" @click="update{{.StructName}}(scope.row)" size="small" type="primary" icon="el-icon-edit">change</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">delete</el-button>
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
      <el-form :model="formData" label-position="right" label-width="80px">
    {{- range .Fields}}
         <el-form-item label="{{.FieldDesc}}:">
      {{- if eq .FieldType "bool" }}
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="Yes" inactive-text="no" v-model="formData.{{.FieldJson}}" clearable ></el-switch>
      {{ end -}}
      {{- if eq .FieldType "string" }}
            <el-input v-model="formData.{{.FieldJson}}" clearable placeholder="please enter" ></el-input>
      {{ end -}}
      {{- if eq .FieldType "int" }}
      {{- if .DictType}}
             <el-select v-model="formData.{{ .FieldJson }}" placeholder="please choose" clearable>
                 <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value"></el-option>
             </el-select>
      {{ else -}}
             <el-input v-model.number="formData.{{ .FieldJson }}" clearable placeholder="please enter"></el-input>
      {{ end -}}
      {{ end -}}
      {{- if eq .FieldType "time.Time" }}
              <el-date-picker type="date" placeholder="Dates" v-model="formData.{{ .FieldJson }}" clearable></el-date-picker>
       {{ end -}}
       {{- if eq .FieldType "float64" }}
              <el-input-number v-model="formData.{{ .FieldJson }}" :precision="2" clearable></el-input-number>
       {{ end -}}
          </el-form-item>
       {{ end -}}
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
    create{{.StructName}},
    delete{{.StructName}},
    delete{{.StructName}}ByIds,
    update{{.StructName}},
    find{{.StructName}},
    get{{.StructName}}List
} from "@/api/{{.PackageName}}";  //  Please replace the address here
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "{{.StructName}}",
  mixins: [infoList],
  data() {
    return {
      listApi: get{{ .StructName }}List,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],

      {{- range .Fields}}
          {{- if .DictType }}
      {{ .DictType }}Options:[],
          {{ end -}}
      {{end -}}

      formData: {
            {{range .Fields}}
            {{- if eq .FieldType "bool" -}}
               {{.FieldJson}}:false,
            {{ end -}}
            {{- if eq .FieldType "string" -}}
               {{.FieldJson}}:"",
            {{ end -}}
            {{- if eq .FieldType "int" -}}
               {{.FieldJson}}:0,
            {{ end -}}
            {{- if eq .FieldType "time.Time" -}}
               {{.FieldJson}}:new Date(),
            {{ end -}}
            {{- if eq .FieldType "float64" -}}
               {{.FieldJson}}:0,
            {{ end -}}
            {{ end }}
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
        return bool ? "Yes" :"no";
      } else {
        return "";
      }
    }
  },
  methods: {
      //Condition Search Front End View this method
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        {{- range .Fields}} {{- if eq .FieldType "bool" }}      
        if (this.searchInfo.{{.FieldJson}}==""){
          this.searchInfo.{{.FieldJson}}=null
        } {{ end }} {{ end }}    
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('You sure you want to delete it?', 'prompt', {
          confirmButtonText: 'determine',
          cancelButtonText: 'cancel',
          type: 'warning'
        }).then(() => {
           this.delete{{.StructName}}(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: 'Please select the data you want to delete'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await delete{{.StructName}}ByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: 'Successfully deleted'
          })
          if (this.tableData.length == ids.length && this.page > 1) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async update{{.StructName}}(row) {
      const res = await find{{.StructName}}({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.re{{.Abbreviation}};
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          {{range .Fields}}
          {{- if eq .FieldType "bool" -}}
             {{.FieldJson}}:false,
          {{ end -}}
          {{- if eq .FieldType "string" -}}
             {{.FieldJson}}:"",
          {{ end -}}
          {{- if eq .FieldType "int" -}}
             {{.FieldJson}}:0,
          {{ end -}}
          {{- if eq .FieldType "time.Time" -}}
             {{.FieldJson}}:new Date(),
          {{ end -}}
          {{- if eq .FieldType "float64" -}}
             {{.FieldJson}}:0,
          {{ end -}}
          {{ end }}
      };
    },
    async delete{{.StructName}}(row) {
      const res = await delete{{.StructName}}({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "Successfully deleted"
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
          res = await create{{.StructName}}(this.formData);
          break;
        case "update":
          res = await update{{.StructName}}(this.formData);
          break;
        default:
          res = await create{{.StructName}}(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"Create / change success"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  {{ range .Fields -}}
    {{- if .DictType }}
    await this.getDict("{{.DictType}}");
    {{ end -}}
  {{- end }}
}
};
</script>

<style>
</style>
