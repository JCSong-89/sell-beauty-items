package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Product represents a product on the site
type Product struct {
	Name     string
	Price    int
	ImageURL string
}

// products is a list of products on the site
var products = []Product{
	{Name: "Product A", Price: 100, ImageURL: "/static/img/product-a.jpg"},
	{Name: "Product B", Price: 200, ImageURL: "/static/img/product-b.jpg"},
	{Name: "Product C", Price: 300, ImageURL: "/static/img/product-c.jpg"},
	{Name: "Product D", Price: 400, ImageURL: "/static/img/product-d.jpg"},
	{Name: "Product E", Price: 500, ImageURL: "/static/img/product-e.jpg"},
	{Name: "Product F", Price: 600, ImageURL: "/static/img/product-f.jpg"},
}

type User struct {
	Username string
	Password string
}

// users is a slice of User objects that will be used for authentication
var users = []User{
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
}

// templates is a global variable that holds all of the parsed templates
var templates = template.Must(template.ParseGlob("static/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Get the theme from the cookie, or set it to "light" if it doesn't exist
	cookieTheme, err := r.Cookie("theme")
	if err != nil {
		if err == http.ErrNoCookie {
			cookieTheme = &http.Cookie{Name: "theme", Value: "light", MaxAge: 60 * 60 * 24 * 365}
			http.SetCookie(w, cookieTheme)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Create a data map to pass to the template

	authenticated := false
	cookieAuthenticated, err := r.Cookie("authenticated")
	if err == nil && cookieAuthenticated.Value == "true" {
		authenticated = true
	}

	data := map[string]interface{}{
		"Title":         "My Awesome Site",
		"Products":      products,
		"Theme":         cookieTheme.Value,
		"Authenticated": authenticated,
	}

	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// If the user is already authenticated, redirect them to the homepage

	fmt.Printf("loginHandler %v\n", r.Method)
	cookie, err := r.Cookie("authenticated")
	if err == nil && cookie.Value == "true" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Printf("loginHandler %v\n", r.Body)

	// If the form was submitted, attempt to authenticate the user
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		fmt.Println("username: ", username)
		fmt.Println("password: ", password)

		// Check if the user exists in the users slice
		for _, user := range users {
			if user.Username == username && user.Password == password {
				// If the user exists, set the authenticated cookie and redirect to the homepage
				cookie := http.Cookie{
					Name:    "authenticated",
					Value:   "true",
					Expires: time.Now().Add(24 * time.Hour),
				}
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

		// If the user doesn't exist, show an error message
		data := map[string]interface{}{
			"Title": "My Awesome Site",
			"Error": "Invalid username or password",
		}
		if err := templates.ExecuteTemplate(w, "login.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// If the form was not submitted, render the login form
	data := map[string]interface{}{
		"Title": "My Awesome Site",
	}
	if err := templates.ExecuteTemplate(w, "login.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Delete the authenticated cookie and redirect to the homepage
	cookie := http.Cookie{
		Name:    "authenticated",
		Value:   "false",
		Expires: time.Now().Add(-24 * time.Hour),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)

	// Start the server and listen for incoming connections
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
