package main

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
	"xiaoyi.xutil.spr/entity"
	"xiaoyi.xutil.spr/util"
)

func main() {
	program := entity.InitProgram()
	config := entity.InitService(program)
	s, err := service.New(program, config)
	util.CheckErr(err)
	if len(os.Args) < 2 {
		err = s.Run()
		util.CheckErr(err)
		return
	}
	cmd := os.Args[1]
	if cmd == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
	}

	if cmd == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	}
	if cmd == "start" {
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
	if cmd == "stop" {
		err = program.Stop(s)
		if err != nil {
			log.Fatal(err)
		}

	}
}
