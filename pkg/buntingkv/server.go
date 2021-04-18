package buntingkv

import (
	"context"

	"bunting.io/bunting/api/buntingkvpb"
)

type ServerConfig struct {
}

type Collection interface {
	Query(*buntingkvpb.KVQueryRequest)
}

type CollectionRouter interface {
	GetCollection([]byte) Collection
}

type TransactionQueue interface {
	Submit(*buntingkvpb.KVTxnRequest)
}

type BuntingKVServer struct {
	collections      Collection
	observers        Collection
	hosts            Collection
	collectionRouter CollectionRouter
	transactions     TransactionQueue
}

// All "internal collections":
// - collections
// - collection observers
// - hosts
// API server manages state in special shards.
// It requires a sync() operation to ensure linearizable reads.
func (s *BuntingKVServer) CreateCollection(ctx context.Context, req *buntingkvpb.CreateCollectionRequest) (*buntingkvpb.CreateCollectionResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) UpdateCollection(ctx context.Context, req *buntingkvpb.UpdateCollectionRequest) (*buntingkvpb.UpdateCollectionResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) DeleteCollection(ctx context.Context, req *buntingkvpb.DeleteCollectionRequest) (*buntingkvpb.DeleteCollectionResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) GetCollection(ctx context.Context, req *buntingkvpb.GetCollectionRequest) (*buntingkvpb.GetCollectionResponse, error) {
	s.collections.Query(nil)
	return nil, nil
}

func (s *BuntingKVServer) ListCollections(ctx context.Context, req *buntingkvpb.ListCollectionsRequest) (*buntingkvpb.ListCollectionsResponse, error) {
	s.collections.Query(nil)
	return nil, nil
}

func (s *BuntingKVServer) CreateCollectionObserver(ctx context.Context, req *buntingkvpb.CreateCollectionObserverRequest) (*buntingkvpb.CreateCollectionObserverResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) UpdateCollectionObserver(ctx context.Context, req *buntingkvpb.UpdateCollectionObserverRequest) (*buntingkvpb.UpdateCollectionObserverResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) DeleteCollectionObserver(ctx context.Context, req *buntingkvpb.DeleteCollectionObserverRequest) (*buntingkvpb.DeleteCollectionObserverResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) GetCollectionObserver(ctx context.Context, req *buntingkvpb.GetCollectionObserverRequest) (*buntingkvpb.GetCollectionObserverResponse, error) {
	s.collections.Query(nil)
	return nil, nil
}

func (s *BuntingKVServer) ListCollectionObservers(ctx context.Context, req *buntingkvpb.ListCollectionObserversRequest) (*buntingkvpb.ListCollectionObserversResponse, error) {
	s.collections.Query(nil)
	return nil, nil
}

func (s *BuntingKVServer) CreateHost(context.Context, *buntingkvpb.CreateHostRequest) (*buntingkvpb.CreateHostResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) UpdateHost(context.Context, *buntingkvpb.UpdateHostRequest) (*buntingkvpb.UpdateHostResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) DeleteHost(context.Context, *buntingkvpb.DeleteHostRequest) (*buntingkvpb.DeleteHostResponse, error) {
	s.transactions.Submit(nil)
	return nil, nil
}

func (s *BuntingKVServer) GetHost(context.Context, *buntingkvpb.GetHostRequest) (*buntingkvpb.GetHostResponse, error) {
	s.collections.Query(nil)
	return nil, nil
}

func (s *BuntingKVServer) ListHosts(context.Context, *buntingkvpb.ListHostsRequest) (*buntingkvpb.ListHostsResponse, error) {
	s.collections.Query(nil)
	return nil, nil
}

func (s *BuntingKVServer) KVTxn(ctx context.Context, req *buntingkvpb.KVTxnRequest) (*buntingkvpb.KVTxnResponse, error) {
	s.transactions.Submit(req)
	return nil, nil
}

func (s *BuntingKVServer) KVQuery(ctx context.Context, req *buntingkvpb.KVQueryRequest) (*buntingkvpb.KVQueryResponse, error) {
	s.collectionRouter.GetCollection(req.Collection).Query(req)
	return nil, nil
}
