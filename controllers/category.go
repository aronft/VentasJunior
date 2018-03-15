package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blazte/VentasJunior/commons"
	"github.com/blazte/VentasJunior/configuration"
	"github.com/blazte/VentasJunior/models"
)

//CategoryCreate crea una categoria par aun producto
func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	Category := models.Category{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer la categoria a registrar: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConection()
	defer db.Close()

	err = db.Create(&Category).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al registrar la categoria: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "categoria creada con éxito"
	m.Code = http.StatusOK
	commons.DisplayMessage(w, m)
}
