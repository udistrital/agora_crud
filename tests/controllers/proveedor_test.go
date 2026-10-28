package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaProveedor = "/v1/proveedor"

func TestProveedorPost(t *testing.T) {
	nuevo := helpers.NuevoProveedor(t)
	w := helpers.Peticion("POST", rutaProveedor+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetProveedorById(r.IdDato(t)); err != nil || v.TerceroId != nuevo.TerceroId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaProveedor+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestProveedorPost Finalizado Correctamente (OK)")
}

func TestProveedorGetOne(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedor, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestProveedorGetOne Finalizado Correctamente (OK)")
}

func TestProveedorGetAll(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	helpers.CrearProveedor(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaProveedor, creado.Id)
	t.Log("TestProveedorGetAll Finalizado Correctamente (OK)")
}

func TestProveedorPut(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	ruta := fmt.Sprintf("%s/%d", rutaProveedor, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.TerceroId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetProveedorById(creado.Id); err != nil || v.TerceroId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaProveedor, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestProveedorPut Finalizado Correctamente (OK)")
}

func TestProveedorDelete(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetProveedorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaProveedor, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestProveedorDelete Finalizado Correctamente (OK)")
}
