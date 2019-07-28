<template>
    <div class="fillContain">
        <head-top>发起交易</head-top>
        <el-row style="margin-top: 20px;">
            <el-col :span="14" :offset="4">
                <header class="form_header">发起交易</header>
                <el-form label-position="left" label-width="10%">
                    <el-form-item label="顾客号">
                        <el-input v-model="customerNo" :maxlength="20"></el-input>
                    </el-form-item>
                    <el-form-item label="商品名称">
                            <el-select v-model="goodsCode" placeholder="输入商品名称/商品码进行检索" filterable remote reserve-keyword :remote-method="(queryString)=>{
                                return selectGoodsFilter(queryString, allSelectList);
                            }">
                                    <el-option
                                            v-for="item in selectGoods"
                                            :key="item.value"
                                            :label="item.label"
                                            :value="item.value">
                                    </el-option>
                            </el-select>
                        </el-form-item>
                    <el-form-item label="市场名称">
                        <el-select v-model="marcketCode" placeholder="输入市场名称进行检索" filterable remote reserve-keyword :remote-method="(queryString)=>{
                            return selectMarcketsFilter(queryString, allSelectList);
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
                goodsCode: "",
                marcketName: "",
                marcketCode: "",
                selectGoods: [{
                    label: '棉衣',
                    value: 'my01'
                }],
                selectMarckets: [{
                    label: '华北01',
                    value:'scn01'
                }],
                allSelectList:[],
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
            async initData(){
                this.customerNo = "rand"+this.randomString(10)
            },
            async getSelectGoods(goodsNameOrCodeLike){},
            selectGoodsFilter(queryString, allSelectList){
            },
            selectMarcketsFilter(queryString, allSelectList){
            },
            async initiateTradeData(){
                try{
                    var tradeResult = await initiateTrade({
                        customerNo: this.customerNo,
                        goodsName: this.selectGoods.filter(select => select.value==this.goodsCode)[0].label,
                        goodsCode: this.goodsCode,
                        marcketName: this.selectMarckets.filter(select => select.value==this.marcketCode)[0].label,
                        tradeTime: this.getCurDateStr('yyyy-MM-dd HH:mm:ss')
                    })
                        
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
            },

            getCurDateStr(fmt){
                var date = new Date()
                var o = {
                    "M+": date.getMonth() + 1, //月份 
                    "d+": date.getDate(), //日 
                    "H+": date.getHours(), //小时 
                    "m+": date.getMinutes(), //分 
                    "s+": date.getSeconds(), //秒 
                    "q+": Math.floor((date.getMonth() + 3) / 3), //季度 
                    "S": date.getMilliseconds() //毫秒 
                };

                if (/(y+)/.test(fmt)) {
                    fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
                }
                for (var k in o){
                    if (new RegExp("(" + k + ")").test(fmt)){
                        fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
                    }
                }
                
                return fmt
            },
            randomString(len) {
            　　len = len || 32;
            　　var $chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678';    /****默认去掉了容易混淆的字符oOLl,9gq,Vv,Uu,I1****/
            　　var maxPos = $chars.length;
            　　var pwd = '';
            　　for (var i = 0; i < len; i++) {
            　　　　pwd += $chars.charAt(Math.floor(Math.random() * maxPos));
            　　}
            　　return pwd;
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