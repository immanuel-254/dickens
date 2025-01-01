package views

import (
	"context"
	"database/sql"
	"dickens/auth"
	"dickens/cyber"
	"dickens/database/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// const AuthRouteGroup = "/auth"

var (
	GetPubKey = View{
		Route: "/getpubkey",
		/* Middlewares: []func(http.Handler) http.Handler{
			helmet.New(helmet.ConfigDefault),
		}, */
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var data map[string]string
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			// Entities To Read; User
			queries := models.New(DB)
			ctx := context.Background()

			user, err := queries.UserLoginRead(ctx, data["email"])

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			check := auth.CheckPasswordHash(data["password"], user.Password)

			if !check {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				err = errors.New("invalid password")
				json.NewEncoder(w).Encode(err)
				return
			}

			err = cyber.CreateRSAKeyFile(fmt.Sprintf("/user/%s", user.Email), "priv.key")

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			priv, err := cyber.LoadPrivateKey(fmt.Sprintf("/user/%s/priv.key", user.Email))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				Pubkey string
			}

			pubkey, err := cyber.PublicKeyPem(&priv.PublicKey)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			output.Pubkey = pubkey

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPost},
	}

	DeletePubKey = View{
		Route: "/deletepubkey",
		/* Middlewares: []func(http.Handler) http.Handler{
			helmet.New(helmet.ConfigDefault),
		}, */
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var data map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			// Entities To Read; User
			queries := models.New(DB)
			ctx := context.Background()

			var input struct {
				id int64
			}

			if _, ok := data["id"].(string); ok {
				id, err := strconv.ParseInt(data["id"].(string), 10, 64)

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(err)
					return
				}

				input.id = id
			}

			if _, ok := data["id"].(float64); ok {
				input.id = int64(data["id"].(float64))
			}

			user, err := queries.UserRead(ctx, input.id)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			err = os.Remove(fmt.Sprintf("/user/%s/priv.key", user.Email))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		}),
		Methods: []string{http.MethodPost},
	}

	ChangeEmail = View{
		Route: "/change-email",
		/* Middlewares: []func(http.Handler) http.Handler{
			helmet.New(helmet.ConfigDefault),
		}, */
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var data map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			// Entities To Read; User
			queries := models.New(DB)
			ctx := context.Background()

			var input struct {
				id int64
			}

			if _, ok := data["id"].(string); ok {
				id, err := strconv.ParseInt(data["id"].(string), 10, 64)

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(err)
					return
				}

				input.id = id
			}

			if _, ok := data["id"].(float64); ok {
				input.id = int64(data["id"].(float64))
			}

			user, err := queries.UserUpdateEmail(ctx, models.UserUpdateEmailParams{
				ID:        input.id,
				Email:     data["newemail"].(string),
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				User models.UserUpdateEmailRow
			}

			output.User = user

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPut},
	}

	ResetPassword = View{
		Route: "/reset-password",
		/* Middlewares: []func(http.Handler) http.Handler{
			helmet.New(helmet.ConfigDefault),
		}, */
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var data map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			// Entities To Read; User
			queries := models.New(DB)
			ctx := context.Background()

			hash, err := auth.HashPassword(data["password"].(string))

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var input struct {
				id int64
			}

			if _, ok := data["id"].(string); ok {
				id, err := strconv.ParseInt(data["id"].(string), 10, 64)

				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(err)
					return
				}

				input.id = id
			}

			if _, ok := data["id"].(float64); ok {
				input.id = int64(data["id"].(float64))
			}

			user, err := queries.UserUpdatePassword(ctx, models.UserUpdatePasswordParams{
				ID:        input.id,
				Password:  hash,
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			})

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			var output struct {
				User models.UserUpdatePasswordRow
			}

			output.User = user

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}),
		Methods: []string{http.MethodPut},
	}
)
