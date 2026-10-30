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
	helpers.VerificarError(t, helpers.Peticion("POST", rutaDeclaracion+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestDeclaracionPost Finalizado Correctamente (OK)")
}

func TestDeclaracionGetOne(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDeclaracion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDeclaracion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaDeclaracion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestDeclaracionGetOne Finalizado Correctamente (OK)")
}

func TestDeclaracionGetAll(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	helpers.CrearDeclaracion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaDeclaracion, creado.Id)
	t.Log("TestDeclaracionGetAll Finalizado Correctamente (OK)")
}

func TestDeclaracionActivo(t *testing.T) {
	activo := helpers.CrearDeclaracion(t)
	inactivo := helpers.CrearDeclaracion(t)
	if err := models.DeleteDeclaracion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaDeclaracion, activo.Id, inactivo.Id)
	t.Log("TestDeclaracionActivo Finalizado Correctamente (OK)")
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
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaDeclaracion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaDeclaracion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
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
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaDeclaracion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestDeclaracionDelete Finalizado Correctamente (OK)")
}
