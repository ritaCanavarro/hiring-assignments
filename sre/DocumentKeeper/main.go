/*
Copyright © 2023 ritaCanavarro
*/
package main

import "DocumentKeeper/http"

const httpPort = 4096

func main() {
	http.StartDocumentFetcher(httpPort)
}
