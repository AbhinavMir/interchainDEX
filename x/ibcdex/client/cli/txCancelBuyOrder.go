package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/abhinav/interchainDEX/x/ibcdex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdCancelBuyOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancelBuyOrder [port] [channel] [amountDenom] [priceDenom] [orderID]",
		Short: "Cancel a buy order",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPort := string(args[0])
			argsChannel := string(args[1])
			argsAmountDenom := string(args[2])
			argsPriceDenom := string(args[3])
			argsOrderID, _ := strconv.ParseInt(args[4], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelBuyOrder(clientCtx.GetFromAddress().String(), string(argsPort), string(argsChannel), string(argsAmountDenom), string(argsPriceDenom), int32(argsOrderID))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
