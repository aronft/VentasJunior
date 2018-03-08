package main

import (
	"flag"
	"log"

	"github.com/blazte/VentasJunior/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizo la migración...")
	}

}
