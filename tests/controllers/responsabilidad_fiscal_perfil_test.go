package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaResponsabilidadFiscalPerfil = "/v1/responsabilidad_fiscal_perfil"

func TestResponsabilidadFiscalPerfilPost(t *testing.T) {
	nuevo := helpers.NuevoResponsabilidadFiscalPerfil(t)
	w := helpers.Peticion("POST", rutaResponsabilidadFiscalPerfil+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetResponsabilidadFiscalPerfilById(r.IdDato(t)); err != nil || v.ResponsabilidadFiscalId != nuevo.ResponsabilidadFiscalId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaResponsabilidadFiscalPerfil+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestResponsabilidadFiscalPerfilPost Finalizado Correctamente (OK)")
}

func TestResponsabilidadFiscalPerfilGetOne(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestResponsabilidadFiscalPerfilGetOne Finalizado Correctamente (OK)")
}

func TestResponsabilidadFiscalPerfilGetAll(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	helpers.ProbarGetAllController(t, rutaResponsabilidadFiscalPerfil, creado.Id)
	t.Log("TestResponsabilidadFiscalPerfilGetAll Finalizado Correctamente (OK)")
}

func TestResponsabilidadFiscalPerfilPut(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	ruta := fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.ResponsabilidadFiscalId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetResponsabilidadFiscalPerfilById(creado.Id); err != nil || v.ResponsabilidadFiscalId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestResponsabilidadFiscalPerfilPut Finalizado Correctamente (OK)")
}

func TestResponsabilidadFiscalPerfilDelete(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetResponsabilidadFiscalPerfilById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestResponsabilidadFiscalPerfilDelete Finalizado Correctamente (OK)")
}
