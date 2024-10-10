package uisearchpageapi

import (
	"encoding/json"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/uisearchpageapp"
	"github.com/gorilla/sessions"
)

type api struct {
	uisearchpageapp *uisearchpageapp.App
	store           sessions.Store
}

func newAPI(cfg Config) *api {

	store := sessions.NewCookieStore([]byte(cfg.StoreKey))
	store.Options = &sessions.Options{
		Path: "/",
	}
	uisearchpageapp := uisearchpageapp.NewApp(cfg.UIsearchpagebus)
	return &api{
		uisearchpageapp: uisearchpageapp,
		store:           store,
	}
}

func (api *api) searchpage(w http.ResponseWriter, r *http.Request) {

	t, err := api.uisearchpageapp.SearchPage(r.Context())
	if err != nil {
		res := struct {
			Error string
		}{
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}
	err = t.Render(r.Context(), w)
	if err != nil {
		res := struct {
			Error string
		}{
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}
}

func (api *api) login(w http.ResponseWriter, r *http.Request) {

	com, err := api.uisearchpageapp.Login(r.Context())
	if err != nil {
		res := struct {
			Error string
		}{
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}
	err = com.Render(r.Context(), w)
	if err != nil {
		res := struct {
			Error string
		}{
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}
}

func (api *api) logout(w http.ResponseWriter, r *http.Request) {

	com, err := api.uisearchpageapp.Logout(r.Context())
	if err != nil {
		res := struct {
			Error string
		}{
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}
	err = com.Render(r.Context(), w)
	if err != nil {
		res := struct {
			Error string
		}{
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}
}
