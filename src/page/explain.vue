<template>
    <div>
		<head-top></head-top>
		<div class="comments-header">
				<div class="empty"></div>
				<label>商品市场价格小工具, 仅作为练习使用, 欢迎给予意见反馈</label>
				<a href="https://github.com/dbsd11/multistore-marcket" target="_blank">github源码地址<i class="el-icon-view el-icon--right"></i></a>
		</div>
		<div class="comments-main">
			<el-card class="box-card">
				<div slot="header" class="clearfix">
				  <span>意见反馈</span>
				  <el-button style="float: right; padding: 3px 0" type="text" @click="noticeContactEmail()">确认</el-button>
				</div>
				<div class="text item is-required">
					<el-input
					type="textarea"
					:rows="10"
					placeholder="请输入内容"
					v-model="comments">
				  </el-input>
				</div>
			  </el-card>
		</div>
    </div>
</template>

<script>
	import headTop from '../components/headTop'
	import {mapState} from 'vuex'
	import { sendEmail } from '@/api/getData'

    export default {
    	data(){
    		return {
				createUser: 'bison',
				contactEmail: 'bitdbsd11@163.com',
				comments: ''
    		}
    	},
    	components: {
    		headTop,
		},
		created() {
			
		},

		computed:{
			...mapState(['adminInfo']),
		},

    	methods: {
    		async noticeContactEmail(){
				var emailData = {
					user: this.adminInfo,
					toAddress: this.noticeContactEmail,
					content: JSON.stringify({
						comments: this.comments
					})
				}
				try{
					sendEmail(emailData)

					this.$message({
						type: 'success',
						message: '谢谢反馈，我会努力改进的哈^_^'
		            });
				}catch(err){
                    console.log('发送意见反馈邮件失败', err);
                }
			}
    	}
    }
</script>

<style lang="less">
	@import '../style/mixin';
	.comments-header {
		height: 40px;
		vertical-align: middle;

		label {
			font-size: 20px;
			margin: 5%;
		}

		.empty{
			height: 30px
		}
	}

	.comments-main {
		padding: 5%
	}

	.text {
    	font-size: 14px;
  	}

	.item {
		margin-bottom: 18px;
	}

	.clearfix:before,
	.clearfix:after {
		display: table;
		content: "";
	}
	.clearfix:after {
		clear: both
	}

	.box-card {
		width: 480px;
	}
</style>
