package conf

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Port             int32
	Supers           []string //默认的super_admin用户列表，uid。
	AllowedOrigins   []string `binding:"required"`
	SessionSecretKey string   `binding:"required"` //加密session cookie的，随便填什么，如需要之前登录的用户立刻失效可以修改这里。
	DataDir          string   `binding:"dir"`      //存放用户数据的根目录(非cache类)，需要考虑此目录的备份工作
	DBType           string   `binding:"oneof=sqlite3 mysql"`
	DBArgs           string   `binding:"required"`
	DBLogDetail      bool
}

func GetProgramPath() string {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		panic(err)
	}
	// fmt.Println("exReal: ", exReal)
	return filepath.Dir(exReal)
}
func loadConfig() Config {
	viper.SetConfigName(".conf")
	viper.SetConfigType("yaml")
	// 在 Home 及当前目录下面查找名为 ".crud-api" 的配置文件
	home, err := os.UserHomeDir()
	if err != nil {
		viper.AddConfigPath(home)
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath(GetProgramPath())
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	// fmt.Println(viper.Get("AllowedOrigins"))
	return config
}

var conf = loadConfig()

func New() *Config {
	return &conf
}
