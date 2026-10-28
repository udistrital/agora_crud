package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddCuentaBancaria(t *testing.T) {
	v := helpers.NuevoCuentaBancaria(t)
	v.Activo = false
	id, err := models.AddCuentaBancaria(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddCuentaBancaria Finalizado Correctamente (OK)")
}

func TestGetCuentaBancariaById(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	v, err := models.GetCuentaBancariaById(creado.Id)
	if err != nil || v.Id != creado.Id || v.NumeroCuenta != creado.NumeroCuenta {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetCuentaBancariaById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetCuentaBancariaById Finalizado Correctamente (OK)")
}

func TestGetAllCuentaBancaria(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	helpers.CrearCuentaBancaria(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllCuentaBancaria, creado.Id)
	t.Log("TestGetAllCuentaBancaria Finalizado Correctamente (OK)")
}

func TestUpdateCuentaBancariaById(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.NumeroCuenta = "987654321"
	if err := models.UpdateCuentaBancariaById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetCuentaBancariaById(creado.Id)
	if v == nil || v.NumeroCuenta != "987654321" {
		t.Fatal("Se esperaba NumeroCuenta actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateCuentaBancariaById(&models.CuentaBancaria{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateCuentaBancariaById Finalizado Correctamente (OK)")
}

func TestDeleteCuentaBancaria(t *testing.T) {
	creado := helpers.CrearCuentaBancaria(t)
	if err := models.DeleteCuentaBancaria(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetCuentaBancariaById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteCuentaBancaria(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteCuentaBancaria Finalizado Correctamente (OK)")
}
