package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Tools
type ToolController struct {
	beego.Controller
}

const (
	PREFIX          = "epull-registry.yunpro.cn"
	DOMAIN_PREFIX   = "ss-"
	HTTPS_PREFIX    = "https://epull-registry.yunpro.cn"
	DOT_SPLITER     = "-dd-"
	ACCOUNT_SPLITER = "-ee-"
	DOCKER_HUB      = "ss-dockerhub-com"
)

// @Title createTool
// @Description create users
// @Param	body		body 	models.Tool	true		"body for user content"
// @Success 200 {int} models.Tool.Id
// @Failure 403 body is empty
// @router /encode-uri [post]
func (w *ToolController) Post() {
	uri := w.Ctx.Request.Form.Get("uri")
	fmt.Println("request uri " + uri)
	targetUri, _, v2Uri, validUri := encodeUri(uri)
	if validUri {
		fmt.Println("target uri:" + targetUri)
		fmt.Println("v2 uri:" + v2Uri)
		w.Ctx.Output.Body([]byte("SUCCESS\n" + targetUri + "\n" + v2Uri))
	} else {
		fmt.Println("Invalid URI: " + uri)
		w.Ctx.Output.Body([]byte("INVALID"))
	}
}

func encodeUri(uri string) (resultUri string, v1Uri string, v2Uri string, validUri bool) {
	uriArr := strings.Split(uri, "/")
	arrLength := len(uriArr)

	//deal with 3 different docker images uri
	if arrLength == 1 {
		//like: ubuntu:latest
		name, version := getNameVersion(uriArr[0])
		resultUri = PREFIX + "/" + DOMAIN_PREFIX + DOCKER_HUB + ACCOUNT_SPLITER + "library/" + name + ":" + version
		v1Uri = HTTPS_PREFIX + "/v1/" + DOMAIN_PREFIX + DOCKER_HUB + ACCOUNT_SPLITER + "library/" + name + "/manifests/" + version
		v2Uri = HTTPS_PREFIX + "/v2/" + DOMAIN_PREFIX + DOCKER_HUB + ACCOUNT_SPLITER + "library/" + name + "/manifests/" + version
	} else if arrLength == 2 {
		//like: tad/tools:latest
		account := uriArr[0]
		name, version := getNameVersion(uriArr[1])
		resultUri = PREFIX + "/" + DOMAIN_PREFIX + DOCKER_HUB + ACCOUNT_SPLITER + account + "/" + name + ":" + version
		v1Uri = HTTPS_PREFIX + "/v1/" + DOMAIN_PREFIX + DOCKER_HUB + ACCOUNT_SPLITER + account + "/" + name + "/manifests/" + version
		v2Uri = HTTPS_PREFIX + "/v2/" + DOMAIN_PREFIX + DOCKER_HUB + ACCOUNT_SPLITER + account + "/" + name + "/manifests/" + version

	} else if arrLength == 3 {
		//like: gcr.io/google/pause:latest
		domain := strings.Replace(uriArr[0], ".", DOT_SPLITER, -1)
		account := uriArr[1]
		name, version := getNameVersion(uriArr[2])
		resultUri = PREFIX + "/" + DOMAIN_PREFIX + domain + ACCOUNT_SPLITER + account + "/" + name + ":" + version
		v1Uri = HTTPS_PREFIX + "/v1/" + DOMAIN_PREFIX + domain + ACCOUNT_SPLITER + account + "/" + name + "/manifests/" + version
		v2Uri = HTTPS_PREFIX + "/v2/" + DOMAIN_PREFIX + domain + ACCOUNT_SPLITER + account + "/" + name + "/manifests/" + version

	} else {
		fmt.Println("ERROR Images. ")
		return "", "", "", false
	}

	return resultUri, v1Uri, v2Uri, true
}

func getNameVersion(fullName string) (name string, version string) {
	arr := strings.Split(fullName, ":")
	arrLength := len(arr)
	if arrLength == 1 {
		return arr[0], "latest"
	} else {
		return arr[0], arr[1]
	}
}
