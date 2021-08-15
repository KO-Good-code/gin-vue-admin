<template>
  <div>
    <el-row :gutter="15">
      <el-form ref="elForm" :model="formData" inline :rules="rules" size="medium" label-width="150px">
        <!-- <el-col :span="12">
          <el-form-item label="单行文本" prop="mobile">
            <el-input v-model="formData.mobile" placeholder="请输入手机号" :maxlength="11" show-word-limit clearable
              prefix-icon='el-icon-mobile' :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-row>
            <el-col :span="9">
              <el-form-item label="日期范围" prop="field118">
                <el-date-picker type="daterange" v-model="formData.field118" format="yyyy-MM-dd"
                  value-format="yyyy-MM-dd" :style="{width: '100%'}" start-placeholder="开始日期"
                  end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="9">
              <el-form-item label="选择" prop="field119">
                <el-select v-model="formData.field119" placeholder="请选择选择" clearable :style="{width: '100%'}">
                  <el-option v-for="(item, index) in field119Options" :key="index" :label="item.label"
                    :value="item.value" :disabled="item.disabled"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
        </el-col> -->
        
        <el-form-item size="large" label="最高限价抽取比例">
            <el-input-number v-model="formData.max" :min="0.01" :max="1" :step="0.01" label="描述文字"></el-input-number>
        </el-form-item>
        <el-form-item size="large" label="平均价抽取比例">
            <el-input-number v-model="formData.min" :min="0.01" :max="1" :step="0.01" label="描述文字"></el-input-number>
        </el-form-item>
        <el-form-item size="large" label="最高限价">
            <el-input v-model="formData.num"></el-input>
        </el-form-item>
        <el-col :span="24">
          <el-form-item size="large">
            <el-button type="primary" @click="submitForm">计算</el-button>
            <el-button @click="resetForm">新增</el-button>
          </el-form-item>
        </el-col>
      </el-form>
    </el-row>
    <el-row>
      双底线：{{query.max}}~{{query.min}}
    </el-row>
    <el-table
      ref="multipleTable"
      :data="list"
      :default-sort = "{prop: 'num', order: 'descending'}"
      >
        <!-- <tableList v-for="(i, t) in listMap" :key="t" :data="i" /> -->
        <el-table-column
          property="name"
          label="公司名称"
          :show-overflow-tooltip="true"
          align="center"
          > </el-table-column>
          <el-table-column
          property="num"
          label="投标价"
          :show-overflow-tooltip="true"
          align="center"
          :filters="[{text:'双底线', value: 1}]"
          :filter-method="filterHandler"
          sortable
          > </el-table-column>
          <el-table-column label="浮动比例" :show-overflow-tooltip="true" align="center" >
            <template slot-scope="scope">
              <el-button icon="el-icon-bottom" type="text" style="color:#67C23A;">{{proportions(scope.row.num)}}</el-button>
            </template>
          </el-table-column>
        <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
          <template slot-scope="scope">
            <el-button type="text" size="mini" @click="add">修改</el-button>
            <el-button type="text" size="mini" @click="remove(scope.$index)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
  </div>
</template>
<script>
// import tableList from "./components/tableList.vue"
export default {
  props: [],
  data() {
    return {
      formData: {
        mobile: undefined,
        field118: null,
        field119: undefined,
        num: 0
      },
      query: {
        max:0,
        min: 0
      },
      rules: {
        mobile: [{
          required: true,
          message: '请输入手机号',
          trigger: 'blur'
        }, {
          pattern: /^1(3|4|5|7|8|9)\d{9}$/,
          message: '手机号格式错误',
          trigger: 'blur'
        }],
        field118: [{
          required: true,
          message: '日期范围不能为空',
          trigger: 'change'
        }],
        field119: [{
          required: true,
          message: '请选择选择',
          trigger: 'change'
        }],
      },
      field119Options: [{
        "label": "选项一",
        "value": 1
      }, {
        "label": "选项二",
        "value": 2
      }],
      list: [
        {
          name: "四川君羊建设集团有限公司",
          num: 31149199
        },
        {
          name: "百年建设集团有限公司",
          num: 31668965
        },
        {
          name: "中恒太鸿建设有限公司",
          num: 31674722
        },
        {
          name: "中建城开环境建设有限公司",
          num: 31962687
        },
        {
          name: "四川富祥建筑工程有限公司",
          num: 31964488
        },
        {
          name: "中诚智翔建设集团有限公司",
          num: 32102915
        },
        {
          name: "新疆交通建设集团股份有限公司",
          num: 42997361
        },
      ]
    }
  },
  computed: {
    listMap(){
      return [
        {
          key: "name",
          label: "公司名称"
        },
        {
          key: "num",
          label: "投标价"
        },
        // {
        //   key: "max",
        //   label: "最高限价抽取比例"
        // },
        // {
        //   key: "min",
        //   label: "平均价抽取比例"
        // },
      ]
    }
  },
  filters: {
    
    
  },
  created() {},
  methods: {
    submitForm() {
      const len = this.list.length;
      const total = this.list.reduce((i, t) => i + t.num,0);
      this.query.min = Math.floor(total/len*this.formData.min);
      this.query.max =Math.floor(this.formData.num*this.formData.max);
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
    add(){},
    proportions(val) {
      const result = (100 - (val/this.formData.num*100)).toFixed(2);
      return `${result}%`
    },
    filterHandler(value, row, column){
      const property = column['property'];
      return row[property] >= this.query.min || row[property]>= this.query.max;
    }
  }
}

</script>
<style>
</style>
