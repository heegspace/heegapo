package heegapo

import (
	"strconv"
)

type Reader struct {
	Conf interface{}
}

func (obj Reader) String(def string) string {
	if _, ok := obj.Conf.(string); ok {
		return obj.Conf.(string)
	}
	return def
}

func (obj Reader) Int(def int) int {
	value, err := strconv.ParseInt(obj.Conf.(string), 10, 64)
	if nil != err {
		return int(def)
	}

	return int(value)
}

func (obj Reader) Int64(def int64) int64 {
	value, err := strconv.ParseInt(obj.Conf.(string), 10, 64)
	if nil != err {
		return int64(def)
	}

	return value
}
