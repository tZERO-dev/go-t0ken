package rules

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/contracts/compliance/rules"
)

var (
	Command = &cobra.Command{
		Use:   "rules",
		Short: "Compliance-rule utilities",
	}

	DeployCommand = &cobra.Command{
		Use:   "deploy <name>",
		Short: "Deploys a new <name> compliance-rule",
		Example: `t0ken rules deploy restrictAll --keystoreAddress owner
Available Rules:
  - restrict
  - restrictAll
  - restrictDuringLockup
  - restrictFromAdvisor
  - restrictFromAffiliate
  - restrictFromInvestor
  - restrictHolderThreshold
  - restrictToAccreditedInvestor
  - restrictToBrokerOrCustodialAccount
  - restrictToContract
  - restrictToCustodianOrCustodialAccountOrBroker
  - restrictTransferFrom`,
		Args:   cobra.ExactArgs(1),
		PreRun: commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				addr common.Address
				tx   *types.Transaction
				err  error
			)

			switch args[0] {
			case "restrict":
				addr, tx, _, err = rules.DeployRestrict(cli.Conn.Opts, cli.Conn.Client)
			case "restrictAll":
				addr, tx, _, err = rules.DeployRestrictAll(cli.Conn.Opts, cli.Conn.Client)
			case "restrictDuringLockup":
				addr, tx, _, err = rules.DeployRestrictDuringLockup(cli.Conn.Opts, cli.Conn.Client)
			case "restrictFromAdvisor":
				addr, tx, _, err = rules.DeployRestrictFromAdvisor(cli.Conn.Opts, cli.Conn.Client)
			case "restrictFromAffiliate":
				addr, tx, _, err = rules.DeployRestrictFromAffiliate(cli.Conn.Opts, cli.Conn.Client)
			case "restrictFromInvestor":
				addr, tx, _, err = rules.DeployRestrictFromInvestor(cli.Conn.Opts, cli.Conn.Client)
			case "restrictHolderThreshold":
				addr, tx, _, err = rules.DeployRestrictHolderThreshold(cli.Conn.Opts, cli.Conn.Client)
			case "restrictToAccreditedInvestor":
				addr, tx, _, err = rules.DeployRestrictToAccreditedInvestor(cli.Conn.Opts, cli.Conn.Client)
			case "restrictToBrokerOrCustodialAccount":
				addr, tx, _, err = rules.DeployRestrictToBrokerOrCustodialAccount(cli.Conn.Opts, cli.Conn.Client)
			case "restrictToContract":
				addr, tx, _, err = rules.DeployRestrictToContract(cli.Conn.Opts, cli.Conn.Client)
			case "restrictToCustodianOrCustodialAccountOrBroker":
				addr, tx, _, err = rules.DeployRestrictToCustodianOrCustodialAccountOrBroker(cli.Conn.Opts, cli.Conn.Client)
			case "restrictTransferFrom":
				addr, tx, _, err = rules.DeployRestrictTransferFrom(cli.Conn.Opts, cli.Conn.Client)
			}

			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "name <address>",
		Short:   "Gets the name of the rule at the <address>",
		Example: "t0ken rules name 0xa16e2faddd9469e5076c33ae53a2c6f40a058413",
		Args:    cli.AddressArgFunc("address", 0),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.Connect(cmd, args)
			caller, err := rules.NewRuleCaller(addr, cli.Conn.Client)
			cli.CheckErr(cmd, err)
			cli.CheckGetter(cmd)(caller.Name(&bind.CallOpts{}))
		},
	},
}

func init() {
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)
}
