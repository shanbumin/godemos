package main

import (
	"log"
	"os/exec"
	"fmt"
)


//cmd:=exec.Command("/bin/bash", "-c", "echo fuck;echo shit;")
func main() {

    log.Println("start ...")
	cmd:=exec.Command("/bin/bash", "-c", "sleep 10")
	err:= cmd.Run()  //<==> cmd.Start()+cmd.Wait()
	if err !=nil{
		fmt.Println(err)
	}
	log.Println("end ...")

}
