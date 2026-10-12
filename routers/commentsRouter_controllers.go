package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DeclaracionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:InformacionFinancieraController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilFinancieroController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PrefijoFacturacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RangoFacturacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ResponsabilidadFiscalPerfilController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TipoDeclaracionProveedorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
