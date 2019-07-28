<template>
    <div class="fillContain">
        <head-top>发起交易</head-top>
        <el-row style="margin-top: 20px;">
            <el-col :span="14" :offset="4">
                <header class="form_header">发起交易</header>
                <el-form label-position="left" label-width="10%">
                    <el-form-item label="顾客号" prop="customerNo">
                        <el-input v-model="customerNo" :maxlength="20"></el-input>
                    </el-form-item>
                    <el-form-item label="商品名称" prop="selectGoods">
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
                    <el-form-item label="市场名称" prop="selectMarckets">
                        <el-select v-model="selectMarckets" placeholder="输入市场名称进行检索" filterable remote reserve-keyword :remote-method="(queryString)=>{
                            selectMacketsFilter(queryString, allSelectList);
                        }">
                                <el-option
                                        v-for="item in selectMarckets"
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
                          <el-button type="primary" @click="initiateTradeData">提 交</el-button>
                        </div>
                    </el-dialog>
              </el-dialog>
            </el-col>
        </el-row>
    </div>

</template>

<script>
    import headTop from '../components/headTop'
    import { initiateTrade } from '@/api/getData'

    export default{
        data(){
            return {
                customerNo: "",
                goodsName: "",
                marcketName: "",
                selectGoods: [{
                    label: '棉衣',
                    value: 'my01'
                }],
                selectMarckets: [{
                    label: '华北01',
                    value:'sc01'
                }],
                dialogFormVisible:false
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
            selectGoodsFilter(queryString, allSelectList){
                return selectGoods
            },
            selectMarcketsFilter(queryString, allSelectList){
                return selectMarckets
            },
            async initiateTradeData(){
                try{
                    var tradeResult = await initiateTrade({
                        customerNo: this.customerNo || "cus01",
                        goodsName: this.goodsName || "棉衣",
                        marcketName: this.marcketName || "华北01"
                    })

                    console.info(tradeResult)
                        
                    if(!tradeResult || tradeResult.error){
                        this.$message({
                            type: 'error',
                            message: tradeResult.error?tradeResult.error.toString():""
                        });
                        return
                    }

                    this.$message({
                        type: 'success',
                        message: '交易成功 result:'+JSON.stringify(tradeResult)
                    });
                    this.dialogFormVisible = false
                }catch(err){
                    console.log('发起交易失败', err);
                    this.$message({
                        type: 'error',
                        message: '发起交易失败, 请联系系统管理员'
                    });
                }
            }
        }
    } 

</script>

<style lang="less">
    @import '../style/mixin';

    .table_container{
        padding: 20px;
    }
</style>