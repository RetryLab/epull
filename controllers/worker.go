package controllers

import (
	"fmt"
	"os/exec"

	"github.com/astaxie/beego"
)

// Operations about Workers
type WorkerController struct {
	beego.Controller
}

// @Title createWorker
// @Description create users
// @Param	body		body 	models.Worker	true		"body for user content"
// @Success 200 {int} models.Worker.Id
// @Failure 403 body is empty
// @router / [post]
func (w *WorkerController) Post() {
	ori := w.Ctx.Request.Form.Get("ori")
	target := w.Ctx.Request.Form.Get("target")
	workerCommand := beego.AppConfig.String("workercommand")

	fmt.Println("command: " + workerCommand + " " + ori + " " + target)

	cmd := exec.Command(workerCommand, ori, target)
	cmd.Start()

	w.Ctx.Output.Body([]byte(""))
}
