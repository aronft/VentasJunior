package routes

import (
	"github.com/blazte/VentasJunior/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetOrderRouter Es una ruta para crear las ordenes
func SetOrderRouter(router *mux.Router) {
	prefix := "/api/order"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.OrderCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)

}
