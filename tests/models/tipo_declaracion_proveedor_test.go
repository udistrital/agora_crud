package models

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestAddTipoDeclaracionProveedor(t *testing.T) {
	v := helpers.NuevoTipoDeclaracionProveedor(t)
	v.Activo = false
	id, err := models.AddTipoDeclaracionProveedor(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() {
		t.Error("Se esperaba Activo=true y FechaCreacion asignada y se obtuvo", v.Activo, v.FechaCreacion)
	}
	t.Log("TestAddTipoDeclaracionProveedor Finalizado Correctamente (OK)")
}

func TestGetTipoDeclaracionProveedorById(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	v, err := models.GetTipoDeclaracionProveedorById(creado.Id)
	if err != nil || v.Id != creado.Id || v.AutorModificacionId != creado.AutorModificacionId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetTipoDeclaracionProveedorById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetTipoDeclaracionProveedorById Finalizado Correctamente (OK)")
}

func TestGetAllTipoDeclaracionProveedor(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	helpers.CrearTipoDeclaracionProveedor(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllTipoDeclaracionProveedor, creado.Id)
	t.Log("TestGetAllTipoDeclaracionProveedor Finalizado Correctamente (OK)")
}

func TestUpdateTipoDeclaracionProveedorById(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.AutorModificacionId = 2
	if err := models.UpdateTipoDeclaracionProveedorById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetTipoDeclaracionProveedorById(creado.Id)
	if v == nil || v.AutorModificacionId != 2 {
		t.Fatal("Se esperaba AutorModificacionId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.IsZero() {
		t.Error("Se esperaba conservar FechaCreacion y asignar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateTipoDeclaracionProveedorById(&models.TipoDeclaracionProveedor{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateTipoDeclaracionProveedorById Finalizado Correctamente (OK)")
}

func TestDeleteTipoDeclaracionProveedor(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	if err := models.DeleteTipoDeclaracionProveedor(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetTipoDeclaracionProveedorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteTipoDeclaracionProveedor(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteTipoDeclaracionProveedor Finalizado Correctamente (OK)")
}
