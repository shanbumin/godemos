package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

//Golang运行cmd命令逐行实时输出执行过程，注意，不是一次输出所有的内容，而是按照执行的过程，逐行逐行的实时显示出来
//cmd:=exec.Command("/bin/bash", "-c", "sleep 10")
func main() {
	command := "bash"
	params := []string{"-c","sleep 1;echo sam001;sleep 2;echo sam002;sleep 3;echo sam003;"}
	//执行cmd命令: ls -l
	execCommand(command, params)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}