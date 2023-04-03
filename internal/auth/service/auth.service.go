package service

import (
	"fmt"
	"html/template"
	"net/http"
	productRepository "sell-beauty-items/internal/products/repository"
	userRepository "sell-beauty-items/internal/users/repository"
	"time"
)

type AuthService struct {
	T *template.Template
	P *productRepository.ProductRepository
	U *userRepository.UserRepository
}

func (a *AuthService) IndexHandler(w http.ResponseWriter, r *http.Request) {
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
		"Products":      a.P.GetProducts(),
		"Theme":         cookieTheme.Value,
		"Authenticated": authenticated,
	}

	err = a.T.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *AuthService) LoginHandler(w http.ResponseWriter, r *http.Request) {
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
		for _, user := range a.U.GetUsers() {
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
		if err := a.T.ExecuteTemplate(w, "login.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// If the form was not submitted, render the login form
	data := map[string]interface{}{
		"Title": "My Awesome Site",
	}
	if err := a.T.ExecuteTemplate(w, "login.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *AuthService) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Delete the authenticated cookie and redirect to the homepage
	cookie := http.Cookie{
		Name:    "authenticated",
		Value:   "false",
		Expires: time.Now().Add(-24 * time.Hour),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
