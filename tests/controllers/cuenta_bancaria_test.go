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
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaCuentaBancaria+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestCuentaBancariaPost Finalizado Correctamente (OK)")
}

func TestCuentaBancariaGetOne(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCuentaBancaria, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCuentaBancaria, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestCuentaBancariaGetOne Finalizado Correctamente (OK)")
}

func TestCuentaBancariaGetAll(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	helpers.CrearCuentaBancaria(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaCuentaBancaria, creado.Id)
	t.Log("TestCuentaBancariaGetAll Finalizado Correctamente (OK)")
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
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaCuentaBancaria, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
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
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestCuentaBancariaDelete Finalizado Correctamente (OK)")
}
