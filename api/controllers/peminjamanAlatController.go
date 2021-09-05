package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/models"
	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/responses"
	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/services"
)

func (a *App) StorepeminjamanAlat(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.PeminjamanAlat{}
	dataMedia := &models.DataAlat{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	cek, _ := dataMedia.GetDataAlatInt(a.DB, "id = ?", int(data.Alat_id))
	if cek == nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Alat tidak di temukan"))
		return
	}

	id := r.Context().Value("userID").(float64)
	data.Users_id = int64(id)

	err = data.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	store, err := data.Save(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) MePeminjamanAlat(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	id := r.Context().Value("userID").(float64)
	jenis := r.URL.Query().Get("jenis")
	data := &models.PeminjamanAlat{}
	users, err := data.FindInt("peminjaman_alats.users_id = ?", int(id), jenis, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}

func (a *App) UserPeminjamanAlat(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	jenis := r.URL.Query().Get("jenis")
	id := r.URL.Query().Get("id")
	data := &models.PeminjamanAlat{}
	users, err := data.FindInt("peminjaman_alats.users_id = ?", services.StringToInt(id), jenis, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) DeletePeminjamanAlat(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("id")
	data := &models.PeminjamanAlat{}
	users, err := data.Delete(param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
