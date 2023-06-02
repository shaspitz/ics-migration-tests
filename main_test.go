package main

import (
	"testing"

	consumertypes "github.com/cosmos/interchain-security/x/ccv/consumer/types"
	"github.com/stretchr/testify/require"
)

func TestICSKeys(t *testing.T) {
	require.NotEmpty(t, consumertypes.StoreKey)
}
