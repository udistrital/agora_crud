package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaPerfilFinanciero = "/v1/perfil_financiero"

func TestPerfilFinancieroPost(t *testing.T) {
	nuevo := helpers.NuevoPerfilFinanciero(t)
	w := helpers.Peticion("POST", rutaPerfilFinanciero+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetPerfilFinancieroById(r.IdDato(t)); err != nil || v.ResolucionFacturacion != nuevo.ResolucionFacturacion || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarError(t, helpers.Peticion("POST", rutaPerfilFinanciero+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestPerfilFinancieroPost Finalizado Correctamente (OK)")
}

func TestPerfilFinancieroGetOne(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaPerfilFinanciero+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestPerfilFinancieroGetOne Finalizado Correctamente (OK)")
}

func TestPerfilFinancieroGetAll(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	helpers.CrearPerfilFinanciero(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaPerfilFinanciero, creado.Id)
	t.Log("TestPerfilFinancieroGetAll Finalizado Correctamente (OK)")
}

func TestPerfilFinancieroActivo(t *testing.T) {
	activo := helpers.CrearPerfilFinanciero(t)
	inactivo := helpers.CrearPerfilFinanciero(t)
	if err := models.DeletePerfilFinanciero(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaPerfilFinanciero, activo.Id, inactivo.Id)
	t.Log("TestPerfilFinancieroActivo Finalizado Correctamente (OK)")
}

func TestPerfilFinancieroPut(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	ruta := fmt.Sprintf("%s/%d", rutaPerfilFinanciero, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.ResolucionFacturacion = "Resolución 002"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetPerfilFinancieroById(creado.Id); err != nil || v.ResolucionFacturacion != "Resolución 002" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaPerfilFinanciero+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestPerfilFinancieroPut Finalizado Correctamente (OK)")
}

func TestPerfilFinancieroDelete(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetPerfilFinancieroById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaPerfilFinanciero+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestPerfilFinancieroDelete Finalizado Correctamente (OK)")
}
