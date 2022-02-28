package heegapo

import (
	"strconv"
)

type aReader struct {
	conf interface{}
}

func (obj aReader) String(def string) string {
	if _, ok := obj.conf.(string); ok {
		return obj.conf.(string)
	}
	return def
}

func (obj aReader) Int(def int) int {
	if _,ok := obj.conf.(int); ok {
		return int(obj.conf.(int))
	}

	if _,ok := obj.conf.(string); !Ok {
		return int(def)
	}

	value, err := strconv.ParseInt(obj.conf.(string), 10, 64)
	if nil != err {
		return int(def)
	}

	return int(value)
}

func (obj aReader) Int64(def int64) int64 {
	if _,ok := obj.conf.(int64); ok {
		return int(obj.conf.(int64))
	}

	if _,ok := obj.conf.(string); !Ok {
		return int64(def)
	}

	value, err := strconv.ParseInt(obj.conf.(string), 10, 64)
	if nil != err {
		return int64(def)
	}

	return value
}
