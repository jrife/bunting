package buntingkv

// The scheduler should decide which shard-hosts should host
// which shards. We need a mechanism for assigning and loading
// shards onto shard hosts.
//
// Log Split - 3 phase split
// What if we want to move processing of a partition in a consumer group from processor-0 to processor-1?
// Easy, switchover happens quickly with new assignment.
// What if we want to move serving of a shard from one shard host to another?
// Mark shard as unavailable.

type SchedulerConfig struct {
	APIServerAddr string
}

func Run() {

}
