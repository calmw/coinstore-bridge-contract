import {
    ProposeCreate as ProposeCreateLog,
    VotePropose as VoteProposeLog,
} from "../generated/Vote/Vote"
import {ProposeCreate,VotePropose} from "../generated/schema"
import {Address, BigInt, Bytes} from "@graphprotocol/graph-ts";
import {timestampToDatetime} from "./utils/constants";

export function handleProposeCreate(event: ProposeCreateLog): void {
    let obj = new ProposeCreate(createEventID(event.block.number, event.logIndex))
    obj.chainId = event.params.chainId
    obj.proposeId = event.params.proposeId
    obj.user = event.params.user
    obj.fee = event.params.fee
    obj.ctime = event.params.ctime
    obj.txHash = event.transaction.hash
    obj.timestamp = event.block.timestamp
    obj.utc_time = timestampToDatetime(event.block.timestamp.toI64())
    obj.save()
}

export function handleVotePropose(event: VoteProposeLog): void {
    let obj = new VotePropose(createEventID(event.block.number, event.logIndex))
    obj.chainId = event.params.chainId
    obj.proposeId = event.params.proposeId
    obj.user = event.params.user
    obj.fee = event.params.fee
    obj.isSupport = event.params.isSupport
    obj.txHash = event.transaction.hash
    obj.timestamp = event.block.timestamp
    obj.utc_time = timestampToDatetime(event.block.timestamp.toI64())
    obj.save()
}
function createEventID(blockNumber: BigInt, logIndex: BigInt): string {
    return blockNumber.toString().concat('-').concat(logIndex.toString())
}

function createResolverID(node: Bytes, resolver: Address): string {
    return resolver.toHexString().concat('-').concat(node.toHexString())
}

