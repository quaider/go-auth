package rdb

type RedisOption struct {
	Addrs        []string `json:"addrs"            mapstructure:"addrs"`
	MasterName   string   `json:"masterName"       mapstructure:"master-name"`
	Password     string   `json:"password"         mapstructure:"password"`
	Database     int      `json:"database"         mapstructure:"database"`
	PoolSize     int      `json:"poolSize"         mapstructure:"pool-size"`      // 10 * runtime.GOMAXPROCS
	MaxRetries   int      `json:"maxRetries"       mapstructure:"max-retries"`    // 命令最大重试次数， 默认为3
	MinIdleConns int      `json:"minIdleConns"     mapstructure:"min-idle-conns"` // 默认为0，不保持
	MaxIdleConns int      `json:"maxIdleConns"     mapstructure:"max-idle-conns"` // 默认为0，不限制
	ReadTimeout  int      `json:"readTimeout"      mapstructure:"read-timeout"`   // 读取数据超时时间 0 - 默认值，3秒
	WriteTimeout int      `json:"writeTimeout"     mapstructure:"write-timeout"`  // 数据写入超时时间 0 - 默认值，3秒
	DialTimeout  int      `json:"dialTimeout"      mapstructure:"dial-timeout"`   // 建立新网络连接时的超时时间, 默认5秒
	PoolTimeout  int      `json:"poolTimeout"      mapstructure:"pool-timeout"`   // 默认是1秒+ReadTimeout, 代表如果连接池所有连接都在使用中，等待获取连接时间，超时将返回错误
}
