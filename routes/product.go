package routes

import (
	"github.com/blazte/VentasJunior/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetProductRouter Es una ruta para crear los productos
func SetProductRouter(router *mux.Router) {
	prefix := "/api/product"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.ProductCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)

}
