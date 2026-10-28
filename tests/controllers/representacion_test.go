package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaRepresentacion = "/v1/representacion"

func TestRepresentacionPost(t *testing.T) {
	nuevo := helpers.NuevoRepresentacion(t)
	w := helpers.Peticion("POST", rutaRepresentacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetRepresentacionById(r.IdDato(t)); err != nil || v.CargoId != nuevo.CargoId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaRepresentacion+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestRepresentacionPost Finalizado Correctamente (OK)")
}

func TestRepresentacionGetOne(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRepresentacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRepresentacion, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestRepresentacionGetOne Finalizado Correctamente (OK)")
}

func TestRepresentacionGetAll(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	helpers.CrearRepresentacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaRepresentacion, creado.Id)
	t.Log("TestRepresentacionGetAll Finalizado Correctamente (OK)")
}

func TestRepresentacionPut(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaRepresentacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.CargoId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetRepresentacionById(creado.Id); err != nil || v.CargoId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaRepresentacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestRepresentacionPut Finalizado Correctamente (OK)")
}

func TestRepresentacionDelete(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRepresentacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetRepresentacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRepresentacion, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestRepresentacionDelete Finalizado Correctamente (OK)")
}
