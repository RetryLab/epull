package main

import (
	"flag"
	"fmt"

	"strings"
)

func main() {
	uri := flag.String("uri", "", "you must input the uri. \n For Example: \n replace --uri gcr.io/google_containers/pause:0.9.3")
	flag.Parse()

	if len(*uri) == 0 {
		fmt.Println("ERROR: uri required")
		return
	}
	resultUri := convertUri(*uri)
	fmt.Println("RESULT: \n" + resultUri)
}

func convertUri(uri string) (resultUri string) {
	uriArr := strings.Split(uri, "/")
	arrLength := len(uriArr)

	PREFIX := "epull-registry.yunpro.cn/_"
	DOCKER_HUB := "hub_docker_com___/"

	isDockerHubImage := false
	if arrLength == 3 {
		isDockerHubImage = false
	} else if arrLength <= 2 && arrLength > 0 {
		isDockerHubImage = true
	} else {
		fmt.Println("ERROR Images. ")
		return
	}

	if isDockerHubImage {
		//TODO: later
		if arrLength == 1 {
			//like: ubuntu:latest
			name := uriArr[0]
			resultUri = PREFIX + DOCKER_HUB + name
		} else {
			//like: goyoo/sb:latest
			account := uriArr[0]
			name := uriArr[1]
			resultUri = PREFIX + DOCKER_HUB + account + "/" + name
		}
	} else {
		domain := uriArr[0]
		account := uriArr[1]
		name := uriArr[2]
		domain = strings.Replace(domain, ".", "_", -1)
		resultUri = PREFIX + domain + "___" + account + "/" + name
	}
	return resultUri
}
