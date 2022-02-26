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
	value, err := strconv.ParseInt(obj.conf.(string), 10, 64)
	if nil != err {
		return int(def)
	}

	return int(value)
}

func (obj aReader) Int64(def int64) int64 {
	value, err := strconv.ParseInt(obj.conf.(string), 10, 64)
	if nil != err {
		return int64(def)
	}

	return value
}
