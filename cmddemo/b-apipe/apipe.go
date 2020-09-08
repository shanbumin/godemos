package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	//①展示如何读取一个命令行的输出内容
	runCmd()
	fmt.Println()
	//②管道可以把一个命令的输出作为另一个命令的输入
	runCmdWithPipe()
}

//这个基于cmd1和cmd2的示例模拟出了操作系统命令     ps aux|grep     apipe
func runCmdWithPipe() {

	//cmd1  cmd2
	fmt.Println("Run command `ps aux | grep apipe`: ")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")

	//设置cmd1的Stdout字段，然后启动cmd1，并等待它运行完毕
	//因为*bytes.Buffer类型实现了io.Writer接口，所以我才能把&outputBuf1赋给cmd1.Stdout。这样命令cmd1启动后的所有输出内容就都会被写入到outputBuf1
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1 //将cmd1的标准输出与&outputBuf1关联起来
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The first command can not be startup %s\n", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the first command: %s\n", err)
		return
	}

	//接下来，再设置cmd2的Stdin和Stdout字段，启动cmd2，并等待它运行完毕:
	cmd2.Stdin = &outputBuf1 //将cmd2的标准输入与&outputBuf1关联起来
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2 //将cmd2的标准输出与&outputBuf2关联起来
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The second command can not be startup: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the second command: %s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}

func runCmd() {
	//cmd0
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")  //echo  -n   My first command comes from golang.

	//本质这里得到的管道是通过os.Pipe函数生成的。只不过，该方法内部又对生成的管道做了少许的附加处理
	stdout0, err := cmd0.StdoutPipe() //创建一个可以获取cmd0输出的管道
	if err != nil {
		fmt.Printf("Error: Couldn't obtain the stdout pipe for command No.0: %s\n", err)
		return
	}
	defer stdout0.Close()
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
		return
	}


	//如何读取cmd0输出内容（一）
	//这种就是比较传统的读取方式了,本质就是io.Reader中的Read方法，只是每次读取之后放到了我们引入的缓冲bytes.Buffer中罢了
	var outputBuf0 bytes.Buffer
	for {
		tempOutput := make([]byte, 5)
		n, err := stdout0.Read(tempOutput)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
				return
			}
		}
		if n > 0 {
			outputBuf0.Write(tempOutput[:n])
		}
	}
	fmt.Printf("%s\n", outputBuf0.String())
	//如何读取cmd0输出内容（二）
	//这种与我们sdk中读取的方式是一样的，是bufio.Reader中的Read方法，里面自封装了缓冲机制
	/*
	outputBuf0 := bufio.NewReader(stdout0)
	//命令的输出本就是一行内容，所以不需要循环判断读取，直接读取一行就over了，所以不需要判断还有没有剩余行额，具体的细节判断封装在ReadLine()内部
	output0, _, err := outputBuf0.ReadLine()
	if err != nil {
		fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
		return
	}
	fmt.Printf("%s\n", string(output0))
	*/

}
