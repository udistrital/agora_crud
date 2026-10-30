package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaPrefijoFacturacion = "/v1/prefijo_facturacion"

func TestPrefijoFacturacionPost(t *testing.T) {
	nuevo := helpers.NuevoPrefijoFacturacion(t)
	w := helpers.Peticion("POST", rutaPrefijoFacturacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetPrefijoFacturacionById(r.IdDato(t)); err != nil || v.Nombre != nuevo.Nombre || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarError(t, helpers.Peticion("POST", rutaPrefijoFacturacion+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestPrefijoFacturacionPost Finalizado Correctamente (OK)")
}

func TestPrefijoFacturacionGetOne(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPrefijoFacturacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaPrefijoFacturacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaPrefijoFacturacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestPrefijoFacturacionGetOne Finalizado Correctamente (OK)")
}

func TestPrefijoFacturacionGetAll(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	helpers.CrearPrefijoFacturacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaPrefijoFacturacion, creado.Id)
	t.Log("TestPrefijoFacturacionGetAll Finalizado Correctamente (OK)")
}

func TestPrefijoFacturacionActivo(t *testing.T) {
	activo := helpers.CrearPrefijoFacturacion(t)
	inactivo := helpers.CrearPrefijoFacturacion(t)
	if err := models.DeletePrefijoFacturacion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaPrefijoFacturacion, activo.Id, inactivo.Id)
	t.Log("TestPrefijoFacturacionActivo Finalizado Correctamente (OK)")
}

func TestPrefijoFacturacionPut(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaPrefijoFacturacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Nombre = "Prefijo actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetPrefijoFacturacionById(creado.Id); err != nil || v.Nombre != "Prefijo actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaPrefijoFacturacion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaPrefijoFacturacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestPrefijoFacturacionPut Finalizado Correctamente (OK)")
}

func TestPrefijoFacturacionDelete(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaPrefijoFacturacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetPrefijoFacturacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaPrefijoFacturacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaPrefijoFacturacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestPrefijoFacturacionDelete Finalizado Correctamente (OK)")
}
