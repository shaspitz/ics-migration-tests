package main

import (
	"testing"

	providertypesv2 "github.com/cosmos/interchain-security/v2/x/ccv/provider/types"
	providertypesv1 "github.com/cosmos/interchain-security/x/ccv/provider/types"
	"github.com/stretchr/testify/require"
)

func TestICSProviderKeyPrefixes(t *testing.T) {

	require.Equal(t, providertypesv1.PortByteKey, providertypesv2.PortByteKey)
	require.Equal(t, providertypesv1.MaturedUnbondingOpsByteKey, providertypesv2.MaturedUnbondingOpsByteKey)
	require.Equal(t, providertypesv1.ValidatorSetUpdateIdByteKey, providertypesv2.ValidatorSetUpdateIdByteKey)
	require.Equal(t, providertypesv1.SlashMeterByteKey, providertypesv2.SlashMeterByteKey)
	require.Equal(t, providertypesv1.SlashMeterReplenishTimeCandidateByteKey, providertypesv2.SlashMeterReplenishTimeCandidateByteKey)
	require.Equal(t, providertypesv1.ChainToChannelBytePrefix, providertypesv2.ChainToChannelBytePrefix)
	require.Equal(t, providertypesv1.ChannelToChainBytePrefix, providertypesv2.ChannelToChainBytePrefix)
	require.Equal(t, providertypesv1.ChainToClientBytePrefix, providertypesv2.ChainToClientBytePrefix)
	require.Equal(t, providertypesv1.InitTimeoutTimestampBytePrefix, providertypesv2.InitTimeoutTimestampBytePrefix)
	require.Equal(t, providertypesv1.PendingCAPBytePrefix, providertypesv2.PendingCAPBytePrefix)
	require.Equal(t, providertypesv1.PendingCRPBytePrefix, providertypesv2.PendingCRPBytePrefix)
	require.Equal(t, providertypesv1.UnbondingOpBytePrefix, providertypesv2.UnbondingOpBytePrefix)
	require.Equal(t, providertypesv1.UnbondingOpIndexBytePrefix, providertypesv2.UnbondingOpIndexBytePrefix)
	require.Equal(t, providertypesv1.ValsetUpdateBlockHeightBytePrefix, providertypesv2.ValsetUpdateBlockHeightBytePrefix)
	require.Equal(t, providertypesv1.ConsumerGenesisBytePrefix, providertypesv2.ConsumerGenesisBytePrefix)
	require.Equal(t, providertypesv1.SlashAcksBytePrefix, providertypesv2.SlashAcksBytePrefix)
	require.Equal(t, providertypesv1.InitChainHeightBytePrefix, providertypesv2.InitChainHeightBytePrefix)
	require.Equal(t, providertypesv1.PendingVSCsBytePrefix, providertypesv2.PendingVSCsBytePrefix)
	require.Equal(t, providertypesv1.VscSendTimestampBytePrefix, providertypesv2.VscSendTimestampBytePrefix)
	require.Equal(t, providertypesv1.ThrottledPacketDataSizeBytePrefix, providertypesv2.ThrottledPacketDataSizeBytePrefix)
	require.Equal(t, providertypesv1.ThrottledPacketDataBytePrefix, providertypesv2.ThrottledPacketDataBytePrefix)
	require.Equal(t, providertypesv1.GlobalSlashEntryBytePrefix, providertypesv2.GlobalSlashEntryBytePrefix)
	require.Equal(t, providertypesv1.ConsumerValidatorsBytePrefix, providertypesv2.ConsumerValidatorsBytePrefix)
	require.Equal(t, providertypesv1.ValidatorsByConsumerAddrBytePrefix, providertypesv2.ValidatorsByConsumerAddrBytePrefix)
	require.Equal(t, providertypesv1.KeyAssignmentReplacementsBytePrefix, providertypesv2.KeyAssignmentReplacementsBytePrefix)
	require.Equal(t, providertypesv1.ConsumerAddrsToPruneBytePrefix, providertypesv2.ConsumerAddrsToPruneBytePrefix)
	require.Equal(t, providertypesv1.SlashLogBytePrefix, providertypesv2.SlashLogBytePrefix)
	require.Equal(t, providertypesv1.ConsumerRewardDenomsBytePrefix, providertypesv2.ConsumerRewardDenomsBytePrefix)

	// no new key prefixes in v2 for provider
}
