import { chainCodeId } from '../config/env'
import { invokeAsync } from '../fabric/invoke.js'
import { queryAsync } from '../fabric/query.js'

export const getGoodsListOnChain = (userId, startId, page) => {
    goodsList = async ()=>{
        return await queryAsync(userId, chainCodeId, "query", ["goods"])
    }
    return goodsList
}

export const getMarcketListOnChain = (userId, startId, page) => {
    marcketList = async ()=>{
        return await queryAsync(userId, chainCodeId, "query", ["marckets"])
    }
    return marcketList
}

export const getTradeListOnChain = (userId, startId, page) => {
    tradeList = async ()=>{
        return await queryAsync(userId, chainCodeId, "query", ["trades"])
    }
    return tradeList
}

export const initiateTradeOnChain = (userId, tradeInfo) => {
    result = async ()=>{
        return await invokeAsync(userId, chainCodeId, "initiateTrade", [JSON.stringify(tradeInfo)])
    }
    return result
}