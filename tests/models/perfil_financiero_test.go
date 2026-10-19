package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddPerfilFinanciero(t *testing.T) {
	v := helpers.NuevoPerfilFinanciero(t)
	v.Activo = false
	id, err := models.AddPerfilFinanciero(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddPerfilFinanciero Finalizado Correctamente (OK)")
}

func TestGetPerfilFinancieroById(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	v, err := models.GetPerfilFinancieroById(creado.Id)
	if err != nil || v.Id != creado.Id || v.ResolucionFacturacion != creado.ResolucionFacturacion {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetPerfilFinancieroById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetPerfilFinancieroById Finalizado Correctamente (OK)")
}

func TestGetAllPerfilFinanciero(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	helpers.ProbarGetAll(t, models.GetAllPerfilFinanciero, creado.Id)
	t.Log("TestGetAllPerfilFinanciero Finalizado Correctamente (OK)")
}

func TestUpdatePerfilFinancieroById(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.ResolucionFacturacion = "Resolución 002"
	if err := models.UpdatePerfilFinancieroById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetPerfilFinancieroById(creado.Id)
	if v == nil || v.ResolucionFacturacion != "Resolución 002" {
		t.Fatal("Se esperaba ResolucionFacturacion actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdatePerfilFinancieroById(&models.PerfilFinanciero{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdatePerfilFinancieroById Finalizado Correctamente (OK)")
}

func TestDeletePerfilFinanciero(t *testing.T) {
	creado := helpers.CrearPerfilFinanciero(t)
	if err := models.DeletePerfilFinanciero(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetPerfilFinancieroById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeletePerfilFinanciero(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeletePerfilFinanciero Finalizado Correctamente (OK)")
}
