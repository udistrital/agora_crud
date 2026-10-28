package helpers

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
)

// IdInexistente es un Id que nunca se crea en la base de datos de pruebas.
const IdInexistente = 999999

// MismaFecha indica si dos fechas corresponden al mismo instante, tolerando
// la pérdida de precisión al guardarlas en la base de datos.
func MismaFecha(a, b time.Time) bool {
	d := a.Sub(b)
	return d > -time.Second && d < time.Second
}

// Las funciones Nuevo<Modelo> retornan un registro de ejemplo sin guardar
// (creando en la base de datos sus llaves foráneas) y las funciones
// Crear<Modelo> lo guardan en la base de datos.

func NuevoActividadEconomicaProveedor(t *testing.T) *models.ActividadEconomicaProveedor {
	t.Helper()
	return &models.ActividadEconomicaProveedor{ProveedorId: CrearProveedor(t), ActividadEconomicaId: 1, EsPrincipal: true}
}

func CrearActividadEconomicaProveedor(t *testing.T) *models.ActividadEconomicaProveedor {
	t.Helper()
	v := NuevoActividadEconomicaProveedor(t)
	if _, err := models.AddActividadEconomicaProveedor(v); err != nil {
		t.Fatal("No se pudo crear ActividadEconomicaProveedor:", err)
	}
	return v
}

func NuevoContacto(t *testing.T) *models.Contacto {
	t.Helper()
	return &models.Contacto{ProveedorId: CrearProveedor(t), NombreContacto: "Contacto prueba", Finalidad: "Notificaciones", TipoCanal: "Correo", ValorEspecifico: "prueba@udistrital.edu.co"}
}

func CrearContacto(t *testing.T) *models.Contacto {
	t.Helper()
	v := NuevoContacto(t)
	if _, err := models.AddContacto(v); err != nil {
		t.Fatal("No se pudo crear Contacto:", err)
	}
	return v
}

func NuevoCuentaBancaria(t *testing.T) *models.CuentaBancaria {
	t.Helper()
	return &models.CuentaBancaria{ProveedorId: CrearProveedor(t), EntidadBancariaId: 1, TipoCuentaId: 1, NumeroCuenta: "123456789", NombreTitular: "Titular prueba", EsPrincipal: true}
}

func CrearCuentaBancaria(t *testing.T) *models.CuentaBancaria {
	t.Helper()
	v := NuevoCuentaBancaria(t)
	if _, err := models.AddCuentaBancaria(v); err != nil {
		t.Fatal("No se pudo crear CuentaBancaria:", err)
	}
	return v
}

func NuevoDeclaracion(t *testing.T) *models.Declaracion {
	t.Helper()
	return &models.Declaracion{TipoDeclaracionId: 1, Texto: "Declaración prueba"}
}

func CrearDeclaracion(t *testing.T) *models.Declaracion {
	t.Helper()
	v := NuevoDeclaracion(t)
	if _, err := models.AddDeclaracion(v); err != nil {
		t.Fatal("No se pudo crear Declaracion:", err)
	}
	return v
}

func NuevoDocumento(t *testing.T) *models.Documento {
	t.Helper()
	return &models.Documento{ProveedorId: CrearProveedor(t), TipoDocumentoId: 1, Enlace: "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b"}
}

func CrearDocumento(t *testing.T) *models.Documento {
	t.Helper()
	v := NuevoDocumento(t)
	if _, err := models.AddDocumento(v); err != nil {
		t.Fatal("No se pudo crear Documento:", err)
	}
	return v
}

func NuevoInformacionFinanciera(t *testing.T) *models.InformacionFinanciera {
	t.Helper()
	return &models.InformacionFinanciera{ProveedorId: CrearProveedor(t), FechaCorte: time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), MonedaId: 1, CapitalAutorizado: 1000000, ActivosTotales: 500000, PasivosTotales: 200000, IndiceLiquidez: 2.5}
}

func CrearInformacionFinanciera(t *testing.T) *models.InformacionFinanciera {
	t.Helper()
	v := NuevoInformacionFinanciera(t)
	if _, err := models.AddInformacionFinanciera(v); err != nil {
		t.Fatal("No se pudo crear InformacionFinanciera:", err)
	}
	return v
}

func NuevoPerfilFinanciero(t *testing.T) *models.PerfilFinanciero {
	t.Helper()
	return &models.PerfilFinanciero{ProveedorId: CrearProveedorJuridico(t), PrefijoFacturacionId: CrearPrefijoFacturacion(t), RangoFacturacionId: CrearRangoFacturacion(t), GranContribuyente: true, ResolucionFacturacion: "Resolución 001"}
}

func CrearPerfilFinanciero(t *testing.T) *models.PerfilFinanciero {
	t.Helper()
	v := NuevoPerfilFinanciero(t)
	if _, err := models.AddPerfilFinanciero(v); err != nil {
		t.Fatal("No se pudo crear PerfilFinanciero:", err)
	}
	return v
}

func NuevoPerfil(t *testing.T) *models.Perfil {
	t.Helper()
	return &models.Perfil{Nombre: "Perfil prueba", Descripcion: "Descripción prueba", CodigoAbreviacion: "PP", NumeroOrden: 1}
}

func CrearPerfil(t *testing.T) *models.Perfil {
	t.Helper()
	v := NuevoPerfil(t)
	if _, err := models.AddPerfil(v); err != nil {
		t.Fatal("No se pudo crear Perfil:", err)
	}
	return v
}

func NuevoPrefijoFacturacion(t *testing.T) *models.PrefijoFacturacion {
	t.Helper()
	return &models.PrefijoFacturacion{Nombre: "Prefijo prueba", Descripcion: "Descripción prueba", CodigoAbreviacion: "PF", NumeroOrden: 1}
}

func CrearPrefijoFacturacion(t *testing.T) *models.PrefijoFacturacion {
	t.Helper()
	v := NuevoPrefijoFacturacion(t)
	if _, err := models.AddPrefijoFacturacion(v); err != nil {
		t.Fatal("No se pudo crear PrefijoFacturacion:", err)
	}
	return v
}

func NuevoProveedor(t *testing.T) *models.Proveedor {
	t.Helper()
	return &models.Proveedor{TerceroId: 1, TipoRegistroId: 1, DescripcionPortafolio: "Portafolio prueba"}
}

func CrearProveedor(t *testing.T) *models.Proveedor {
	t.Helper()
	v := NuevoProveedor(t)
	if _, err := models.AddProveedor(v); err != nil {
		t.Fatal("No se pudo crear Proveedor:", err)
	}
	return v
}

func NuevoProveedorJuridico(t *testing.T) *models.ProveedorJuridico {
	t.Helper()
	return &models.ProveedorJuridico{NombreComercial: "Empresa prueba", MatriculaMercantil: 12345, FechaConstitucion: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC), FechaRenovacion: time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)}
}

func CrearProveedorJuridico(t *testing.T) *models.ProveedorJuridico {
	t.Helper()
	v := NuevoProveedorJuridico(t)
	if _, err := models.AddProveedorJuridico(v); err != nil {
		t.Fatal("No se pudo crear ProveedorJuridico:", err)
	}
	return v
}

func NuevoProveedorNatural(t *testing.T) *models.ProveedorNatural {
	t.Helper()
	return &models.ProveedorNatural{PerfilDeclarado: CrearPerfil(t), ExperienciaLaboralMeses: 12, ExperienciaProfesionalMeses: 6}
}

func CrearProveedorNatural(t *testing.T) *models.ProveedorNatural {
	t.Helper()
	v := NuevoProveedorNatural(t)
	if _, err := models.AddProveedorNatural(v); err != nil {
		t.Fatal("No se pudo crear ProveedorNatural:", err)
	}
	return v
}

func NuevoRangoFacturacion(t *testing.T) *models.RangoFacturacion {
	t.Helper()
	return &models.RangoFacturacion{Nombre: "Rango prueba", Descripcion: "Descripción prueba", CodigoAbreviacion: "RF", NumeroOrden: 1}
}

func CrearRangoFacturacion(t *testing.T) *models.RangoFacturacion {
	t.Helper()
	v := NuevoRangoFacturacion(t)
	if _, err := models.AddRangoFacturacion(v); err != nil {
		t.Fatal("No se pudo crear RangoFacturacion:", err)
	}
	return v
}

func NuevoRepresentacion(t *testing.T) *models.Representacion {
	t.Helper()
	return &models.Representacion{ProveedorId: CrearProveedorJuridico(t), RepresentanteId: 1, TipoRepresentacionId: 1, CargoId: 1, DescripcionFacultades: "Facultades prueba"}
}

func CrearRepresentacion(t *testing.T) *models.Representacion {
	t.Helper()
	v := NuevoRepresentacion(t)
	if _, err := models.AddRepresentacion(v); err != nil {
		t.Fatal("No se pudo crear Representacion:", err)
	}
	return v
}

func NuevoResponsabilidadFiscalPerfil(t *testing.T) *models.ResponsabilidadFiscalPerfil {
	t.Helper()
	return &models.ResponsabilidadFiscalPerfil{PerfilFinancieroId: CrearPerfilFinanciero(t), ResponsabilidadFiscalId: 1}
}

func CrearResponsabilidadFiscalPerfil(t *testing.T) *models.ResponsabilidadFiscalPerfil {
	t.Helper()
	v := NuevoResponsabilidadFiscalPerfil(t)
	if _, err := models.AddResponsabilidadFiscalPerfil(v); err != nil {
		t.Fatal("No se pudo crear ResponsabilidadFiscalPerfil:", err)
	}
	return v
}

func NuevoTipoDeclaracionProveedor(t *testing.T) *models.TipoDeclaracionProveedor {
	t.Helper()
	return &models.TipoDeclaracionProveedor{ProveedorId: CrearProveedor(t), DeclaracionId: CrearDeclaracion(t)}
}

func CrearTipoDeclaracionProveedor(t *testing.T) *models.TipoDeclaracionProveedor {
	t.Helper()
	v := NuevoTipoDeclaracionProveedor(t)
	if _, err := models.AddTipoDeclaracionProveedor(v); err != nil {
		t.Fatal("No se pudo crear TipoDeclaracionProveedor:", err)
	}
	return v
}
