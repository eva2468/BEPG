package routes

import (
	"github.com/eva2468/bepg/pkg/controllers"
	"github.com/gorilla/mux"
)

var PaymentRoutes = func(router *mux.Router) {

	router.HandleFunc("/bepg_checkbin", controllers.Bepg_checkbin).Methods("POST")
	//router.HandleFunc("/application/json", controllers.GetCardAuth).Methods("POST")
	router.HandleFunc("/generate_otp", controllers.Generate_otp).Methods("POST")
	router.HandleFunc("/verify_otp", controllers.Verify_otp).Methods("POST")
	router.HandleFunc("/bepg_authorize", controllers.Verify_otp).Methods("POST")
	router.HandleFunc("/bepg_reverse", controllers.Verify_otp).Methods("POST")
	router.HandleFunc("/check_EMIAvailability", controllers.Check_EMIAvailability).Methods("POST")
	router.HandleFunc("/initiate_SIRegistration", controllers.Initiate_SIRegistration).Methods("POST")
	router.HandleFunc("/complete_SIRegistration", controllers.Verify_otp).Methods("POST")
	router.HandleFunc("/deregister_SI", controllers.Verify_otp).Methods("POST")
}
