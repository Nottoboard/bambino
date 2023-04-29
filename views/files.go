package views

import (
	"github.com/joegasewicz/gomek"
	"net/http"
)

type File struct {
}

func (f *File) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	var data = struct {
		Name string
	}{
		Name: "test1",
	}
	gomek.JSON(w, data, http.StatusOK)
}
func (f *File) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {

}
func (f *File) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (f *File) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
