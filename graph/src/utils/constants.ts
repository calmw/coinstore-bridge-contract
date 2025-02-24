/* eslint-disable prefer-const */
import {BigDecimal, BigInt} from '@graphprotocol/graph-ts'


export const ADDRESS_ZERO = '0x0000000000000000000000000000000000000000'


export let ZERO_BI = BigInt.fromI32(0)
export let ONE_BI = BigInt.fromI32(1)
export let ZERO_BD = BigDecimal.fromString('0')
export let ONE_BD = BigDecimal.fromString('1')
export let BI_18 = BigInt.fromI32(18)

export function getDayID(timestamp: BigInt): i32 {
    let newTP = timestamp.toI32()
    let dayID = newTP / 86400
    return dayID
}

export function getDetaTimestamp(dayID: i32): i32 {
    let dayStartTimestamp = dayID * 86400
    return dayStartTimestamp
}

export function getEndDay(dayID: i32, days: i32): i32 {
    let dayStartTimestamp = getDetaTimestamp(dayID)

    let endTS = dayStartTimestamp + days * 86400
    return endTS


}

export function timestampToDatetime(time: i64): string {
    let date = new Date(time * 1000);

    let YY = date.getUTCFullYear().toString();
    let MM = date.getUTCMonth() + 1 < 10 ? "0" + (date.getUTCMonth() + 1).toString() : (date.getUTCMonth() + 1).toString();
    let DD = date.getUTCDate() < 10 ? "0" + date.getUTCDate().toString() : date.getUTCDate().toString();
    let hh = date.getUTCHours() < 10 ? "0" + date.getUTCHours().toString() : date.getUTCHours().toString();
    let mm = date.getUTCMinutes() < 10 ? "0" + date.getUTCMinutes().toString() : date.getUTCMinutes().toString();
    let ss = date.getUTCSeconds() < 10 ? "0" + date.getUTCSeconds().toString() : date.getUTCSeconds().toString();

    // 这里修改返回时间的格式
    return YY + "-" + MM + "-" + DD + " " + hh + ":" + mm + ":" + ss;
}