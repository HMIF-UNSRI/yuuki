package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(writer http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: writer}
}

func (writer responseWriter) Status() int {
	return writer.status
}

func (writer *responseWriter) WriteHeader(code int) {
	if writer.wroteHeader {
		return
	}

	writer.status = code
	writer.ResponseWriter.WriteHeader(int(code))
	writer.wroteHeader = true

	return
}

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	start := time.Now()
	wrapped := wrapResponseWriter(writer)

	middleware.Handler.ServeHTTP(wrapped,request)

	log.WithFields(log.Fields{
		"status":   wrapped.Status(),
		"method":   request.Method,
		"path":     request.URL.EscapedPath(),
		"duration": time.Since(start),
	}).Infoln()
}

//func LoggingMiddleware(logger *log.Logger) func(handler http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		fn := func(writer http.ResponseWriter, request *http.Request) {
//			start := time.Now()
//			wrapped := wrapResponseWriter(writer)
//
//			next.ServeHTTP(wrapped, request)
//
//			logger.WithFields(log.Fields{
//				"status":   wrapped.Status(),
//				"method":   request.Method,
//				"path":     request.URL.EscapedPath(),
//				"duration": time.Since(start),
//			}).Infoln()
//		}
//
//		return http.HandlerFunc(fn)
//	}
//}
