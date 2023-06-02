package main

import (
	"testing"

	consumertypesv2 "github.com/cosmos/interchain-security/v2/x/ccv/consumer/types"
	consumertypesv1 "github.com/cosmos/interchain-security/x/ccv/consumer/types"
	"github.com/stretchr/testify/require"
)

func TestICSConsumerKeyPrefixes(t *testing.T) {

	require.Equal(t, consumertypesv1.PortByteKey, consumertypesv2.PortByteKey)
	require.Equal(t, consumertypesv1.LastDistributionTransmissionByteKey, consumertypesv2.LastDistributionTransmissionByteKey)
	require.Equal(t, consumertypesv1.UnbondingTimeByteKey, consumertypesv2.UnbondingTimeByteKey)
	require.Equal(t, consumertypesv1.ProviderClientByteKey, consumertypesv2.ProviderClientByteKey)
	require.Equal(t, consumertypesv1.ProviderChannelByteKey, consumertypesv2.ProviderChannelByteKey)
	require.Equal(t, consumertypesv1.PendingChangesByteKey, consumertypesv2.PendingChangesByteKey)
	require.Equal(t, consumertypesv1.PendingDataPacketsByteKey, consumertypesv2.PendingDataPacketsByteKey)
	require.Equal(t, consumertypesv1.PreCCVByteKey, consumertypesv2.PreCCVByteKey)
	require.Equal(t, consumertypesv1.InitialValSetByteKey, consumertypesv2.InitialValSetByteKey)
	require.Equal(t, consumertypesv1.LastStandaloneHeightByteKey, consumertypesv2.LastStandaloneHeightByteKey)
	require.Equal(t, consumertypesv1.SmallestNonOptOutPowerByteKey, consumertypesv2.SmallestNonOptOutPowerByteKey)
	require.Equal(t, consumertypesv1.HistoricalInfoBytePrefix, consumertypesv2.HistoricalInfoBytePrefix)
	require.Equal(t, consumertypesv1.PacketMaturityTimeBytePrefix, consumertypesv2.PacketMaturityTimeBytePrefix)
	require.Equal(t, consumertypesv1.HeightValsetUpdateIDBytePrefix, consumertypesv2.HeightValsetUpdateIDBytePrefix)
	require.Equal(t, consumertypesv1.OutstandingDowntimeBytePrefix, consumertypesv2.OutstandingDowntimeBytePrefix)
	require.Equal(t, consumertypesv1.PendingDataPacketsBytePrefix, consumertypesv2.PendingDataPacketsBytePrefix)
	require.Equal(t, consumertypesv1.CrossChainValidatorBytePrefix, consumertypesv2.CrossChainValidatorBytePrefix)

	// three new keys in v2, need to be higher enum value than highest existing prefix
	highestV1prefix := consumertypesv1.CrossChainValidatorBytePrefix
	require.Greater(t, consumertypesv2.InitGenesisHeightByteKey, highestV1prefix)
	require.Greater(t, consumertypesv2.StandaloneTransferChannelIDByteKey, consumertypesv2.InitGenesisHeightByteKey)
	require.Greater(t, consumertypesv2.PrevStandaloneChainByteKey, consumertypesv2.StandaloneTransferChannelIDByteKey)
}
