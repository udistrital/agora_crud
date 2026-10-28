package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaActividadEconomicaProveedor = "/v1/actividad_economica_proveedor"

func TestActividadEconomicaProveedorPost(t *testing.T) {
	nuevo := helpers.NuevoActividadEconomicaProveedor(t)
	w := helpers.Peticion("POST", rutaActividadEconomicaProveedor+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetActividadEconomicaProveedorById(r.IdDato(t)); err != nil || v.ActividadEconomicaId != nuevo.ActividadEconomicaId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaActividadEconomicaProveedor+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestActividadEconomicaProveedorPost Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorGetOne(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaActividadEconomicaProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaActividadEconomicaProveedor, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestActividadEconomicaProveedorGetOne Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorGetAll(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	helpers.CrearActividadEconomicaProveedor(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaActividadEconomicaProveedor, creado.Id)
	t.Log("TestActividadEconomicaProveedorGetAll Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorPut(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	ruta := fmt.Sprintf("%s/%d", rutaActividadEconomicaProveedor, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.ActividadEconomicaId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetActividadEconomicaProveedorById(creado.Id); err != nil || v.ActividadEconomicaId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaActividadEconomicaProveedor, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestActividadEconomicaProveedorPut Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorDelete(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaActividadEconomicaProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetActividadEconomicaProveedorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaActividadEconomicaProveedor, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestActividadEconomicaProveedorDelete Finalizado Correctamente (OK)")
}
