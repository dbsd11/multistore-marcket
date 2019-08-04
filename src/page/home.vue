<template>
    <div>
        <head-top></head-top>
		<section class="data_section">
			<header class="section_title">数据统计</header>
			<el-row :gutter="20" style="margin-bottom: 10px;">
                <el-col :span="4"><div class="data_list today_head"><span class="data_num head">当日数据：</span></div></el-col>
				<el-col :span="4"><div class="data_list"><span class="data_num">{{tradeAmount}}</span> 交易量(元)</div></el-col>
				<el-col :span="4"><div class="data_list"><span class="data_num">{{tradeCount}}</span> 交易数</div></el-col>
                <el-col :span="4"><div class="data_list"><span class="data_num">{{tradeCustomerCount}}</span> 客户数</div></el-col>
			</el-row>
            <el-row :gutter="20">
                <el-col :span="4"><div class="data_list all_head"><span class="data_num head">总数据：</span></div></el-col>
                <el-col :span="4"><div class="data_list"><span class="data_num">{{totalTradeAmount}}</span>交易量(元)</div></el-col>
				<el-col :span="4"><div class="data_list"><span class="data_num">{{totalTradeCount}}</span>交易数</div></el-col>
				<el-col :span="4"><div class="data_list"><span class="data_num">{{totalGoodsCount}}</span>商品数量</div></el-col>
                <el-col :span="4"><div class="data_list"><span class="data_num">{{totalMarcketsCount}}</span>交易市场数量</div></el-col>

            </el-row>
		</section>
    </div>
</template>

<script>
	import headTop from '../components/headTop'
	import { getGoodsList, getMarcketList , getTradeList } from '@/api/getData'
    export default {
    	data(){
    		return {
				tradeAmount: 0.00,
				tradeCount: 0,
				tradeCustomerCount: 0,
				totalTradeAmount: 0.00,
				totalTradeCount: 0,
				totalGoodsCount: 0,
				totalMarcketsCount: 0
    		}
    	},
    	components: {
    		headTop,
    	},
    	mounted(){
    		this.initData();
    	},
        computed: {

        },
    	methods: {
    		async initData(){
				try{
					var goodList = await getGoodsList();
					var marcketsList = await getMarcketList();
					var tradeList = await getTradeList();

					this.totalGoodsCount = goodList.length
					this.totalMarcketsCount = marcketsList.length
					this.totalTradeCount = tradeList.length

					this.totalTradeAmount = 0.00
					tradeList.forEach(trade => {
						this.totalTradeAmount += parseFloat(trade["tradePrice"])
					});

					this.tradeAmount = 0.00
					this.tradeCount = 0
					this.tradeCustomerCount = 0
					var customerNos = []
					tradeList.filter(trade => trade["tradeTime"].startsWith(this.getCurDateStr('yyyy-MM-dd'))).forEach(trade => {
						this.tradeAmount += parseFloat(trade["tradePrice"])
						this.tradeCount += 1
						if(!(trade["customerNo"] in customerNos)){
							this.tradeCustomerCount += 1
							customerNos.push(trade["customerNo"])
						}
					})
				}catch(err){
					console.info("获取统计数据失败:"+err)
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
    	}
    }
</script>

<style lang="less">
	@import '../style/mixin';
	.data_section{
		padding: 20px;
		margin-bottom: 40px;
		.section_title{
			text-align: center;
			font-size: 30px;
			margin-bottom: 10px;
		}
		.data_list{
			text-align: center;
			font-size: 14px;
			color: #666;
            border-radius: 6px;
            background: #E5E9F2;
            .data_num{
                color: #333;
                font-size: 26px;

            }
            .head{
                border-radius: 6px;
                font-size: 22px;
                padding: 4px 0;
                color: #fff;
                display: inline-block;
            }
        }
        .today_head{
            background: #FF9800;
        }
        .all_head{
            background: #20A0FF;
        }
	}
    .wan{
        .sc(16px, #333)
    }
</style>
