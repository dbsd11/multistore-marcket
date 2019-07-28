<template>
    <div class="fillContain">
        <head-top>发起交易</head-top>
        <el-row style="margin-top: 20px;">
            <el-col :span="14" :offset="4">
                <header class="form_header">发起交易</header>
                <el-form>
                    <el-form-item label="顾客号" prop="customerId">
                        <el-input v-model="customerId"></el-input>
                    </el-form-item>
                    <el-form-item label="商品名称" prop="todo2">
                            <el-select v-model="selectGoods" placeholder="输入商品名称/商品码进行检索" filterable remote reserve-keyword :remote-method="(queryString)=>{
                                selectGoodsFilter(queryString, allSelectList);
                            }">
                                    <el-option
                                            v-for="item in selectGoods"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value">
                                    </el-option>
                            </el-select>
                        </el-form-item>
                </el-form>
                <el-button type="primary" @click="dialogFormVisible = true">确 定</el-button>
                <el-dialog title="确定发起交易" v-model="dialogFormVisible">
                        <div slot="footer" class="dialog-footer">
                          <el-button @click="dialogFormVisible = false">取 消</el-button>
                          <el-button type="primary" @click="initiateTrade">提 交</el-button>
                        </div>
                    </el-dialog>
              </el-dialog>
            </el-col>
        </el-row>
    </div>

</template>

<script>
	import headTop from '../components/headTop'

    export default{
        data(){
            return {
                customerId: "",
                goodsCode: "",
                selectGoods: []
            }
        },
        components:{
            headTop
        },
        created(){
            this.initData()
        },
        methods: {
            async initData(){},
            async getSelectGoods(goodsNameOrCodeLike){},
            selectGoodsFilter(queryString, allSelectList){},
            initiateTrade(){}
        }
    } 

</script>

<style lang="less">
    @import '../style/mixin';

    .table_container{
        padding: 20px;
    }
</style>