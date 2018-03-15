package migration

import (
	"github.com/blazte/VentasJunior/configuration"
	"github.com/blazte/VentasJunior/models"
)

//Migrate se crean las tablas en la base de datos
func Migrate() {
	// se conecta a la base de datos
	db := configuration.GetConection()
	defer db.Close() //cerramos la conexxion

	//crear los modelos

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Order{})
	db.CreateTable(&models.Category{})
	db.CreateTable(&models.Product{})
	db.CreateTable(&models.OrderDetail{})

	db.Model(&models.Product{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Model(&models.Order{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.OrderDetail{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
	db.Model(&models.OrderDetail{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")

	db.Model(&models.OrderDetail{}).AddUniqueIndex("product_id_order_id", "product_id", "order_id")
}
