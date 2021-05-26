<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addUser" type="primary">New users</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="Avatar" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <CustomPic :picSrc="scope.row.headerImg" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="uuid" min-width="250" prop="uuid"></el-table-column>
      <el-table-column label="username" min-width="150" prop="userName"></el-table-column>
      <el-table-column label="nickname" min-width="150" prop="nickName"></el-table-column>
      <el-table-column label="User role" min-width="150">
        <template slot-scope="scope">
          <el-cascader
            @change="changeAuthority(scope.row)"
            v-model="scope.row.authority.authorityId"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            filterable
          ></el-cascader>
        </template>
      </el-table-column>
      <el-table-column label="operating" min-width="150">
        <template slot-scope="scope">
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>Be sure to delete this user?</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">cancel</el-button>
              <el-button type="primary" size="mini" @click="deleteUser(scope.row)">determine</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="small" slot="reference">delete</el-button>
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

    <el-dialog :visible.sync="addUserDialog" custom-class="user-dialog" title="New users">
      <el-form :rules="rules" ref="userForm" :model="userInfo">
        <el-form-item label="username" label-width="80px" prop="username">
          <el-input v-model="userInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="password" label-width="80px" prop="password">
          <el-input v-model="userInfo.password"></el-input>
        </el-form-item>
        <el-form-item label="Alias" label-width="80px" prop="nickName">
          <el-input v-model="userInfo.nickName"></el-input>
        </el-form-item>
        <el-form-item label="Avatar" label-width="80px">
          <div style="display:inline-block" @click="openHeaderChange">
            <img class="header-img-box" v-if="userInfo.headerImg" :src="userInfo.headerImg" />
            <div v-else class="header-img-box">Select from the media library</div>
          </div>
        </el-form-item>
        <el-form-item label="User role" label-width="80px" prop="authorityId">
          <el-cascader
            v-model="userInfo.authorityId"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            filterable
          ></el-cascader>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddUserDialog">Take</el-button>
        <el-button @click="enterAddUserDialog" type="primary">Confirm</el-button>
      </div>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :targetKey="`headerImg`"/>
  </div>
</template>


<script>
// Get list content package In Mixins internal GetTableData method initialized packaged
const path = process.env.VUE_APP_BASE_API;
import {
  getUserList,
  setUserAuthority,
  register,
  deleteUser
} from "@/api/user";
import { getAuthorityList } from "@/api/authority";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
import CustomPic from "@/components/customPic";
import ChooseImg from "@/components/chooseImg";
export default {
  name: "Api",
  mixins: [infoList],
  components: { CustomPic,ChooseImg },
  data() {
    return {
      listApi: getUserList,
      path: path,
      authOptions: [],
      addUserDialog: false,
      userInfo: {
        username: "",
        password: "",
        nickName: "",
        headerImg: "",
        authorityId: ""
      },
      rules: {
        username: [
          { required: true, message: "please enter user name", trigger: "blur" },
          { min: 5, message: "Minimum 5 characters", trigger: "blur" }
        ],
        password: [
          { required: true, message: "Please enter the user password", trigger: "blur" },
          { min: 6, message: "最低6位字符", trigger: "blur" }
        ],
        nickName: [
          { required: true, message: "Please enter the user nickname", trigger: "blur" }
        ],
        authorityId: [
          { required: true, message: "Please select a user role", trigger: "blur" }
        ]
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"])
  },
  methods: {
    openHeaderChange(){
      this.$refs.chooseImg.open()
    },
    setOptions(authData) {
      this.authOptions = [];
      this.setAuthorityOptions(authData, this.authOptions);
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
        AuthorityData.map(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            };
            this.setAuthorityOptions(item.children, option.children);
            optionsData.push(option);
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            };
            optionsData.push(option);
          }
        });
    },
    async deleteUser(row) {
      const res = await deleteUser({ id: row.ID });
      if (res.code == 0) {
        this.getTableData();
        row.visible = false;
      }
    },
    async enterAddUserDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo);
          if (res.code == 0) {
            this.$message({ type: "success", message: "创建成功" });
          }
          await this.getTableData();
          this.closeAddUserDialog();
        }
      });
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields();
      this.addUserDialog = false;
    },
    handleAvatarSuccess(res) {
      this.userInfo.headerImg = res.data.file.url;
    },
    addUser() {
      this.addUserDialog = true;
    },
    async changeAuthority(row) {
      const res = await setUserAuthority({
        uuid: row.uuid,
        authorityId: row.authority.authorityId
      });
      if (res.code == 0) {
        this.$message({ type: "success", message: "角色设置成功" });
      }
    }
  },
  async created() {
    this.getTableData();
    const res = await getAuthorityList({ page: 1, pageSize: 999 });
    this.setOptions(res.data.list);
  }
};
</script>
<style lang="scss">

.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>