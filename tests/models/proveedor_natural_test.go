package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddProveedorNatural(t *testing.T) {
	v := helpers.NuevoProveedorNatural(t)
	v.Activo = false
	id, err := models.AddProveedorNatural(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddProveedorNatural Finalizado Correctamente (OK)")
}

func TestGetProveedorNaturalById(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	v, err := models.GetProveedorNaturalById(creado.Id)
	if err != nil || v.Id != creado.Id || v.ExperienciaLaboralMeses != creado.ExperienciaLaboralMeses {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetProveedorNaturalById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetProveedorNaturalById Finalizado Correctamente (OK)")
}

func TestGetAllProveedorNatural(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	helpers.ProbarGetAll(t, models.GetAllProveedorNatural, creado.Id)
	t.Log("TestGetAllProveedorNatural Finalizado Correctamente (OK)")
}

func TestUpdateProveedorNaturalById(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.ExperienciaLaboralMeses = 24
	if err := models.UpdateProveedorNaturalById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetProveedorNaturalById(creado.Id)
	if v == nil || v.ExperienciaLaboralMeses != 24 {
		t.Fatal("Se esperaba ExperienciaLaboralMeses actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateProveedorNaturalById(&models.ProveedorNatural{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateProveedorNaturalById Finalizado Correctamente (OK)")
}

func TestDeleteProveedorNatural(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	if err := models.DeleteProveedorNatural(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetProveedorNaturalById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteProveedorNatural(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteProveedorNatural Finalizado Correctamente (OK)")
}
