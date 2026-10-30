package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaCuentaBancaria = "/v1/cuenta_bancaria"

func TestCuentaBancariaPost(t *testing.T) {
	nuevo := helpers.NuevoCuentaBancaria(t)
	w := helpers.Peticion("POST", rutaCuentaBancaria+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetCuentaBancariaById(r.IdDato(t)); err != nil || v.NumeroCuenta != nuevo.NumeroCuenta || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarError(t, helpers.Peticion("POST", rutaCuentaBancaria+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestCuentaBancariaPost Finalizado Correctamente (OK)")
}

func TestCuentaBancariaGetOne(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCuentaBancaria, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCuentaBancaria, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaCuentaBancaria+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestCuentaBancariaGetOne Finalizado Correctamente (OK)")
}

func TestCuentaBancariaGetAll(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	helpers.CrearCuentaBancaria(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaCuentaBancaria, creado.Id)
	t.Log("TestCuentaBancariaGetAll Finalizado Correctamente (OK)")
}

func TestCuentaBancariaActivo(t *testing.T) {
	activo := helpers.CrearCuentaBancaria(t)
	inactivo := helpers.CrearCuentaBancaria(t)
	if err := models.DeleteCuentaBancaria(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaCuentaBancaria, activo.Id, inactivo.Id)
	t.Log("TestCuentaBancariaActivo Finalizado Correctamente (OK)")
}

func TestCuentaBancariaPut(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	ruta := fmt.Sprintf("%s/%d", rutaCuentaBancaria, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.NumeroCuenta = "987654321"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetCuentaBancariaById(creado.Id); err != nil || v.NumeroCuenta != "987654321" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaCuentaBancaria+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaCuentaBancaria, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestCuentaBancariaPut Finalizado Correctamente (OK)")
}

func TestCuentaBancariaDelete(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaCuentaBancaria, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetCuentaBancariaById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaCuentaBancaria, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaCuentaBancaria+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestCuentaBancariaDelete Finalizado Correctamente (OK)")
}
