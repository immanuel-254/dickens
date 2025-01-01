package views

import (
	"context"
	"database/sql"
	"dickens/database/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const UserRouteGroup = "/user"

var (
	UseCreate = View{
		Route: fmt.Sprintf("%s/create", UserRouteGroup),
		/* Middlewares: []func(http.Handler) http.Handler{
			helmet.New(helmet.ConfigDefault),
		}, */
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Entities To be Created; User and Profile

			var data map[string]string
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			queries := models.New(DB)
			ctx := context.Background()

			user, err := queries.UserCreate(ctx, models.UserCreateParams{
				Surname:   data["surname"],
				FirstName: data["firstname"],
				LastName:  data["lastname"],
				Password:  data["password"],
				Email:     data["email"],
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			profile, err := queries.ProfileCreate(ctx, models.ProfileCreateParams{
				UserID:    sql.NullInt64{Int64: user.ID, Valid: true},
				Username:  data["username"],
				Image:     sql.NullString{String: data["image"], Valid: true},
				Bio:       sql.NullString{String: data["bio"], Valid: true},
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				User    models.UserCreateRow `json:"user"`
				Profile models.Profile       `json:"profile"`
			}

			output.User = user
			output.Profile = profile

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPost},
	}

	UserUpdate = View{
		Route: fmt.Sprintf("%s/update/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/update/", UserRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err)
				return
			}

			var data map[string]string
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			var userupdate models.UserUpdateParams
			surname, surnameExists := data["surname"]
			firstname, firstnameExists := data["firstname"]
			lastname, lastnameExists := data["lastname"]

			userupdate.ID = int64(id)
			userupdate.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

			if !surnameExists && !firstnameExists && !lastnameExists {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNoContent)
				json.NewEncoder(w).Encode(errors.New("empty data"))
				return
			}

			if surnameExists {
				userupdate.Surname = surname
			}
			if firstnameExists {
				userupdate.FirstName = firstname
			}
			if lastnameExists {
				userupdate.LastName = lastname
			}

			// Entities To Update; User
			queries := models.New(DB)
			ctx := context.Background()

			user, err := queries.UserUpdate(ctx, userupdate)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				User models.UserUpdateRow `json:"user"`
			}

			output.User = user

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(output)

		}),
		Methods: []string{http.MethodPut},
	}

	UserDelete = View{
		Route: fmt.Sprintf("%s/delete/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := strings.TrimPrefix(r.URL.Path, fmt.Sprintf("%s/delete/", UserRouteGroup))
			idStr = strings.TrimLeft(idStr, "/")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}
			// Entities To Delete; User, Profile

			queries := models.New(DB)
			ctx := context.Background()

			err = queries.UserDelete(ctx, int64(id))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}),
		Methods: []string{http.MethodDelete},
	}
)
