package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankByWordCount(t *testing.T) {
	var tests = []struct {
		input  string
		wanted []struct {
			Word       string
			Occurences int
		}
	}{
		{
			input: "fox fox fox", wanted: []struct {
				Word       string
				Occurences int
			}{
				{Word: "fox", Occurences: 3},
			},
		},
		{
			input: "A a a a a b b B c", wanted: []struct {
				Word       string
				Occurences int
			}{
				{Word: "a", Occurences: 5},
				{Word: "b", Occurences: 3},
				{Word: "c", Occurences: 1},
			},
		},
	}

	for _, tc := range tests {
		testName := fmt.Sprintf("Testing %+v", tc)
		t.Run(testName, func(t *testing.T) {
			result := RankByWordCount(tc.input)
			assert.Equal(t, len(tc.wanted), len(result))

			for i, wordRank := range result {
				assert.Equal(t, wordRank.Word, tc.wanted[i].Word)
				assert.Equal(t, wordRank.Occurences, tc.wanted[i].Occurences)
			}
		})
	}
}

func TestMergeWordCount(t *testing.T) {
	var tests = []struct {
		a      PairList
		b      PairList
		wanted []struct {
			Word       string
			Occurences int
		}
	}{
		{
			a: PairList{Pair{Word: "a", Occurences: 1}},
			b: PairList{Pair{Word: "a", Occurences: 1}},
			wanted: []struct {
				Word       string
				Occurences int
			}{
				{Word: "a", Occurences: 2},
			},
		},
		{
			a: PairList{Pair{Word: "a", Occurences: 1}, Pair{Word: "b", Occurences: 2}},
			b: PairList{Pair{Word: "b", Occurences: 3}, Pair{Word: "c", Occurences: 1}},
			wanted: []struct {
				Word       string
				Occurences int
			}{
				{Word: "a", Occurences: 1},
				{Word: "b", Occurences: 5},
				{Word: "c", Occurences: 1},
			},
		},
	}

	for _, tc := range tests {
		testName := fmt.Sprintf("Testing %+v", tc)
		t.Run(testName, func(t *testing.T) {
			result := MergeWordCounts(tc.a, tc.b)
			assert.Equal(t, len(tc.wanted), len(result))

			for i, wordRank := range result {
				assert.Equal(t, wordRank.Word, tc.wanted[i].Word)
				assert.Equal(t, wordRank.Occurences, tc.wanted[i].Occurences)
			}
		})
	}
}
