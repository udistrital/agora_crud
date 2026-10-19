package models

import (
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddRangoFacturacion(t *testing.T) {
	v := helpers.NuevoRangoFacturacion(t)
	v.Activo = false
	id, err := models.AddRangoFacturacion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo {
		t.Error("Se esperaba Activo=true al crear el registro")
	}
	t.Log("TestAddRangoFacturacion Finalizado Correctamente (OK)")
}

func TestGetRangoFacturacionById(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	v, err := models.GetRangoFacturacionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Nombre != creado.Nombre {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetRangoFacturacionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetRangoFacturacionById Finalizado Correctamente (OK)")
}

func TestGetAllRangoFacturacion(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	helpers.ProbarGetAll(t, models.GetAllRangoFacturacion, creado.Id)
	t.Log("TestGetAllRangoFacturacion Finalizado Correctamente (OK)")
}

func TestUpdateRangoFacturacionById(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	cambios := *creado
	cambios.Nombre = "Rango actualizado"
	if err := models.UpdateRangoFacturacionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetRangoFacturacionById(creado.Id)
	if v == nil || v.Nombre != "Rango actualizado" {
		t.Fatal("Se esperaba Nombre actualizado y se obtuvo", v)
	}
	if err := models.UpdateRangoFacturacionById(&models.RangoFacturacion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateRangoFacturacionById Finalizado Correctamente (OK)")
}

func TestDeleteRangoFacturacion(t *testing.T) {
	creado := helpers.CrearRangoFacturacion(t)
	if err := models.DeleteRangoFacturacion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetRangoFacturacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteRangoFacturacion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteRangoFacturacion Finalizado Correctamente (OK)")
}
