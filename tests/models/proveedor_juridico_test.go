package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddProveedorJuridico(t *testing.T) {
	v := helpers.NuevoProveedorJuridico(t)
	v.Activo = false
	id, err := models.AddProveedorJuridico(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddProveedorJuridico Finalizado Correctamente (OK)")
}

func TestGetProveedorJuridicoById(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	v, err := models.GetProveedorJuridicoById(creado.Id)
	if err != nil || v.Id != creado.Id || v.NombreComercial != creado.NombreComercial {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetProveedorJuridicoById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetProveedorJuridicoById Finalizado Correctamente (OK)")
}

func TestGetAllProveedorJuridico(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	helpers.ProbarGetAll(t, models.GetAllProveedorJuridico, creado.Id)
	t.Log("TestGetAllProveedorJuridico Finalizado Correctamente (OK)")
}

func TestUpdateProveedorJuridicoById(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.NombreComercial = "Empresa actualizada"
	if err := models.UpdateProveedorJuridicoById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetProveedorJuridicoById(creado.Id)
	if v == nil || v.NombreComercial != "Empresa actualizada" {
		t.Fatal("Se esperaba NombreComercial actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateProveedorJuridicoById(&models.ProveedorJuridico{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateProveedorJuridicoById Finalizado Correctamente (OK)")
}

func TestDeleteProveedorJuridico(t *testing.T) {
	creado := helpers.CrearProveedorJuridico(t)
	if err := models.DeleteProveedorJuridico(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetProveedorJuridicoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteProveedorJuridico(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteProveedorJuridico Finalizado Correctamente (OK)")
}
