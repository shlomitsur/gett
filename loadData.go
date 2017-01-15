package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const drivers_file string = "./data/drivers.json"
const metrics_file string = "./data/metrics.json"

//using plural noun because of the db tables 'drivers' and 'metrics' (the ORM assumes it)
type Drivers struct {
	Id            int    `orm:"column(id);auto" json:"id"`
	Name          string `orm:"column(name);size(255);null" json:"name"`
	LicenseNumber string `orm:"column(license_number);size(255);null" json:"license_number"`
}

type Metrics struct {
	Id        int     `orm:"column(id);auto"`
	Name      string  `orm:"column(name);size(255);null" json:"metric_name"`
	Value     string  `orm:"column(value);null" json:"value"`
	Lon       float64 `orm:"column(lon);digits(18);decimals(12)" json:"lon"`
	Timestamp int64   `orm:"column(timestamp);type(datetime)" json:"timestamp"`
	Lat       float64 `orm:"column(lat);digits(18);decimals(12)" json:"lat"`
	DriverId  int     `orm:"column(driver_id)" json:"driver_id,string"`
}

type ModuleSource interface {
	location() string
}

func (this Drivers) location() string {
	return drivers_file
}

func (this Metrics) location() string {
	return metrics_file
}

func main() {

	var ormInstance orm.Ormer = initORM()
	var drivers Drivers
	var metrics Metrics
	processModule(ormInstance, drivers)
	processModule(ormInstance, metrics)
}

func initORM() orm.Ormer {
	// register models
	orm.RegisterModel(new(Drivers))
	orm.RegisterModel(new(Metrics))
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/gett?charset=utf8", 30)
	o := orm.NewOrm()
	return o
}

func hasJsonBrackets(r *bufio.Reader) bool {
	untilBraces, err := r.Peek(10)
	if err != nil {
		fmt.Println("Error peeking reader")
		return false
	}
	return strings.Contains(string(untilBraces), "[")
}

func processModule(ormInstance orm.Ormer, moduleSource ModuleSource) error {
	file, err := os.Open(moduleSource.location())
	if err != nil {
		fmt.Println("Error opening file ", moduleSource.location())
		return err
	}

	defer file.Close()
	r := bufio.NewReader(file)
	hasBrackets := hasJsonBrackets(r)
	dec := json.NewDecoder(r)

	//metrics.json does not have surrounding brackets so we're not using json.unmarshal (although we can do in the its callback)
	if hasBrackets {
		// read open bracket
		_, err := dec.Token()
		if err != nil {
			fmt.Printf("Error while parsing json ", err)
		}
	}
	for dec.More() {
		switch v := moduleSource.(type) {
		case Drivers:
			decodeErr := dec.Decode(&v)
			_, ormError := ormInstance.Insert(&v)
			if decodeErr != nil {
				fmt.Printf("Decode err %v\n", decodeErr)
				continue
			}
			if ormError != nil {
				fmt.Printf("ORM err %v\n", ormError)
				continue
			}
			fmt.Printf("%v\n", v.Name)
		//There is duplicate code here because the ORM does not allow Interface{}, Go allows 'case Drivers, Metrics:' but
		//That does not preseves the cast, I'm stackoverflowing this, now answer yet..
		case Metrics:
			decodeErr := dec.Decode(&v)
			fmt.Printf("name %v, value %v , lon %v , timestamp, %v, lat %v , driverId %v\n", v.Name, v.Value, v.Lon, v.Timestamp, v.Lat, v.DriverId)
			if v.DriverId == 0 {
				fmt.Printf("Skipping empty driver_id item")
				continue
			}
			_, ormError := ormInstance.Insert(&v)
			if decodeErr != nil {
				fmt.Printf("Decode err %v\n", decodeErr)
				continue
			}
			if ormError != nil {
				fmt.Printf("ORM err %v\n", ormError)
				continue
			}
			fmt.Printf("%v\n", v.Name)
		default:
			fmt.Println("unknown type")
		}
	}
	if hasBrackets {
		// read closing bracket
		_, err := dec.Token()
		if err != nil {
			fmt.Printf("Error while parsing json ", err)
		}
	}
	return nil
}

