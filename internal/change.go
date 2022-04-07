package internal

import (
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
	filename := os.Args[1]
	inputFileHandle, err := os.OpenFile(filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Open file for execution exception:%v", err.Error()))
	}
	defer func() {
		_, err = inputFileHandle.WriteAt([]byte(consts.EnableScriptHead), 0)
		if err != nil {
			panic(fmt.Sprintf("Rewrite %v :%v", filename, err.Error()))
		}
		err = inputFileHandle.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close file %v :%v", filename, err.Error()))
		}
	}()
	_, err = inputFileHandle.WriteAt([]byte(consts.DisableScriptHead), 0)
	if err != nil {
		panic(fmt.Sprintf("Rewrite %v :%v", filename, err.Error()))
	}
	cmd := exec.Command("go", "run", filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to execute terminal command:%v", err.Error()))
	}
	output, _ := cmd.Output()
	io.Copy(os.Stdout, bytes.NewReader(output))
	_ = cmd.Run()
}
