package views

import (
	"context"
	"database/sql"
	"dickens/database/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const CategoryRouteGroup = "/category"

var (
	CategoryCreate = View{
		Route: fmt.Sprintf("%s/create", CategoryRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Entities To be Created; Category
			var data map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			queries := models.New(DB)
			ctx := context.Background()

			var input struct {
				userid int64
			}

			if _, ok := data["userid"].(string); ok {
				id, err := strconv.ParseInt(data["userid"].(string), 10, 64)

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(err)
					return
				}

				input.userid = id
			}

			if _, ok := data["userid"].(float64); ok {
				input.userid = int64(data["userid"].(float64))
			}

			category, err := queries.CategoryCreate(ctx, models.CategoryCreateParams{
				UserID:    sql.NullInt64{Int64: input.userid, Valid: true},
				Name:      data["name"].(string),
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Category models.Category `json:"category"`
			}

			output.Category = category

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPost},
	}

	CategoryRead = View{
		Route: fmt.Sprintf("%s/read/", CategoryRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/read/", CategoryRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(err)
				return
			}
			// Entities To Read; Category
			queries := models.New(DB)
			ctx := context.Background()

			category, err := queries.CategoryRead(ctx, int64(id))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Category models.Category `json:"category"`
			}

			output.Category = category

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodGet},
	}

	CategoryList = View{
		Route: fmt.Sprintf("%s/list", CategoryRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Entities To List; Category
			queries := models.New(DB)
			ctx := context.Background()

			categories, err := queries.CategoryBlogList(ctx)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Categories []models.CategoryBlogListRow `json:"categories"`
			}

			output.Categories = categories

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodGet},
	}

	CategoryUpdate = View{
		Route: fmt.Sprintf("%s/update/", CategoryRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/update/", CategoryRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(err)
				return
			}
			// Entities To Update; Category
			var data map[string]string
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			queries := models.New(DB)
			ctx := context.Background()

			category, err := queries.CategoryUpdate(ctx, models.CategoryUpdateParams{
				ID:        int64(id),
				Name:      data["name"],
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Category models.Category `json:"category"`
			}

			output.Category = category

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPut},
	}

	CategoryDelete = View{
		Route: fmt.Sprintf("%s/delete/", CategoryRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Entities To Delete; Category
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/delete/", CategoryRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(err)
				return
			}

			queries := models.New(DB)
			ctx := context.Background()

			err = queries.CategoryDelete(ctx, int64(id))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			// get all categories
			categories, err := queries.CategoryBlogList(ctx)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			// delete many to many relations
			for _, category := range categories {
				if category.CategoryID == int64(id) {
					err = queries.CategoryBlogDelete(ctx, models.CategoryBlogDeleteParams{
						BlogID:     category.BlogID,
						CategoryID: int64(id),
					})

					if err != nil {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusInternalServerError)
						json.NewEncoder(w).Encode(err)
						return
					}
				}
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		}),
		Methods: []string{http.MethodDelete},
	}
)
