package main

import (
	"net/http"

	"github.com/klausks/bootdev-chirpy/admin"
	"github.com/klausks/bootdev-chirpy/api"
)

var serverMux = http.NewServeMux()

var server = http.Server{
	Handler: serverMux,
	Addr:    ":8080",
}

func StartServer() error {
	metricsHandler := admin.NewHandler()
	serverMux.Handle("/app/", http.StripPrefix("/app", metricsHandler.MiddlewareFileServerHitsIncrement(http.FileServer(http.Dir(".")))))
	serverMux.HandleFunc("GET /api/healthz", api.HandleReadinessReq)
	serverMux.HandleFunc("POST /api/validate_chirp", api.HandleValidateChirpReq)

	serverMux.HandleFunc("GET /admin/metrics", metricsHandler.HandleGetFileServerHits)
	serverMux.HandleFunc("POST /admin/reset", metricsHandler.HandleResetMetricsReq)

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
