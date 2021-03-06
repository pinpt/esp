package cmd

import (
	"fmt"
	"github.com/pinpt/esp/internal/common"
	"os"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func display(p common.EspParam, detail bool) {
	if detail {
		detailDisplay(p)
	} else {
		displayParam(p)
	}
}

func displayParam(p common.EspParam) {
	name := aurora.BrightYellow(p.Name)
	fmt.Printf("%s: %s\n", name, p.Value)
}

func detailDisplay(p common.EspParam) {
	data := [][]string{
		{aurora.BrightYellow("ID").String(), p.Id},
		{aurora.BrightYellow("Last_Modified").String(), p.LastModifiedDate.String()},
		{aurora.BrightYellow("Name").String(), p.Name},
		{aurora.BrightYellow("Type").String(), p.Type},
		{aurora.BrightYellow("Value").String(), p.Value},
		{aurora.BrightYellow("Version").String(), strconv.FormatInt(p.Version, 10)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Keys", "Value"})
	table.AppendBulk(data)
	table.Render()
}

func getParamPath(p string) string {
	if strings.HasPrefix(p, "/") {
		return p
	}
	return esp.GetAppParamPath(p)
}

// getCmd gets the parameter from the backend store
var getCmd = &cobra.Command{
	Use:   "get [path]",
	Short: "Query path for SSM",
	Long:  `Allows you to get a specific ssm parameter with an exact path or recursively get params.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		decrypt, _ := cmd.Flags().GetBool("decrypt")
		details, _ := cmd.Flags().GetBool("details")

		param := c.GetParam(common.GetOneInput{
			Name:    getParamPath(args[0]),
			Decrypt: decrypt,
		})
		display(param, details)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("details", "t", false, "Show all of the attributes of a parameter.")
	getCmd.Flags().BoolP("decrypt", "d", false, "Decrypt SSM secure strings.")
}
