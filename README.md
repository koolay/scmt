# swagger-cli
=============

[swagger](http://swagger.io/specification/#pathsObject) 命令行工具

## 功能

> 使用` scmt --help ` 命令查看指令说明．  


```
Usage:
  scmt [command]

Available Commands:
  create      Output json of swagger from special file or directinary
  validate    validate swagger content

```

```

Usage:
  scmt create [flags]

Flags:
      --config string         config file (default is $HOME/.scmt.yaml)
  -H, --headers stringArray   http headers.
	eg:
	-H Authorization="Bearer mytoken"
	
  -l, --language string       language, php,pytho,go etc.
      --name string           name of swagger project.
  -o, --output stringArray    Where to output, can be json/api/yml.
	eg:
	output to a json file: -o /home/koolay/swagger.json
	output to a yml file: -o /home/koolay/swagger.yml
	output to POST an api: -o http://myhost.com/swagger
	 (default [json])
  -s, --sources stringArray   full path of special directory or file
      --version string        version of swagger project.

```

```

Usage:
  scmt validate [flags]

Flags:
      --url string   swagger url

```

#### 1. 从指定位置的代码文件生成swagger json 文件(支持的语言有php,js,python等)  

- 输出保存到到json文件

`scmt create -s /home/koolay/.go/src/github.com/koolay/scmt -l php --name myswagger --version 1.0.0 -o myswagger.json

> json文件位置可以是当前目录，或者绝对路径．如a.json则保存到当前scmt可执行文件当前目录

- 在终端输出json

`scmt create -s /home/koolay/.go/src/github.com/koolay/scmt -l php --name myswagger -o stdout

- 输出结果调用api (http PUT)

`scmt create -s /home/koolay/.go/src/github.com/koolay/scmt -l php --name myswagger -o http://localhost:1337/api/open/swagger -H x-ticket="xxx"`

- 输出到终端和保存到json

`scmt create -s /home/koolay/.go/src/github.com/koolay/scmt -l php --name myswagger -o stdout -o a.json


#### 2. 验证swagger json文件格式正确性(还没完成)


## TODO

- 验证swagger json文件的正确性
- 生成 yaml file
- 自定义content-type



