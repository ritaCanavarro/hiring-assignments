package api

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/ritaCanavarro/hiring-assignments/sre/DocumentKeeper/internal/http"
)

// -------------------- Structures ----------------------

type HttpServerSettings struct{
	HttpPort string 
}

func InitializeHttpServerSettings(httpPort string) HttpServerSettings{
	return HttpServerSettings{
		HttpPort: httpPort,
	}
}

func StartDocumentFetcher(){
	httpServer, err := http.NewHttpServer(HttpServerSettings.httpPort)

	if err != nil{
		logrus.Errorf("http server error %v", err)
		os.Exit(1)
	}

	httpServer.Handler.mux.Handle("/document/{id}",GetDocument).Methods("GET")


	wg.Go(func() error{
		if err := httpServer.ListenAndServer(); err != net_http.ErrServerClosed {
			logrus.Errorf("http server error %v", err)
			return err
		}

		return nil
	})
}