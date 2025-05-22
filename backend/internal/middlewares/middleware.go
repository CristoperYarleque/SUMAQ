package middlewares

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/cyarleque/sumaq/internal/database"
	"github.com/cyarleque/sumaq/internal/utils"
	"github.com/google/uuid"
)

const (
	REQUESTID string = "requestId"
)

type middleware struct {
	Connection database.Connections
	utils      utils.UtilsInterface
}

type MiddlewareInterface interface {
	MyMiddleware(next http.Handler) http.Handler
}

func NewMiddleware(db database.Connections) MiddlewareInterface {
	return &middleware{
		Connection: db,
		utils:      &utils.Utils{},
	}
}

func (m *middleware) MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/ping" {
			next.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/v1/login" {
			next.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/v1/users" {
			next.ServeHTTP(w, r)
			return
		}

		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		claims, err := m.utils.ParseJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		requestId := uuid.New()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", claims["user_id"])
		ctx = context.WithValue(ctx, "role", claims["role"])
		ctx = context.WithValue(ctx, REQUESTID, requestId)
		r = r.WithContext(ctx)
		m.printRequestBody(r)

		// defer m.releaseConnection(requestId)
		next.ServeHTTP(w, r)
	})
}
func (m *middleware) printRequestBody(r *http.Request) {
	// Leer el cuerpo de la solicitud
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
	}

	// Restaurar el cuerpo de la solicitud para que pueda ser leído nuevamente por otros manejadores
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	ctx := r.Context()
	// // Imprimir la solicitud en la consola con formato JSON
	log.Printf("User ID %v. Role %v. Request ID %v. Received request: %s - %s\n", ctx.Value("user_id"), ctx.Value("role"), ctx.Value(REQUESTID), r.URL.Path, string(body))
}
