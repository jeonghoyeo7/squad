package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/cosmosquad-labs/squad/v3/x/farm/client/cli"
	"github.com/cosmosquad-labs/squad/v3/x/farm/client/rest"
)

// ProposalHandler is the public plan command handler.
// Note that rest.ProposalRESTHandler will be deprecated in the future.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitFarmingPlanProposal, rest.ProposalRESTHandler)
)
