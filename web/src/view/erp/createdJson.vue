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
        <el-button type="primary">新增</el-button>
        <el-button >重置</el-button>
      </el-form-item>
    </el-form>
    <el-table
      ref="multipleTable"
      :data="Json.map"
      >
        <tableList v-for="(i, t) in listMap" :key="t" :data="i" />
        <el-table-column label="审核操作" align="center" class-name="small-padding fixed-width">
          <template slot-scope="scope">
            <el-button type="text" size="mini" >通过</el-button>
            <el-button type="text" size="mini" >拒绝</el-button>
          </template>
        </el-table-column>
      </el-table>
  </div>
</template>

<script>
import tableList from "./components/tableList.vue"
export default {
  data () {
    return {
      formData: {
        mobile: undefined,
        field118: null,
        field119: undefined,
        
      },
      field119Options: [],
      Json: {
        name: "filed.json",
        map: [

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
          key: "uiName",
          label: "展示名",
        },
        {
          key: "name",
          label: "字段名",
        },
        {
          key: "dataType",
          label: "字段类型",
        },
        {
          key: "defaultValue",
          label: "默认值/枚举值",
        },
        {
          key: "name",
          label: "默认值/枚举值",
        },
        {
          key: "name",
          label: "展示名",
        },
      ];
      return result;
    }
  },

  mounted() {},

  methods: {}
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