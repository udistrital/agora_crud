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
	helpers.VerificarError(t, helpers.Peticion("POST", rutaContacto+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestContactoPost Finalizado Correctamente (OK)")
}

func TestContactoGetOne(t *testing.T) {
	creado := helpers.CrearContacto(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaContacto, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaContacto, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaContacto+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestContactoGetOne Finalizado Correctamente (OK)")
}

func TestContactoGetAll(t *testing.T) {
	creado := helpers.CrearContacto(t)
	helpers.CrearContacto(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaContacto, creado.Id)
	t.Log("TestContactoGetAll Finalizado Correctamente (OK)")
}

func TestContactoActivo(t *testing.T) {
	activo := helpers.CrearContacto(t)
	inactivo := helpers.CrearContacto(t)
	if err := models.DeleteContacto(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaContacto, activo.Id, inactivo.Id)
	t.Log("TestContactoActivo Finalizado Correctamente (OK)")
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
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaContacto+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaContacto, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
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
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaContacto+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestContactoDelete Finalizado Correctamente (OK)")
}
