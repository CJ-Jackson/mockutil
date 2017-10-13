package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const buildFlag = `// +build debug`

func main() {
	if len(os.Args) < 2 {
		return
	}

	sourceFile := os.Args[1]
	sourceFileSplit := strings.Split(sourceFile, ".")

	destFile := strings.Join(sourceFileSplit[:len(sourceFileSplit)-1], ".") + "_mock.go"
	packageName := os.Args[2]

	buf := &bytes.Buffer{}
	fmt.Fprintln(buf, buildFlag)
	fmt.Fprintln(buf)

	dir, err := os.Getwd()
	if nil != err {
		return
	}

	cmd := exec.Command("mockgen", "-source", sourceFile, "-package", packageName, "-write_package_comment=false")
	cmd.Stdout = buf

	if nil != cmd.Run() {
		return
	}

	file, err := os.Create(dir + "/" + destFile)
	if nil != err {
		return
	}
	defer file.Close()

	io.Copy(file, buf)
}
