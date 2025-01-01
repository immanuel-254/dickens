package models

import (
	"context"
	"database/sql"
	"dickens/database"
	"testing"
	"time"
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

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	// migrate to DB
	database.Migrate(db, "../migrations")

	queries := New(db)
	ctx := context.Background()

	_, err = queries.UserCreate(ctx, UserCreateParams{
		Surname:   "dickens",
		FirstName: "test",
		LastName:  "example",
		Email:     "test@example.com",
		Password:  "password",
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.UserList(ctx)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserRead(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.UserRead(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	user, err := queries.UserUpdate(ctx, UserUpdateParams{
		ID:        1,
		FirstName: "testupdated",
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if user.FirstName != "testupdated" {
		t.Errorf("%s and testupdated are not the same", user.FirstName)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserUpdatePassword(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.UserUpdatePassword(ctx, UserUpdatePasswordParams{
		ID:        1,
		Password:  "passwordupdated",
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserUpdateEmail(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.UserUpdateEmail(ctx, UserUpdateEmailParams{
		ID:        1,
		Email:     "testupdated@example.com",
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestCategoryCreate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.CategoryCreate(ctx, CategoryCreateParams{
		UserID:    sql.NullInt64{Int64: 1, Valid: true},
		Name:      "Name",
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestCategoryList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.CategoryList(ctx)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCategoryRead(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.CategoryRead(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCategoryUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.CategoryUpdate(ctx, CategoryUpdateParams{
		ID:        1,
		Name:      "NameUpdated",
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestBlogCreate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	blog, err := queries.BlogCreate(ctx, BlogCreateParams{
		UserID:    sql.NullInt64{Int64: 1, Valid: true},
		Title:     "TestBlog",
		Body:      "TestBody",
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}

	err = queries.AssignBlogToCategory(ctx, AssignBlogToCategoryParams{
		BlogID:     blog.ID,
		CategoryID: 1,
		CreatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestBlogList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.BlogList(ctx)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBlogCategoryList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.BlogCategoriesList(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCategoryBlogList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.CategoryBlogList(ctx)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBlogRead(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.BlogRead(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBlogUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.BlogUpdate(ctx, BlogUpdateParams{
		ID:    1,
		Title: "TestBlogUpdated",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestProfileCreate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.ProfileCreate(ctx, ProfileCreateParams{
		UserID:    sql.NullInt64{Int64: 1, Valid: true},
		Username:  "TestUsername",
		Image:     sql.NullString{String: "TestImage", Valid: true},
		Bio:       sql.NullString{String: "TestBio", Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestProfileList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.ProfileList(ctx)

	if err != nil {
		t.Fatal(err)
	}
}

func TestProfileRead(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.ProfileRead(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestProfileUpdate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.ProfileUpdate(ctx, ProfileUpdateParams{
		ID:        1,
		Username:  "TestUsernameUpdated",
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestLogCreate(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	err = queries.LogCreate(ctx, LogCreateParams{
		DbTable:   "TestTable",
		Action:    "TestAction",
		ObjectID:  1,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestLogList(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	_, err = queries.LogList(ctx)

	if err != nil {
		t.Fatal(err)
	}
}

func TestProfileDelete(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	err = queries.ProfileDelete(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBlogDelete(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	err = queries.BlogDelete(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCategoryDelete(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	err = queries.CategoryDelete(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDelete(t *testing.T) {
	err := database.LoadEnv("../../bin/env.json", &env)

	if err != nil {
		t.Fatalf("%v", err)
	}

	db := database.ConnectToDB(env.TURSO_DATABASE_url, env.TURSO_AUTH_TOKEN)
	defer db.Close()

	if db == nil {
		t.Fatal("Failed to connect to database")
	}

	queries := New(db)
	ctx := context.Background()

	err = queries.UserDelete(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}

	database.DropMigrate(db, "../migrations")
}
