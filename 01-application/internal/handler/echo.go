package handler

import (
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"time"

	util "CLOUDRU/internal/utils"

	"github.com/sirupsen/logrus"
)

var templateFS embed.FS

var tmpl = template.Must(
	template.ParseFS(templateFS, "templates/index.html"),
)

type PageData struct {
	Hostname  string
	IP        string
	Author    string
	Time      string
	UserAgent string
}

func EchoHandler(log *logrus.Logger, author string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Errorf("cannot get hostname: %v", err)
			hostname = "unknown"
		}
		ip := util.GetLocalIP(log)

		data := PageData{
			Hostname:  hostname,
			IP:        ip,
			Author:    author,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
			UserAgent: r.UserAgent(),
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			log.Errorf("template execute error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.WithFields(logrus.Fields{
			"host": hostname,
			"ip":   ip,
		}).Info("served /")
	}
}

func HealthHandler(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Debug("served /healthz")
	}
}
