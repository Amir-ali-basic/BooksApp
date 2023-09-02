package main

import (
	"fmt"
	"net/http"
	"time"
	"vue-api/internal/data"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// routes generates our routes and attaches them to handlers, using the chi router
// note that we return type http.Handler, and not *chi.Mux; since chi.Mux satisfies
// the interface requirements for http.Handler, it makes sense to return the type
// that is part of the standard library.
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)

	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
		var users data.User
		all, err := users.GetAll()

		fmt.Println("get all uesrs")

		if err != nil {
			app.errorLog.Println(err)
			return
		}
		app.writeJSON(w, http.StatusOK, all)
	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = data.User{
			Email:     "you@there.com",
			FirstName: "You",
			LastName:  "There",
			Password:  "password",
		}

		app.infoLog.Println("Adding user...")

		id, err := app.userModels.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.infoLog.Println("Got back id of", id)
		newUser, _ := app.userModels.GetById(id)
		app.writeJSON(w, http.StatusOK, newUser)
	})

	mux.Get("/test-generate-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.tokenModels.GenerateToken(2, 60*time.Minute)

		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.Email = "admin@example.com"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-save-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.tokenModels.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		user, err := app.userModels.GetById(2)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.UserID = user.ID
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.InsertToken(*token, *user)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-validate-token", func(w http.ResponseWriter, r *http.Request) {
		tokenToValidate := r.URL.Query().Get("token")
		valid, err := app.tokenModels.ValidToken(tokenToValidate)
		if err != nil {
			app.errorJSON(w, err, http.StatusUnauthorized) // Return an unauthorized status on validation failure
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "Token is valid",
			Data:    valid,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})
	mux.Get("/get-by-token", func(w http.ResponseWriter, r *http.Request) {
		userToken := r.URL.Query().Get("token")
		fmt.Println("uwer token", userToken)

		// Call the GetByToken function to retrieve the token details
		retrievedToken, err := app.tokenModels.GetByToken(userToken)
		if err != nil {
			// If there's an error or no matching token found, return an unauthorized status
			app.errorJSON(w, err, http.StatusUnauthorized)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "Token is valid",
			Data:    retrievedToken, // Return the retrieved token details
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	return mux
}
