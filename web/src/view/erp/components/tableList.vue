<!--  -->
<template>
    <el-table-column
      :property="data.key"
      :label="data.label"
      :show-overflow-tooltip="true"
      align="center"
      >
      <template slot-scope="scope">
        <el-select v-if="data.type === 'select'" v-model="scope.row[data.key]" placeholder="请选择">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
        <el-input v-else v-model="scope.row[data.key]" :placeholder="data.placeholder"></el-input>
      </template>
    </el-table-column>
</template>

<script>
export default {
  data () {
    return {
      options: [
        {
          label: "文字",
          value: "txt"
        },
        {
          label: "图片",
          value: "img"
        },
        {
          label: "枚举",
          value: "map"
        },
        {
          label: "时间",
          value: "time"
        },
      ]
    };
  },

  model: {
    prop: "value",
  },

  props: ["data", "value"],

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