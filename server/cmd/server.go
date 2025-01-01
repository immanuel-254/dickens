package main

import (
	"dickens/database"
	"dickens/server/middlewares/helmet"
	"dickens/server/views"
	"fmt"
	"net/http"
	"time"
)

var (
	env database.Config
)

func init() {
	err := database.LoadEnv("env.json", &env)

	if err != nil {
		panic(err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	// migrate to DB
	database.Migrate(db, "database/migrations")
}

func main() {
	mux := http.NewServeMux()

	allviews := []views.View{
		views.BlogCreate,
		views.BlogDelete,
		views.BlogList,
		views.BlogRead,
		views.BlogUpdate,
		views.CategoryCreate,
		views.CategoryDelete,
		views.CategoryRead,
		views.CategoryList,
		views.CategoryUpdate,
		views.ChangeEmail,
		views.DeletePubKey,
		views.GetPubKey,
		views.ProfileList,
		views.ProfileRead,
		views.ProfileUpdate,
		views.ResetPassword,
		views.UseCreate,
		views.UserDelete,
		views.UserUpdate,
	}
	views.Routes(mux, allviews)

	server := &http.Server{
		Addr:         ":8080",                               // Custom port
		Handler:      helmet.New(helmet.ConfigDefault)(mux), // Attach the mux as the handler
		ReadTimeout:  10 * time.Second,                      // Set read timeout
		WriteTimeout: 10 * time.Second,                      // Set write timeout
		IdleTimeout:  30 * time.Second,                      // Set idle timeout
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
