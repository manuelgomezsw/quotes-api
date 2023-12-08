package repository

import (
	"log"
	"quotes-api/internal/infraestructure"
)

func CreateQuery() {
	_, err := infraestructure.ClientDB.Exec(
		"INSERT INTO `quotes-db`.`quotes` (author, work, phrase) VALUES (?, ?, ?)",
		"Manuel Gomez 2",
		"El despertar en las mañanas 2",
		"Las ideas son la semilla del progreso 2")
	if err != nil {
		log.Fatalf("Something happened %s", err)
	}
}
