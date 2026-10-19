package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddResponsabilidadFiscalPerfil(t *testing.T) {
	v := helpers.NuevoResponsabilidadFiscalPerfil(t)
	v.Activo = false
	id, err := models.AddResponsabilidadFiscalPerfil(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddResponsabilidadFiscalPerfil Finalizado Correctamente (OK)")
}

func TestGetResponsabilidadFiscalPerfilById(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	v, err := models.GetResponsabilidadFiscalPerfilById(creado.Id)
	if err != nil || v.Id != creado.Id || v.ResponsabilidadFiscalId != creado.ResponsabilidadFiscalId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetResponsabilidadFiscalPerfilById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetResponsabilidadFiscalPerfilById Finalizado Correctamente (OK)")
}

func TestGetAllResponsabilidadFiscalPerfil(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	helpers.ProbarGetAll(t, models.GetAllResponsabilidadFiscalPerfil, creado.Id)
	t.Log("TestGetAllResponsabilidadFiscalPerfil Finalizado Correctamente (OK)")
}

func TestUpdateResponsabilidadFiscalPerfilById(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.ResponsabilidadFiscalId = 2
	if err := models.UpdateResponsabilidadFiscalPerfilById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetResponsabilidadFiscalPerfilById(creado.Id)
	if v == nil || v.ResponsabilidadFiscalId != 2 {
		t.Fatal("Se esperaba ResponsabilidadFiscalId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateResponsabilidadFiscalPerfilById(&models.ResponsabilidadFiscalPerfil{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateResponsabilidadFiscalPerfilById Finalizado Correctamente (OK)")
}

func TestDeleteResponsabilidadFiscalPerfil(t *testing.T) {
	creado := helpers.CrearResponsabilidadFiscalPerfil(t)
	if err := models.DeleteResponsabilidadFiscalPerfil(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetResponsabilidadFiscalPerfilById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteResponsabilidadFiscalPerfil(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteResponsabilidadFiscalPerfil Finalizado Correctamente (OK)")
}
