package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaPerfil = "/v1/perfil"

func TestPerfilPost(t *testing.T) {
	nuevo := helpers.NuevoPerfil(t)
	w := helpers.Peticion("POST", rutaPerfil+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetPerfilById(r.IdDato(t)); err != nil || v.Nombre != nuevo.Nombre || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaPerfil+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestPerfilPost Finalizado Correctamente (OK)")
}

func TestPerfilGetOne(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPerfil, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPerfil, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestPerfilGetOne Finalizado Correctamente (OK)")
}

func TestPerfilGetAll(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	helpers.CrearPerfil(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaPerfil, creado.Id)
	t.Log("TestPerfilGetAll Finalizado Correctamente (OK)")
}

func TestPerfilPut(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	ruta := fmt.Sprintf("%s/%d", rutaPerfil, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Nombre = "Perfil actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetPerfilById(creado.Id); err != nil || v.Nombre != "Perfil actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaPerfil, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestPerfilPut Finalizado Correctamente (OK)")
}

func TestPerfilDelete(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaPerfil, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetPerfilById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaPerfil, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestPerfilDelete Finalizado Correctamente (OK)")
}
