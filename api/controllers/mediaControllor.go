package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/models"
	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/responses"
)

func (a *App) StoreDataMedia(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.DataMedia{}
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

	err = data.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	id := r.Context().Value("userID").(float64)

	data.Created_by = int64(id)
	store, err := data.SaveDataMedia(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}
func (a *App) UpdateDataMedia(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.DataMedia{}
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

	err = data.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	id := r.Context().Value("userID").(float64)

	data.Updated_by = int64(id)
	store, err := data.UpdateDataMedia(data.ID, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) AllDataMedia(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("seacrh")
	param1 := r.URL.Query().Get("value")
	data := &models.DataMedia{}
	users, err := data.GetAll(param1, param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) DeleteDataMedia(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("id")
	data := &models.DataMedia{}
	users, err := data.Delete(param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
