<template>
    <div class="fillContain">
        <head-top>商品列表</head-top>
        <div class="table_container">
            <el-table
                :data="tableData"
                @expand="expand"
                :expand-row-keys="expendRow"
                :row-key="row=>row.index"
                style="width:100%">
                <!-- <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" inline>
                            <el-form-item label="商品名称">
                                <span>{{ props.row.name }}</span>
                            </el-form-item>
                            <el-form-item label="商品码">
                                <span>{{ props.row.code }}</span>
                            </el-form-item>
                            <el-form-item label="商品类型">
                                <span>{{ props.row.type }}</span>
                            </el-form-item>
                            <el-form-item label="商品图片">
                                <span>{{ props.row.smallUrl }}</span>
                            </el-form-item>
                            <el-form-item label="商品价格">
                                <span>{{ props.row.price }}</span>
                            </el-form-item>
                            <el-form-item label="所属集市">
                                <span>{{ props.row.marcketName }}</span>
                            </el-form-item>
                            <el-form-item label="库存剩余">
                                <span>{{ props.row.inventory }}</span>
                            </el-form-item>
                            <el-form-item label="商品标签">
                                <span>{{ props.row.tag }}</span>
                            </el-form-item>
                            <el-form-item label="评星">
                                <span>{{ props.row.star }}</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column> -->
                
                <el-table-column type="index" width="100"></el-table-column>
                <el-table-column label="商品名称" prop="name"></el-table-column>
                <el-table-column label="商品码" prop="code"></el-table-column>
                <el-table-column label="商品类型" prop="type"></el-table-column>
                <el-table-column label="商品图片" prop="smallUrl"></el-table-column>
                <el-table-column label="商品价格(单位元)" prop="price"></el-table-column>
                <el-table-column label="所属集市" prop="marcketName"></el-table-column>
                <el-table-column label="库存剩余" prop="inventory"></el-table-column>
                <el-table-column label="商品标签" prop="tag"></el-table-column>
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
    import { getGoodsList } from '@/api/getData'

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
                    this.getGoods()
                }catch(err){
                    console.log('获取数据失败', err);
                }
            },
            async getGoods(){
                try{
                    var goodsList = await getGoodsList(null, {
                        page: this.currentPage,
                        pageSize: this.currentPageSize
                    })
                        
                    if(!goodsList){
                        return
                    }

                    this.tableData = [];
                    goodsList.forEach(item => {
                        const goodsItem = {
                            name: item.name,
                            code: item.code,
                            type: item.type,
                            smallUrl: item.smallUrl,
                            price: item.price,
                            marcketName: item.marcketName,
                            inventory: item.inventory,
                            tag: item.tag,
                            star: item.star
                        }
                        this.tableData.push(goodsItem)
                    })

                    this.count = goodsList.length
                }catch(err){
                    console.log('获取数据失败', err);
                }
            },
            handleSizeChange(val) {
                console.log(`每页 ${val} 条`);
                this.currentPageSize = val
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