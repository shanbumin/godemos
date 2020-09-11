package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)



// 修改进程的执行用户
func withUserAttr(cmd *exec.Cmd, name string) error {
	// 检测用户是否存在
	user, err := user.Lookup(name)
	if err != nil {
		return fmt.Errorf( "invalid user %s", name)
	}

	// set process attr
	// 获取用户 id
	uid, err := strconv.ParseUint(user.Uid, 10, 32)
	if err != nil {
		return err
	}
	// 获取用户组 id
	gid, err := strconv.ParseUint(user.Gid, 10, 32)
	if err != nil {
		return err
	}

	//创建一个系统属性实例(注意)
	sysProcAttr := &syscall.SysProcAttr{
		Setpgid: true, //是否设定组id给新进程
		Credential:&syscall.Credential{
			Uid:         uint32(uid),
			Gid:         uint32(gid),
			Groups:      nil,
			NoSetGroups: true,
		},
	}

	//设置该命令行的执行用户
	cmd.SysProcAttr = sysProcAttr

	return nil
}


func main() {

	cmd:=exec.Command("/bin/bash", "-c", "sleep 10;")
	err:=withUserAttr(cmd,"sam") //我们设置该命令行必须由sam用户来执行，否则报不允许
	if err !=nil{
		fmt.Println(err)
		return
	}
	err= cmd.Run()
	if err !=nil{
		fmt.Println(err)
	}


	
}
