<template>
  <el-table :data="data">
    <el-table-column
      fixed
      prop="name"
      label="日期">
    </el-table-column>
    <el-table-column
      fixed
      prop="address"
      label="日期">
    </el-table-column>
  </el-table>
</template>

<script lang="ts">
  import { Component, Vue } from "vue-property-decorator"
  import { Table, TableColumn } from "element-ui"

  @Component({
    components: {
      ElTable: Table,
      ElTableColumn: TableColumn
    }
  })
  export default class List extends Vue {
    private data = [
      {
        name: 'John Brown',
        age: 18,
        address: 'New York No. 1 Lake Park',
        date: '2016-10-03'
      },
      {
        name: 'Jim Green',
        age: 24,
        address: 'London No. 1 Lake Park',
        date: '2016-10-01'
      },
      {
        name: 'Joe Black',
        age: 30,
        address: 'Sydney No. 1 Lake Park',
        date: '2016-10-02'
      },
      {
        name: 'Jon Snow',
        age: 26,
        address: 'Ottawa No. 2 Lake Park',
        date: '2016-10-04'
      }
    ]
  }
</script>

<style scoped>

</style>