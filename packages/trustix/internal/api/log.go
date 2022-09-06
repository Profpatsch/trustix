// Copyright (C) 2021 Tweag IO
// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: GPL-3.0-only

package api

// Common log operations

import (
	"context"
	"fmt"
	"hash"

	"github.com/nix-community/trustix/packages/trustix-proto/api"
	"github.com/nix-community/trustix/packages/trustix-proto/schema"
	vlog "github.com/nix-community/trustix/packages/trustix/internal/log"
	"github.com/nix-community/trustix/packages/trustix/internal/storage"
)

func getLogConsistencyProof(hashFn func() hash.Hash, txn *storage.BucketTransaction, ctx context.Context, req *api.GetLogConsistencyProofRequest) (resp *api.ProofResponse, err error) {
	resp = &api.ProofResponse{}

	vLog, err := vlog.NewVerifiableLog(hashFn, txn, *req.SecondSize)
	if err != nil {
		return nil, err
	}

	proof, err := vLog.ConsistencyProof(*req.FirstSize, *req.SecondSize)
	if err != nil {
		return nil, err
	}

	resp.Proof = proof

	return resp, nil
}

func getLogAuditProof(hashFn func() hash.Hash, txn *storage.BucketTransaction, ctx context.Context, req *api.GetLogAuditProofRequest) (resp *api.ProofResponse, err error) {

	vLog, err := vlog.NewVerifiableLog(hashFn, txn, *req.TreeSize)
	if err != nil {
		return nil, err
	}

	proof, err := vLog.AuditProof(*req.Index, *req.TreeSize)
	if err != nil {
		return nil, err
	}

	resp.Proof = proof

	return resp, nil
}

func getLogEntries(txn *storage.BucketTransaction, ctx context.Context, req *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error) {

	resp := &api.LogEntriesResponse{
		Leaves: []*schema.LogLeaf{},
	}

	logStorage := vlog.NewLogStorage(txn)

	for i := int(*req.Start); i <= int(*req.Finish); i++ {
		leaf, err := logStorage.Get(0, uint64(i))
		if err != nil {
			return nil, err
		}
		resp.Leaves = append(resp.Leaves, leaf)
	}

	total := *req.Finish - *req.Start + 1
	if uint64(len(resp.Leaves)) != total {
		return nil, fmt.Errorf("%d != %d", len(resp.Leaves), total)
	}

	return resp, nil
}
