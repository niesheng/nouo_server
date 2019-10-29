package main

type config struct {
	Postgres postgresConfig `json:"postgres"` //postgresql's config
	Server   serverConfig   `json:"server"`   //server's config
}

type serverConfig struct {
	Tls    bool         `json:"tls"`    //tls
	Cert   string       `json:"cert"`   //ssl certificate file
	Key    string       `json:"key"`    //ssl certificate's key
	Port   string       `json:"port"`   //listen port of the Nouo Web Server
	Upload uploadConfig `json:"upload"` //upload folder of the Nouo Web Server
}

type uploadConfig struct {
	Path  string   `json:"path"`
	Allow []string `json:"allow"`
}

type postgresConfig struct {
	Server   string `json:"server"`   //host's ip of postgresql's server
	Port     string `json:"port"`     //listsen's port of the postgresql's server
	Username string `json:"username"` //username of the postgresql's server
	Password string `json:"password"` //password
	Database string `json:"database"` //database
	Admin    string `json:"admin"`    //数据库运行管理用户
	Router   string `json:"router"`   //plv8 main of postgresql's server
}

type webRequest struct {
	Path   string                 `json:"path"`   //访问路径
	Host   string                 `json:"host"`   //访问路径
	Tls    bool                   `json:"tls"`    //访问路径
	Ip     string                 `json:"ip"`     //客户端IP
	Method string                 `json:"method"` //提交模式
	Get    map[string]interface{} `json:"get"`    //用户传输的参数
	Post   map[string]interface{} `json:"post"`   //用户传输的参数
	Cookie map[string]interface{} `json:"cookie"` //cookie信息
	Files  map[string]interface{} `json:"files"`  //传递文件相关信息
	Header map[string]interface{} `json:"header"` //所有客户端信息、包含用户信息
}

type webFile struct {
	Name string `json:"name"` //文件名
	Size int64  `json:"size"` //文件大小
	Path string `json:"path"` //文件路径
	Type string `json:"Type"` //文件类型
}

type webResponse struct {
	Body   string                 `json:"body"`   //返回页面内容
	Header map[string]interface{} `json:"header"` //返回页面头
	Cookie map[string]interface{} `json:"cookie"` //返回cookie
	File   webFile                `json:"file"`   //返回文件
}
