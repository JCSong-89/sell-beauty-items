package middlewares

import "net/http"

func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		authenticated := "false"
		cookieAuthenticated, err := r.Cookie("authenticated")
		if err == nil && cookieAuthenticated.Value == "true" {
			authenticated = "true"
		}

		r.Header.Set("authenticated", authenticated)
		r.Header.Set("cookieTheme", cookieTheme.Value)

		println(r.Header)

		next.ServeHTTP(w, r)
	})
}
