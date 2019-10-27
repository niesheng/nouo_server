package main

type config struct {
	Postgres postgresConfig `json:"postgres"` //数据库配置
	Ssl      bool           `json:"ssl"`      //是否使用HTTPS
	Cert     string         `json:"cert"`     //SSL证书文件位置
	Key      string         `json:"key"`      //SSL证书密钥文件位置
	Port     string         `json:"port"`     //端口
	Upload   string         `json:"upload"`   //上传目录
	Schema   string         `json:"schema"`   //平台名

}

type postgresConfig struct {
	Host     string `json:"host"`     //地址
	Port     string `json:"port"`     //端口
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
	Database string `json:"database"` //数据库
	Admin    string `json:"admin"`    //运行管理用户
}

//向后端传送的结构体
type webRequest struct {
	Url    string                 `json:"url"`    //当前主机
	Ip     string                 `json:"ip"`     //客户端IP
	Method string                 `json:"method"` //提交模式
	Get    map[string]interface{} `json:"get"`    //用户传输的参数
	Post   map[string]interface{} `json:"post"`   //用户传输的参数
	Cookie map[string]interface{} `json:"cookie"` //cookie信息
	Files  map[string][]webFile   `json:"files"`  //传递文件相关信息
	Header map[string]interface{} `json:"header"` //所有客户端信息、包含用户信息
}

type webFile struct {
	Name string `json:"name"` //文件名
	Size int64  `json:"size"` //文件大小
	Path string `json:"path"` //文件路径
}
