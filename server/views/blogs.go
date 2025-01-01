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

const BlogRouteGroup = "/blog"

var (
	BlogCreate = View{
		Route: fmt.Sprintf("%s/create", BlogRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Entities To be Created; Blog
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

			blog, err := queries.BlogCreate(ctx, models.BlogCreateParams{
				UserID:    sql.NullInt64{Int64: input.userid, Valid: true},
				Title:     data["title"].(string),
				Body:      data["body"].(string),
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var categories []int64

			if _, ok := data["categories"].(string); ok {

				rawCategories := data["categories"].(string)
				parts := strings.Split(rawCategories, ",")

				for _, part := range parts {
					trimmedPart := strings.TrimSpace(part) // Remove any extra spaces
					if part == "" {
						continue // Skip empty parts (e.g., due to a trailing comma)
					}
					num, err := strconv.ParseInt(trimmedPart, 10, 64) // Convert to int64
					if err != nil {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusInternalServerError)
						json.NewEncoder(w).Encode(err)
						return
					}
					categories = append(categories, num)
				}
			}

			if _, ok := data["categories"].([]interface{}); ok {
				rawCategories := data["categories"].([]interface{})

				for i, v := range rawCategories {
					if num, ok := v.(float64); ok {
						categories[i] = int64(num)
					}
				}
			}

			for _, value := range categories {
				err = queries.AssignBlogToCategory(ctx, models.AssignBlogToCategoryParams{
					BlogID:     blog.ID,
					CategoryID: value,
					CreatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
				})

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(err)
					return
				}
			}

			var output struct {
				Blog models.Blog `json:"blog"`
			}

			output.Blog = blog

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPost},
	}

	BlogRead = View{
		Route: fmt.Sprintf("%s/read/", BlogRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/read/", BlogRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(err)
				return
			}

			// Entities To Read; Blog
			queries := models.New(DB)
			ctx := context.Background()

			blog, err := queries.BlogRead(ctx, int64(id))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Blog models.Blog `json:"blog"`
			}

			output.Blog = blog

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodGet},
	}

	BlogList = View{
		Route: fmt.Sprintf("%s/list", BlogRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Entities To Read; Blog, Category
			queries := models.New(DB)
			ctx := context.Background()

			blogs, err := queries.BlogList(ctx)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Blogs []models.Blog `json:"blogs"`
			}

			output.Blogs = blogs

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodGet},
	}

	BlogUpdate = View{
		Route: fmt.Sprintf("%s/update/", BlogRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/update/", BlogRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(err)
				return
			}

			// Entities To Update; Blog, Category
			var data map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			queries := models.New(DB)
			ctx := context.Background()

			blog, err := queries.BlogUpdate(ctx, models.BlogUpdateParams{
				ID:        int64(id),
				Title:     data["title"].(string),
				Body:      data["body"].(string),
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Blog models.BlogUpdateRow `json:"blog"`
			}

			output.Blog = blog

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPut},
	}

	BlogDelete = View{
		Route: fmt.Sprintf("%s/delete/", BlogRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/delete/", BlogRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(err)
				return
			}

			// Entities To Delete; Blog
			queries := models.New(DB)
			ctx := context.Background()

			err = queries.BlogDelete(ctx, int64(id))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			// get blog categories
			categories, err := queries.BlogCategoriesList(ctx, int64(id))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			// delete many to many relations
			for _, category := range categories {
				err = queries.CategoryBlogDelete(ctx, models.CategoryBlogDeleteParams{
					BlogID:     int64(id),
					CategoryID: category.CategoryID,
				})

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(err)
					return
				}

			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		}),
		Methods: []string{http.MethodDelete},
	}
)
