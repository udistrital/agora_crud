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
	helpers.VerificarError(t, helpers.Peticion("POST", rutaProveedor+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestProveedorPost Finalizado Correctamente (OK)")
}

func TestProveedorGetOne(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedor, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaProveedor+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestProveedorGetOne Finalizado Correctamente (OK)")
}

func TestProveedorGetAll(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	helpers.CrearProveedor(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaProveedor, creado.Id)
	t.Log("TestProveedorGetAll Finalizado Correctamente (OK)")
}

func TestProveedorActivo(t *testing.T) {
	activo := helpers.CrearProveedor(t)
	inactivo := helpers.CrearProveedor(t)
	if err := models.DeleteProveedor(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaProveedor, activo.Id, inactivo.Id)
	t.Log("TestProveedorActivo Finalizado Correctamente (OK)")
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
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaProveedor+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaProveedor, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
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
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaProveedor+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestProveedorDelete Finalizado Correctamente (OK)")
}
