package models

import (
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddPerfil(t *testing.T) {
	v := helpers.NuevoPerfil(t)
	v.Activo = false
	id, err := models.AddPerfil(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo {
		t.Error("Se esperaba Activo=true al crear el registro")
	}
	t.Log("TestAddPerfil Finalizado Correctamente (OK)")
}

func TestGetPerfilById(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	v, err := models.GetPerfilById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Nombre != creado.Nombre {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetPerfilById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetPerfilById Finalizado Correctamente (OK)")
}

func TestGetAllPerfil(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	helpers.ProbarGetAll(t, models.GetAllPerfil, creado.Id)
	t.Log("TestGetAllPerfil Finalizado Correctamente (OK)")
}

func TestUpdatePerfilById(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	cambios := *creado
	cambios.Nombre = "Perfil actualizado"
	if err := models.UpdatePerfilById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetPerfilById(creado.Id)
	if v == nil || v.Nombre != "Perfil actualizado" {
		t.Fatal("Se esperaba Nombre actualizado y se obtuvo", v)
	}
	if err := models.UpdatePerfilById(&models.Perfil{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdatePerfilById Finalizado Correctamente (OK)")
}

func TestDeletePerfil(t *testing.T) {
	creado := helpers.CrearPerfil(t)
	if err := models.DeletePerfil(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetPerfilById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeletePerfil(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeletePerfil Finalizado Correctamente (OK)")
}
