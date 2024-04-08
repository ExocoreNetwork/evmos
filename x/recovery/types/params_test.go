package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	testCases := []struct {
		name     string
		params   Params
		expError bool
	}{
		{
			"empty params",
			Params{},
			false,
		},
		{
			"default params",
			DefaultParams(),
			false,
		},
		{
			"custom params",
			NewParams(true, time.Hour, []string(nil), []string(nil)),
			false,
		},
		{
			"invalid duration",
			NewParams(true, -1, []string(nil), []string(nil)),
			true,
		},
		{
			"invalid authorized channel",
			NewParams(true, time.Hour, []string{""}, []string(nil)),
			true,
		},
		{
			"invalid EVM channel",
			NewParams(true, time.Hour, []string(nil), []string{""}),
			true,
		},
	}

	for _, tc := range testCases {
		err := tc.params.Validate()
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
		}
	}
}

func TestValidate(t *testing.T) {
	require.Error(t, validateBool(""))
	require.NoError(t, validateBool(true))

	require.Error(t, validateDuration(true))
	require.NoError(t, validateDuration(time.Hour))

	require.Error(t, ValidateChannels(true))
	require.NoError(t, ValidateChannels([]string{"channel-0"}))
}
