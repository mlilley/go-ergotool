package cmd

import (
	"fmt"

	"github.com/mlilley/go-ergotool"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Confirm ability to parse and display a given kicad_pcb file",
	Long:  "Confirm ability to parse and display a given kicad_pcb file",
	RunE: func(cmd *cobra.Command, args []string) error {

		srcFilename, _ := cmd.Flags().GetString("src")

		// read and parse input file
		srcDoc, err := ergotool.ReadDoc(srcFilename)
		if err != nil {
			return fmt.Errorf("unable to read src file '%s': %v", srcFilename, err)
		}

		// print parsed document
		fmt.Println(srcDoc.Root().String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringP("src", "s", "", "source kicad_pcb file")
	parseCmd.MarkFlagFilename("src")
	parseCmd.MarkFlagRequired("src")
}
