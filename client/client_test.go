package main

import (
	"dickens/auth"
	"dickens/database"
	"dickens/server/middlewares/helmet"
	"dickens/server/views"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AllViews() []views.View {

	return []views.View{
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
}

var (
	env database.Config
	// ctx context.Context
)

func TestSignUp(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	hash, err := auth.HashPassword("password")

	if err != nil {
		t.Fatal(err)
	}

	data := map[string]string{
		"username":  "testusername",
		"bio":       "testbio",
		"email":     "test@example.com",
		"firstname": "test",
		"image":     "/test.png",
		"lastname":  "example",
		"password":  hash,
		"surname":   "dickens",
	}

	response := SignUp(server.URL, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		t.Fatal(response["body"])
	}
}

func TestLogin(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"email":    "test@example.com",
		"password": "password",
	}

	response := LogIn(server.URL, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		t.Fatal(response["body"])
	}
}

func TestLogout(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"id": "1",
	}

	response := LogOut(server.URL, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		t.Fatal(response["body"])
	}
}

func TestChangeEmail(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"id":       "1",
		"newemail": "testupdated@example.com",
	}

	response := ChangeEmail(server.URL, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		t.Fatal(response["body"])
	}
}

func TestResetPassword(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"id":       "1",
		"password": "passwordupdated",
	}

	response := ResetPassword(server.URL, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		t.Fatal(response["body"])
	}
}

func TestReadProfile(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := ReadProfile(server.URL, 1)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		t.Fatal(views.ProfileRead.Route)
		t.Fatal(response["body"])
	}
}

func TestListProfile(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := ListProfile(server.URL)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		t.Fatal(views.ProfileRead.Route)
		t.Fatal(response["body"])
	}
}

func TestUpdateUser(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"firstname": "testupdated",
		"lastname":  "exampleupdated",
		"surname":   "dickensupdated",
	}

	response := UpdateUser(server.URL, 1, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.UserUpdate.Route)
		t.Fatal(response["body"])
	}
}

func TestUpdateProfile(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"username": "testusernameupdated",
		"image":    "/testupdated.png",
		"bio":      "testbioupdated",
	}

	response := UpdateProfile(server.URL, 1, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.UserUpdate.Route)
		t.Fatal(response["body"])
	}
}

func TestCreateCategory(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"userid": "1",
		"name":   "testname",
	}

	response := CreateCategory(server.URL, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.UserUpdate.Route)
		t.Fatal(response["body"])
	}
}

func TestReadCategory(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := ReadCategory(server.URL, 1)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestListCategory(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := ListCategory(server.URL)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestUpdatedCategory(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"name": "testnameupdated",
	}

	response := UpdateCategory(server.URL, 1, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestCreateBlog(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"userid":     "1",
		"title":      "testtitle",
		"body":       "testbody",
		"categories": "1,",
	}

	response := CreateBlog(server.URL, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestReadBlog(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := ReadBlog(server.URL, 1)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestListBlog(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := ListBlog(server.URL)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestUpdateBlog(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	data := map[string]string{
		"title": "testtitleupdated",
		"body":  "testbodyupdated",
	}

	response := UpdateBlog(server.URL, 1, data)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestDeleteBlog(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := DeleteBlog(server.URL, 1)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}
}

func TestDeleteCategory(t *testing.T) {
	err := database.LoadEnv("../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	views.DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer views.DB.Close()

	if views.DB == nil {
		t.Fatal("Failed to connect to database")
	}

	database.Migrate(views.DB, "../database/migrations")

	mux := http.NewServeMux()

	allviews := AllViews()
	views.Routes(mux, allviews)

	server := httptest.NewServer(helmet.New(helmet.ConfigDefault)(mux))
	defer server.Close()

	response := DeleteCategory(server.URL, 1)

	if response == nil {
		t.Fatal("Something is broken")
	}

	if response["error"] != "" {
		t.Fatal(response["error"])
	}

	if response["status"] != "200 OK" {
		// t.Fatal(views.CategoryRead.Route)
		t.Fatal(response["body"])
	}

	database.DropMigrate(views.DB, "../database/migrations")
}
