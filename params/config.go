package params

type Zap struct {
	Level         string   `toml:"level"`
	JsonFormat    bool     `toml:"jsonFormat"`
	LogDirectory  string   `toml:"logDirectory"`
	LogFilePrefix string   `toml:"logFilePrefix"`
	LinkName      string   `toml:"linkName"`
	ShowCaller    bool     `toml:"showCaller"`
	EncodeLevel   string   `toml:"encodeLevel"`
	LogInConsole  bool     `toml:"logInConsole"`
	MaxAge        int      `toml:"maxAge"`
	RotationTime  int      `toml:"rotationTime"`
	GinSkipPath   []string `toml:"ginSkipPath"`
}

type Mysql struct {
	Path         string `toml:"path"`
	Config       string `toml:"config"`
	DbName       string `toml:"dbName"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	MaxIdleConns int    `toml:"maxIdleConns"`
	MaxOpenConns int    `toml:"maxOpenConns"`
	LogMode      string `toml:"logMode"`
}

type Redis struct {
	DB       int    `toml:"db"`
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
}

type JWT struct {
	SigningKey  string `toml:"signingKey"`  // jwt签名
	ExpiresTime int64  `toml:"expiresTime"` // 过期时间
}

type Eureka struct {
	Username              string   `toml:"username"`
	Password              string   `toml:"password"`
	Urls                  []string `toml:"urls"`
	AppName               string   `toml:"appName"`
	RegisterWithEureka    bool     `toml:"registerWithEureka"`
	InstanceID            string   `toml:"instanceID"`
	HostName              string   `toml:"hostName"`
	RenewalIntervalInSecs int      `toml:"renewalIntervalInSecs"`
	DurationInSecs        int      `toml:"durationInSecs"`
}
