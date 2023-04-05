package middlewares

import "net/http"

var RequestChan = make(chan *http.Request, 10)

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청을 채널에 전달
		RequestChan <- r
		// 요청을 다음 핸들러에게 전달
		next.ServeHTTP(w, r)
	})
}
