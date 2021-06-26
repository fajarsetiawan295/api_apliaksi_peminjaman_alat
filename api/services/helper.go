package services

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func UploadFile(n *http.Request) (string, error) {
	var filename string
	uploadedFile, handler, err := n.FormFile("foto")
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filename = "foto" + RandStringBytes(5) + " - " + handler.Filename
	fileLocation := filepath.Join(dir, "img/", filename)
	println("ini file nya "+dir, "img/", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return "", err
	}

	return "img/" + filename, nil
}

func StringToInt(n string) int {
	i, err := strconv.Atoi(n)
	if err != nil {
		log.Println(err)
	}
	return i
}
