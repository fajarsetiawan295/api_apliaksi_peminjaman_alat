package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tokoumat/api/models"
	"github.com/tokoumat/api/responses"
	"github.com/tokoumat/api/services"
	"github.com/tokoumat/utils"
)

var (
	resp = map[string]interface{}{"status": true, "message": "Succes", "code": 200}
)

func (a *App) UserSignUp(w http.ResponseWriter, r *http.Request) {
	erro := r.ParseMultipartForm(100000)
	if erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return
	}
	upload, err := services.UploadFile(r)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	cek := &models.User{}

	usr, _ := cek.GetUser(a.DB, "email  = ?", r.FormValue("email"))
	if usr != nil {
		responses.FAILED(w, http.StatusBadRequest, "email already registered, please login", nil)
		return
	}

	cekhandpone, _ := cek.GetUser(a.DB, " nomor_hp  = ?", r.FormValue("nomor_hp"))
	if cekhandpone != nil {
		responses.FAILED(w, http.StatusBadRequest, "nomor handphone already registered, please login", nil)
		return
	}

	user := &models.User{
		Name:     r.FormValue("name"),
		Nomor_hp: r.FormValue("nomor_hp"),
		Npm:      r.FormValue("npm"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
		Role:     r.FormValue("role"),
		Foto:     upload,
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userCreated, err := user.SaveUser(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = userCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}

// Login signs in users
func (a *App) Login(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	usr, err := user.GetUser(a.DB, "email = ?", user.Email)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if usr == nil { // user is not registered
		responses.FAILED(w, http.StatusBadRequest, "email not registered", nil)
		return
	}

	err = models.CheckPasswordHash(user.Password, usr.Password)
	if err != nil {
		responses.FAILED(w, http.StatusBadRequest, "Login failed, check your password", nil)
		return
	}

	token, err := utils.EncodeAuthToken(usr.ID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["token"] = token
	resp["data"] = usr
	responses.JSON(w, http.StatusOK, resp)
	return

}

func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetAllUsers(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) getProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(float64)
	users, err := models.Getfinduser(uint(userId), a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
func (a *App) getbyrole(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("role")
	users, err := models.GetAllbyRole(param1, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	resp["data"] = users
	responses.JSON(w, http.StatusOK, resp)
	return
}
