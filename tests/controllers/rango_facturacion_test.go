package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaRangoFacturacion = "/v1/rango_facturacion"

func TestRangoFacturacionPost(t *testing.T) {
	nuevo := helpers.NuevoRangoFacturacion(t)
	w := helpers.Peticion("POST", rutaRangoFacturacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetRangoFacturacionById(r.IdDato(t)); err != nil || v.Nombre != nuevo.Nombre || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaRangoFacturacion+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestRangoFacturacionPost Finalizado Correctamente (OK)")
}

func TestRangoFacturacionGetOne(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRangoFacturacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRangoFacturacion, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestRangoFacturacionGetOne Finalizado Correctamente (OK)")
}

func TestRangoFacturacionGetAll(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	helpers.CrearRangoFacturacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaRangoFacturacion, creado.Id)
	t.Log("TestRangoFacturacionGetAll Finalizado Correctamente (OK)")
}

func TestRangoFacturacionPut(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaRangoFacturacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Nombre = "Rango actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetRangoFacturacionById(creado.Id); err != nil || v.Nombre != "Rango actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaRangoFacturacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestRangoFacturacionPut Finalizado Correctamente (OK)")
}

func TestRangoFacturacionDelete(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRangoFacturacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetRangoFacturacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRangoFacturacion, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestRangoFacturacionDelete Finalizado Correctamente (OK)")
}
