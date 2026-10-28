package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

// Respuesta representa el cuerpo JSON que retornan los controladores.
type Respuesta struct {
	Success bool
	Status  string
	Message string
	Total   *int64 // solo en las respuestas de GetAll
	Data    interface{}
}

// Peticion envía una petición HTTP al enrutador de beego y retorna la
// respuesta. Requiere que el paquete routers haya sido importado.
func Peticion(metodo, ruta, cuerpo string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(metodo, ruta, strings.NewReader(cuerpo))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	return w
}

// JSON serializa v para usarlo como cuerpo de una petición.
func JSON(t *testing.T, v interface{}) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal("No se pudo serializar el cuerpo de la petición:", err)
	}
	return string(b)
}

// VerificarRespuesta comprueba el código HTTP de la respuesta y, cuando es
// exitosa, decodifica y retorna su cuerpo.
func VerificarRespuesta(t *testing.T, w *httptest.ResponseRecorder, estadoEsperado int) *Respuesta {
	t.Helper()
	if w.Code != estadoEsperado {
		t.Fatalf("Se esperaba el estado %d y se obtuvo %d: %s", estadoEsperado, w.Code, w.Body.String())
	}
	if estadoEsperado >= http.StatusBadRequest {
		return nil
	}
	var r Respuesta
	if err := json.Unmarshal(w.Body.Bytes(), &r); err != nil {
		t.Fatal("La respuesta no es un JSON válido:", err, w.Body.String())
	}
	if !r.Success || r.Status != strconv.Itoa(estadoEsperado) {
		t.Error("Se esperaba Success=true y Status", estadoEsperado, "y se obtuvo", r.Success, r.Status)
	}
	return &r
}

// Dato retorna el campo Data de la respuesta como un mapa.
func (r *Respuesta) Dato(t *testing.T) map[string]interface{} {
	t.Helper()
	m, ok := r.Data.(map[string]interface{})
	if !ok {
		t.Fatal("Se esperaba un objeto en Data y se obtuvo", r.Data)
	}
	return m
}

// IdDato retorna el Id del objeto contenido en Data.
func (r *Respuesta) IdDato(t *testing.T) int {
	t.Helper()
	id, ok := r.Dato(t)["Id"].(float64)
	if !ok {
		t.Fatal("La respuesta no contiene Id:", r.Data)
	}
	return int(id)
}

// ProbarGetAllController verifica el comportamiento común del servicio GetAll
// de un controlador, cuya ruta base es ruta (p. ej. /v1/perfil). Requiere al
// menos dos registros en la tabla, uno de ellos con Id idExistente.
func ProbarGetAllController(t *testing.T, ruta string, idExistente int) {
	t.Helper()
	id := strconv.Itoa(idExistente)

	t.Run("ConParametros", func(t *testing.T) {
		w := Peticion("GET", ruta+"/?query=Id:"+id+"&fields=Id,Activo&sortby=Id&order=desc&limit=5&offset=0", "")
		r := VerificarRespuesta(t, w, http.StatusOK)
		l, ok := r.Data.([]interface{})
		if !ok || len(l) != 1 {
			t.Fatal("Se esperaba 1 registro y se obtuvo", r.Data)
		}
		if m, ok := l[0].(map[string]interface{}); !ok || m["Id"] != float64(idExistente) {
			t.Error("Se esperaba el registro", id, "y se obtuvo", l[0])
		}
		if r.Total == nil || *r.Total != 1 {
			t.Error("Se esperaba Total 1 con el filtro por Id y se obtuvo", r.Total)
		}
	})

	t.Run("SinResultados", func(t *testing.T) {
		w := Peticion("GET", ruta+"/?query=Id:"+strconv.Itoa(IdInexistente), "")
		r := VerificarRespuesta(t, w, http.StatusOK)
		l, ok := r.Data.([]interface{})
		if !ok || len(l) != 1 {
			t.Error("Se esperaba una lista con un objeto vacío y se obtuvo", r.Data)
		}
		if r.Total == nil || *r.Total != 0 {
			t.Error("Se esperaba Total 0 sin resultados y se obtuvo", r.Total)
		}
	})

	// El total cuenta los registros que cumplen los filtros sin aplicar limit
	t.Run("TotalIgnoraLimit", func(t *testing.T) {
		r := VerificarRespuesta(t, Peticion("GET", ruta+"/?limit=1", ""), http.StatusOK)
		l, ok := r.Data.([]interface{})
		if !ok || len(l) != 1 || r.Total == nil || *r.Total < 2 {
			t.Error("Se esperaba 1 registro y un Total mayor o igual a 2 y se obtuvo", r.Data, r.Total)
		}
	})

	t.Run("TotalIgnoraOffset", func(t *testing.T) {
		r := VerificarRespuesta(t, Peticion("GET", ruta+"/?offset="+strconv.Itoa(IdInexistente), ""), http.StatusOK)
		if r.Total == nil || *r.Total < 2 {
			t.Error("Se esperaba un Total mayor o igual a 2 en una página vacía y se obtuvo", r.Total)
		}
	})

	t.Run("QueryInvalido", func(t *testing.T) {
		VerificarRespuesta(t, Peticion("GET", ruta+"/?query=Id", ""), http.StatusBadRequest)
	})

	t.Run("OrdenInvalido", func(t *testing.T) {
		VerificarRespuesta(t, Peticion("GET", ruta+"/?sortby=Id&order=abc", ""), http.StatusNotFound)
	})
}
