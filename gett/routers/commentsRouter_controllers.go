package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gett/controllers:DriversController"] = append(beego.GlobalControllerRouter["gett/controllers:DriversController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:DriversController"] = append(beego.GlobalControllerRouter["gett/controllers:DriversController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:DriversController"] = append(beego.GlobalControllerRouter["gett/controllers:DriversController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:DriversController"] = append(beego.GlobalControllerRouter["gett/controllers:DriversController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:DriversController"] = append(beego.GlobalControllerRouter["gett/controllers:DriversController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:MetricsController"] = append(beego.GlobalControllerRouter["gett/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:MetricsController"] = append(beego.GlobalControllerRouter["gett/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:MetricsController"] = append(beego.GlobalControllerRouter["gett/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:MetricsController"] = append(beego.GlobalControllerRouter["gett/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["gett/controllers:MetricsController"] = append(beego.GlobalControllerRouter["gett/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
