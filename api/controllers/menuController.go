package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tokoumat/api/models"
	"github.com/tokoumat/api/responses"
)

func (a *App) StoreMenu(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.Menu{}
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

	store, err := data.Save(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}
func (a *App) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.Menu{}
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

	store, err := data.Update(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) AllMenu(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.Menu{}
	users, err := data.All(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) ByRoleMenu(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("role")
	data := &models.Menu{}
	users, err := data.FindString("role = ?", param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
