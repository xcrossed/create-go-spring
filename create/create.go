package create

import (
	"fmt"
	"io/ioutil"
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

type TemplateOption struct {
	IsSupportWeb   bool
	IsSupportMySql bool
	IsSupportRedis bool
	ProjectName    string
	OutputBaseDir  string
}

func (c *create) Run() error {
	// 检查目录是否存在
	fileInfo, err := os.Stat(c.OutputDir)
	if os.IsExist(err) {
		panic(fmt.Sprintf("dir:%v is exist.", c.OutputDir))
	}
	if !fileInfo.IsDir() {
		panic(fmt.Sprintf("dir:%v is a file.", c.OutputDir))
	}
	//创建目录
	err = os.MkdirAll(c.OutputDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("create %v is err %v.", c.OutputDir, err))
	}

	//初始化go.mod

	//

	for _, v := range c.SupportModule {
		fmt.Printf("v:%v\n", v)
	}

	return nil
}

func WriteFile(dstPath string, data []byte) error {
	dir, _ := filepath.Split(dstPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Printf("mkdir %v err:%v", dstPath, err)
		return err
	}
	if err := ioutil.WriteFile(dstPath, data, os.ModePerm); err != nil {
		fmt.Printf("write file err:%v", err)
		return err
	}
	return nil
}

func NewCreate(opts *Options) Create {
	if len(opts.SupportModule) == 0 {
		opts.SupportModule = "all"
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
