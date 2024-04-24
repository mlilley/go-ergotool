package cmd

import (
	"fmt"

	"github.com/mlilley/go-ergotool"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update footprint positions in an existing kicad_pcb file with those from a newly generated kicad_pcb output file",
	Long:  "Update footprint positions in an existing kicad_pcb file with those from a newly generated kicad_pcb output file",
	RunE: func(cmd *cobra.Command, args []string) error {

		srcFilename, _ := cmd.Flags().GetString("src")
		destFilename, _ := cmd.Flags().GetString("dest")

		// read and parse input files
		srcDoc, err := ergotool.ReadDoc(srcFilename)
		if err != nil {
			return fmt.Errorf("unable to read src file '%s': %v", srcFilename, err)
		}
		destDoc, err := ergotool.ReadDoc(destFilename)
		if err != nil {
			return fmt.Errorf("failed to read dest file '%s': %v", destFilename, err)
		}

		// perform operation
		err = ergotool.UpdateFootprintLocations(srcDoc, destDoc)
		if err != nil {
			return err
		}

		// write output to destination file
		err = ergotool.WriteDoc(destFilename, destDoc)
		if err != nil {
			return fmt.Errorf("failed to write dest file '%s': %v", destFilename, err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
	updateCmd.Flags().StringP("src", "s", "", "source ergogen kicad_pcb file")
	updateCmd.MarkFlagFilename("src")
	updateCmd.MarkFlagRequired("src")

	updateCmd.Flags().StringP("dest", "d", "", "destination kicad_pcb file to update")
	updateCmd.MarkFlagFilename("dest")
	updateCmd.MarkFlagRequired("dest")

	//updateCmd.MarkFlagsRequiredTogether("src", "dest")

	// Todo: Add flags for specifying a reference-value filter regex
	// Todo: Add flags for specifying a footprint-name filter regex
}
