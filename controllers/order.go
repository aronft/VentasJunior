package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blazte/VentasJunior/commons"
	"github.com/blazte/VentasJunior/configuration"
	"github.com/blazte/VentasJunior/models"
)

//OrderCreate crea una orden de compra
func OrderCreate(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer la orden a registrar: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	// fmt.Printf("%+v\n", order)

	db := configuration.GetConection()
	defer db.Close()

	err = db.Create(&order).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al registrar la orden: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Orden creada con éxito"
	m.Code = http.StatusOK
	commons.DisplayMessage(w, m)
}
