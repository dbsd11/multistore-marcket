<template>
    <div class="fillContain">
        <head-top>集市列表</head-top>
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
                            <el-form-item label="集市名称">
                                <span>{{ props.row.name }}</span>
                            </el-form-item>
                            <el-form-item label="集市规模">
                                <span>{{ props.row.code }}</span>
                            </el-form-item>
                            <el-form-item label="集市地址">
                                <span>{{ props.row.address }}</span>
                            </el-form-item>
                            <el-form-item label="每日顾客量">
                                <span>{{ props.row.customerNum }}</span>
                            </el-form-item>
                            <el-form-item label="每日交易量">
                                <span>{{ props.row.tradeAmount }}</span>
                            </el-form-item>
                            <el-form-item label="集市标签">
                                <span>{{ props.row.tag }}</span>
                            </el-form-item>    
                            <el-form-item label="评星">
                                <span>{{ props.row.star }}</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>
                
                <el-table-column type="index" width="100"></el-table-column>
                <el-table-column label="集市名称" prop="name"></el-table-column>
                <el-table-column label="集市规模" prop="code"></el-table-column>
                <el-table-column label="集市地址" prop="address"></el-table-column>
                <el-table-column label="每日顾客量" prop="customerNum"></el-table-column>
                <el-table-column label="每日交易量" prop="tradeAmount"></el-table-column>
                <el-table-column label="集市标签" prop="tag"></el-table-column>
                <el-table-column label="评星" prop="star"></el-table-column>
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

    export default{
        data(){
            return {
                count: 0,
                currentPage: 1,
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
            async initData(){},
            async getGoods(){},
            handleSizeChange(val) {
                console.log(`每页 ${val} 条`);
            },
            handleCurrentChange(val) {
                this.currentPage = val;
                this.offset = (val - 1)*this.limit;
                this.getFoods()
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