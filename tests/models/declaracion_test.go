package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddDeclaracion(t *testing.T) {
	v := helpers.NuevoDeclaracion(t)
	v.Activo = false
	id, err := models.AddDeclaracion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddDeclaracion Finalizado Correctamente (OK)")
}

func TestGetDeclaracionById(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	v, err := models.GetDeclaracionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Texto != creado.Texto {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetDeclaracionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetDeclaracionById Finalizado Correctamente (OK)")
}

func TestGetAllDeclaracion(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	helpers.ProbarGetAll(t, models.GetAllDeclaracion, creado.Id)
	t.Log("TestGetAllDeclaracion Finalizado Correctamente (OK)")
}

func TestUpdateDeclaracionById(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Texto = "Declaración actualizada"
	if err := models.UpdateDeclaracionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetDeclaracionById(creado.Id)
	if v == nil || v.Texto != "Declaración actualizada" {
		t.Fatal("Se esperaba Texto actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateDeclaracionById(&models.Declaracion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateDeclaracionById Finalizado Correctamente (OK)")
}

func TestDeleteDeclaracion(t *testing.T) {
	creado := helpers.CrearDeclaracion(t)
	if err := models.DeleteDeclaracion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetDeclaracionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteDeclaracion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteDeclaracion Finalizado Correctamente (OK)")
}
