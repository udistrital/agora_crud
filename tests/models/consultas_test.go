package models

import (
	"strconv"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

// TestRecuperarConsulta verifica que el panic que lanza el ORM con campos
// inexistentes se convierta en error en todas las funciones GetAll.
func TestRecuperarConsulta(t *testing.T) {
	funciones := map[string]helpers.FuncionGetAll{
		"ActividadEconomicaProveedor": models.GetAllActividadEconomicaProveedor,
		"Contacto":                    models.GetAllContacto,
		"CuentaBancaria":              models.GetAllCuentaBancaria,
		"Declaracion":                 models.GetAllDeclaracion,
		"Documento":                   models.GetAllDocumento,
		"InformacionFinanciera":       models.GetAllInformacionFinanciera,
		"Perfil":                      models.GetAllPerfil,
		"PerfilFinanciero":            models.GetAllPerfilFinanciero,
		"PrefijoFacturacion":          models.GetAllPrefijoFacturacion,
		"Proveedor":                   models.GetAllProveedor,
		"ProveedorJuridico":           models.GetAllProveedorJuridico,
		"ProveedorNatural":            models.GetAllProveedorNatural,
		"RangoFacturacion":            models.GetAllRangoFacturacion,
		"Representacion":              models.GetAllRepresentacion,
		"ResponsabilidadFiscalPerfil": models.GetAllResponsabilidadFiscalPerfil,
		"TipoDeclaracionProveedor":    models.GetAllTipoDeclaracionProveedor,
	}
	for nombre, getAll := range funciones {
		t.Run(nombre, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("No se esperaba un panic y se obtuvo", r)
				}
			}()
			if _, _, err := getAll(map[string]string{"CampoInexistente": "1"}, nil, nil, nil, 0, 10); err == nil {
				t.Error("Se esperaba error al filtrar por un campo inexistente")
			}
		})
	}
	t.Log("TestRecuperarConsulta Finalizado Correctamente (OK)")
}

// TestFiltroBooleano verifica que los filtros del parámetro query sobre campos
// booleanos, propios o de una relación, se comparen como booleanos.
func TestFiltroBooleano(t *testing.T) {
	actividad := helpers.CrearActividadEconomicaProveedor(t) // EsPrincipal = true
	inactiva := helpers.CrearActividadEconomicaProveedor(t)
	if err := models.DeleteActividadEconomicaProveedor(inactiva.Id); err != nil {
		t.Fatal("No se pudo inactivar la actividad:", err)
	}
	id := strconv.Itoa(actividad.Id)
	idInactiva := strconv.Itoa(inactiva.Id)

	casos := []struct {
		nombre   string
		query    map[string]string
		esperado int64
	}{
		{"CampoPropioTrue", map[string]string{"Id": id, "Activo": "true"}, 1},
		{"CampoPropioFalse", map[string]string{"Id": idInactiva, "Activo": "false"}, 1},
		{"CampoPropioEnMinuscula", map[string]string{"Id": idInactiva, "activo": "false"}, 1},
		{"CampoPropioNoCoincide", map[string]string{"Id": id, "Activo": "false"}, 0},
		{"OtroCampoBooleano", map[string]string{"Id": id, "EsPrincipal": "true"}, 1},
		{"CampoDeRelacion", map[string]string{"Id": id, "ProveedorId__Activo": "true"}, 1},
		{"CampoDeRelacionNoCoincide", map[string]string{"Id": id, "ProveedorId__Activo": "false"}, 0},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			if _, total, err := models.GetAllActividadEconomicaProveedor(c.query, nil, nil, nil, 0, 10); err != nil || total != c.esperado {
				t.Errorf("Se esperaba un total de %d con %v y se obtuvo %d %v", c.esperado, c.query, total, err)
			}
		})
	}
	t.Log("TestFiltroBooleano Finalizado Correctamente (OK)")
}
