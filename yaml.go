package heegapo

import (
	"gopkg.in/yaml.v2"
)

type aYaml struct {
	conf map[interface{}]interface{}
}

func newYaml(conf interface{}) *aYaml {
	obj := &aYaml{
		conf: make(map[interface{}]interface{}),
	}

	if _, ok := conf.(string); ok {
		yaml.Unmarshal([]byte(conf.(string)), &obj.conf)
	}

	return obj
}

func (this *aYaml) Get(args ...string) aReader {
	if 0 == len(args) || nil == this.conf {
		return aReader{}
	}

	var value aReader
	temp := this.conf
	for k, arg := range args {
		if _, ok := temp[arg]; !ok {
			return aReader{}
		}

		value.conf = temp[arg]
		if k < (len(args) - 1) {
			if _, ok := temp[arg].(map[interface{}]interface{}); !ok {
				return aReader{}
			}

			temp = temp[arg].(map[interface{}]interface{})
		}
	}

	return value
}
