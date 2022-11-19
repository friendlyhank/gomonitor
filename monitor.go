package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func PromHandler()http.Handler{
	return promhttp.Handler()
}