module.exports = (env) => {
    var express = require('express')
    var api = express.Router();

    var invokeAsync = require('./invoke').invokeAsync
    var queryAsync = require('./query').queryAsync

    chainCodeId = env.chainCodeId
    channelName = env.channelName
    peerGrpcUrl = env.peerGrpcUrl
    peerGrpcUrl2 = env.peerGrpcUrl2
    orderrGrpcUrl = env.orderrGrpcUrl

    var Fabric_Client = require('fabric-client');
    var path = require('path');
    var util = require('util');
    var os = require('os');

    //
    var fabric_client = new Fabric_Client();

    // setup the fabric network
    var channel = fabric_client.newChannel(channelName);
    var peer = fabric_client.newPeer(peerGrpcUrl);
    var order = fabric_client.newOrderer(orderrGrpcUrl);
    channel.addPeer(peer);
    channel.addOrderer(order);

    var store_path = path.join(__dirname, 'hfc-key-store');
    console.log('Store path:'+store_path);
    
    // create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
    Fabric_Client.newDefaultKeyValueStore({ path: store_path
    }).then((state_store) => {
        // assign the store to the fabric client
        fabric_client.setStateStore(state_store);
        var crypto_suite = Fabric_Client.newCryptoSuite();
        // use the same location for the state store (where the users' certificate are kept)
        // and the crypto store (where the users' keys are kept)
        var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
        crypto_suite.setCryptoKeyStore(crypto_store);
        fabric_client.setCryptoSuite(crypto_suite);
    })

	// perhaps expose some API metadata at the root
	api.get('/getGoodsListOnChain', (req, res) => {
        queryAsync(fabric_client, channel, "admin@bison.group", chainCodeId, "query", ["goods"]).then((result)=> {
            res.json(JSON.parse(result));
        })
    })
    api.get('/getMarchetListOnChain', (req, res) => {
        queryAsync(fabric_client, channel, "admin@bison.group", chainCodeId, "query", ["marckets"]).then((result)=> {
            res.json(JSON.parse(result));
        })
    })
    api.get('/getTradeListOnChain', (req, res) => {
        queryAsync(fabric_client, channel, "admin@bison.group", chainCodeId, "query", ["trades"]).then((result)=> {
            res.json(JSON.parse(result));
        })
    })
    api.post('/initiateTradeOnChain', (req, res) => {
        invokeAsync(fabric_client, channel, "admin@bison.group", chainCodeId, "initiateTrade", [new Date().getTime().toString(), JSON.stringify(req.body)]).then((result)=> {
            console.info("bbbb")
            res.json(JSON.parse(result));
        })

        console.info("aaaa")
    })

	return api;
}