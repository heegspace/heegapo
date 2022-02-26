package heegapo

import (
	"github.com/shima-park/agollo"
)

type aDefault struct {
	apoll agollo.Agollo
}

func newDefault(apoll agollo.Agollo) *aDefault {
	obj := &aDefault{
		apoll: apoll,
	}

	return obj
}

func (this *aDefault) Get(args ...string) Reader {
	value := this.apoll.Get(args[0])
	return Reader{
		Conf: value,
	}
}
