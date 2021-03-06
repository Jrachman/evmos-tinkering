package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

func queryState() error {
	myAddress, err := sdk.AccAddressFromBech32("eth...")
	if err != nil {
		return err
	}

	// Create a connection to the gRPC server.
	grpcConn := grpc.Dial(
		"127.0.0.1:9090",    // your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: myAddress, Denom: "atom"},
	)
	if err != nil {
		return err
	}

	fmt.Println(bankRes.GetBalance()) // Prints the account balance

	return nil
}
