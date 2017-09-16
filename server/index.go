package server

import (
	"net/http"
	"strconv"

	"github.com/nathan-osman/go-ldserver/presenter"
)

// TODO: redirect after POST

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(2 ^ 16); err != nil {
			data["Error"] = err.Error()
		}
		switch r.Form.Get("action") {
		case "upload":
			if s.presenter != nil {
				s.presenter.Stop()
			}
			r, _, err := r.FormFile("file")
			if err != nil {
				data["Error"] = err.Error()
				break
			}
			defer r.Close()
			p, err := presenter.NewPresenter(s.manager, r)
			if err != nil {
				data["Error"] = err.Error()
				break
			}
			s.presenter = p
		case "start":
			if s.presenter != nil {
				s.presenter.Start()
			}
		case "stop":
			if s.presenter != nil {
				s.presenter.Stop()
			}
		default:
			data["Error"] = "invalid action specified"
		}
	}
	data["Loaded"] = s.presenter != nil
	b := renderTemplate(data)
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
