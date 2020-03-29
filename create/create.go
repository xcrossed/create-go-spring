package create

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Create interface {
	Run() error
}

type create struct {
	SupportModule []string
	ProjectName   string
	OutputDir     string
}

func (c *create) Run() error {
	//创建目录

	//初始化go.mod

	return nil
}

func (g *create) Usage() string {
	return Usage()
}

func NewCreate(opts *Options) Create {
	if len(opts.SupportModule) == 0 {
		opts.SupportModule = "web"
	}
	supportModules := strings.Split(opts.SupportModule, ",")
	if len(opts.OutputDir) == 0 {
		opts.OutputDir = filepath.Join(GetCurrentDirectory(), opts.ProjectName)
	}

	return &create{
		SupportModule: supportModules,
		ProjectName:   opts.ProjectName,
		OutputDir:     opts.OutputDir,
	}
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
