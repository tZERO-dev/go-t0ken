package rules

//import (
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/spf13/cobra"
//
//	"github.com/tzero-dev/go-t0ken/cli"
//	"github.com/tzero-dev/go-t0ken/contracts/compliance/rules"
//)
//
//var RestrictCommand = *cobra.Command{
//	Use:     "name <address>",
//	Short:   "Gets the name of the rule at the <address>",
//	Example: "t0ken rules name 0xa16e2faddd9469e5076c33ae53a2c6f40a058413",
//	Args:    cli.AddressArgFunc("address", 0),
//	Run: func(cmd *cobra.Command, args []string) {
//		addr := common.HexToAddress(args[0])
//		cli.Connect(cmd, args)
//		caller, err := rules.NewRuleCaller(addr, cli.Conn.Client)
//		cli.CheckErr(cmd, err)
//		cli.CheckGetter(cmd)(caller.Name(&bind.CallOpts{}))
//	},
//}
