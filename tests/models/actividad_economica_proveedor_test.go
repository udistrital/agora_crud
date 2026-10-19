package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddActividadEconomicaProveedor(t *testing.T) {
	v := helpers.NuevoActividadEconomicaProveedor(t)
	v.Activo = false
	id, err := models.AddActividadEconomicaProveedor(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddActividadEconomicaProveedor Finalizado Correctamente (OK)")
}

func TestGetActividadEconomicaProveedorById(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	v, err := models.GetActividadEconomicaProveedorById(creado.Id)
	if err != nil || v.Id != creado.Id || v.ActividadEconomicaId != creado.ActividadEconomicaId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetActividadEconomicaProveedorById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetActividadEconomicaProveedorById Finalizado Correctamente (OK)")
}

func TestGetAllActividadEconomicaProveedor(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	helpers.ProbarGetAll(t, models.GetAllActividadEconomicaProveedor, creado.Id)
	t.Log("TestGetAllActividadEconomicaProveedor Finalizado Correctamente (OK)")
}

func TestUpdateActividadEconomicaProveedorById(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.ActividadEconomicaId = 2
	if err := models.UpdateActividadEconomicaProveedorById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetActividadEconomicaProveedorById(creado.Id)
	if v == nil || v.ActividadEconomicaId != 2 {
		t.Fatal("Se esperaba ActividadEconomicaId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateActividadEconomicaProveedorById(&models.ActividadEconomicaProveedor{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateActividadEconomicaProveedorById Finalizado Correctamente (OK)")
}

func TestDeleteActividadEconomicaProveedor(t *testing.T) {
	creado := helpers.CrearActividadEconomicaProveedor(t)
	if err := models.DeleteActividadEconomicaProveedor(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetActividadEconomicaProveedorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteActividadEconomicaProveedor(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteActividadEconomicaProveedor Finalizado Correctamente (OK)")
}
