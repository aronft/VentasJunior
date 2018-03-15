package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blazte/VentasJunior/commons"
	"github.com/blazte/VentasJunior/configuration"
	"github.com/blazte/VentasJunior/models"
)

//ProductCreate crea una producto
func ProductCreate(w http.ResponseWriter, r *http.Request) {
	Product := models.Product{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&Product)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el producto  a registrar: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConection()
	defer db.Close()

	err = db.Create(&Product).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al registrar el producto: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Producto creada con éxito"
	m.Code = http.StatusOK
	commons.DisplayMessage(w, m)
}
