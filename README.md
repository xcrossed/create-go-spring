# create-go-spring

语法示例:

```
create-go-spring -m 'github.com/go-spring/create-go-spring' [ web http-rpc redis gorm ... ]
```

-m 指定项目的 module 名称，存在 -m 标记时表示创建初始化项目，没有 -m 标记时表示给现有项目添加新的 go-spring 模块。

最后面是 go-spring 的模块名，用户选择了什么模块，初始项目中就有那个模块的简单示例；如果什么参数都不加就显示 help 信息，列出所有支持的模块。