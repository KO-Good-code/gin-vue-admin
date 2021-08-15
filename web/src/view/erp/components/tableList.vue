<!--  -->
<template>
    <el-table-column
      :property="data.key"
      :label="data.label"
      :show-overflow-tooltip="true"
      align="center"
      sortable
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
        <el-input v-else-if="data.type === 'input'" v-model="scope.row[data.key]" :placeholder="data.placeholder"></el-input>
        <span v-else>{{scope.row[data.key]}}</span>
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
    filterHandler(value, row, column){

    }
  }
}

</script>
<style lang='scss' scoped>
</style>