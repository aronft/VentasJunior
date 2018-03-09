package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blazte/VentasJunior/models"
)

//DisplayMessage devuelve un mensaje de error al cliente
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al covertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
