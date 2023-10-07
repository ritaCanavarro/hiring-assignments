/*
Copyright © 2023 ritaCanavarro
*/
package cmd

import (
	"os"
	"os/signal"

	"github.com/hiring-assignments/sre/DocumentKeeper/internal/command"
)

var documentFetcherFlags = &command.documentFetcherFlags{}

var documentfetcherCmd = &cobra.Command{
	Use: "documentfetcher",
	Short: "Document fetcher will fetch you either a PNG or PDF."

	Run: func(cmd *cobra.Command, args []string){
		ctx, cancel := context.WithCancel(context.Backgroung())
		wg, ctx := errgroup.WithContext(ctx)
		term := make(chan os.Signal, 1)
		signal.Notify(term, os.Interrupt, syscall.SIGTERM)

		//TODO: Initialize HTTP SERVER

		//TODO: Put HTTP Server to run
	}
}

func init() {
	documentfetcherCmd.Flags().Int("http.port", "", "Defines the http port of the server.")
	rootCmd.AddCommand(documentfetcherCmd)
}
