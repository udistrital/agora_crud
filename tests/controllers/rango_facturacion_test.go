package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaRangoFacturacion = "/v1/rango_facturacion"

func TestRangoFacturacionPost(t *testing.T) {
	nuevo := helpers.NuevoRangoFacturacion(t)
	w := helpers.Peticion("POST", rutaRangoFacturacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetRangoFacturacionById(r.IdDato(t)); err != nil || v.Nombre != nuevo.Nombre || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarError(t, helpers.Peticion("POST", rutaRangoFacturacion+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestRangoFacturacionPost Finalizado Correctamente (OK)")
}

func TestRangoFacturacionGetOne(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRangoFacturacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRangoFacturacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaRangoFacturacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestRangoFacturacionGetOne Finalizado Correctamente (OK)")
}

func TestRangoFacturacionGetAll(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	helpers.CrearRangoFacturacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaRangoFacturacion, creado.Id)
	t.Log("TestRangoFacturacionGetAll Finalizado Correctamente (OK)")
}

func TestRangoFacturacionActivo(t *testing.T) {
	activo := helpers.CrearRangoFacturacion(t)
	inactivo := helpers.CrearRangoFacturacion(t)
	if err := models.DeleteRangoFacturacion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaRangoFacturacion, activo.Id, inactivo.Id)
	t.Log("TestRangoFacturacionActivo Finalizado Correctamente (OK)")
}

func TestRangoFacturacionPut(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaRangoFacturacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Nombre = "Rango actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetRangoFacturacionById(creado.Id); err != nil || v.Nombre != "Rango actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaRangoFacturacion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaRangoFacturacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestRangoFacturacionPut Finalizado Correctamente (OK)")
}

func TestRangoFacturacionDelete(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRangoFacturacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetRangoFacturacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRangoFacturacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaRangoFacturacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestRangoFacturacionDelete Finalizado Correctamente (OK)")
}
