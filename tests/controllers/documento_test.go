package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

const rutaDocumento = "/v1/documento"

func TestDocumentoPost(t *testing.T) {
	nuevo := helpers.NuevoDocumento(t)
	w := helpers.Peticion("POST", rutaDocumento+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetDocumentoById(r.IdDato(t)); err != nil || v.TipoDocumentoId != nuevo.TipoDocumentoId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	helpers.VerificarRespuesta(t, helpers.Peticion("POST", rutaDocumento+"/", `{"Id":"texto"}`), http.StatusBadRequest)
	t.Log("TestDocumentoPost Finalizado Correctamente (OK)")
}

func TestDocumentoGetOne(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDocumento, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDocumento, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestDocumentoGetOne Finalizado Correctamente (OK)")
}

func TestDocumentoGetAll(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	helpers.CrearDocumento(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaDocumento, creado.Id)
	t.Log("TestDocumentoGetAll Finalizado Correctamente (OK)")
}

func TestDocumentoPut(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	ruta := fmt.Sprintf("%s/%d", rutaDocumento, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.TipoDocumentoId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetDocumentoById(creado.Id); err != nil || v.TipoDocumentoId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	cambios.Id = creado.Id + 1
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusBadRequest)
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest)
	cambios.Id = 0
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaDocumento, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarRespuesta(t, w, http.StatusBadRequest)
	t.Log("TestDocumentoPut Finalizado Correctamente (OK)")
}

func TestDocumentoDelete(t *testing.T) {
	creado := helpers.CrearDocumento(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaDocumento, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetDocumentoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaDocumento, helpers.IdInexistente), "")
	helpers.VerificarRespuesta(t, w, http.StatusNotFound)
	t.Log("TestDocumentoDelete Finalizado Correctamente (OK)")
}
