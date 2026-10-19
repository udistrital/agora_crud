package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaProveedorJuridico = "/v1/proveedor_juridico"

func TestProveedorJuridicoPost(t *testing.T) {
	nuevo := helpers.NuevoProveedorJuridico(t)
	w := helpers.Peticion("POST", rutaProveedorJuridico+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetProveedorJuridicoById(r.IdDato(t)); err != nil || v.NombreComercial != nuevo.NombreComercial || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaProveedorJuridico+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestProveedorJuridicoPost Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoGetOne(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedorJuridico, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedorJuridico, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestProveedorJuridicoGetOne Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoGetAll(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	helpers.ProbarGetAllController(t, rutaProveedorJuridico, creado.Id)
	t.Log("TestProveedorJuridicoGetAll Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoPut(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	ruta := fmt.Sprintf("%s/%d", rutaProveedorJuridico, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.NombreComercial = "Empresa actualizada"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetProveedorJuridicoById(creado.Id); err != nil || v.NombreComercial != "Empresa actualizada" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaProveedorJuridico, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestProveedorJuridicoPut Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoDelete(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaProveedorJuridico, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetProveedorJuridicoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaProveedorJuridico, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestProveedorJuridicoDelete Finalizado Correctamente (OK)")
}
