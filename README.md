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

### 从指定位置的代码文件生成swagger json 文件(支持的语言有php,js,python等)  

- 输出保存到到json文件

`scmt create -s /home/koolay/myApp -l php --name myswagger --version 1.0.0 -o myswagger.json`

> json文件位置可以是当前目录，或者绝对路径．如a.json则保存到当前scmt可执行文件当前目录

- 在终端输出json

`scmt create -s /home/koolay/myApp -l php --name myswagger -o stdout`

- 输出结果调用api (http PUT)

`scmt create -s /home/koolay/myApp/product.php -l php --name myswagger -o http://localhost:1337/api/open/swagger -H x-ticket="xxx"`

或者指定文件夹

`scmt create -s /home/koolay/myApp -l php --name myswagger -o http://localhost:1337/api/open/swagger -H x-ticket="xxx"`

- 输出到终端和保存到json

`scmt create -s /home/koolay/myApp -l php --name myswagger -o stdout -o a.json`


### 代码注释格式

- php语言请[参考](fixture/template.php)

- python语言请[参考](fixture/template.py)

  js与php类似


### 注释说明

- [@api](#api)
- [@apiGroup](#apiGroup)
- [@apiParam](#apiParam)
- [@apiResponse](#apiResponse)

**api属性 <a name="api">**

> @api {method} path [title]  

|  Name          |  Description  |
|:-----------|:------------|
|  method  |  Request method name: DELETE, GET, POST, PUT, ...  |
| path  | Request Path. |
|  title  | optional	A short title. (used for navigation and article header)  |

**@apiGroup <a name="apiGroup">**

api分组或所属模块

> @apiGroup product

| Name          | Description |
|:-----------|:------------|
| name | 模块名称　|


**@apiParam <a name="apiParam">**

api参数

> @apiParam [(group)] [{type}] [field=defaultValue] [description]

| Name         | Description |
|:-----------|:------------|
| {type} `optional` | 参数类型.　如: {Boolean}, {Number}, {String} |
| {string{..5}} |  不超过５个字符的字符串. |
| {string{2..5}} | 2到５个字符长度  |
| {number{100-999}} | 数字100到999.  |
| field	 | 字段名　|
| [field] | 可选填字段. |
| =defaultValue `optinal` | 参数默认值  |
| descriptionoptional     | 参数描述 |



**@apiResponse <a name="apiResponse">**

api响应输出

> @apiResponse statusCode  {   data
> }

| Name         | Description |
|:-----------|:------------|
| statusCode | http状态码　如: 200, 201 ... |
| {data} |  响应数据, 用相应的类型值代码类型. |

> **Note:** 最后的大括号需要单独一行

### 验证swagger json文件格式正确性(还没完成)


## TODO 

- 验证swagger json文件的正确性
- 生成 yaml file
- 自定义content-type



