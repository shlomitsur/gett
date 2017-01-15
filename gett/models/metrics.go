package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Metrics struct {
	Id        int       `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(255);null"`
	Value     int       `orm:"column(value);null"`
	Lon       float64   `orm:"column(lon);digits(18);decimals(12)"`
	Timestamp time.Time `orm:"column(timestamp);type(datetime)"`
	Lat       float64   `orm:"column(lat);digits(18);decimals(12)"`
	DriverId  int       `orm:"column(driver_id)"`
}

func (t *Metrics) TableName() string {
	return "metrics"
}

func init() {
	orm.RegisterModel(new(Metrics))
}

// AddMetrics insert a new Metrics into database and returns
// last inserted Id on success.
func AddMetrics(m *Metrics) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMetricsById retrieves Metrics by Id. Returns error if
// Id doesn't exist
func GetMetricsById(id int) (v *Metrics, err error) {
	o := orm.NewOrm()
	v = &Metrics{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMetrics retrieves all Metrics matches certain condition. Returns empty list if
// no records exist
func GetAllMetrics(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Metrics))
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

	var l []Metrics
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

// UpdateMetrics updates Metrics by Id and returns error if
// the record to be updated doesn't exist
func UpdateMetricsById(m *Metrics) (err error) {
	o := orm.NewOrm()
	v := Metrics{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMetrics deletes Metrics by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMetrics(id int) (err error) {
	o := orm.NewOrm()
	v := Metrics{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Metrics{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
