package controllers

import (
	"errors"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// leerActivo lee el parametro activo de las consultas: true (predeterminado) para
// registros activos o false para registros inactivos (borrado logico).
func leerActivo(ctrl *beego.Controller) (bool, error) {
	switch ctrl.GetString("activo") {
	case "", "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, errors.New("el parametro activo debe ser true o false")
}

// agregarFiltroActivo agrega a query el filtro por Activo de las consultas GetAll.
// El parametro activo tiene prioridad; si no se envia, se respeta Activo dentro del
// parametro query y, si tampoco esta, se consultan solo los registros activos.
func agregarFiltroActivo(ctrl *beego.Controller, query map[string]string) error {
	if ctrl.GetString("activo") == "" {
		for k := range query {
			if strings.EqualFold(k, "Activo") {
				return nil
			}
		}
	}
	activo, err := leerActivo(ctrl)
	if err != nil {
		return err
	}
	for k := range query {
		if strings.EqualFold(k, "Activo") {
			delete(query, k)
		}
	}
	query["Activo"] = strconv.FormatBool(activo)
	return nil
}
