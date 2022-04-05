package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// NewSetting 返回一个保存了配置文件相关信息的viper对象
func NewSetting() (*Setting, error) {
	vp := viper.New()
	//设置配置文件的文件名
	vp.SetConfigName("config")
	//设置配置文件所在的路径
	vp.AddConfigPath("configs/")
	//设置配置文件的后缀名
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}
