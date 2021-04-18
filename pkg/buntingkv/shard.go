package buntingkv

import (
	"context"

	"bunting.io/bunting/api/buntingkvpb"
)

type CommandIterator interface {
	Next() buntingkvpb.KVRequestOp
}

type ShardImpl struct {
	Commands  CommandIterator
	KeyValues <-chan buntingkvpb.KeyValue
	KVStore   KVStore
	Table     Table
}

func (s *ShardImpl) Run(ctx context.Context) error {
	go func() {
		for {
			select {
			case kv := <-s.KeyValues:
				s.syncKV(kv)
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		r := s.Commands.Next()

		switch r.Request.(type) {
		case *buntingkvpb.KVRequestOp_Create:
			s.handleCreate(r.GetCreate())
		case *buntingkvpb.KVRequestOp_Update:
			s.handleUpdate(r.GetUpdate())
		case *buntingkvpb.KVRequestOp_Patch:
			s.handlePatch(r.GetPatch())
		case *buntingkvpb.KVRequestOp_Batch:
			s.handleBatch(r.GetBatch())
		case *buntingkvpb.KVRequestOp_Delete:
			s.handleDelete(r.GetDelete())
		case *buntingkvpb.KVRequestOp_Query:
			s.handleQuery(r.GetQuery())
		}
	}
}

func (s *ShardImpl) syncKV(kv buntingkvpb.KeyValue) {
	// put or delete from table. fence with version
	s.KVStore.Put(kv.Key.Key, kv.Value) // iff curr_index < kv.Index
	s.KVStore.Delete(kv.Key.Key)        // iff curr_index < kv.Index
}

func (s *ShardImpl) handleCreate(r *buntingkvpb.KVCreateRequest) {
}

func (s *ShardImpl) handleUpdate(r *buntingkvpb.KVUpdateRequest) {
}

func (s *ShardImpl) handlePatch(r *buntingkvpb.KVPatchRequest) {
}

func (s *ShardImpl) handleBatch(r *buntingkvpb.KVBatchRequest) {
}

func (s *ShardImpl) handleDelete(r *buntingkvpb.KVDeleteRequest) {
}

func (s *ShardImpl) handleQuery(r *buntingkvpb.KVQueryRequest) {
}
