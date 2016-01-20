package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["poster/controllers:ToolController"] = append(beego.GlobalControllerRouter["poster/controllers:ToolController"],
		beego.ControllerComments{
			"Post",
			`/encode-uri`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["poster/controllers:WorkerController"] = append(beego.GlobalControllerRouter["poster/controllers:WorkerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
