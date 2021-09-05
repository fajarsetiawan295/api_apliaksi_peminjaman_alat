package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tokoumat/api/models"
	"github.com/tokoumat/api/responses"
	"github.com/tokoumat/api/services"
)

func (a *App) StoreKerusakan(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.KerusakanAlat{}
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
func (a *App) UpdateKerusakan(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.KerusakanAlat{}
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

	store, err := data.UpdateDataKerusakan(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) MeKerusakan(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	id := r.Context().Value("userID").(float64)
	data := &models.KerusakanAlat{}
	users, err := data.FindInt("kerusakan_alats.users_id = ?", int(id), a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) UserKerusakan(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	id := r.URL.Query().Get("id")
	data := &models.KerusakanAlat{}
	users, err := data.FindInt("kerusakan_alats.users_id = ?", services.StringToInt(id), a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) AllKerusakan(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	id := r.URL.Query().Get("id")
	data := &models.KerusakanAlat{}
	users, err := data.FindAll("kerusakan_alats.users_id = ?", services.StringToInt(id), a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) DeleteKerusakan(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("id")
	data := &models.KerusakanAlat{}
	users, err := data.Delete(param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
