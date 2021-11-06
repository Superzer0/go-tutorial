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
			input: "a a a a a b b b c", wanted: []struct {
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
