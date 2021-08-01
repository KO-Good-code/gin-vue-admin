<!-- 创建JSON表单 -->
<template>
  <div class="createdJson">
    <el-form ref="elForm" class="form" :model="formData" :inline="true" label-width="80px">
      <el-form-item label="选择" prop="field119">
        <el-select v-model="formData.field119" placeholder="请选择选择" clearable>
          <el-option v-for="(item, index) in field119Options" :key="index" :label="item.label"
            :value="item.value" :disabled="item.disabled"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="addJson">新增表单</el-button>
      </el-form-item>
    </el-form>
    <el-table
      ref="multipleTable"
      :data="Json.map"
      >
        <tableList v-for="(i, t) in listMap" :key="t" :data="i" />
        <el-table-column label="审核操作" align="center" class-name="small-padding fixed-width">
          <template slot-scope="scope">
            <el-button type="text" size="mini" @click="add">新增</el-button>
            <el-button type="text" size="mini" @click="remove(scope.$index)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog
        title="提示"
        :visible.sync="dialogVisible"
        width="30%"
        >
        <el-form ref="JsonForm" class="form" :model="Json" :inline="true" label-width="80px">
          <el-form-item prop="name" label="新建表命">
            <el-input v-model="Json.name"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="submit">确 定</el-button>
        </span>
      </el-dialog>

  </div>
</template>

<script>
import tableList from "./components/tableList.vue"
import { creatdJsonName } from "@/api/erp";
export default {
  data () {
    return {
      formData: {
        mobile: undefined,
        field119: undefined,
      },
      dialogVisible: false,
      field119Options: [],
      Json: {
        name: "",
        map: [
            {
              label: "sdsa",
              name: "132",
              dataType: "sdsa",
              defaultValue: "sdsa",
            }
        ]
      }
    };
  },

  components: {
    tableList
  },

  computed: {
    listMap() {
      const result = [
        {
          key: "label",
          label: "展示名",
          placeholder: "请输入展示名"
        },
        {
          key: "name",
          label: "字段名",
          placeholder: "请输入字段名"
        },
        {
          key: "dataType",
          label: "字段类型",
          type: "select"
        },
        {
          key: "defaultValue",
          label: "默认值/枚举值",
          placeholder: "0|男,1|女"
        },
      ];
      return result;
    }
  },

  mounted() {},

  methods: {
    add() {
      this.Json.map.push({
        uiName: "",
        name: "",
        dataType: "",
        defaultValue: "",
      })
    },
    remove(index) {
      this.Json.map.splice(index, 1)
    },
    addJson() {
      this.dialogVisible = true;
    },
    submit() {
      this.creatdJsonName()
    },
    async creatdJsonName() {
      try {
        const res = await creatdJsonName();
        console.log(res)
      } catch (error) {
        console.log(error)
      }
    }
  }
}

</script>
<style lang='scss' scoped>
.createdJson{
  .form{
    .el-form-item{
      width: 400px;
    }
  }
}
</style>