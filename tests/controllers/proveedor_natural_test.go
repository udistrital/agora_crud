package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaProveedorNatural = "/v1/proveedor_natural"

func TestProveedorNaturalPost(t *testing.T) {
	nuevo := helpers.NuevoProveedorNatural(t)
	w := helpers.Peticion("POST", rutaProveedorNatural+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetProveedorNaturalById(r.IdDato(t)); err != nil || v.ExperienciaLaboralMeses != nuevo.ExperienciaLaboralMeses || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarError(t, helpers.Peticion("POST", rutaProveedorNatural+"/", `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestProveedorNaturalPost Finalizado Correctamente (OK)")
}

func TestProveedorNaturalGetOne(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedorNatural, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaProveedorNatural, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaProveedorNatural+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestProveedorNaturalGetOne Finalizado Correctamente (OK)")
}

func TestProveedorNaturalGetAll(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	helpers.CrearProveedorNatural(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaProveedorNatural, creado.Id)
	t.Log("TestProveedorNaturalGetAll Finalizado Correctamente (OK)")
}

func TestProveedorNaturalActivo(t *testing.T) {
	activo := helpers.CrearProveedorNatural(t)
	inactivo := helpers.CrearProveedorNatural(t)
	if err := models.DeleteProveedorNatural(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaProveedorNatural, activo.Id, inactivo.Id)
	t.Log("TestProveedorNaturalActivo Finalizado Correctamente (OK)")
}

func TestProveedorNaturalPut(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	ruta := fmt.Sprintf("%s/%d", rutaProveedorNatural, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.ExperienciaLaboralMeses = 24
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetProveedorNaturalById(creado.Id); err != nil || v.ExperienciaLaboralMeses != 24 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest, "el Id del cuerpo no coincide con el id de la ruta")
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaProveedorNatural+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaProveedorNatural, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestProveedorNaturalPut Finalizado Correctamente (OK)")
}

func TestProveedorNaturalDelete(t *testing.T) {
	creado := helpers.CrearProveedorNatural(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaProveedorNatural, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetProveedorNaturalById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaProveedorNatural, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaProveedorNatural+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestProveedorNaturalDelete Finalizado Correctamente (OK)")
}
