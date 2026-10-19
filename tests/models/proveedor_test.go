package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddProveedor(t *testing.T) {
	v := helpers.NuevoProveedor(t)
	v.Activo = false
	id, err := models.AddProveedor(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddProveedor Finalizado Correctamente (OK)")
}

func TestGetProveedorById(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	v, err := models.GetProveedorById(creado.Id)
	if err != nil || v.Id != creado.Id || v.TerceroId != creado.TerceroId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetProveedorById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetProveedorById Finalizado Correctamente (OK)")
}

func TestGetAllProveedor(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	helpers.ProbarGetAll(t, models.GetAllProveedor, creado.Id)
	t.Log("TestGetAllProveedor Finalizado Correctamente (OK)")
}

func TestUpdateProveedorById(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.TerceroId = 2
	if err := models.UpdateProveedorById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetProveedorById(creado.Id)
	if v == nil || v.TerceroId != 2 {
		t.Fatal("Se esperaba TerceroId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateProveedorById(&models.Proveedor{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateProveedorById Finalizado Correctamente (OK)")
}

func TestDeleteProveedor(t *testing.T) {
	creado := helpers.CrearProveedor(t)
	if err := models.DeleteProveedor(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetProveedorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteProveedor(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteProveedor Finalizado Correctamente (OK)")
}
