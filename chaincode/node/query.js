'use strict';
/*
* Copyright IBM Corp All Rights Reserved
*
* SPDX-License-Identifier: Apache-2.0
*/
/*
 * Chaincode query
 */

process.on('unhandledRejection', error => {
});

exports.queryAsync = function(fabric_client, channel, userId, chaincodeId, fcn, args){
	return Promise.resolve((()=>{
		return fabric_client.getUserContext(userId, true);
	})()).then((userFromStore)=>{
		if (userFromStore && userFromStore.isEnrolled()) {
			console.log('Successfully loaded %s from persistence', userId);
		} else {
			throw new Error('Failed to get %s .... run registerUser.js', userId);
		}
	
		var request = {
			chaincodeId: chaincodeId,
			fcn: fcn,
			args: args
		}

		return channel.queryByChaincode(request);
	}).then((query_responses) => {
		console.log("Query has completed, checking results");
		// query_responses could have more than one  results if there multiple peers were used as targets
		if (query_responses && query_responses.length == 1) {
			if (query_responses[0] instanceof Error) {
				console.error("error from query = ", query_responses[0]);
			} else {
				return query_responses[0].toString()
			}
		} else {
			console.log("No payloads were returned from query");
		}
	}).catch((err) => {
		console.error('Failed to query successfully :: ' + err);
	});

}