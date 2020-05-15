package entity

import (
	"github.com/kardianos/service"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"xiaoyi.xutil.spr/util"
)

const (
	initFile = "init.yaml"
)

type Program struct {

	//服务名
	Name string `yaml:name`
	//服务器显示
	DisplayName string `yaml:displayName`
	//服务描述
	Description string `yaml:description`
	//路径
	Path string `yaml:path`
	//程序
	Program string `yaml:program`
	//程序参数
	Args string `yaml:args`
}

func (p *Program) Run() {
	path := filepath.Join(p.Path, p.Program)
	// 此处编写具体的服务代码
	util.Start(p.Name, p.Path, path, p.Args)
}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	util.Logs("开始服务")
	go p.Run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	util.Logs("停止服务")
	util.Stop(p.Name, p.Path)
	return nil
}

//初始化服务
func InitService(p *Program) *service.Config {
	var serviceConfig = &service.Config{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
	}
	return serviceConfig
}

//读取配置文件
func InitProgram() *Program {
	i := new(Program)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, "config", initFile))
	util.CheckErr(err)
	err = yaml.Unmarshal(yamlFile, i)
	util.CheckErr(err)
	return i
}
