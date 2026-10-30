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
	helpers.VerificarError(t, helpers.Peticion("POST", rutaRepresentacion+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestRepresentacionPost Finalizado Correctamente (OK)")
}

func TestRepresentacionGetOne(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRepresentacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRepresentacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaRepresentacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestRepresentacionGetOne Finalizado Correctamente (OK)")
}

func TestRepresentacionGetAll(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	helpers.CrearRepresentacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaRepresentacion, creado.Id)
	t.Log("TestRepresentacionGetAll Finalizado Correctamente (OK)")
}

func TestRepresentacionActivo(t *testing.T) {
	activo := helpers.CrearRepresentacion(t)
	inactivo := helpers.CrearRepresentacion(t)
	if err := models.DeleteRepresentacion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaRepresentacion, activo.Id, inactivo.Id)
	t.Log("TestRepresentacionActivo Finalizado Correctamente (OK)")
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
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaRepresentacion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaRepresentacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
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
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaRepresentacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestRepresentacionDelete Finalizado Correctamente (OK)")
}
