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
	helpers.VerificarError(t, helpers.Peticion("POST", rutaResponsabilidadFiscalPerfil+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestResponsabilidadFiscalPerfilPost Finalizado Correctamente (OK)")
}

func TestResponsabilidadFiscalPerfilGetOne(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaResponsabilidadFiscalPerfil+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestResponsabilidadFiscalPerfilGetOne Finalizado Correctamente (OK)")
}

func TestResponsabilidadFiscalPerfilGetAll(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	helpers.CrearResponsabilidadFiscalPerfil(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaResponsabilidadFiscalPerfil, creado.Id)
	t.Log("TestResponsabilidadFiscalPerfilGetAll Finalizado Correctamente (OK)")
}

func TestResponsabilidadFiscalPerfilActivo(t *testing.T) {
	activo := helpers.CrearResponsabilidadFiscalPerfil(t)
	inactivo := helpers.CrearResponsabilidadFiscalPerfil(t)
	if err := models.DeleteResponsabilidadFiscalPerfil(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaResponsabilidadFiscalPerfil, activo.Id, inactivo.Id)
	t.Log("TestResponsabilidadFiscalPerfilActivo Finalizado Correctamente (OK)")
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
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaResponsabilidadFiscalPerfil+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaResponsabilidadFiscalPerfil, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
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
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaResponsabilidadFiscalPerfil+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestResponsabilidadFiscalPerfilDelete Finalizado Correctamente (OK)")
}
