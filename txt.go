package heegapo

type aTxt struct {
	conf interface{}
}

func newTxt(conf interface{}) *aTxt {
	obj := &aTxt{}

	return obj
}

func (this *aTxt) Get(args ...string) Reader {
	if 0 == len(args) || nil == this.conf {
		return Reader{}
	}

	var value Reader
	value.Conf = this.conf
	return value
}
