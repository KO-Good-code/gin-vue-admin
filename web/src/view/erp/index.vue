<template>
  <div>
    <el-row :gutter="15">
      <el-form ref="elForm" :model="formData" inline size="medium" >
        <el-form-item size="large" label="最高限价抽取比例">
            <el-input-number v-model="formData.max" :min="0.01" :max="1" :step="0.01" label="描述文字"></el-input-number>
        </el-form-item>
        <el-form-item size="large" label="平均价抽取比例">
            <el-input-number v-model="formData.min" :min="0.01" :max="1" :step="0.01" label="描述文字"></el-input-number>
        </el-form-item>
        <el-form-item size="large" label="最高限价">
            <el-input v-model="formData.num"></el-input>
        </el-form-item>
        <el-row>
          <el-form-item size="large" label="双底线">
            {{query.max}}~{{query.min}}
          </el-form-item>
        </el-row>
      </el-form>
      <el-form ref="form"  :model="form" inline :rules="rules" size="medium" @key.enter.native="submitForm('form')">
        <el-form-item size="large" label="公司名称" prop="name">
            <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item size="large" label="投标价" prop="num">
            <el-input v-model.number="form.num"></el-input>
        </el-form-item>
        <el-form-item size="large">
          <!-- <el-button type="primary" @click="submitForm">计算</el-button> -->
          <el-button icon="el-icon-edit" type="primary" native-type="submit" @click="submitForm('form')">新增</el-button>
        </el-form-item>
      </el-form>
    </el-row>
    <el-table
      ref="multipleTable"
      :data="list"
      :default-sort = "{prop: 'num', order: 'descending'}"
      :row-class-name="tableRowClassName"
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
            <el-button type="text" size="mini" @click="edit(scope.row)">修改</el-button>
            <el-button type="text" size="mini" @click="remove(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        title="提示"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="handleClose">
        <el-form ref="newform"  :model="newform" :rules="rules" size="medium">
          <el-form-item size="large" label="公司名称" prop="name">
              <el-input v-model="newform.name"></el-input>
          </el-form-item>
          <el-form-item size="large" label="投标价" prop="num">
              <el-input v-model.number="newform.num"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="add">确 定</el-button>
        </span>
      </el-dialog>
  </div>
</template>
<script>
// import tableList from "./components/tableList.vue"
export default {
  props: [],
  data() {
    return {
      dialogVisible: false,
      formData: {
        max: null,
        min: null,
        num: 0
      },
      newform: {
        name: "",
        num: null
      },
      form: {
        name: "",
        num: ""
      },
      query: {
        max:0,
        min: 0
      },
      rules: {
        name: [{
          required: true,
          message: '请输入公司名称',
          trigger: 'blur'
        }],
        num: [{
          required: true,
          message: '请输入公司投标价',
          trigger: 'change'
        }],
      },
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
  watch: {
    formData: {
      handler(){
        this.computedNum();
      },
      deep: true
    }
    
  },
  created() {},
  methods: {
    // 计算
    computedNum() {
      const len = this.list.length;
      const total = this.list.reduce((i, t) => i + t.num,0);
      this.query.min = Math.floor(total/len*this.formData.min);
      console.log(this.query.min)
      this.query.max =Math.floor(this.formData.num*this.formData.max);
    },
    // 新增数据
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (!valid) return
        // TODO 提交表单
        this.list.push({
          ...this.form
        });
        this.resetForm(formName)
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    proportions(val) {
      const result = (100 - (val/this.formData.num*100)).toFixed(2);
      return `${result}%`
    },
    // 确定
    add(){
      this.$confirm('此操作将永久修改该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // TODO 提交修改数据接口
        this.list[this.newform.index].name = this.newform.name;
        this.list[this.newform.index].num = this.newform.num;
        console.log(this.newform.index)
        this.dialogVisible = false;
      })
      
    },
    // 删除数据按钮
    remove(row) {
      this.$confirm('此操作将永久删除该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // TODO 删除数据接口
        this.list = this.list.filter( i => i.name !== row.name)
      })
    },
    // 修改数据按钮
    edit(row){
      this.dialogVisible = true;
      this.newform.name = row.name;
      this.newform.num = row.num;
      this.newform.id = id;
    },
    // 关闭弹窗回调
    handleClose(){
      this.resetForm("newform")
    },
    // 筛选双底线
    filterHandler(value, row, column){
      const property = column['property'];
      return row[property] >= this.query.min || row[property]>= this.query.max;
    },
    // 前三名颜色
    tableRowClassName({row, rowIndex}) {
      if (rowIndex <= 2) {
        return 'top';
      }
      return '';
    }
  }
}

</script>
<style lang="scss">
.top{
  background-color:oldlace !important;
}
</style>
