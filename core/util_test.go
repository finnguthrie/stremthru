package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeMagnetHash(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		output string
	}{
		{
			"40 Upper",
			"DB42AF2B6AE098558B98A14CD7874EF64162CAC8",
			"db42af2b6ae098558b98a14cd7874ef64162cac8",
		},
		{
			"40 Lower",
			"db42af2b6ae098558b98a14cd7874ef64162cac8",
			"db42af2b6ae098558b98a14cd7874ef64162cac8",
		},
		{
			"32 Upper",
			"3NBK6K3K4CMFLC4YUFGNPB2O6ZAWFSWI",
			"db42af2b6ae098558b98a14cd7874ef64162cac8",
		},
		{
			"32 Lower",
			"3nbk6k3k4cmflc4yufgnpb2o6zawfswi",
			"db42af2b6ae098558b98a14cd7874ef64162cac8",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, NormalizeMagnetHash(tc.input), tc.output)
		})
	}
}
