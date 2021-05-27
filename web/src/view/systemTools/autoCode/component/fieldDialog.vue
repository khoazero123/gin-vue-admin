<template>
  <div>
    <span style="color:red">If the condition is LIKE, only the string is supported.</span>
    <el-form
      :model="dialogMiddle"
      ref="fieldDialogFrom"
      label-width="120px"
      label-position="left"
      :rules="rules"
    >
      <el-form-item label="Field name" prop="fieldName">
        <el-col :span="6">
          <el-input v-model="dialogMiddle.fieldName" autocomplete="off"></el-input>
        </el-col>
        <el-col :offset="1" :span="2">
          <el-button @click="autoFill">Automatic filling</el-button>
        </el-col>
      </el-form-item>
      <el-form-item label="Field description" prop="fieldDesc">
        <el-col :span="6">
          <el-input v-model="dialogMiddle.fieldDesc" autocomplete="off"></el-input>
        </el-col>
      </el-form-item>
      <el-form-item label="FieldJSON" prop="fieldJson">
        <el-col :span="6">
          <el-input v-model="dialogMiddle.fieldJson" autocomplete="off"></el-input>
        </el-col>
      </el-form-item>
      <el-form-item label="Database field name" prop="columnName">
        <el-col :span="6">
          <el-input v-model="dialogMiddle.columnName" autocomplete="off"></el-input>
        </el-col>
      </el-form-item>
      <el-form-item label="Database field description" prop="comment">
        <el-col :span="6">
          <el-input v-model="dialogMiddle.comment" autocomplete="off"></el-input>
        </el-col>
      </el-form-item>
      <el-form-item label="Field data type" prop="fieldType">
        <el-col :span="8">
          <el-select
            v-model="dialogMiddle.fieldType"
            placeholder="Please select Field data type"
            @change="getDbfdOptions"
            clearable
          >
            <el-option
              v-for="item in typeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-col>
      </el-form-item>

      <el-form-item label="Database field type" prop="dataType">
        <el-col :span="8">
          <el-select
            :disabled="!dialogMiddle.fieldType"
            v-model="dialogMiddle.dataType"
            placeholder="Please select the database field type"
            clearable
          >
            <el-option
              v-for="item in dbfdOptions"
              :key="item.label"
              :label="item.label"
              :value="item.label"
            ></el-option>
          </el-select>
        </el-col>
      </el-form-item>
      <el-form-item label="Database field length" prop="dataTypeLong">
        <el-col :span="8">
          <el-input placeholder="Customized type must be specified" :disabled="!dialogMiddle.dataType" v-model="dialogMiddle.dataTypeLong"></el-input>
        </el-col>
      </el-form-item>
      <el-form-item label="Field query conditions" prop="fieldSearchType">
        <el-col :span="8">
          <el-select v-model="dialogMiddle.fieldSearchType" placeholder="Please select Field Query Conditions" clearable>
            <el-option
              v-for="item in typeSearchOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-col>
      </el-form-item>

      <el-form-item label="Associate dictionary" prop="dictType">
        <el-col :span="8">
          <el-select :disabled="dialogMiddle.fieldType!='int'" v-model="dialogMiddle.dictType" placeholder="Please select a dictionary" clearable>
            <el-option
              v-for="item in dictOptions"
              :key="item.type"
              :label="`${item.type}(${item.name})`"
              :value="item.type"
            ></el-option>
          </el-select>
        </el-col>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { getDict } from "@/utils/dictionary";
import { toSQLLine , toLowerCase } from "@/utils/stringFun.js";
import { getSysDictionaryList } from "@/api/sysDictionary";
export default {
  name: "FieldDialog",
  props: {
    dialogMiddle: {
      type: Object,
      default: function() {
        return {};
      }
    }
  },
  data() {
    return {
      dbfdOptions: [],
      dictOptions: [],
      typeSearchOptions: [
        {
          label: "=",
          value: "="
        },
        {
          label: "<>",
          value: "<>"
        },
        {
          label: ">",
          value: ">"
        },
        {
          label: "<",
          value: "<"
        },
        {
          label: "LIKE",
          value: "LIKE"
        }
      ],
      typeOptions: [
        {
          label: "String",
          value: "string"
        },
        {
          label: "Integrity",
          value: "int"
        },
        {
          label: "Boolean value",
          value: "bool"
        },
        {
          label: "Floating point",
          value: "float64"
        },
        {
          label: "time",
          value: "time.Time"
        }
      ],
      rules: {
        fieldName: [
          { required: true, message: "Please enter the Field name", trigger: "blur" }
        ],
        fieldDesc: [
          { required: true, message: "Please enter the Field Description", trigger: "blur" }
        ],
        fieldJson: [
          { required: true, message: "Please enter Field format JSON", trigger: "blur" }
        ],
        columnName: [
          { required: true, message: "Please enter the database field", trigger: "blur" }
        ],
        fieldType: [
          { required: true, message: "Please select Field data type", trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    autoFill(){
        this.dialogMiddle.fieldJson = toLowerCase(this.dialogMiddle.fieldName)
        this.dialogMiddle.columnName = toSQLLine(this.dialogMiddle.fieldJson)
    },
    async getDbfdOptions() {
        this.dialogMiddle.dataType = ""
        this.dialogMiddle.dataTypeLong = ""
        this.dialogMiddle.fieldSearchType = ""
        this.dialogMiddle.dictType = ""
      if (this.dialogMiddle.fieldType) {
        const res = await getDict(this.dialogMiddle.fieldType);
        this.dbfdOptions = res;
      }
    }
  },
  async created() {
    const dictRes = await getSysDictionaryList({
      page: 1,
      pageSize: 999999
    });

    this.dictOptions = dictRes.data.list
  },
};
</script>
<style lang="scss">
</style>
