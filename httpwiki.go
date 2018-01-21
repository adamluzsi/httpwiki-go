package httpwiki

import (
	"io/ioutil"
	"net/http"

	"github.com/adamluzsi/httpwiki-go/status"
)

type HTTPWiki struct {
	Storage   Storage
	Templates Templates
}

func New(s Storage, t Templates) *HTTPWiki {
	return &HTTPWiki{Storage: s, Templates: t}
}

func (wiki HTTPWiki) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		if r.URL.Query().Get("edit") != "" {
			wiki.edit(w, r)
		} else {
			wiki.show(w, r)
		}

	case http.MethodPost:
		wiki.save(w, r)

	default:
		status.NotFound(w)
	}
}

func (wiki *HTTPWiki) writeResponse(w http.ResponseWriter, code int, body []byte) {
	w.Header().Set("Content-Type", http.DetectContentType(body))
	w.WriteHeader(code)
	w.Write(body)
}

func (wiki *HTTPWiki) edit(w http.ResponseWriter, r *http.Request) {
	data, err := wiki.Storage.Load(r.URL.Path)

	if err != nil {
		data = []byte{}
	}

	wiki.writeResponse(w, http.StatusOK, wiki.Templates.EditPage(data))
}

func (wiki *HTTPWiki) show(w http.ResponseWriter, r *http.Request) {

	data, err := wiki.Storage.Load(r.URL.Path)

	if err != nil {
		query := r.URL.Query()
		query.Set("edit", "TRUE")
		http.Redirect(w, r, r.URL.Path+"?"+query.Encode(), 307)
		return
	}

	wiki.writeResponse(w, http.StatusOK, wiki.Templates.ShowPage(data))

}

func (wiki *HTTPWiki) save(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		status.InternalServerError(w)
		return
	}

	if err := wiki.Storage.Save(r.URL.Path, data); err != nil {
		status.InternalServerError(w)
		return
	}

	http.Redirect(w, r, r.URL.Path, http.StatusOK)

}
