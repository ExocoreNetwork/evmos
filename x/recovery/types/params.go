// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package types

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	claimstypes "github.com/evmos/evmos/v14/x/claims/types"
)

// ParamsKey params store key
var ParamsKey = []byte("Params")

// DefaultPacketTimeoutDuration defines the default packet timeout for outgoing
// IBC transfers
var (
	DefaultEnableRecovery        = true
	DefaultPacketTimeoutDuration = 4 * time.Hour
	DefaultAuthorizedChannels    = claimstypes.DefaultAuthorizedChannels
	DefaultEVMChannels           = claimstypes.DefaultEVMChannels
)

// NewParams creates a new Params instance
func NewParams(
	enableRecovery bool, timeoutDuration time.Duration,
	authorizedChannels []string, evmChannels []string,
) Params {
	return Params{
		EnableRecovery:        enableRecovery,
		PacketTimeoutDuration: timeoutDuration,
		AuthorizedChannels:    authorizedChannels,
		EVMChannels:           evmChannels,
	}
}

// DefaultParams defines the default params for the recovery module
func DefaultParams() Params {
	return Params{
		EnableRecovery:        DefaultEnableRecovery,
		PacketTimeoutDuration: DefaultPacketTimeoutDuration,
		AuthorizedChannels:    DefaultAuthorizedChannels,
		EVMChannels:           DefaultEVMChannels,
	}
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateDuration(i interface{}) error {
	duration, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if duration < 0 {
		return fmt.Errorf("packet timout duration cannot be negative")
	}

	return nil
}

// ValidateChannels checks if channels ids are valid
func ValidateChannels(i interface{}) error {
	channels, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, channel := range channels {
		if err := host.ChannelIdentifierValidator(channel); err != nil {
			return errorsmod.Wrap(
				channeltypes.ErrInvalidChannelIdentifier, err.Error(),
			)
		}
	}

	return nil
}

// Validate checks that the fields have valid values
func (p Params) Validate() error {
	if err := validateDuration(p.PacketTimeoutDuration); err != nil {
		return err
	}
	if err := ValidateChannels(p.AuthorizedChannels); err != nil {
		return err
	}
	if err := ValidateChannels(p.EVMChannels); err != nil {
		return err
	}
	return validateBool(p.EnableRecovery)
}

// IsAuthorizedChannel returns true if the channel provided is in the list of
// authorized channels
func (p Params) IsAuthorizedChannel(channel string) bool {
	for _, authorizedChannel := range p.AuthorizedChannels {
		if channel == authorizedChannel {
			return true
		}
	}

	return false
}

// IsEVMChannel returns true if the channel provided is in the list of
// EVM channels
func (p Params) IsEVMChannel(channel string) bool {
	for _, evmChannel := range p.EVMChannels {
		if channel == evmChannel {
			return true
		}
	}

	return false
}
