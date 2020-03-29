package create

type Module struct {
	Name      string
	Desc      string
	ImportLib []string
	TplPath   string
	OutPath   string
}

var ModuleMap = make(map[string]*Module)

// InitAllSupportedModules init all support modules info
func InitAllSupportedModules() {
	web := &Module{
		Name:      "web",
		Desc:      "gospring web",
		ImportLib: make([]string, 0),
		TplPath:   "",
		OutPath:   "",
	}
	ModuleMap[web.Name] = web

	redis := &Module{
		Name:      "redis",
		Desc:      "gospring redis",
		ImportLib: make([]string, 0),
		TplPath:   "",
		OutPath:   "",
	}

	ModuleMap[redis.Name] = redis
	mysql := &Module{
		Name:      "mysql",
		Desc:      "gospring mysql",
		ImportLib: make([]string, 0),
		TplPath:   "",
		OutPath:   "",
	}
	ModuleMap[mysql.Name] = mysql
}
