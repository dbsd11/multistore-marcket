<template>
    <div class="fillContain">
        <head-top>交易历史列表</head-top>
        <div class="table_container">
            <el-table
                :data="tableData"
                @expand="expand"
                :expand-row-keys="expendRow"
                :row-key="row=>row.index"
                style="width:100%">
                <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" inline>
                            <el-form-item label="商品名称">
                                <span>{{ props.row.goodsName }}</span>
                            </el-form-item>
                            <el-form-item label="商品码">
                                <span>{{ props.row.goodsCode }}</span>
                            </el-form-item>
                            <el-form-item label="所属集市">
                                <span>{{ props.row.marcketName }}</span>
                            </el-form-item>
                            <el-form-item label="顾客号">
                                <span>{{ props.row.customerNo }}</span>
                            </el-form-item>
                            <el-form-item label="交易价格">
                                <span>{{ props.row.tradePrice }}</span>
                            </el-form-item>
                            <el-form-item label="交易时间">
                                <span>{{ props.row.tradeTime }}</span>
                            </el-form-item> 
                        </el-form>
                    </template>
                </el-table-column>
                
                <el-table-column label="商品名称" prop="goodsName"></el-table-column>
                <el-table-column label="商品码" prop="goodsCode"></el-table-column>
                <el-table-column label="所属集市" prop="marcketName"></el-table-column>
                <el-table-column label="顾客号" prop="customerNo"></el-table-column>
                <el-table-column label="交易价格" prop="tradePrice"></el-table-column>
                <el-table-column label="交易时间" prop="tradeTime"></el-table-column>
            </el-table>
            <div class="Pagination">
                <el-pagination
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    :current-page="currentPage"
                    :page-size="20"
                    layout="total, prev, pager, next"
                    :total="count">
                </el-pagination>
            </div>
        </div>
    </div>

</template>

<script>
    import headTop from '../components/headTop'
    import { getTradeList } from '@/api/getData'

    export default{
        data(){
            return {
                count: 0,
                currentPage: 1,
                currentPageSize: 20,
                tableData:[],
                selectTable: {},
                expendRow: []
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
                try{
                    this.count = 100000
                    this.getTradeListData()
                }catch(err){
                    console.log('获取数据失败', err);
                }
            },
            async getTradeListData(){
                tradeList = getTradeList(null, {
                        page: this.currentPage,
                        pageSize: this.currentPageSize
                    })

                this.tableData = [];
                tradeList.forEach(item => {
                    const tradeItem = {
                        goodsName: item.goodsName,
                        goodsCode: item.goodsCode,
                        marcketName: item.marcketName,
                        customerNo: item.customerNo,
                        tradePrice: item.tradePrice,
                        tradeTime: item.tradeTime
                    }
                    this.tableData.push(tradeItem)
                })
            },
            handleSizeChange(val) {
                console.log(`每页 ${val} 条`);
                this.currentPageSize = val
            },
            handleCurrentChange(val) {
                this.currentPage = val;
                this.offset = (val - 1)*this.limit;
                this.getTradeListData()
            },
            expand(row, status){
            	if (status) {
	                this.tableData.splice(row.index, 1, {...row, ...{}});
                    this.$nextTick(() => {
                        this.expendRow.push(row.index);
                    })
	            }else{
                    const index = this.expendRow.indexOf(row.index);
                    this.expendRow.splice(index, 1)
                }
            },
        }
    } 

</script>

<style lang="less">
    @import '../style/mixin';

    .table_container{
        padding: 20px;
    }
</style>