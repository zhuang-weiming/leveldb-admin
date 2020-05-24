<template>
  <div class="key-list">
    <el-row class="filter-header">
      <el-col :span="11" class="filter">
        <el-input
          placeholder="通过前缀搜索"
          v-model="prefix">
          <i slot="prefix" class="el-input__icon el-icon-search"></i>
        </el-input>
      </el-col>
      <el-col :span="2" :offset="9">
        <el-select v-model="format" placeholder="格式">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-col>
    </el-row>
    <el-row :gutter="40">
      <el-col :span="11">
        <el-table
          :data="data"
          class="list"
          :border="true"
          ref="table"
          :height="tableHeight"
          :stripe="true"
          @row-click="handleItemClick"
        >
          <el-table-column
            prop="keyName"
            label="key">
          </el-table-column>
          <el-table-column
            width="80"
            fixed="right"
            label="操作">
            <template slot-scope="scope">
              <el-button @click="handleDelete(scope.row)" size="mini" type="danger" icon="el-icon-delete"
                         round></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
      <el-col :span="11">
        <el-input
          :rows="tableHeight / 21"
          type="textarea"
          placeholder="value"
          v-model="value">
        </el-input>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
  import {Component, Prop, Vue, Watch} from "vue-property-decorator"
  import {Table, TableColumn, Message, Input, Row, Col, Button, Select, Option} from "element-ui"
  import {getKeys, getValue} from "@/api/leveldb_web"

  interface Item {
    keyName: string;
  }

  @Component({
    components: {
      ElTable: Table,
      ElTableColumn: TableColumn,
      ElInput: Input,
      ElRow: Row,
      ElCol: Col,
      ElButton: Button,
      ElSelect: Select,
      ElOption: Option,
    }
  })
  export default class List extends Vue {
    @Prop({default: ""})
    private db!: string

    @Prop({default: 0})
    private offsetTop!: number

    private data: Item[] = []

    private prefix = ""
    private tableHeight = 100
    private value = ""
    private format = ""
    private options = [
      {
        label: 'Json',
        value: 'Json',
      }
    ]

    mounted() {
      this.$nextTick(function () {
        this.tableHeight = window.innerHeight - this.offsetTop - 60 - 20

        // 监听窗口大小变化
        // eslint-disable-next-line @typescript-eslint/no-this-alias
        const self = this
        window.onresize = function () {
          self.tableHeight = window.innerHeight - self.offsetTop - 60 - 20
        }
      })
    }

    created() {
      this.loadKeys()
    }

    handleItemClick(row: Item) {
      console.log(row)
      const db = this.db
      if (!db || db == "") {
        return Message.error("无效的db name")
      }
      getValue({db: this.db, key: row.keyName}).then(res => {
        this.value = res.data.Value
      })
    }

    handleDelete(row: Item) {
      console.log(row)
    }

    @Watch("prefix")
    onPrefixChange() {
      this.loadKeys()
    }

    loadKeys() {
      const db = this.db
      if (!db || db == "") {
        return Message.error("无效的db name")
      }

      getKeys({
        db: db,
        prefix: this.prefix,
        searchText: ""
      }).then(res => {
        this.data = res.data.Items.map((item: string) => {
          return {
            keyName: item
          }
        })
      })
    }
  }
</script>

<style>
  .key-list {
    height: 100%;
  }

  .filter-header {
    margin-bottom: 20px;
    height: 40px;
  }
  .filter {
    padding-right: 20px;
  }
</style>