package uisearchpageapi

import (
	"encoding/json"
	"net/http"

	"github.com/David-Kalashir/crs-front/app/domain/uisearchpageapp"
)

type api struct {
	uisearchpageapp *uisearchpageapp.App
}

func newAPI(uisearchpageapp *uisearchpageapp.App) *api {
	return &api{
		uisearchpageapp: uisearchpageapp,
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