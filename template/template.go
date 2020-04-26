package template

import (
	"fmt"
	"os"
	"sync"
	"text/template"
)

type Template struct {
	TplKey     string
	DstPath    string
	TplContent string
}

type TemplateFactory interface {
	Add(Template) error
	Exec() error
}

type templateFactory struct {
	templateMap map[string]Template
	value       interface{}
	mutx        sync.RWMutex
}

func New(v interface{}) TemplateFactory {
	return &templateFactory{
		templateMap: make(map[string]Template),
		value:       v,
	}
}

func (tf *templateFactory) Add(template Template) error {
	tf.mutx.Lock()
	defer tf.mutx.Unlock()
	if _, ok := tf.templateMap[template.TplKey]; ok {
		return fmt.Errorf("repeat add template:%v", template.TplKey)
	}
	return nil
}

func (tf *templateFactory) Exec() error {
	tf.mutx.RLock()
	defer tf.mutx.RUnlock()
	for _, v := range tf.templateMap {
		t := template.New(v.TplKey)
		t = template.Must(t.Parse(v.TplContent))
		if err := t.Execute(os.Stdout, tf.value); err != nil {
			fmt.Printf("exec file %v err:%v", v.DstPath, err)
		}
	}

	return nil
}
