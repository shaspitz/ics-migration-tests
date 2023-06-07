package main

import (
	"testing"
	"time"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v4/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	"github.com/cosmos/interchain-security/testutil/crypto"
	ccvtestutilv1 "github.com/cosmos/interchain-security/testutil/keeper"
	ccvtestutilv2 "github.com/cosmos/interchain-security/v2/testutil/keeper"
	consumertypesv2 "github.com/cosmos/interchain-security/v2/x/ccv/consumer/types"
	"github.com/cosmos/interchain-security/x/ccv/consumer/types"
	consumertypesv1 "github.com/cosmos/interchain-security/x/ccv/consumer/types"
	ccvtypesv1 "github.com/cosmos/interchain-security/x/ccv/types"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	chainID                      = "gaia"
	trustingPeriod time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod      time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift  time.Duration = time.Second * 10
)

var (
	height      = clienttypes.NewHeight(0, 4)
	upgradePath = []string{"upgrade", "upgradedIBCState"}
)

func TestConsumerGenesisOnProvider(t *testing.T) {
	keeperParamsv1 := ccvtestutilv1.NewInMemKeeperParams(t)
	providerKeeperv1, ctx, ctrl, _ := ccvtestutilv1.GetProviderKeeperAndCtx(t, keeperParamsv1)
	defer ctrl.Finish()

	// generate validator public key
	cId := crypto.NewCryptoIdentityFromIntSeed(238934)
	pubKey := cId.TMCryptoPubKey()

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})
	valHash := valSet.Hash()
	valUpdates := tmtypes.TM2PB.ValidatorUpdates(valSet)

	cs := ibctmtypes.NewClientState("chainID", ibctmtypes.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath, false, false)
	consensusState := ibctmtypes.NewConsensusState(time.Now(), commitmenttypes.NewMerkleRoot([]byte("apphash")), valHash)

	providerKeeperv1.SetConsumerGenesis(ctx, "chainID", consumertypesv1.GenesisState{
		Params:                      consumertypesv1.DefaultParams(),
		ProviderClientId:            "",
		ProviderChannelId:           "ccvchannel",
		NewChain:                    true,
		ProviderClientState:         cs,
		ProviderConsensusState:      consensusState,
		MaturingPackets:             nil,
		InitialValSet:               valUpdates,
		HeightToValsetUpdateId:      nil,
		OutstandingDowntimeSlashing: nil,
		PendingConsumerPackets:      ccvtypesv1.ConsumerPacketDataList{},
		LastTransmissionBlockHeight: types.LastTransmissionBlockHeight{},
		// No preCCV flag!
	})

	// Transfer store ownership to v2 keeper
	keeperParamsv2 := ccvtestutilv2.InMemKeeperParams{
		Cdc:            keeperParamsv1.Cdc,
		StoreKey:       keeperParamsv1.StoreKey,
		ParamsSubspace: keeperParamsv1.ParamsSubspace,
		Ctx:            keeperParamsv1.Ctx,
	}

	providerKeeperv2, ctx, ctrl, _ := ccvtestutilv2.GetProviderKeeperAndCtx(t, keeperParamsv2)
	defer ctrl.Finish()

	// v2 keeper should be able to read v1 data
	consumerGenesisGot, found := providerKeeperv2.GetConsumerGenesis(ctx, "chainID")
	require.True(t, found)
	require.Equal(t, "ccvchannel", consumerGenesisGot.ProviderChannelId)

	// Default for preCCV should be false
	require.False(t, consumerGenesisGot.PreCCV)
}

func TestSetPendingPackets(t *testing.T) {

	keeperParamsv1 := ccvtestutilv1.NewInMemKeeperParams(t)
	consumerKeeperv1, ctx, ctrl, _ := ccvtestutilv1.GetConsumerKeeperAndCtx(t, keeperParamsv1)
	defer ctrl.Finish()

	require.Empty(t, consumerKeeperv1.GetPendingPackets(ctx).List)
	consumerKeeperv1.SetPendingPackets(ctx, ccvtypesv1.ConsumerPacketDataList{List: []ccvtypesv1.ConsumerPacketData{{}, {}}})
	require.Len(t, consumerKeeperv1.GetPendingPackets(ctx).List, 2)

	// Transfer store ownership to v2 keeper
	keeperParamsv2 := ccvtestutilv2.InMemKeeperParams{
		Cdc:            keeperParamsv1.Cdc,
		StoreKey:       keeperParamsv1.StoreKey,
		ParamsSubspace: keeperParamsv1.ParamsSubspace,
		Ctx:            keeperParamsv1.Ctx,
	}

	consumerKeeperv2, ctx, ctrl, _ := ccvtestutilv2.GetConsumerKeeperAndCtx(t, keeperParamsv2)
	defer ctrl.Finish()

	// v2 keeper should be able to read v1 data
	require.Len(t, consumerKeeperv2.GetPendingPackets(ctx).List, 2)
}

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
