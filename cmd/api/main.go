package main;

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/stierprogrammer/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
);

func main() {
	log.SetReportCaller(true);

	var r *chi.Mux = chi.NewRouter();

	handlers.Handler(r);
}