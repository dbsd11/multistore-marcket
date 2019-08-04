<template>
	<div class="manage_page fillcontain">
		<el-row style="height: 100%;">
	  		<el-col :span="4"  style="min-height: 100%; background-color: #324057;">
				<el-menu :default-active="defaultActive" style="min-height: 100%;" theme="dark" router>
					<el-menu-item index="manage"><i class="el-icon-menu"></i>首页</el-menu-item>
					<el-menu-item index="initiateTrade"><i class="el-icon-plus"></i>发起交易</el-menu-item>
					<el-submenu index="2">
						<template slot="title"><i class="el-icon-star-on"></i>商品管理</template>
						<el-menu-item index="goodList">商品列表</el-menu-item>
						<el-menu-item index="marcketList">集市列表</el-menu-item>
					</el-submenu>
					<el-submenu index="3">
						<template slot="title"><i class="el-icon-document"></i>交易</template>
						<el-menu-item index="tradeList">交易记录</el-menu-item>
					</el-submenu>
					<el-menu-item index="explain"><i class="el-icon-warning"></i>说明</el-menu-item>
				</el-menu>
			</el-col>
			<el-col :span="20" style="height: 100%;overflow: auto;">
				<keep-alive>
				    <router-view></router-view>
				</keep-alive>
			</el-col>
		</el-row>
  	</div>
</template>

<script>
    export default {
		computed: {
			defaultActive: function(){
				return this.$route.path.replace('/', '');
			}
		},
    }
</script>


<style lang="less" scoped>
	@import '../style/mixin';
	.manage_page{
		
	}
</style>
