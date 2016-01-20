package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["poster/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poster/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poster/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poster/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poster/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:ObjectController"] = append(beego.GlobalControllerRouter["poster/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:UserController"] = append(beego.GlobalControllerRouter["poster/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:UserController"] = append(beego.GlobalControllerRouter["poster/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:UserController"] = append(beego.GlobalControllerRouter["poster/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:UserController"] = append(beego.GlobalControllerRouter["poster/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:UserController"] = append(beego.GlobalControllerRouter["poster/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:UserController"] = append(beego.GlobalControllerRouter["poster/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:UserController"] = append(beego.GlobalControllerRouter["poster/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:WorkerController"] = append(beego.GlobalControllerRouter["poster/controllers:WorkerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:WorkerController"] = append(beego.GlobalControllerRouter["poster/controllers:WorkerController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

}
