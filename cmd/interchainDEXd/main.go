package main

import (
	"os"

	"github.com/abhinav/interchainDEX/app"
	"github.com/abhinav/interchainDEX/cmd/interchainDEXd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
