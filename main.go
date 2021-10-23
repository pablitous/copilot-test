//start an api with jwt auth
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/profile", controllers.Profile).Methods("GET")
	router.HandleFunc("/api/user/all", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/get", controllers.FollowUser).Methods("POST")
	//add jwt middleware
	router.Use(middleware.JwtAuthentication)
	log.Fatal(http.ListenAndServe(":8000", router))

}
