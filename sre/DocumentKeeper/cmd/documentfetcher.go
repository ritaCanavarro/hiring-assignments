/*
Copyright © 2023 ritaCanavarro
*/
package cmd

import (
	"DocumentKeeper/internal/http"

	"github.com/spf13/cobra"
)

var httpPort string

var documentfetcherCmd = &cobra.Command{
	Use:   "documentfetcher",
	Short: "Document fetcher will fetch you either a PNG or PDF.",

	Run: func(cmd *cobra.Command, args []string) {
		http.StartDocumentFetcher(httpPort)
	},
}

func init() {
	documentfetcherCmd.Flags().IntVarP(&httpPort, "http.port", "", "Defines the http port of the server.")
	rootCmd.AddCommand(documentfetcherCmd)
}
