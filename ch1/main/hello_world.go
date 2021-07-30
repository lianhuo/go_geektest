package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// fmt.Println("Hello world Git")
	// adbconnect()
	// startOuO()
	startflipclock()
}

func startOuO(){
	RunCommand("adb","shell","am","start","com.elpatrixf.OuO/com.elpatrixf.OuO.Main")
}

func startflipclock(){
	RunCommand("adb","shell","am","start","one.alynx.flipclock/one.alynx.flipclock.MainActivity")
}

func adbconnect(){
	ip := "192.168.3.6:5555"
	RunCommand("adb","connect",ip)
	// RunCommand("adb","disconnect",ip)
	// RunCommand("scrcpy","-s",ip)
}

func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
    // 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
    // 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}