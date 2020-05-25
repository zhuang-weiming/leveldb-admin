<template>
  <el-container class="home">
    <el-header>
    LevelDB Web Admin
    </el-header>
    <el-main class="main">
      <el-tabs v-model="showTab" @tab-click="handleClick" tab-position="left" ref="tabs">
        <el-tab-pane v-for="(db, index) in dbs" :key="index" :label="db" :name="db">
          <key-list :db.sync="db" :offset-top.sync="tabsOffsetTop"></key-list>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<script lang="ts">
  import { Component, Vue } from "vue-property-decorator"
  import KeyList from "@/components/KeyList.vue"
  import {dbs} from "@/api/leveldb_admin"
  import { Tabs, TabPane, Header, Main, Container } from "element-ui"

  @Component({
    components: {
      KeyList,
      ElTabs: Tabs,
      ElTabPane: TabPane,
      ElHeader: Header,
      ElMain: Main,
      ElContainer: Container,
    }
  })
  export default class Home extends Vue {
    private showTab = 'temp'

    private dbs: Array<string> = []

    private tabsOffsetTop = 0

    mounted() {
      this.$nextTick(function () {
        this.tabsOffsetTop = ((this.$refs.tabs as Vue).$el as HTMLElement).offsetTop
      })
    }

    created() {
      dbs().then(res => {
        this.dbs = res.data
      })
    }
    handleClick() {
      console.log("todo")
    }
  }
</script>

<style scoped>
  .el-header {
    background-color: #d8e4e4;
    color: #333;
    line-height: 60px;
    margin-bottom: 5px;
  }
  .el-main {
    padding-top: 20px;
  }
</style>
