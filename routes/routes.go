
package routes

import (
	"github.com/gorilla/mux"
	"go-phonecase-crud/controllers"
)


func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()


	router.HandleFunc("/phonecases", controllers.CreatePhoneCase).Methods("POST")
	router.HandleFunc("/phonecases", controllers.GetAllPhoneCases).Methods("GET")
	router.HandleFunc("/phonecases/{id}", controllers.GetPhoneCaseByID).Methods("GET")
	router.HandleFunc("/phonecases/{id}", controllers.UpdatePhoneCase).Methods("PUT")
	router.HandleFunc("/phonecases/{id}", controllers.DeletePhoneCase).Methods("DELETE")

	return router
}
