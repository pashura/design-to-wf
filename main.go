package main

import (
	"flag"
	"github.com/pashura/design-to-wf/api"
)

func main() {
	javaPackageName := flag.String("repo", "testPackage", "Java Package Name")
	flag.Parse()

	api.Run(*javaPackageName)
}
