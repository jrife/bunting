package buntingkv

import (
	"context"

	"bunting.io/bunting/api/buntingkvpb"
)

type ShardHostConfig struct {
	APIServerAddr         string
	IsClusterMetadataHost bool
}

type Shard interface {
	Query(*buntingkvpb.KVQueryRequest) (buntingkvpb.KVQueryResponse, error)
	Run(context.Context, KVIterator) error
}

type KVIterator interface {
	Next() KVUpdate
}

type KVUpdate interface {
	IsPut()
	IsDelete()
	Key() buntingkvpb.Key
	Value() []byte
	Ack()
	Nack()
}

type ShardHost struct {
	config ShardHostConfig
	shards map[int64]Shard
}

func (sh *ShardHost) KVQuery(req *buntingkvpb.KVQueryRequest) (*buntingkvpb.KVQueryResponse, error) {
	// for _, id := range req.ShardIds {
	// 	//resp, _ := sh.shards[id].Query(req)
	// }

	return nil, nil
}

func (sh *ShardHost) Run(ctx context.Context) {
	if sh.config.IsClusterMetadataHost {
		sh.runClusterMetadataShard()
	}

	for _, shard := range sh.shards {
		var shardUpdateIter KVIterator

		// shardUpdateIter wraps kafka operations
		// consuming partition mapping to this shard.

		shard.Run(ctx, shardUpdateIter)
	}
}

func (sh *ShardHost) runClusterMetadataShard() {
	// Shard should:
	// 1) Read current offset from persistent state.
	// 2) Start consuming from last offset.
}

type KVStore interface {
	Query(q *buntingkvpb.KVQueryRequest) (*buntingkvpb.KVQueryResponse, error)
	Put(key, value []byte) error
	Delete(key []byte) error
}

type Table interface {
	Put(buntingkvpb.KeyValue) error
	Delete(buntingkvpb.Key) error
}

type CommandProcessor struct {
	Transactions <-chan buntingkvpb.KVTxnRequest
	KVStore      KVStore
	Table        Table
}

func (p *CommandProcessor) Run(ctx context.Context) error {
	for {
		select {
		case txn := <-p.Transactions:
			p.handleTxn(txn)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (p *CommandProcessor) handleTxn(txn buntingkvpb.KVTxnRequest) error {
	// fans out based on partitions on each operation.
	return nil
}
