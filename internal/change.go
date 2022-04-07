package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/self-denying/goscript/consts"
	"io"
	"os"
	"os/exec"
)

/**
 * Created by : GoLand
 * User: ruohuai
 * Date: 2022/4/7
 * Time: 14:43
 */

func BecomeScripts() {
	if len(os.Args) != 2 {
		panic("Please enter the name of the file to execute")
	}
	inputFileHandle, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Open file for execution exception:%v", err.Error()))
	}
	defer inputFileHandle.Close()
	reader := bufio.NewReader(inputFileHandle)
	tempFileHandle, err := os.Create(consts.TempFile)
	if err != nil {
		panic(fmt.Sprintf("Create temp file exception:%v", err.Error()))
	}
	defer os.Remove(consts.TempFile)
	defer tempFileHandle.Close()
	var lineNumber int
	for {
		lb, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if lineNumber == 1 {
			lb = append(lb, consts.NewLine)
			_, err := tempFileHandle.Write(lb)
			if err != nil {
				panic(fmt.Sprintf("write temp file exception:%v", err.Error()))
			}
		} else {
			lineNumber++
		}
	}
	cmd := exec.Command("go", "run", consts.TempFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to execute terminal command:%v", err.Error()))
	}
	output, _ := cmd.Output()
	io.Copy(os.Stdout, bytes.NewReader(output))
	_ = cmd.Run()
}
