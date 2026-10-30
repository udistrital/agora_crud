package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddInformacionFinanciera(t *testing.T) {
	v := helpers.NuevoInformacionFinanciera(t)
	v.Activo = false
	id, err := models.AddInformacionFinanciera(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddInformacionFinanciera Finalizado Correctamente (OK)")
}

func TestAddInformacionFinancieraFechaCorte(t *testing.T) {
	v := helpers.NuevoInformacionFinanciera(t)
	v.FechaCorte = time.Time{}
	if _, err := models.AddInformacionFinanciera(v); err != nil {
		t.Fatal("Se esperaba el registro creado y se obtuvo", err)
	}
	esperada := time.Date(time.Now().UTC().Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	if !v.FechaCorte.Equal(esperada) {
		t.Error("Se esperaba FechaCorte", esperada, "y se obtuvo", v.FechaCorte)
	}

	v = helpers.NuevoInformacionFinanciera(t)
	fechaCorte := v.FechaCorte
	if _, err := models.AddInformacionFinanciera(v); err != nil || !v.FechaCorte.Equal(fechaCorte) {
		t.Error("Se esperaba conservar la FechaCorte enviada y se obtuvo", v.FechaCorte, err)
	}
	t.Log("TestAddInformacionFinancieraFechaCorte Finalizado Correctamente (OK)")
}

func TestGetInformacionFinancieraById(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	v, err := models.GetInformacionFinancieraById(creado.Id)
	if err != nil || v.Id != creado.Id || v.MonedaId != creado.MonedaId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetInformacionFinancieraById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetInformacionFinancieraById Finalizado Correctamente (OK)")
}

func TestGetAllInformacionFinanciera(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	helpers.CrearInformacionFinanciera(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllInformacionFinanciera, creado.Id)
	t.Log("TestGetAllInformacionFinanciera Finalizado Correctamente (OK)")
}

func TestUpdateInformacionFinancieraById(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.MonedaId = 2
	if err := models.UpdateInformacionFinancieraById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetInformacionFinancieraById(creado.Id)
	if v == nil || v.MonedaId != 2 {
		t.Fatal("Se esperaba MonedaId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateInformacionFinancieraById(&models.InformacionFinanciera{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateInformacionFinancieraById Finalizado Correctamente (OK)")
}

func TestDeleteInformacionFinanciera(t *testing.T) {
	creado := helpers.CrearInformacionFinanciera(t)
	if err := models.DeleteInformacionFinanciera(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetInformacionFinancieraById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteInformacionFinanciera(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteInformacionFinanciera Finalizado Correctamente (OK)")
}
