<!--  -->
<template>
    <el-table-column
      :property="data.key"
      :label="data.label"
      :show-overflow-tooltip="true"
      align="center"
      >
      <template slot-scope="scope">
        <el-image style="width: 25px; height: 25px" v-if="data.type == 'img' && scope.row[data.key]" :src="scope.row[data.key][0]" :preview-src-list="[scope.row[data.key]]">
        </el-image>
        <span v-else-if="data.filter">{{ Allfilters(data.filter, scope.row[data.key], scope.row) }}</span>
        <span v-else>{{ scope.row[data.key] }}</span>
      </template>
    </el-table-column>
</template>

<script>
export default {
  data () {
    return {
    };
  },

  props: ["data"],

  filters: {
    
  },

  components: {},

  computed: {},

  mounted() {},

  methods: {
    applyStatusFilters(val) {
      const result = {
        0: "待审核",
        1: "审核通过",
        2: "审核失败",
      }
      return result[val] || result[0];
    },
    expiredFilters(val) {
      const result = {
        0: "无效",
        1: "有效",
      }
      return result[val] || "";
    },
    authSourceFilters(val) {
      const result = {
        1: "我的",
        2: "圈子",
      }
      return result[val] || "";
    },
    workCityFilters(val, row) {
      return `${val} - ${row.workCity}${row.workTown}`
    },
    Allfilters(name, val, row) {
      return this[name] && this[name](val, row)
    }
  }
}

</script>
<style lang='scss' scoped>
</style>