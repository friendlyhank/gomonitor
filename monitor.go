package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// PromHandler - 设置handler信息
func PromHandler()http.Handler{
	return promhttp.Handler()
}