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

		err := initializeDocumentFetcherFlags(cmd)

		if err != nil {
			logrus.Errorf("couldn't initialize Document Fetcher cmd flags %v", err)
			os.Exit(1)
		}

		//TODO: Initialize HTTP SERVER

		//TODO: Put HTTP Server to run
	}
}

func initializeDocumentFetcherFlags(cmd *cobra.Command){
	httpPort, err := cmd.Flags().GetString("http.port")

	if err != nil {
		logrus.Errorf("couldn't fetch http.port config %v", err)
		return err
	}

	documentFetcherFlags.HttpPort = httpPort

	return nil
}

func init() {
	documentfetcherCmd.Flags().Int("document.id", "", "Defines the identifier of the document to be fetched.")
	rootCmd.AddCommand(documentfetcherCmd)
}
