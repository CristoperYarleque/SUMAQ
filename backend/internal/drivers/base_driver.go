package drivers

import (
	"encoding/json"
	"net/http"
)

type baseDriver struct {
}

type BaseDriverInterface interface {
	handleError(w http.ResponseWriter, err error, code int)
	handleSuccess(w http.ResponseWriter, v interface{}, code int)
}

func (c *baseDriver) respond(w http.ResponseWriter, response map[string]interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	respJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error al serializar la respuesta", http.StatusInternalServerError)
		return
	}
	w.Write(respJSON)
}

func (c *baseDriver) handleError(w http.ResponseWriter, err error, code int) {
	response := map[string]interface{}{
		"code":   code,
		"errors": []string{err.Error()},
	}
	c.respond(w, response, code)
}

func (c *baseDriver) handleSuccess(w http.ResponseWriter, v interface{}, code int) {
	response := map[string]interface{}{
		"code": code,
		"data": v,
	}
	c.respond(w, response, code)
}
