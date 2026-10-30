package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaInformacionFinanciera = "/v1/informacion_financiera"

func TestInformacionFinancieraPost(t *testing.T) {
	nuevo := helpers.NuevoInformacionFinanciera(t)
	w := helpers.Peticion("POST", rutaInformacionFinanciera+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetInformacionFinancieraById(r.IdDato(t)); err != nil || v.MonedaId != nuevo.MonedaId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarError(t, helpers.Peticion("POST", rutaInformacionFinanciera+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestInformacionFinancieraPost Finalizado Correctamente (OK)")
}

func TestInformacionFinancieraGetOne(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaInformacionFinanciera, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaInformacionFinanciera, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaInformacionFinanciera+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestInformacionFinancieraGetOne Finalizado Correctamente (OK)")
}

func TestInformacionFinancieraGetAll(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	helpers.CrearInformacionFinanciera(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaInformacionFinanciera, creado.Id)
	t.Log("TestInformacionFinancieraGetAll Finalizado Correctamente (OK)")
}

func TestInformacionFinancieraActivo(t *testing.T) {
	activo := helpers.CrearInformacionFinanciera(t)
	inactivo := helpers.CrearInformacionFinanciera(t)
	if err := models.DeleteInformacionFinanciera(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaInformacionFinanciera, activo.Id, inactivo.Id)
	t.Log("TestInformacionFinancieraActivo Finalizado Correctamente (OK)")
}

func TestInformacionFinancieraPut(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	ruta := fmt.Sprintf("%s/%d", rutaInformacionFinanciera, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.MonedaId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetInformacionFinancieraById(creado.Id); err != nil || v.MonedaId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaInformacionFinanciera+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaInformacionFinanciera, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestInformacionFinancieraPut Finalizado Correctamente (OK)")
}

func TestInformacionFinancieraDelete(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaInformacionFinanciera, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetInformacionFinancieraById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaInformacionFinanciera, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaInformacionFinanciera+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestInformacionFinancieraDelete Finalizado Correctamente (OK)")
}
