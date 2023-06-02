package main

import (
	"testing"

	consumertypesv2 "github.com/cosmos/interchain-security/v2/x/ccv/consumer/types"
	consumertypesv1 "github.com/cosmos/interchain-security/x/ccv/consumer/types"
	"github.com/stretchr/testify/require"
)

func TestICSKeys(t *testing.T) {
	require.NotEmpty(t, consumertypesv1.StoreKey)
	require.NotEmpty(t, consumertypesv2.RouterKey)
}
