package helpers

import (
	"strconv"
	"testing"
)

// FuncionGetAll corresponde a la firma de las funciones GetAll<Modelo> del
// paquete models.
type FuncionGetAll func(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) ([]interface{}, error)

// ProbarGetAll verifica el comportamiento común de una función GetAll<Modelo>:
// filtros, selección de campos, ordenamiento, paginación y validaciones.
// Requiere que exista en la base de datos el registro con Id idExistente.
func ProbarGetAll(t *testing.T, getAll FuncionGetAll, idExistente int) {
	t.Helper()
	id := strconv.Itoa(idExistente)

	t.Run("SinFiltros", func(t *testing.T) {
		l, err := getAll(map[string]string{}, nil, nil, nil, 0, 0)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros y se obtuvo", len(l), err)
		}
	})

	t.Run("FiltroPorId", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id": id}, nil, nil, nil, 0, 10)
		if err != nil || len(l) != 1 {
			t.Error("Se esperaba 1 registro y se obtuvo", len(l), err)
		}
	})

	t.Run("FiltroIsNull", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id__isnull": "false"}, nil, nil, nil, 0, 10)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros y se obtuvo", len(l), err)
		}
	})

	t.Run("SinResultados", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id": strconv.Itoa(IdInexistente)}, nil, nil, nil, 0, 10)
		if err != nil || l != nil {
			t.Error("Se esperaba una lista vacía y se obtuvo", l, err)
		}
	})

	t.Run("SeleccionDeCampos", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id": id}, []string{"Id", "Activo"}, nil, nil, 0, 10)
		if err != nil || len(l) != 1 {
			t.Fatal("Se esperaba 1 registro y se obtuvo", len(l), err)
		}
		m, ok := l[0].(map[string]interface{})
		if !ok || len(m) != 2 || m["Id"] != idExistente {
			t.Error("Se esperaba un mapa con los campos Id y Activo y se obtuvo", l[0])
		}
	})

	t.Run("OrdenPorCampo", func(t *testing.T) {
		l, err := getAll(map[string]string{}, nil, []string{"Id", "Activo"}, []string{"desc", "asc"}, 0, 10)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros ordenados y se obtuvo", len(l), err)
		}
	})

	t.Run("OrdenUnicoParaVariosCampos", func(t *testing.T) {
		for _, orden := range []string{"asc", "desc"} {
			l, err := getAll(map[string]string{}, nil, []string{"Id", "Activo"}, []string{orden}, 0, 10)
			if err != nil || len(l) == 0 {
				t.Error("Se esperaban registros ordenados", orden, "y se obtuvo", len(l), err)
			}
		}
	})

	t.Run("Paginacion", func(t *testing.T) {
		l, err := getAll(map[string]string{}, nil, nil, nil, 0, 1)
		if err != nil || len(l) != 1 {
			t.Error("Se esperaba 1 registro con limit=1 y se obtuvo", len(l), err)
		}
	})

	t.Run("CampoInexistente", func(t *testing.T) {
		if _, err := getAll(map[string]string{}, []string{"CampoInexistente"}, nil, nil, 0, 10); err == nil {
			t.Error("Se esperaba error al solicitar un campo inexistente")
		}
	})

	casosOrdenInvalido := []struct {
		nombre string
		sortby []string
		order  []string
	}{
		{"OrdenInvalidoPorCampo", []string{"Id"}, []string{"abc"}},
		{"OrdenUnicoInvalido", []string{"Id", "Activo"}, []string{"abc"}},
		{"TamanosNoCoinciden", []string{"Id", "Activo"}, []string{"asc", "desc", "asc"}},
		{"OrderSinSortby", nil, []string{"asc"}},
	}
	for _, c := range casosOrdenInvalido {
		t.Run(c.nombre, func(t *testing.T) {
			if _, err := getAll(map[string]string{}, nil, c.sortby, c.order, 0, 10); err == nil {
				t.Error("Se esperaba error con sortby", c.sortby, "y order", c.order)
			}
		})
	}
}
