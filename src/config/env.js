/**
 * 配置编译环境和线上环境之间的切换
 * 
 * baseUrl: 域名地址
 * routerMode: 路由模式
 * baseImgPath: 图片存放地址
 * 
 */
let baseUrl = ''; 
let routerMode = 'hash';
let baseImgPath;
let channelName = 'mychannel';
let peerGrpcUrl = 'grpc://119.3.167.139:7051';
let peerGrpcUrl2 = 'grpc://119.3.167.139:7053';
let orderrGrpcUrl = 'grpc://119.3.167.139:7050';
let caUrl = 'http://119.3.167.139:7054';
let chaincodeId = "multistore-marcket:2.8"

if (process.env.NODE_ENV == 'development') {
	baseUrl = '';
    baseImgPath = '/img/';
}else{
	baseUrl = '//elm.cangdu.org';
    baseImgPath = '//elm.cangdu.org/img/';
}

export {
	baseUrl,
	routerMode,
	baseImgPath
}