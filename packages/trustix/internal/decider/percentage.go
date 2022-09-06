// Copyright (C) 2021 Tweag IO
// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: GPL-3.0-only

package decider

import (
	"fmt"
	"sort"
)

type minimumPercent struct {
	minimumPct int
}

func NewMinimumPercentDecider(minimumPct int) (LogDecider, error) {
	return &minimumPercent{
		minimumPct: minimumPct,
	}, nil
}

func (q *minimumPercent) Name() string {
	return "percentage"
}

func (q *minimumPercent) Decide(inputs []*DeciderInput) (*DeciderOutput, error) {
	numInputs := len(inputs)
	pctPerEntry := 100 / numInputs

	// Map Value to list of matches
	entries := make(map[string][]*DeciderInput)
	for i := range inputs {
		input := inputs[i]
		l := entries[input.Value]
		l = append(l, input)
		entries[input.Value] = l
	}

	type sortStruct struct {
		key string
		pct int
	}

	makeReturn := func(m *sortStruct) (*DeciderOutput, error) {
		ret := &DeciderOutput{
			Value:      m.key,
			Confidence: m.pct,
		}
		return ret, nil
	}

	// Filter out any keys with less than minimum percentage match and put in list
	var matchesMinimum []*sortStruct
	for k, v := range entries {
		pct := len(v) * pctPerEntry
		if pct >= q.minimumPct {
			m := &sortStruct{
				key: k,
				pct: pct,
			}
			matchesMinimum = append(matchesMinimum, m)
			if pct >= 50 || len(inputs) == 1 || len(entries) == 1 || len(inputs) == len(v) {
				return makeReturn(m)
			}
		}
	}

	switch len(matchesMinimum) {
	case 0:
		return nil, fmt.Errorf("Could not reach the minimum %d quorum", q.minimumPct)
	case 1:
		return makeReturn(matchesMinimum[0])
	default:
		// Sort list by the highest match percentage
		sort.SliceStable(matchesMinimum, func(i, j int) bool {
			return matchesMinimum[i].pct > matchesMinimum[j].pct
		})

		return makeReturn(matchesMinimum[0])
	}
}
