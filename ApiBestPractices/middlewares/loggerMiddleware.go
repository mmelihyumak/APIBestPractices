package middlewares

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

// ResponseWriterWrapper bir HTTP yanıtını yakalamamıza olanak tanır
type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK, body: &bytes.Buffer{}}
}

func (rw *ResponseWriterWrapper) Write(b []byte) (int, error) {
	rw.body.Write(b) // Yanıt gövdesini loglamak için kaydet
	return rw.ResponseWriter.Write(b)
}

func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// LoggerMiddleware HTTP isteğini ve yanıtını kaydeder
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Request gövdesini oku ve logla
		var reqBody bytes.Buffer
		if r.Body != nil {
			// Request gövdesini kopyala ve yeniden kullan
			reqBody.ReadFrom(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(reqBody.Bytes()))
		}
		log.Printf("Request - Method: %s, URI: %s, RemoteAddr: %s, Body: %s", r.Method, r.RequestURI, r.RemoteAddr, reqBody.String())

		// Yanıtı sarmalayan ResponseWriter'ı kullan
		wrappedWriter := NewResponseWriterWrapper(w)
		next.ServeHTTP(wrappedWriter, r)

		// Yanıt süresini ve gövdesini logla
		duration := time.Since(start)
		log.Printf("Response - Status: %d, Duration: %v, Body: %s", wrappedWriter.statusCode, duration, wrappedWriter.body.String())
	})
}