package routes

import (
	"github.com/blazte/VentasJunior/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetUserRouter es la ruta para el registro de usuario
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			//negroni.HandlerFunc(controller.ValidateToken)
			negroni.Wrap(subRouter),
		),
	)
}
