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
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaPerfilFinanciero+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestPerfilFinancieroPost Finalizado Correctamente (OK)")
}

func TestPerfilFinancieroGetOne(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestPerfilFinancieroGetOne Finalizado Correctamente (OK)")
}

func TestPerfilFinancieroGetAll(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	helpers.ProbarGetAllController(t, rutaPerfilFinanciero, creado.Id)
	t.Log("TestPerfilFinancieroGetAll Finalizado Correctamente (OK)")
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
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaPerfilFinanciero, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
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
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestPerfilFinancieroDelete Finalizado Correctamente (OK)")
}
