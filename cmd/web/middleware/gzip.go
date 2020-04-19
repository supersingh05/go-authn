package middleware

import (
	"net/http"
	"strings"

	"compress/gzip"

	"github.com/supersingh05/go-authn/pkg/common"
)

type GzipMiddleware struct {
	next http.Handler
	app  common.Application
}

type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)
}

func NewGzipMiddleware(app common.Application, next http.Handler) http.Handler {

	return GzipMiddleware{
		next: next,
		app:  app,
	}
}

func (g GzipMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	encoding := r.Header.Get("Accept-Encoding")
	if strings.Contains(encoding, "gzip") {
		g.serveGzip(rw, r)
	} else if strings.Contains(encoding, "deflate") {
		panic("deflate not implemented")
	} else {
		g.next.ServeHTTP(rw, r)
	}

}

func (g GzipMiddleware) serveGzip(rw http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(rw)
	defer gzw.Close()

	rw.Header().Set("Content-Encoding", "gzip")
	g.next.ServeHTTP(GzipResponseWriter{gzw, rw}, r)

}

func (g GzipMiddleware) servePlain(rw http.ResponseWriter, r *http.Request) {
	g.next.ServeHTTP(rw, r)
}
