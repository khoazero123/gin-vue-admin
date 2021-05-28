<template>
  <div class="system">
    <el-form :model="config" label-width="100px" ref="form" class="system">
      <!--  System start  -->
      <h2>System Configuration</h2>
      <el-form-item label="Environmental value">
        <el-input v-model="config.system.env"></el-input>
      </el-form-item>
      <el-form-item label="Port value">
        <el-input v-model.number="config.system.addr"></el-input>
      </el-form-item>
      <el-form-item label="Database type">
        <el-select v-model="config.system.dbType">
          <el-option value="mysql"></el-option>
          <el-option value="sqlite"></el-option>
          <el-option value="sqlserver"></el-option>
          <el-option value="postgresql"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="OSS type">
        <el-select v-model="config.system.ossType">
          <el-option value="local"></el-option>
          <el-option value="qiniu"></el-option>
          <el-option value="tencent-cos"></el-option>
          <el-option value="aliyun-oss"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Profile environment variable name">
        <el-input v-model.number="config.system.configEnv"></el-input>
      </el-form-item>
      <el-form-item label="Data initialization">
        <el-checkbox v-model="config.system.needInitData">Open</el-checkbox>
      </el-form-item>
      <el-form-item label="Multiple login interception">
        <el-checkbox v-model="config.system.useMultipoint">Open</el-checkbox>
      </el-form-item>
      <!--  System end  -->

      <!--  JWT start  -->
      <h2>JWT signature</h2>
      <el-form-item label="JWT signature">
        <el-input v-model="config.jwt.signingKey"></el-input>
      </el-form-item>
      <!--  JWT end  -->

      <!--  Zap start  -->
      <h2>Zap log configuration</h2>
      <el-form-item label="level">
        <el-input v-model.number="config.zap.level"></el-input>
      </el-form-item>
      <el-form-item label="Output">
        <el-input v-model="config.zap.format"></el-input>
      </el-form-item>
      <el-form-item label="Log prefix">
        <el-input v-model="config.zap.prefix"></el-input>
      </el-form-item>
      <el-form-item label="Log folder">
        <el-input v-model="config.zap.director"></el-input>
      </el-form-item>
      <el-form-item label="Soft link name">
        <el-input v-model="config.zap.linkName"></el-input>
      </el-form-item>
      <el-form-item label="Encoding stage">
        <el-input v-model="config.zap.encodeLevel"></el-input>
      </el-form-item>
      <el-form-item label="Stack name">
        <el-input v-model="config.zap.stacktraceKey"></el-input>
      </el-form-item>
      <el-form-item label="Display">
        <el-checkbox v-model="config.zap.showLine"></el-checkbox>
      </el-form-item>
      <el-form-item label="Output console">
        <el-checkbox v-model="config.zap.logInConsole"></el-checkbox>
      </el-form-item>
      <!--  Zap end  -->

      <!--  Redis start  -->
      <h2>Redis Admin Database Configuration</h2>
      <el-form-item label="db">
        <el-input v-model="config.redis.db"></el-input>
      </el-form-item>
      <el-form-item label="addr">
        <el-input v-model="config.redis.addr"></el-input>
      </el-form-item>
      <el-form-item label="password">
        <el-input v-model="config.redis.password"></el-input>
      </el-form-item>
      <!--  Redis end  -->

      <!--  Email start  -->
      <h2>Mailbox configuration</h2>
      <el-form-item label="Receiver mailbox">
        <el-input v-model="config.email.to" placeholder="Can be separated by a comma"></el-input>
      </el-form-item>
      <el-form-item label="port">
        <el-input v-model.number="config.email.port"></el-input>
      </el-form-item>
      <el-form-item label="Sender mailbox">
        <el-input v-model="config.email.from"></el-input>
      </el-form-item>
      <el-form-item label="host">
        <el-input v-model="config.email.host"></el-input>
      </el-form-item>
      <el-form-item label="Is it SSL?">
        <el-checkbox v-model="config.email.isSSL"></el-checkbox>
      </el-form-item>
      <el-form-item label="secret">
        <el-input v-model="config.email.secret"></el-input>
      </el-form-item>
      <el-form-item label="Test mail">
        <el-button @click="email">Test mail</el-button>
      </el-form-item>
      <!--  Email end  -->

      <!--  Casbin start  -->
      <h2>Casbin configuration</h2>
      <el-form-item label="Model address">
        <el-input v-model="config.casbin.modelPath"></el-input>
      </el-form-item>
      <!--  Casbin end  -->

      <!--  Captcha start  -->
      <h2>Verification code configuration</h2>
      <el-form-item label="keyLong">
        <el-input v-model.number="config.captcha.keyLong"></el-input>
      </el-form-item>
      <el-form-item label="imgWidth">
        <el-input v-model.number="config.captcha.imgWidth"></el-input>
      </el-form-item>
      <el-form-item label="imgHeight">
        <el-input v-model.number="config.captcha.imgHeight"></el-input>
      </el-form-item>
      <!--  Captcha end  -->

      <!--  dbType start  -->
      <template v-if="config.system.dbType == 'mysql'">
        <h2>mysql Admin Database Configuration</h2>
        <el-form-item label="username">
          <el-input v-model="config.mysql.username"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.mysql.password"></el-input>
        </el-form-item>
        <el-form-item label="path">
          <el-input v-model="config.mysql.path"></el-input>
        </el-form-item>
        <el-form-item label="dbname">
          <el-input v-model="config.mysql.dbname"></el-input>
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.mysql.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.mysql.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="logMode">
          <el-checkbox v-model="config.mysql.logMode"></el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.dbType == 'sqlite'">
        <h2>sqlite Admin Database Configuration</h2>
        <el-form-item label="path">
          <el-input v-model="config.mysql.path"></el-input>
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.mysql.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.mysql.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.mysql.logger"></el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.dbType == 'sqlserver'">
        <h2>sqlserver Admin Database Configuration</h2>
        <el-form-item label="username">
          <el-input v-model="config.sqlserver.username"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.sqlserver.password"></el-input>
        </el-form-item>
        <el-form-item label="path">
          <el-input v-model="config.sqlserver.path"></el-input>
        </el-form-item>
        <el-form-item label="dbname">
          <el-input v-model="config.sqlserver.dbname"></el-input>
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.sqlserver.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.sqlserver.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.sqlserver.logger"></el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.dbType == 'postgresql'">
        <h2>postgresql Admin Database Configuration</h2>
        <el-form-item label="username">
          <el-input v-model="config.mysql.username"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.mysql.password"></el-input>
        </el-form-item>
        <el-form-item label="dbName">
          <el-input v-model="config.mysql.dbName"></el-input>
        </el-form-item>
        <el-form-item label="port">
          <el-input v-model="config.mysql.port"></el-input>
        </el-form-item>
        <el-form-item label="config">
          <el-input v-model="config.mysql.config"></el-input>
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.mysql.maxIdleConns"></el-input>
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.mysql.maxOpenConns"></el-input>
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.mysql.logger"></el-checkbox>
        </el-form-item>
        <el-form-item label="prefer-simple-protocol">
          <el-checkbox v-model="config.mysql.preferSimpleProtocol"></el-checkbox>
        </el-form-item>
      </template>
      <!--  dbType end  -->

      <!--  ossType start  -->
      <template v-if="config.system.ossType == 'local'">
        <h2>Local upload configuration</h2>
        <el-form-item label="Local file path">
          <el-input v-model="config.local.path"></el-input>
        </el-form-item>
      </template>
      <template v-if="config.system.ossType == 'qiniu'">
        <h2>QiniU upload configuration</h2>
        <el-form-item label="Storage area">
          <el-input v-model="config.qiniu.zone"></el-input>
        </el-form-item>
        <el-form-item label="Spatial name">
          <el-input v-model="config.qiniu.bucket"></el-input>
        </el-form-item>
        <el-form-item label="CDN accelerated domain name">
          <el-input v-model="config.qiniu.imgPath"></el-input>
        </el-form-item>
        <el-form-item label="Whether to use HTTPS">
          <el-checkbox v-model="config.qiniu.useHttps">Open</el-checkbox>
        </el-form-item>
        <el-form-item label="accessKey">
          <el-input v-model="config.qiniu.accessKey"></el-input>
        </el-form-item>
        <el-form-item label="secretKey">
          <el-input v-model="config.qiniu.secretKey"></el-input>
        </el-form-item>
        <el-form-item label="Uploading whether to use CDN upload acceleration">
          <el-checkbox v-model="config.qiniu.useCdnDomains">Open</el-checkbox>
        </el-form-item>
      </template>
       <template v-if="config.system.ossType == 'tencent-cos'">
        <h2>Tencent Cloud CoS Upload Configuration</h2>
        <el-form-item label="bucket">
          <el-input v-model="config.tencentCOS.bucket"></el-input>
        </el-form-item>
        <el-form-item label="region">
          <el-input v-model="config.tencentCOS.region"></el-input>
        </el-form-item>
        <el-form-item label="secretID">
          <el-input v-model="config.tencentCOS.secretID"></el-input>
        </el-form-item>
        <el-form-item label="secretKey">
          <el-input v-model="config.tencentCOS.secretKey"></el-input>
        </el-form-item>
        <el-form-item label="pathPrefix">
          <el-input v-model="config.tencentCOS.pathPrefix"></el-input>
        </el-form-item>
        <el-form-item label="baseURL">
          <el-input v-model="config.tencentCOS.baseURL"></el-input>
        </el-form-item>
      </template>
       <template v-if="config.system.ossType == 'aliyun-oss'">
        <h2>Ali Cloud OSS upload configuration</h2>
        <el-form-item label="endpoint">
          <el-input v-model="config.aliyunOSS.endpoint"></el-input>
        </el-form-item>
        <el-form-item label="accessKeyId">
          <el-input v-model="config.aliyunOSS.accessKeyId"></el-input>
        </el-form-item>
        <el-form-item label="accessKeySecret">
          <el-input v-model="config.aliyunOSS.accessKeySecret"></el-input>
        </el-form-item>
        <el-form-item label="bucketName">
          <el-input v-model="config.aliyunOSS.bucketName"></el-input>
        </el-form-item>
        <el-form-item label="bucketUrl">
          <el-input v-model="config.aliyunOSS.bucketUrl"></el-input>
        </el-form-item>
      </template>
      <!--  ossType end  -->

      <el-form-item>
        <el-button @click="update" type="primary">Update immediately</el-button>
        <el-button @click="reload" type="primary">Restart service (in development)</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from "@/api/system";
import { emailTest } from "@/api/email";
export default {
  name: "Config",
  data() {
    return {
      config: {
        system: {},
        jwt: {},
        casbin: {},
        mysql: {},
        sqlite: {},
        redis: {},
        qiniu: {},
        tencentCOS:{},
        aliyunOSS:{},
        captcha: {},
        zap: {},
        local: {},
        email: {}
      }
    };
  },
  async created() {
    await this.initForm();
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig();
      if (res.code == 0) {
        this.config = res.data.config;
      }
    },
    reload() {},
    async update() {
      const res = await setSystemConfig({ config: this.config });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "Profile setting success"
        });
        await this.initForm();
      }
    },
    async email() {
      const res = await emailTest();
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "Mail sent successfully"
        });
        await this.initForm();
      } else {
        this.$message({
          type: "error",
          message: "Mail delivery failed"
        });
      }
    }
  }
};
</script>
<style lang="scss">
.system {
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 1px 3px 0px #e7e8e8;
  }
}
</style>
