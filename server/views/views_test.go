package views

import (
	"bytes"
	"context"
	"dickens/auth"
	"dickens/database"
	"dickens/database/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	env database.Config
	// ctx context.Context
)

func TestUserCreate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(DB, "../../database/migrations")

	// Create a request to pass to the handler
	var Input struct {
		Username  string `json:"username"`
		Image     string `json:"image"`
		Bio       string `json:"bio"`
		Surname   string `json:"surname"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	hash, err := auth.HashPassword("password")

	if err != nil {
		t.Fatal(err)
	}

	Input.Username = "testusername"
	Input.Bio = "testbio"
	Input.Email = "test@example.com"
	Input.Firstname = "test"
	Input.Image = "/test.png"
	Input.Lastname = "example"
	Input.Password = hash
	Input.Surname = "dickens"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPost, UseCreate.Route, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	UseCreate.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		User    models.UserCreateRow `json:"user"`
		Profile models.Profile       `json:"profile"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.User.Email != Input.Email {
		t.Errorf("%s and %s are not the same", output.User.Email, Input.Email)
	}

	if output.User.FirstName != Input.Firstname {
		t.Errorf("%s and %s are not the same", output.User.FirstName, Input.Firstname)
	}

	if output.User.LastName != Input.Lastname {
		t.Errorf("%s and %s are not the same", output.User.LastName, Input.Lastname)
	}

	if output.User.Surname != Input.Surname {
		t.Errorf("%s and %s are not the same", output.User.Surname, Input.Surname)
	}

	if output.Profile.Bio.String != Input.Bio {
		t.Errorf("%s and %s are not the same", output.Profile.Bio.String, Input.Bio)
	}

	if output.Profile.Image.String != Input.Image {
		t.Errorf("%s and %s are not the same", output.Profile.Image.String, Input.Image)
	}

	if output.Profile.Username != Input.Username {
		t.Errorf("%s and %s are not the same", output.Profile.Username, Input.Username)
	}
}

func TestUserProfileRead(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s1", ProfileRead.Route), nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	ProfileRead.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		User    models.UserReadRow `json:"user"`
		Profile models.Profile     `json:"profile"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func TestUserProfileList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodGet, ProfileList.Route, nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	ProfileList.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	type useroutput struct {
		User    models.UserListRow `json:"user"`
		Profile models.Profile     `json:"profile"`
	}

	var output []useroutput

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func TestUserUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler
	var Input struct {
		Surname   string `json:"surname"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	Input.Firstname = "testupdated"
	Input.Lastname = "exampleupdated"
	Input.Surname = "dickensupdated"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s1", UserUpdate.Route), reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	UserUpdate.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		User models.UserUpdateRow `json:"user"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.User.FirstName != Input.Firstname {
		t.Errorf("%s and %s are not the same", output.User.FirstName, Input.Firstname)
	}

	if output.User.LastName != Input.Lastname {
		t.Errorf("%s and %s are not the same", output.User.LastName, Input.Lastname)
	}

	if output.User.Surname != Input.Surname {
		t.Errorf("%s and %s are not the same", output.User.Surname, Input.Surname)
	}
}

func TestProfileUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler
	var Input struct {
		Bio   string `json:"bio"`
		Image string `json:"image"`
	}

	Input.Bio = "testbioupdated"
	Input.Image = "/testupdated.png"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s1", ProfileUpdate.Route), reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	ProfileUpdate.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		Profile models.Profile `json:"profile"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.Profile.Bio.String != Input.Bio {
		t.Errorf("%s and %s are not the same", output.Profile.Bio.String, Input.Bio)
	}

	if output.Profile.Image.String != Input.Image {
		t.Errorf("%s and %s are not the same", output.Profile.Image.String, Input.Image)
	}
}

func TestCategoryCreate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(DB, "../../database/migrations")

	// Create a request to pass to the handler
	var Input struct {
		UserID int64  `json:"userid"`
		Name   string `json:"name"`
	}

	Input.UserID = int64(1)
	Input.Name = "testname"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPost, CategoryCreate.Route, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	CategoryCreate.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	// Decode the JSON response
	var output struct {
		Category models.Category `json:"category"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.Category.UserID.Int64 != Input.UserID {
		t.Errorf("%v and %v are not the same", output.Category.UserID.Int64, Input.UserID)
	}

	if output.Category.Name != Input.Name {
		t.Errorf("%s and %s are not the same", output.Category.Name, Input.Name)
	}
}

func TestCategoryRead(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s1", CategoryRead.Route), nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	CategoryRead.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		Category models.Category `json:"category"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func TestCategoryList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodGet, CategoryList.Route, nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	CategoryList.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		Categories []models.CategoryBlogListRow `json:"categories"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func TestCategoryUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler
	var Input struct {
		Name string `json:"name"`
	}

	Input.Name = "testnameupdated"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s1", CategoryUpdate.Route), reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	CategoryUpdate.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	var output struct {
		Category models.Category `json:"category"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.Category.Name != Input.Name {
		t.Errorf("%s and %s are not the same", output.Category.Name, Input.Name)
	}
}

func TestBlogCreate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(DB, "../../database/migrations")

	// Create a request to pass to the handler
	var Input struct {
		UserID     int64   `json:"userid"`
		Categories []int64 `json:"categories"`
		Title      string  `json:"title"`
		Body       string  `json:"body"`
	}

	Input.UserID = 1
	Input.Categories = []int64{1}
	Input.Title = "testtitle"
	Input.Body = "testbody"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPost, BlogCreate.Route, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	BlogCreate.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	// Decode the JSON response
	var output struct {
		Blog models.Blog `json:"blog"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.Blog.UserID.Int64 != Input.UserID {
		t.Errorf("%v and %v are not the same", output.Blog.UserID.Int64, Input.UserID)
	}

	if output.Blog.Title != Input.Title {
		t.Errorf("%s and %s are not the same", output.Blog.Title, Input.Title)
	}

	if output.Blog.Body != Input.Body {
		t.Errorf("%s and %s are not the same", output.Blog.Body, Input.Body)
	}
}

func TestBlogRead(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s1", BlogRead.Route), nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	BlogRead.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		Blog models.Blog `json:"blog"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func TestBlogList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodGet, BlogList.Route, nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	BlogList.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		Blogs []models.Blog `json:"blogs"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func TestBlogUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler
	var Input struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	Input.Title = "testtitleupdated"
	Input.Body = "testbodyupdated"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s1", BlogUpdate.Route), reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	BlogUpdate.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	var output struct {
		Blog models.BlogUpdateRow `json:"blog"`
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.Blog.Title != Input.Title {
		t.Errorf("%s and %s are not the same", output.Blog.Title, Input.Title)
	}

	if output.Blog.Body != Input.Body {
		t.Errorf("%s and %s are not the same", output.Blog.Body, Input.Body)
	}
}

func TestBlogDelete(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s1", BlogDelete.Route), nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	BlogDelete.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	queries := models.New(DB)
	ctx := context.Background()

	blogs, err := queries.BlogList(ctx)

	if err != nil {
		t.Fatal(err)
	}

	if len(blogs) != 0 {
		t.Fatalf("The blog was not deleted: %v", blogs)
	}

	categoryBlogs, err := queries.CategoryBlogList(ctx)

	if err != nil {
		t.Fatal(err)
	}

	if len(categoryBlogs) != 0 {
		t.Fatalf("The category blogs was not deleted: %v", categoryBlogs)
	}
}

func TestCategoryDelete(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s1", CategoryDelete.Route), nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	CategoryDelete.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	queries := models.New(DB)
	ctx := context.Background()

	categories, err := queries.CategoryList(ctx)

	if err != nil {
		t.Fatal(err)
	}

	if len(categories) != 0 {
		t.Fatalf("The categories was not deleted: %v", categories)
	}

	categoryBlogs, err := queries.CategoryBlogList(ctx)

	if err != nil {
		t.Fatal(err)
	}

	if len(categoryBlogs) != 0 {
		t.Fatalf("The category blogs was not deleted: %v", categoryBlogs)
	}
}

func TestGetPubKey(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(DB, "../../database/migrations")

	// Create a request to pass to the handler
	var Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	Input.Email = "test@example.com"
	Input.Password = "password"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPost, GetPubKey.Route, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	GetPubKey.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	// Decode the JSON response
	var output struct {
		Pubkey string
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.Pubkey == "" {
		t.Fatalf("pub key was not sent: %v", err)
	}
}

func TestDeletePubKey(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(DB, "../../database/migrations")

	// Create a request to pass to the handler
	var Input struct {
		ID int64 `json:"id"`
	}

	Input.ID = 1

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPost, DeletePubKey.Route, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	DeletePubKey.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}
}

func TestChangeEmail(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(DB, "../../database/migrations")

	// Create a request to pass to the handler
	var Input struct {
		ID    int64  `json:"id"`
		Email string `json:"newemail"`
	}

	Input.ID = 1
	Input.Email = "testupdated@example.com"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPut, ChangeEmail.Route, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	ChangeEmail.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	var output struct {
		User models.UserUpdateEmailRow
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if output.User.Email != Input.Email {
		t.Errorf("%v and %v are not the same", output.User.Email, Input.Email)
	}
}

func TestResetPassword(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(DB, "../../database/migrations")

	// Create a request to pass to the handler
	var Input struct {
		ID       int64  `json:"id"`
		Password string `json:"password"`
	}

	Input.ID = 1
	Input.Password = "passwordupdated"

	bodyBytes, err := json.Marshal(Input)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	reqBody := bytes.NewBuffer(bodyBytes)

	req := httptest.NewRequest(http.MethodPut, ResetPassword.Route, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	ResetPassword.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	var output struct {
		User models.UserUpdatePasswordRow
	}

	err = json.NewDecoder(rec.Body).Decode(&output)
	if err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func TestUserDelete(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	DB = database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer DB.Close()

	if DB == nil {
		t.Fatal("Failed to connect to database")
	}

	// Create a request to pass to the handler

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s1", UserDelete.Route), nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler
	UserDelete.Handler.ServeHTTP(rec, req)

	// Get the recorded response
	res := rec.Result()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		bodyString := strings.TrimSpace(rec.Body.String())
		t.Errorf("expected status OK; got %v\n%s", res.Status, bodyString)
	}

	queries := models.New(DB)
	ctx := context.Background()

	profiles, err := queries.ProfileList(ctx)

	if err != nil {
		t.Fatal(err)
	}

	if len(profiles) != 0 {
		t.Fatalf("The profile was not deleted: %v", profiles)
	}

	database.DropMigrate(DB, "../../database/migrations")
}

// database.DropMigrate(DB, "../../database/migrations")
