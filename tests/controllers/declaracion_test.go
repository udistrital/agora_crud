package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaDeclaracion = "/v1/declaracion"

func TestDeclaracionPost(t *testing.T) {
	nuevo := helpers.NuevoDeclaracion(t)
	w := helpers.Peticion("POST", rutaDeclaracion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetDeclaracionById(r.IdDato(t)); err != nil || v.Texto != nuevo.Texto || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaDeclaracion+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestDeclaracionPost Finalizado Correctamente (OK)")
}

func TestDeclaracionGetOne(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDeclaracion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDeclaracion, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestDeclaracionGetOne Finalizado Correctamente (OK)")
}

func TestDeclaracionGetAll(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	helpers.CrearDeclaracion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaDeclaracion, creado.Id)
	t.Log("TestDeclaracionGetAll Finalizado Correctamente (OK)")
}

func TestDeclaracionPut(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	ruta := fmt.Sprintf("%s/%d", rutaDeclaracion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Texto = "Declaración actualizada"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetDeclaracionById(creado.Id); err != nil || v.Texto != "Declaración actualizada" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaDeclaracion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestDeclaracionPut Finalizado Correctamente (OK)")
}

func TestDeclaracionDelete(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaDeclaracion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetDeclaracionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaDeclaracion, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestDeclaracionDelete Finalizado Correctamente (OK)")
}
