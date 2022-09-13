package log

type LogParameter struct {
	logDirectory  string
	logFilePrefix string
	linkName      string
	level         string
	encodeLevel   string
	jsonFormat    bool
	showCaller    bool
	logInConsole  bool
	maxAge        int
	rotationTime  int
}

// 订单选项
type LogOption func(p *LogParameter)

func LogDirectoryOption(logDirectory string) LogOption {
	return func(p *LogParameter) {
		p.logDirectory = logDirectory
	}
}

func LogFilePrefixOption(logFilePrefix string) LogOption {
	return func(p *LogParameter) {
		p.logFilePrefix = logFilePrefix
	}
}

func LinkNameOption(linkName string) LogOption {
	return func(p *LogParameter) {
		p.linkName = linkName
	}
}

func LevelOption(level string) LogOption {
	return func(p *LogParameter) {
		p.level = level
	}
}

func JsonFormatOption(jsonFormat bool) LogOption {
	return func(p *LogParameter) {
		p.jsonFormat = jsonFormat
	}
}

func EncodeLevelOption(encodeLevel string) LogOption {
	return func(p *LogParameter) {
		p.encodeLevel = encodeLevel
	}
}

func ShowCallerOption(showCaller bool) LogOption {
	return func(p *LogParameter) {
		p.showCaller = showCaller
	}
}

func LogInConsoleOption(logInCosole bool) LogOption {
	return func(p *LogParameter) {
		p.logInConsole = logInCosole
	}
}

func MaxAgeOption(maxAge int) LogOption {
	return func(p *LogParameter) {
		p.maxAge = maxAge
	}
}

func RotationTimeOption(rotationTime int) LogOption {
	return func(p *LogParameter) {
		p.rotationTime = rotationTime
	}
}

func ParseLogParameter(opts ...LogOption) *LogParameter {
	p := &LogParameter{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
