package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

var TEST bool
var TemplateDir = "./templates"

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(NotFound)
	router.PathPrefix("/pawebcontrol").HandlerFunc(ShowPaWebControl)
	router.PathPrefix("/sinks").HandlerFunc(ShowSinks)
	//router.Path("/notfound").Handler(http.HandlerFunc(NotFound))
	router.PathPrefix("/volumes").HandlerFunc(ShowVolumes)
	router.PathPrefix("/pacmd").HandlerFunc(ShowPaCmd)
	router.PathPrefix("/api").HandlerFunc(HandleAPI)
	//router.PathPrefix("/workflows").HandlerFunc(ShowWorkflows)

	// static file server
	router.PathPrefix("/assets").Handler(
		http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))),
	)
	router.PathPrefix("/fonts").Handler(
		http.StripPrefix("/fonts/", http.FileServer(http.Dir("./assets/fonts"))),
	)
	router.Path("/favicon.ico").Methods("GET").HandlerFunc(Favicon)
	router.Path("/").HandlerFunc(Dashboard)

	return
}
