package contract

const AppKey = "hade:app"

type App interface {
	// AppID 表示当前这个app的唯一id，可以用于分布式锁等
	AppID() string
	// Version 定义当前版本
	Version() string
	// BaseFold 定义项目基础地址
	BaseFolder() string
	// ConfigFolder 配置文件地址
	ConfigFolder() string
	// LogFolder 日志文件地址
	LogFolder() string
	// ProviderFolder 业务自己提供的服务地址
	ProviderFolder() string
	// MiddlewareFolder 中间件地址
	MiddlewareFolder() string
	// CommandFolder 业务命令地址
	CommandFolder() string
	// RuntimeFolder 业务中间态文件位置
	RuntimeFolder() string
	// 存放测试所需要的信息
	TestFolder()  string
}
