package heegapo

type aTxt struct {
	conf interface{}
}

func newTxt(conf interface{}) *aTxt {
	obj := &aTxt{}

	return obj
}

func (this *aTxt) Get(args ...string) aReader {
	if 0 == len(args) || nil == this.conf {
		return aReader{}
	}

	var value aReader
	value.conf = this.conf
	return value
}
