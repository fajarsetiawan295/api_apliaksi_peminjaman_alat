package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/tokoumat/api/models"
	"github.com/tokoumat/api/responses"
	"github.com/tokoumat/api/services"
)

func (a *App) StorePeminjaman(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.PeminjamanBarang{}
	dataMedia := &models.DataMedia{}
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

	cek, _ := dataMedia.GetDataMediaInt(a.DB, "id = ?", data.Media_id)
	if cek == nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("media tidak di temukan"))
		return
	}

	id := r.Context().Value("userID").(float64)
	data.Users_id = int64(id)
	data.Status = 0

	err = data.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if data.Jenis == 0 {

		user := &models.User{
			Status_pembayaran: 2,
		}
		cekupdate, err := user.UpdateStatusPembayaran(int(data.Users_id), a.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		resp["cekupdate"] = cekupdate

	}
	if data.Jenis == 1 {
		user := &models.User{
			Status_penelitian: 2,
		}
		cekupdate, err := user.UpdateStatusPenelitian(uint(data.Users_id), 2, a.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		resp["cekupdate"] = cekupdate
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
func (a *App) UpdatePeminjaman(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	var sta int
	data := &models.PeminjamanBarang{}
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

	if data.Jenis == 0 {

		if data.Status == 1 {
			sta = 1
		} else {
			sta = 2
		}

		println("masuk pembayaran")
		println(sta)
		user := &models.User{
			Status_pembayaran: int64(sta),
		}
		cekupdate, err := user.UpdateStatusPembayaran(int(data.ID), a.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		resp["cekupdate"] = cekupdate

	}
	if data.Jenis == 1 {

		if data.Status == 1 {
			sta = 1
		} else {
			sta = 2
		}
		user := &models.User{
			Status_penelitian: int64(sta),
		}

		println("masuk penelitian")
		println(sta)
		cekupdate, err := user.UpdateStatusPenelitian(uint(data.ID), 2, a.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		resp["cekupdate"] = cekupdate
	}

	store, err := data.Update(data.ID, data.Jenis, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["data"] = store
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) MePeminjaman(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	id := r.Context().Value("userID").(float64)
	data := &models.PeminjamanBarang{}
	jenis := r.URL.Query().Get("jenis")
	users, err := data.FindInt("peminjaman_barangs.users_id = ?", int(id), jenis, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}

func (a *App) StatusPeminjaman(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	data := &models.PeminjamanBarang{}
	jenis := r.URL.Query().Get("jenis")
	status := r.URL.Query().Get("status")
	id := r.URL.Query().Get("id")
	users, err := data.FindStatus("peminjaman_barangs.users_id = ?", services.StringToInt(id), jenis, status, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}

func (a *App) DeletePeminjaman(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Sukses", "code": 200}
	param2 := r.URL.Query().Get("id")
	data := &models.PeminjamanBarang{}
	users, err := data.Delete(param2, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
