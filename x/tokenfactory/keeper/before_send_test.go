package keeper_test

import (
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v16/x/tokenfactory/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type SendMsgTestCase struct {
	desc       string
	msg        func(denom string) *banktypes.MsgSend
	expectPass bool
}

func (s *KeeperTestSuite) TestBeforeSendHook() {
	s.SkipIfWSL()
	for _, tc := range []struct {
		desc     string
		wasmFile string
		sendMsgs []SendMsgTestCase
	}{
		{
			desc:     "should not allow sending 100 amount of *any* denom",
			wasmFile: "./testdata/no100.wasm",
			sendMsgs: []SendMsgTestCase{
				{
					desc: "sending 1 of factorydenom should not error",
					msg: func(factorydenom string) *banktypes.MsgSend {
						return banktypes.NewMsgSend(
							s.TestAccs[0],
							s.TestAccs[1],
							sdk.NewCoins(sdk.NewInt64Coin(factorydenom, 1)),
						)
					},
					expectPass: true,
				},
				{
					desc: "sending 100 of non-factorydenom should not error",
					msg: func(factorydenom string) *banktypes.MsgSend {
						return banktypes.NewMsgSend(
							s.TestAccs[0],
							s.TestAccs[1],
							sdk.NewCoins(sdk.NewInt64Coin(factorydenom, 1)),
						)
					},
					expectPass: true,
				},
				{
					desc: "sending 100 of factorydenom should not work",
					msg: func(factorydenom string) *banktypes.MsgSend {
						return banktypes.NewMsgSend(
							s.TestAccs[0],
							s.TestAccs[1],
							sdk.NewCoins(sdk.NewInt64Coin(factorydenom, 100)),
						)
					},
					expectPass: false,
				},
				{
					desc: "sending 100 of factorydenom should not work",
					msg: func(factorydenom string) *banktypes.MsgSend {
						return banktypes.NewMsgSend(
							s.TestAccs[0],
							s.TestAccs[1],
							sdk.NewCoins(sdk.NewInt64Coin("foo", 100)),
						)
					},
					expectPass: false,
				},
				{
					desc: "having 100 coin within coins should not work",
					msg: func(factorydenom string) *banktypes.MsgSend {
						return banktypes.NewMsgSend(
							s.TestAccs[0],
							s.TestAccs[1],
							sdk.NewCoins(sdk.NewInt64Coin(factorydenom, 100), sdk.NewInt64Coin("foo", 1)),
						)
					},
					expectPass: false,
				},
			},
		},
	} {
		s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			// setup test
			s.SetupTest()

			// upload and instantiate wasm code
			wasmCode, err := os.ReadFile(tc.wasmFile)
			s.Require().NoError(err, "test: %v", tc.desc)
			codeID, _, err := s.contractKeeper.Create(s.Ctx, s.TestAccs[0], wasmCode, nil)
			s.Require().NoError(err, "test: %v", tc.desc)
			cosmwasmAddress, _, err := s.contractKeeper.Instantiate(s.Ctx, codeID, s.TestAccs[0], s.TestAccs[0], []byte("{}"), "", sdk.NewCoins())
			s.Require().NoError(err, "test: %v", tc.desc)

			// create new denom
			res, err := s.msgServer.CreateDenom(sdk.WrapSDKContext(s.Ctx), types.NewMsgCreateDenom(s.TestAccs[0].String(), "bitcoin"))
			s.Require().NoError(err, "test: %v", tc.desc)
			denom := res.GetNewTokenDenom()

			// mint enough coins to the creator
			_, err = s.msgServer.Mint(sdk.WrapSDKContext(s.Ctx), types.NewMsgMint(s.TestAccs[0].String(), sdk.NewInt64Coin(denom, 1000000000)))
			s.Require().NoError(err)
			// mint some non token factory denom coins for testing
			s.FundAcc(sdk.AccAddress(s.TestAccs[0].String()), sdk.Coins{sdk.NewInt64Coin("foo", 100000000000)})

			// set beforesend hook to the new denom
			_, err = s.msgServer.SetBeforeSendHook(sdk.WrapSDKContext(s.Ctx), types.NewMsgSetBeforeSendHook(s.TestAccs[0].String(), denom, cosmwasmAddress.String()))
			s.Require().NoError(err, "test: %v", tc.desc)

			for _, sendTc := range tc.sendMsgs {
				_, err := s.bankMsgServer.Send(sdk.WrapSDKContext(s.Ctx), sendTc.msg(denom))
				if sendTc.expectPass {
					s.Require().NoError(err, "test: %v", sendTc.desc)
				} else {
					s.Require().Error(err, "test: %v", sendTc.desc)
				}
			}
		})
	}
}
