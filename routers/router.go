// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/gestion_personas_crud/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/proveedor_natural",
			beego.NSInclude(
				&controllers.ProveedorNaturalController{},
			),
		),

		beego.NSNamespace("/proveedor",
			beego.NSInclude(
				&controllers.ProveedorController{},
			),
		),

		beego.NSNamespace("/perfil",
			beego.NSInclude(
				&controllers.PerfilController{},
			),
		),

		beego.NSNamespace("/proveedor_juridico",
			beego.NSInclude(
				&controllers.ProveedorJuridicoController{},
			),
		),

		beego.NSNamespace("/contacto",
			beego.NSInclude(
				&controllers.ContactoController{},
			),
		),

		beego.NSNamespace("/cuenta_bancaria",
			beego.NSInclude(
				&controllers.CuentaBancariaController{},
			),
		),

		beego.NSNamespace("/actividad_economica_proveedor",
			beego.NSInclude(
				&controllers.ActividadEconomicaProveedorController{},
			),
		),

		beego.NSNamespace("/perfil_financiero",
			beego.NSInclude(
				&controllers.PerfilFinancieroController{},
			),
		),

		beego.NSNamespace("/prefijo_facturacion",
			beego.NSInclude(
				&controllers.PrefijoFacturacionController{},
			),
		),

		beego.NSNamespace("/rango_facturacion",
			beego.NSInclude(
				&controllers.RangoFacturacionController{},
			),
		),

		beego.NSNamespace("/responsabilidad_fiscal_perfil",
			beego.NSInclude(
				&controllers.ResponsabilidadFiscalPerfilController{},
			),
		),

		beego.NSNamespace("/informacion_financiera",
			beego.NSInclude(
				&controllers.InformacionFinancieraController{},
			),
		),

		beego.NSNamespace("/representacion",
			beego.NSInclude(
				&controllers.RepresentacionController{},
			),
		),

		beego.NSNamespace("/documento",
			beego.NSInclude(
				&controllers.DocumentoController{},
			),
		),

		beego.NSNamespace("/tipo_declaracion_proveedor",
			beego.NSInclude(
				&controllers.TipoDeclaracionProveedorController{},
			),
		),

		beego.NSNamespace("/declaracion",
			beego.NSInclude(
				&controllers.DeclaracionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
