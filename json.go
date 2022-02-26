package heegapo

import (
	"encoding/json"
)

type aJson struct {
	conf map[string]interface{}
}

func newJson(conf interface{}) *aJson {
	obj := &aJson{
		conf: make(map[string]interface{}),
	}

	if _, ok := conf.(string); ok {
		json.Unmarshal([]byte(conf.(string)), &obj.conf)
	}

	return obj
}

func (this *aJson) Get(args ...string) aReader {
	if 0 == len(args) || nil == this.conf {
		return aReader{}
	}

	var value aReader
	temp := this.conf
	for k, arg := range args {
		if _, ok := temp[arg]; !ok {
			return aReader{}
		}

		value.Conf = temp[arg]
		if k < (len(args) - 1) {
			if _, ok := temp[arg].(map[string]interface{}); !ok {
				return aReader{}
			}

			temp = temp[arg].(map[string]interface{})
		}
	}

	return value
}
