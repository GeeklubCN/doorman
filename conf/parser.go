package conf

import (
	"bytes"
	"io/ioutil"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Parser interface {
	Parse(configPath string) (Config, error)
}

type ParserFunc func(string) (Config, error)

func (f ParserFunc) Parse(configPath string) (Config, error) {
	return f(configPath)
}

func NewParser() Parser {
	return NewParserWithFileReader(ioutil.ReadFile)
}

func NewParserWithFileReader(f FileReaderFunc) Parser {
	return parser{f}
}

type FileReaderFunc func(string) ([]byte, error)

type parser struct {
	fileReader FileReaderFunc
}

func (p parser) Parse(configPath string) (Config, error) {
	var res Config
	data, err := p.fileReader(configPath)
	if err != nil {
		logrus.Errorf("read config file content fail! path:%s, err:%s", configPath, err)
		return res, err
	}

	viper.SetConfigType(path.Ext(configPath)[1:])
	if err = viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		logrus.Errorf("read config fail! path:%s, data:%s, err:%s", configPath, data, err)
		return res, err
	}
	if err = viper.Unmarshal(&res); err != nil {
		logrus.Errorf("unmarshal config fail! path:%s, data:%s, err:%s", configPath, data, err)
		return res, err
	}
	if err = res.Init(); err != nil {
		logrus.Errorf("config Init fail! path:%s, res:%s, err:%s", configPath, res, err)
		return res, err
	}
	return res, nil
}
