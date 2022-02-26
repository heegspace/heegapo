package heegapo

import (
	"strings"

	"github.com/tinyhubs/tinydom"
)

type aXml struct {
	xmldoc tinydom.XMLDocument
}

func newXml(conf interface{}) *aXml {
	obj := &aXml{}

	if _, ok := conf.(string); ok {
		xmldoc, _ := tinydom.LoadDocument(strings.NewReader(conf.(string)))
		obj.xmldoc = xmldoc
	}

	return obj
}

func (this *aXml) Get(args ...string) Reader {
	if 0 == len(args) || nil == this.xmldoc {
		return Reader{}
	}

	var xml tinydom.XMLElement
	for _, arg := range args {
		if nil == xml {
			xml = this.xmldoc.FirstChildElement(arg)

			continue
		}

		if nil == xml {
			break
		}

		xml = xml.FirstChildElement(arg)
	}

	if nil == xml {
		return Reader{}
	}

	return Reader{
		Conf: xml.Text(),
	}
}
