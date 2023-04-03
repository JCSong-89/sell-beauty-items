package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	authService "sell-beauty-items/internal/auth/service"
	productRepository "sell-beauty-items/internal/products/repository"
	userRepository "sell-beauty-items/internal/users/repository"
)

// templates is a global variable that holds all of the parsed templates
var templates = template.Must(template.ParseGlob("../static/*.html"))

/*
func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body to get the username and password
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	token := &jwt.Token{}
	// Check if the username and password are correct
	if username == "admin" && password == "password123" {
		// If they are correct, generate a JWT and send it back to the client
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	w.Header().Set("Content-Type", "application/json")
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

		else {
			// If they are not correct, return a 401 Unauthorized error
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
}
*/

func main() {
	// Register the handlers for the server
	authServiceObject := authService.AuthService{
		T: templates,
		U: &userRepository.UserRepository{},
		P: &productRepository.ProductRepository{},
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}
	staticDir := filepath.Join(wd, "..", "static")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	http.HandleFunc("/", authServiceObject.IndexHandler)
	http.HandleFunc("/login", authServiceObject.LoginHandler)

	// Start the server and listen for incoming connections
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
