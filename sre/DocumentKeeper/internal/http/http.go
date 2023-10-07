package http

import()

func NewHttpServer(port string) (*http.Server, error){
	err := prometheus.DefaultRegisterer.Register(version.NewCollector("documentkeeper"))

	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/-/health", func(rw http.ResponseWriter, _ *http.Request){
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")

		resp := map[string]string{
			"message": "Healthy",
		}

		jsonResp, err := json.Marshal(resp)

		if err != nil {
			log.FatalF("Error happened in JSON marshal. Err: %s", err)
			return
		}

		rw.Write(jsonResp)
	})

	mux.HandleFunc("/-/ready", func(rw http.ResponseWriter, _ *http.Request){
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")

		resp := map[string]string{
			"message": "Ready",
		}

		jsonResp, err := json.Marshal(resp)

		if err != nil {
			log.FatalF("Error happened in JSON marshal. Err: %s", err)
			return
		}

		rw.Write(jsonResp)
	})
}