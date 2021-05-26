<template>
  <div class="init">
    <p class="in-one a-fadeinT">Welcome to GIN-VUE-ADMIN</p>
    <p class="in-two a-fadeinT">You need to initialize your database and populate the initial data</p>
    <div class="form-card in-three a-fadeinB">
      <el-form ref="form" :model="form" label-width="100px">
        <el-form-item label="Database type">
          <el-select disabled v-model="form.sqlType" placeholder="please choose">
            <el-option key="mysql" label="mysql(I only support mysql)" value="mysql">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="host">
          <el-input v-model="form.host" placeholder="Please enter the database link"></el-input>
        </el-form-item>
        <el-form-item label="port">
          <el-input v-model="form.port" placeholder="Please enter the database port"></el-input>
        </el-form-item>
        <el-form-item label="userName">
          <el-input v-model="form.userName" placeholder="Please enter the database username"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input
            v-model="form.password"
            placeholder="Please enter the database password (no empty)"
          ></el-input>
        </el-form-item>
        <el-form-item label="dbName">
          <el-input v-model="form.dbName" placeholder="Please enter the name of the database"></el-input>
        </el-form-item>
        <el-form-item>
          <div style="text-align: right">
            <el-button type="primary" @click="onSubmit">Immediate initialization</el-button>
          </div>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { initDB } from "@/api/initdb";
export default {
  name: "Init",
  data() {
    return {
      form: {
        sqlType: "mysql",
        host: "127.0.0.1",
        port: "3306",
        userName: "root",
        password: "",
        dbName: "gva",
      },
    };
  },
  methods: {
    async onSubmit() {
      const loading = this.$loading({
        lock: true,
        text: "Initializing the database, please wait",
        spinner: "el-icon-loading",
        background: "rgba(0, 0, 0, 0.7)",
      });
      try {
        const res = await initDB(this.form);
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: res.msg,
          });
          this.$router.push({name:"login"})
        }
          loading.close();
      } catch (err) {
          loading.close();
      }
    },
  },
};
</script>
<style lang="scss">
.init {
  height: 100vh;
  flex-direction: column;
  display: flex;
  align-items: center;
  background: #fff;
}
.in-one {
  font-size: 26px;
  line-height: 98px;
}
.in-two {
  font-size: 22px;
}
.form-card {
  margin-top: 60px;
  box-shadow: 0px 0px 5px 0px rgba(5, 12, 66, 0.15);
  width: 60vw;
  height: auto;
  background: #fff;
  padding: 20px;
  border-radius: 6px;
}
</style>
