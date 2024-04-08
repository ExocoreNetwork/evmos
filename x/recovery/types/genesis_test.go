package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenesisValidate(t *testing.T) {
	testCases := []struct {
		name     string
		genesis  GenesisState
		expError bool
	}{
		{
			"empty genesis",
			GenesisState{},
			false,
		},
		{
			"default genesis",
			*DefaultGenesisState(),
			false,
		},
		{
			"custom genesis",
			NewGenesisState(NewParams(true, time.Hour, []string(nil), []string(nil))),
			false,
		},
		{
			"custom genesis with channels",
			NewGenesisState(NewParams(true, time.Hour, []string{"channel-0"}, []string{"channel-1"})),
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.genesis.Validate()
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
		}
	}
}
