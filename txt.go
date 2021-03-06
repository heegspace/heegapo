package heegapo

type aTxt struct {
	conf interface{}
}

func newTxt(conf interface{}) *aTxt {
	obj := &aTxt{
		conf: conf,
	}

	return obj
}

func (this *aTxt) Get(args ...string) aReader {
	if nil == this.conf {
		return aReader{}
	}

	var value aReader
	value.conf = this.conf
	return value
}
