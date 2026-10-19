package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddRepresentacion(t *testing.T) {
	v := helpers.NuevoRepresentacion(t)
	v.Activo = false
	id, err := models.AddRepresentacion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddRepresentacion Finalizado Correctamente (OK)")
}

func TestGetRepresentacionById(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	v, err := models.GetRepresentacionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.CargoId != creado.CargoId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetRepresentacionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetRepresentacionById Finalizado Correctamente (OK)")
}

func TestGetAllRepresentacion(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	helpers.ProbarGetAll(t, models.GetAllRepresentacion, creado.Id)
	t.Log("TestGetAllRepresentacion Finalizado Correctamente (OK)")
}

func TestUpdateRepresentacionById(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.CargoId = 2
	if err := models.UpdateRepresentacionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetRepresentacionById(creado.Id)
	if v == nil || v.CargoId != 2 {
		t.Fatal("Se esperaba CargoId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateRepresentacionById(&models.Representacion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateRepresentacionById Finalizado Correctamente (OK)")
}

func TestDeleteRepresentacion(t *testing.T) {
	creado := helpers.CrearRepresentacion(t)
	if err := models.DeleteRepresentacion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetRepresentacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteRepresentacion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteRepresentacion Finalizado Correctamente (OK)")
}
