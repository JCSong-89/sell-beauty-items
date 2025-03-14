package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	authService "sell-beauty-items/internal/auth/service"
	productRepository "sell-beauty-items/internal/products/repository"
	productService "sell-beauty-items/internal/products/service"
	userRepository "sell-beauty-items/internal/users/repository"
	"sell-beauty-items/pkg/middlewares"
	"sell-beauty-items/pkg/utils"
)

// templates is a global variable that holds all of the parsed templates
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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)
	templatesPath := filepath.Join(dir, "static", "*.html")
	templates := template.Must(template.ParseGlob(templatesPath))

	// Register the handlers for the server
	authServiceObject := authService.AuthService{
		T: templates,
		U: &userRepository.UserRepository{},
		P: &productRepository.ProductRepository{},
	}
	productServiceObject := productService.ProductService{
		Repository: &productRepository.ProductRepository{},
	}

	staticDir := filepath.Join(dir, "static")
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	handler := middlewares.CookieMiddleware(mux)

	mux.HandleFunc("/", authServiceObject.IndexHandler)
	mux.HandleFunc("/login", authServiceObject.LoginHandler)
	mux.HandleFunc("/logout", authServiceObject.LogoutHandler)
	mux.HandleFunc("/shop", productServiceObject.GetAllProducts)

	handler = middlewares.RequestMiddleware(mux)
	alias := ""
	aliasFlug := false

	go func() {
		for {
			select {
			case req := <-middlewares.RequestChan:
				alias, aliasFlug = utils.ShopURLStringController(req.URL.String())
				/* 싱글톤으로 만들어서 해당 싱글톤 객체로 처리도록 변경*/
				fmt.Printf("URL: %s\nMethod: %s\n\n", req.URL.String(), req.Method)
			}
		}
	}()

	if aliasFlug {
		mux.HandleFunc("/shop/"+alias, productServiceObject.GetHandler)
	}
	// Start the server and listen for incoming connections
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
