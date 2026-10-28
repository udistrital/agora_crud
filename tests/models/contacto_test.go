package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddContacto(t *testing.T) {
	v := helpers.NuevoContacto(t)
	v.Activo = false
	id, err := models.AddContacto(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddContacto Finalizado Correctamente (OK)")
}

func TestGetContactoById(t *testing.T) {
	creado := helpers.CrearContacto(t)
	v, err := models.GetContactoById(creado.Id)
	if err != nil || v.Id != creado.Id || v.ValorEspecifico != creado.ValorEspecifico {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetContactoById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetContactoById Finalizado Correctamente (OK)")
}

func TestGetAllContacto(t *testing.T) {
	creado := helpers.CrearContacto(t)
	helpers.CrearContacto(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllContacto, creado.Id)
	t.Log("TestGetAllContacto Finalizado Correctamente (OK)")
}

func TestUpdateContactoById(t *testing.T) {
	creado := helpers.CrearContacto(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.ValorEspecifico = "otro@udistrital.edu.co"
	if err := models.UpdateContactoById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetContactoById(creado.Id)
	if v == nil || v.ValorEspecifico != "otro@udistrital.edu.co" {
		t.Fatal("Se esperaba ValorEspecifico actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateContactoById(&models.Contacto{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateContactoById Finalizado Correctamente (OK)")
}

func TestDeleteContacto(t *testing.T) {
	creado := helpers.CrearContacto(t)
	if err := models.DeleteContacto(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetContactoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteContacto(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteContacto Finalizado Correctamente (OK)")
}
