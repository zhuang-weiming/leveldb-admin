<template>
  <div class="home">
    <el-tabs v-model="showTab" @tab-click="handleClick">
      <el-tab-pane v-for="(db, index) in dbs" :key="index" :label="db" :name="db">
        <key-list></key-list>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts">
  import { Component, Vue } from "vue-property-decorator"
  import KeyList from "@/components/KeyList.vue"
  import { getDbs } from "@/api/leveldb_web"
  import { Tabs, TabPane } from "element-ui"

  @Component({
    components: {
      KeyList,
      ElTabs: Tabs,
      ElTabPane: TabPane
    }
  })
  export default class Home extends Vue {
    private showTab = 'key'

    private dbs: Array<string> = []

    created() {
      getDbs().then(res => {
        this.dbs = res.data
      })
    }
    handleClick() {
      console.log("todo")
    }
  }
</script>

<style scoped>
  .home {
    padding: 20px;
  }
</style>
