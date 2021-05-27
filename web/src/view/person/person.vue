<template>
  <div>
    <el-row>
      <el-col :span="6">
        <div class="fl-left avatar-box">
          <div class="user-card">
              <div class="user-headpic-update" :style="{ 'background-image': `url(${(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg})`,'background-repeat':'no-repeat','background-size':'cover' }" >
              <span class="update" @click="openChooseImg">
                <i class="el-icon-edit"></i>
                re-upload</span>
              </div>
            <div class="user-personality">
              <p class="nickname">{{userInfo.nickName}}</p>
              <p class="person-info">This guy is too lazy, nothing left</p>
            </div>
            <div class="user-information">
              <ul>
                <li>
                   <i class="el-icon-user"></i>{{userInfo.nickName}}
                </li>
                <li>
                  <i class="el-icon-data-analysis"></i>Beijing Inverse Aurora Technology Co., Ltd. - Technical Department - Front End Business Group
                </li>
                <li>
                  <i class="el-icon-video-camera-solid"></i>China · Beijing Chaoyang District
                </li>
                <li>
                  <i class="el-icon-medal-1"></i>goLang/JavaScript/Vue/Gorm
                </li>
              </ul>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="18">
        <div class="user-addcount">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="账号绑定" name="second">
              <ul>
                <li>
                  <p class="title">Secret mobile phone</p>
                  <p class="desc">
                    Bind mobile phone:1245678910
                    <a href="#">Immediate modification</a>
                  </p>
                </li>
                <li>
                  <p class="title">Secret mailbox</p>
                  <p class="desc">
                    Binding mailbox：gin-vue-admin@google.com.cn
                    <a href="#">Immediate modification</a>
                  </p>
                </li>
                <li>
                  <p class="title">Security Question</p>
                  <p class="desc">
                    No security issues are set
                    <a href="#">Set</a>
                  </p>
                </li>
                <li>
                  <p class="title">change Password</p>
                  <p class="desc">
                    Modify personal password
                    <a href="#" @click="showPassword=true">change Password</a>
                  </p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-col>
    </el-row>

    <ChooseImg ref="chooseImg" @enter-img="enterImg" />

    <el-dialog :visible.sync="showPassword" @close="clearPassword" title="change Password" width="360px">
      <el-form :model="pwdModify" :rules="rules" label-width="80px" ref="modifyPwdForm">
        <el-form-item :minlength="6" label="old password" prop="password">
          <el-input show-password v-model="pwdModify.password"></el-input>
        </el-form-item>
        <el-form-item :minlength="6" label="new password" prop="newPassword">
          <el-input show-password v-model="pwdModify.newPassword"></el-input>
        </el-form-item>
        <el-form-item :minlength="6" label="confirm password" prop="confirmPassword">
          <el-input show-password v-model="pwdModify.confirmPassword"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="showPassword=false">Take</el-button>
        <el-button @click="savePassword" type="primary">Confirm</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import ChooseImg from "@/components/chooseImg";
import { setUserInfo,changePassword } from "@/api/user";

import { mapGetters, mapMutations } from "vuex";
const path = process.env.VUE_APP_BASE_API;
export default {
  name: "Person",
  data() {
    return {
      path: path,
      activeName: "second",
      showPassword: false,
      pwdModify: {},
      rules: {
        password: [
          { required: true, message: "Please enter your password", trigger: "blur" },
          { min: 6, message: "At least 6 characters", trigger: "blur" }
        ],
        newPassword: [
          { required: true, message: "Please enter a new password", trigger: "blur" },
          { min: 6, message: "At least 6 characters", trigger: "blur" }
        ],
        confirmPassword: [
          { required: true, message: "Please enter a confirmation password", trigger: "blur" },
          { min: 6, message: "At least 6 characters", trigger: "blur" },
          {
            validator: (rule, value, callback) => {
              if (value !== this.pwdModify.newPassword) {
                callback(new Error("Two passwords are inconsistent"));
              } else {
                callback();
              }
            },
            trigger: "blur"
          }
        ]
      }
    };
  },
  components: {
    ChooseImg
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"])
  },
  methods: {
    ...mapMutations("user", ["ResetUserInfo"]),
    savePassword() {
      this.$refs.modifyPwdForm.validate(valid => {
        if (valid) {
          changePassword({
            username: this.userInfo.userName,
            password: this.pwdModify.password,
            newPassword: this.pwdModify.newPassword
          }).then((res) => {
            if(res.code == 0){
              this.$message.success("successfully change password!");
            }
            this.showPassword = false;
          });
        } else {
          return false;
        }
      });
    },
    clearPassword() {
      this.pwdModify = {
        password: "",
        newPassword: "",
        confirmPassword: ""
      };
      this.$refs.modifyPwdForm.clearValidate();
    },
    openChooseImg() {
      this.$refs.chooseImg.open();
    },
    async enterImg(url) {
      const res = await setUserInfo({ headerImg: url, ID: this.userInfo.ID });
      if (res.code == 0) {
        this.ResetUserInfo({ headerImg: url });
        this.$message({
          type: "success",
          message: "Set success"
        });
      }
    },
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>
<style lang="scss">
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
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
.avatar-box {
  box-shadow: -2px 0 20px -16px;
  width: 80%;
  height: 100%;
  .user-card {
    min-height: calc(90vh - 200px);
    padding: 30px 20px;
    text-align: center;
    .el-avatar {
      border-radius: 50%;
    }
    .user-personality {
      padding: 24px 0;
      text-align: center;
      p {
        font-size: 16px;
      }
      .nickname {
        font-size: 26px;
      }
      .person-info{
        margin-top: 6px;
        font-size: 14px;
        color:#999
      }
    }
    .user-information {
      width: 100%;
      height: 100%;
      text-align: left;
      ul {
        display: inline-block;
        height: 100%;
        li {
          i {
            margin-right: 8px;
          }
          padding: 20px 0;
          font-size: 16px;
          font-weight: 400;
          color: #606266;
        }
      }
    }
  }
}
.user-addcount {
  ul {
    li {
      .title {
        padding: 10px;
        font-size: 18px;
        color: #696969;
      }
      .desc {
        font-size: 16px;
        padding: 0 10px 20px 10px;
        color: #a9a9a9;
        a {
          color: rgb(64, 158, 255);
          float: right;
        }
      }
      border-bottom: 2px solid #f0f2f5;
    }
  }
}
.user-headpic-update{
    width: 120px;
    height: 120px;
    line-height: 120px;
    margin: 0 auto;
    display: flex;
    justify-content: center;
    border-radius: 20px;
     &:hover{
      color: #fff;
      background: linear-gradient(to bottom, rgba(255,255,255,0.15) 0%, rgba(0,0,0,0.15) 100%), radial-gradient(at top center, rgba(255,255,255,0.40) 0%, rgba(0,0,0,0.40) 120%) #989898;
      background-blend-mode: multiply,multiply;
      .update{
        color:#fff ;
      }
    }
    .update{
      height: 120px;
      width: 120px;
      text-align: center;
      color:transparent;
    }
  }
</style>