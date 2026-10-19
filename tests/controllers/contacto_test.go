package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaContacto = "/v1/contacto"

func TestContactoPost(t *testing.T) {
	nuevo := helpers.NuevoContacto(t)
	w := helpers.Peticion("POST", rutaContacto+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetContactoById(r.IdDato(t)); err != nil || v.ValorEspecifico != nuevo.ValorEspecifico || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaContacto+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestContactoPost Finalizado Correctamente (OK)")
}

func TestContactoGetOne(t *testing.T) {
	creado := helpers.CrearContacto(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaContacto, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaContacto, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestContactoGetOne Finalizado Correctamente (OK)")
}

func TestContactoGetAll(t *testing.T) {
	creado := helpers.CrearContacto(t)
	helpers.ProbarGetAllController(t, rutaContacto, creado.Id)
	t.Log("TestContactoGetAll Finalizado Correctamente (OK)")
}

func TestContactoPut(t *testing.T) {
	creado := helpers.CrearContacto(t)
	ruta := fmt.Sprintf("%s/%d", rutaContacto, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.ValorEspecifico = "otro@udistrital.edu.co"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetContactoById(creado.Id); err != nil || v.ValorEspecifico != "otro@udistrital.edu.co" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaContacto, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestContactoPut Finalizado Correctamente (OK)")
}

func TestContactoDelete(t *testing.T) {
	creado := helpers.CrearContacto(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaContacto, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetContactoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaContacto, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestContactoDelete Finalizado Correctamente (OK)")
}
