package app

import (
	"finanzas/misc/handlers"
	"finanzas/misc/models"
	"fmt"
	"log"
	"net/http"
	"os"

	gorillah "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func connectToDB(conn string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	// Run Migrations
	runMigrations(db)

	return db, nil
}

func runMigrations(db *gorm.DB) {

	db.AutoMigrate(&models.ItemType{})

	db.AutoMigrate(&models.Item{})

	db.AutoMigrate(&models.ItemProcess{})
}

func Main() error {

	db, err := connectToDB("postgresql://localhost/finanzas?user=postgres&password=gadmin&sslmode=disable")
	if err != nil {
		log.Println("Hey!", err)
		return err
	}

	defer db.Close()

	itemHandler := handlers.NewItemHandler(db)

	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/item/{id}", itemHandler.FindByID).Methods("GET")
	router.HandleFunc("/item/find", itemHandler.Find).Methods("POST")
	router.HandleFunc("/item", itemHandler.Create).Methods("POST")
	router.HandleFunc("/item", itemHandler.Update).Methods("PUT")

	listen := fmt.Sprintf(":%d", 9000)

	return http.ListenAndServe(listen, gorillah.CombinedLoggingHandler(os.Stdout, router))
}
