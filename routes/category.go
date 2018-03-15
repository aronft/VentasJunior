package routes

import (
	"github.com/blazte/VentasJunior/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetCategoryRouter Es una ruta para crear las categoria
func SetCategoryRouter(router *mux.Router) {
	prefix := "/api/category"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CategoryCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)

}
