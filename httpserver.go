package main

import (
	"fmt"
	"net/http"
)

var serverMux = http.NewServeMux()

var server = http.Server{
	Handler: serverMux,
	Addr:    ":8080",
}

func (cfg *metrics) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(rw, req)
	})
}

func StartServer() error {
	apiCfg := metrics{}
	serverMux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(".")))))
	serverMux.HandleFunc("GET /api/healthz", readinessHandler)
	serverMux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	serverMux.HandleFunc("POST /admin/reset", apiCfg.resetMetricsHandler)

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func readinessHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("OK"))
}

func (apiCfg *metrics) metricsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	hits := apiCfg.fileserverHits.Load()
	rw.Write([]byte(fmt.Sprintf("Hits: %v", hits)))
}

func (metrics *metrics) resetMetricsHandler(rw http.ResponseWriter, req *http.Request) {
	metrics.fileserverHits.Store(0)
	rw.Header().Add("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("OK"))
}
