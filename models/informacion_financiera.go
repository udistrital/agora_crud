package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type InformacionFinanciera struct {
	Id                  int        `orm:"column(id);pk;auto"`
	ProveedorId         *Proveedor `orm:"column(proveedor_id);rel(fk)"`
	FechaCorte          time.Time  `orm:"column(fecha_corte);type(date)"`
	MonedaId            int        `orm:"column(moneda_id);null"`
	CapitalAutorizado   float64    `orm:"column(capital_autorizado)"`
	ActivosTotales      float64    `orm:"column(activos_totales);null"`
	PasivosTotales      float64    `orm:"column(pasivos_totales);null"`
	IndiceLiquidez      float64    `orm:"column(indice_liquidez)"`
	Activo              bool       `orm:"column(activo)"`
	FechaCreacion       time.Time  `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion   time.Time  `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	AutorModificacionId int        `orm:"column(autor_modificacion_id);null"`
}

func (t *InformacionFinanciera) TableName() string {
	return "informacion_financiera"
}

func init() {
	orm.RegisterModel(new(InformacionFinanciera))
}

// AddInformacionFinanciera insert a new InformacionFinanciera into database and returns
// last inserted Id on success.
func AddInformacionFinanciera(m *InformacionFinanciera) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	m.FechaCreacion = time.Now().UTC()
	if m.FechaCorte.IsZero() {
		var currentYear = time.Now().UTC().Year()
		m.FechaCorte = time.Date(currentYear, 12, 31, 0, 0, 0, 0, time.UTC)
	}
	id, err = o.Insert(m)
	return
}

// GetInformacionFinancieraById retrieves InformacionFinanciera by Id. Returns error if
// Id doesn't exist
func GetInformacionFinancieraById(id int) (v *InformacionFinanciera, err error) {
	o := orm.NewOrm()
	v = &InformacionFinanciera{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInformacionFinanciera retrieves all InformacionFinanciera matches certain condition. Returns empty list if
// no records exist
func GetAllInformacionFinanciera(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InformacionFinanciera)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []InformacionFinanciera
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateInformacionFinanciera updates InformacionFinanciera by Id and returns error if
// the record to be updated doesn't exist
func UpdateInformacionFinancieraById(m *InformacionFinanciera) (err error) {
	o := orm.NewOrm()
	v := InformacionFinanciera{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.FechaCreacion = v.FechaCreacion
		m.FechaModificacion = time.Now().UTC()
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInformacionFinanciera deletes InformacionFinanciera by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInformacionFinanciera(id int) (err error) {
	o := orm.NewOrm()
	v := InformacionFinanciera{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		v.Activo = false
		if num, err = o.Update(&v); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
