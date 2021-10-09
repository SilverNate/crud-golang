package controllers

import (
	"net/http"

	"github.com/SilverNate/crud/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Selamat datang di Aplikasi CRUD dengan SOLID Principles")

}
