package views

import (
	"github.com/joegasewicz/gomek"
	"net/http"
)

type Health struct{}

func (h *Health) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	templateData := make(gomek.Data)
	templateData["health"] = map[string]string{
		"status": "OK",
	}
	gomek.JSON(w, templateData, http.StatusOK)
}
func (h *Health) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (h *Health) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (h *Health) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
