package models

import (
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddPrefijoFacturacion(t *testing.T) {
	v := helpers.NuevoPrefijoFacturacion(t)
	v.Activo = false
	id, err := models.AddPrefijoFacturacion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo {
		t.Error("Se esperaba Activo=true al crear el registro")
	}
	t.Log("TestAddPrefijoFacturacion Finalizado Correctamente (OK)")
}

func TestGetPrefijoFacturacionById(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	v, err := models.GetPrefijoFacturacionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Nombre != creado.Nombre {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetPrefijoFacturacionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetPrefijoFacturacionById Finalizado Correctamente (OK)")
}

func TestGetAllPrefijoFacturacion(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	helpers.ProbarGetAll(t, models.GetAllPrefijoFacturacion, creado.Id)
	t.Log("TestGetAllPrefijoFacturacion Finalizado Correctamente (OK)")
}

func TestUpdatePrefijoFacturacionById(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	cambios := *creado
	cambios.Nombre = "Prefijo actualizado"
	if err := models.UpdatePrefijoFacturacionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetPrefijoFacturacionById(creado.Id)
	if v == nil || v.Nombre != "Prefijo actualizado" {
		t.Fatal("Se esperaba Nombre actualizado y se obtuvo", v)
	}
	if err := models.UpdatePrefijoFacturacionById(&models.PrefijoFacturacion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdatePrefijoFacturacionById Finalizado Correctamente (OK)")
}

func TestDeletePrefijoFacturacion(t *testing.T) {
	creado := helpers.CrearPrefijoFacturacion(t)
	if err := models.DeletePrefijoFacturacion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetPrefijoFacturacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeletePrefijoFacturacion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeletePrefijoFacturacion Finalizado Correctamente (OK)")
}
