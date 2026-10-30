package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaTipoDeclaracionProveedor = "/v1/tipo_declaracion_proveedor"

func TestTipoDeclaracionProveedorPost(t *testing.T) {
	nuevo := helpers.NuevoTipoDeclaracionProveedor(t)
	w := helpers.Peticion("POST", rutaTipoDeclaracionProveedor+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetTipoDeclaracionProveedorById(r.IdDato(t)); err != nil || v.AutorModificacionId != nuevo.AutorModificacionId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarError(t, helpers.Peticion("POST", rutaTipoDeclaracionProveedor+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestTipoDeclaracionProveedorPost Finalizado Correctamente (OK)")
}

func TestTipoDeclaracionProveedorGetOne(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaTipoDeclaracionProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaTipoDeclaracionProveedor, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaTipoDeclaracionProveedor+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestTipoDeclaracionProveedorGetOne Finalizado Correctamente (OK)")
}

func TestTipoDeclaracionProveedorGetAll(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	helpers.CrearTipoDeclaracionProveedor(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaTipoDeclaracionProveedor, creado.Id)
	t.Log("TestTipoDeclaracionProveedorGetAll Finalizado Correctamente (OK)")
}

func TestTipoDeclaracionProveedorActivo(t *testing.T) {
	activo := helpers.CrearTipoDeclaracionProveedor(t)
	inactivo := helpers.CrearTipoDeclaracionProveedor(t)
	if err := models.DeleteTipoDeclaracionProveedor(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaTipoDeclaracionProveedor, activo.Id, inactivo.Id)
	t.Log("TestTipoDeclaracionProveedorActivo Finalizado Correctamente (OK)")
}

func TestTipoDeclaracionProveedorPut(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	ruta := fmt.Sprintf("%s/%d", rutaTipoDeclaracionProveedor, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.AutorModificacionId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetTipoDeclaracionProveedorById(creado.Id); err != nil || v.AutorModificacionId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaTipoDeclaracionProveedor+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaTipoDeclaracionProveedor, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestTipoDeclaracionProveedorPut Finalizado Correctamente (OK)")
}

func TestTipoDeclaracionProveedorDelete(t *testing.T) {
	creado := helpers.CrearTipoDeclaracionProveedor(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaTipoDeclaracionProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetTipoDeclaracionProveedorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaTipoDeclaracionProveedor, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaTipoDeclaracionProveedor+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestTipoDeclaracionProveedorDelete Finalizado Correctamente (OK)")
}
