package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/models"
	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/responses"
	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/services"
)

func (a *App) StoreDataPenelitian(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.DataPenelitian{}
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

	modelalat := &models.DataAlat{}

	cek, err := modelalat.GetDataAlatInt(a.DB, "id = ?", int(data.Data_alats_id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	id := r.Context().Value("userID").(float64)
	data.Nama_alat = cek.Name_alat
	data.Created_by = int64(id)
	store, err := data.SaveDataPenelitian(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}
func (a *App) UpdateDataPenelitian(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.DataPenelitian{}
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
	store, err := data.UpdateStatusDataPenelitian(data.ID, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) AllDataPenelitian(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("seacrh")
	param1 := r.URL.Query().Get("value")
	data := &models.DataPenelitian{}
	users, err := data.GetAll(param1, param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) UserDataPenelitian(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("id")
	data := &models.DataPenelitian{}
	dataint := services.StringToInt(param2)
	users, err := data.Getint("created_by", dataint, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) DeleteDataPenelitian(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("id")
	data := &models.DataPenelitian{}
	users, err := data.Delete(param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
