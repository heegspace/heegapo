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

func (this *aYaml) Get(args ...string) Reader {
	if 0 == len(args) || nil == this.conf {
		return Reader{}
	}

	var value Reader
	temp := this.conf
	for k, arg := range args {
		if _, ok := temp[arg]; !ok {
			return Reader{}
		}

		value.Conf = temp[arg]
		if k < (len(args) - 1) {
			if _, ok := temp[arg].(map[interface{}]interface{}); !ok {
				return Reader{}
			}

			temp = temp[arg].(map[interface{}]interface{})
		}
	}

	return value
}
