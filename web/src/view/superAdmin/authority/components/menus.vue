<template>
  <div>
    <div class="clearflex">
      <el-button @click="relation" class="fl-right" size="small" type="primary">Confirm</el-button>
    </div>
    <el-tree
      :data="menuTreeData"
      :default-checked-keys="menuTreeIds"
      :props="menuDefaultProps"
      @check="nodeChange"
      default-expand-all
      highlight-current
      node-key="ID"
      ref="menuTree"
      show-checkbox
    >
     <span class="custom-tree-node" slot-scope="{ node , data }">
        <span>{{ node.label }}</span>
        <span>
          <el-button
            type="text"
            size="mini"
            :style="{color:row.defaultRouter == data.name?'#E6A23C':'#85ce61'}"
            :disabled="!node.checked"
            @click="() => setDefault(data)">
            {{row.defaultRouter == data.name?"Homepage":"Setup as front page"}}
          </el-button>
        </span>
      </span>
    </el-tree>
  </div>
</template>
<script>
import { getBaseMenuTree, getMenuAuthority, addMenuAuthority } from '@/api/menu'
import {
  updateAuthority,
} from "@/api/authority";
export default {
  name: 'Menus',
  props: {
    row: {
      default: function() {
        return {}
      },
      type: Object
    }
  },
  data() {
    return {
      menuTreeData: [],
      menuTreeIds: [],
      needConfirm:false,
      menuDefaultProps: {
        children: 'children',
        label: function(data){
          return data.meta.title
        }
      }
    }
  },
  methods: {
    async setDefault(data){
      const res = await updateAuthority({authorityId: this.row.authorityId,AuthorityName: this.row.authorityName,parentId: this.row.parentId,defaultRouter:data.name})
      if(res.code == 0){
        this.$message({type:"success",message:"Set success"})
        this.row.defaultRouter = res.data.authority.defaultRouter
      }
    },
    nodeChange(){
      this.needConfirm = true
    },
    // Switching interception method for exposure to outer layers
    enterAndNext(){
      this.relation()
    },
    // Related tree confirmation method
    async relation() {
      const checkArr = this.$refs.menuTree.getCheckedNodes(false, true)
      const res = await addMenuAuthority({
        menus: checkArr,
        authorityId: this.row.authorityId
      })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: 'Menu setup success!'
        })
      }
    }
  },
  async created() {
    // Get all menus trees
    const res = await getBaseMenuTree()
    this.menuTreeData = res.data.menus

    const res1 = await getMenuAuthority({ authorityId: this.row.authorityId })
    const menus = res1.data.menus
    const arr = []
    menus.map(item => {
      // Prevent direct selection of parents
      if (!menus.some(same => same.parentId === item.menuId)) {
        arr.push(Number(item.menuId))
      }
    })
    this.menuTreeIds = arr
  }
}
</script>
<style lang="scss">
</style>