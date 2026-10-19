package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddDocumento(t *testing.T) {
	v := helpers.NuevoDocumento(t)
	v.Activo = false
	id, err := models.AddDocumento(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddDocumento Finalizado Correctamente (OK)")
}

func TestGetDocumentoById(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	v, err := models.GetDocumentoById(creado.Id)
	if err != nil || v.Id != creado.Id || v.TipoDocumentoId != creado.TipoDocumentoId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetDocumentoById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetDocumentoById Finalizado Correctamente (OK)")
}

func TestGetAllDocumento(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	helpers.ProbarGetAll(t, models.GetAllDocumento, creado.Id)
	t.Log("TestGetAllDocumento Finalizado Correctamente (OK)")
}

func TestUpdateDocumentoById(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.TipoDocumentoId = 2
	if err := models.UpdateDocumentoById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetDocumentoById(creado.Id)
	if v == nil || v.TipoDocumentoId != 2 {
		t.Fatal("Se esperaba TipoDocumentoId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateDocumentoById(&models.Documento{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateDocumentoById Finalizado Correctamente (OK)")
}

func TestDeleteDocumento(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	if err := models.DeleteDocumento(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetDocumentoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteDocumento(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteDocumento Finalizado Correctamente (OK)")
}
