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
      <el-col :span="2">
        <el-select v-model="format" placeholder="格式化">
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
          highlight-current-row
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
        <i v-show="searchText" class="el-icon-arrow-down next" @click="next"></i>
      </el-col>
      <el-col :span="11">
        <el-button @click="handleUpdate" size="small" type="warning" round class="update">更新</el-button>
        <el-input
          :rows="tableHeight / 21"
          type="textarea"
          resize="none"
          placeholder="value"
          v-model="currentValue">
        </el-input>
      </el-col>
    </el-row>

  </div>
</template>

<script lang="ts">
  import {Component, Prop, Vue, Watch} from "vue-property-decorator"
  import {Table, TableColumn, Message, Input, Row, Col, Button, Select, Option} from "element-ui"
  import {keys, keyInfo, keyDelete, keyUpdate} from "@/api/leveldb_admin"

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
    private currentOldValue = ''
    private currentValue = ''
    private currentKey = ""
    private format = ""
    private searchText = ""
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

    next() {
      this.loadKeys()
    }

    created() {
      this.loadKeys()
    }

    handleItemClick(row: Item) {
      const db = this.db
      if (!db || db == "") {
        return Message.error("无效的db name")
      }
      this.currentKey = row.keyName
      keyInfo({db: this.db, key: row.keyName}).then(res => {
        this.currentOldValue = res.data.Value
        this.currentValue = res.data.Value
      })
    }

    handleUpdate() {
      if (this.currentKey && (this.currentOldValue != this.currentValue)) {
          keyUpdate({db: this.db, key: this.currentKey, value: this.currentValue}).then(res => {
            if (res.data.Success) {
              Message.success("更新成功!")
            } else {
              Message.success("更新失败!")
            }
          })
      } else {
        Message.info("无变化")
      }
    }

    handleDelete(row: Item) {
      keyDelete({db: this.db, key: row.keyName}).then(res => {
        if (res.data.Success) {
          Message.success("删除成功!")
        } else {
          Message.success("删除失败!")
        }
      })
      delete this.data[this.data.indexOf(row)]

      this.data = this.data.map(item => {
        return item
      })
    }

    @Watch("prefix")
    onPrefixChange() {
      this.data = []
      this.loadKeys()
    }

    loadKeys() {
      const db = this.db
      if (!db || db == "") {
        return Message.error("无效的db name")
      }

      keys({
        db: db,
        prefix: this.prefix,
        searchText: this.searchText
      }).then(res => {
        this.searchText = res.data.IsPart ? res.data.SearchText : ""
        this.data = this.data.concat(res.data.Items.map((item: string) => {
          return {
            keyName: item
          }
        }))
      })
    }
  }
</script>

<style>
  .key-list {
    cursor: pointer;
    height: 100%;
  }

  .filter-header {
    margin-bottom: 20px;
    height: 40px;
  }
  .filter {
    padding-right: 20px;
    margin-right: 20px;
  }
  .update {
    position: absolute;
    z-index: 10;
    right: 10%;
    bottom: 4%;
  }
  .next {
    z-index: 10;
    width: 100%;
    text-align: center;
  }
  .el-table__body tr.current-row>td {
    background: #6ffff7 !important;
  }
</style>