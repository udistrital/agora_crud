package controllers

import (
	"log"
	"os"
	"testing"

	beego "github.com/beego/beego/v2/server/web"

	// Registra las rutas /v1/... de los controladores
	_ "github.com/udistrital/gestion_personas_crud/routers"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

func TestMain(m *testing.M) {
	limpiar, err := helpers.IniciarBaseDatos()
	if err != nil {
		log.Fatal("No se pudo iniciar la base de datos de pruebas: ", err)
	}
	// Equivale a copyrequestbody = true de conf/app.conf
	beego.BConfig.CopyRequestBody = true
	codigo := m.Run()
	limpiar()
	os.Exit(codigo)
}
