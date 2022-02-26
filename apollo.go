package heegapo

import (
	"fmt"
	"strings"
	"sync"

	"github.com/shima-park/agollo"
)

type Apolloer interface {
	Get(args ...string) aReader
}

type ConfigType int

const (
	DEFAULT_Type = ConfigType(0x000)
	XML_Type     = ConfigType(0x001)
	JSON_Type    = ConfigType(0x002)
	YML_Type     = ConfigType(0x003)
	YAML_Type    = ConfigType(0x004)
	TXT_Type     = ConfigType(0x005)
)

var DefaultApollo Apollo

type Apollo struct {
	apollo agollo.Agollo
	conf   map[interface{}]interface{}
	rwlock sync.RWMutex

	opts Options
}

func NewApollo() *Apollo {
	obj := &Apollo{
		conf:   make(map[interface{}]interface{}),
		rwlock: sync.RWMutex{},
	}

	return obj
}

// 配置刷新
//
// @param	name	namespace名字
// @param 	conf	配置数据
//
func (this *Apollo) refresh(name, conf interface{}) (err error) {
	if nil == conf {
		return
	}

	this.rwlock.Lock()
	defer this.rwlock.Unlock()

	this.conf[name] = conf
	return
}

// 根据空间名获取配置的类型信息
//
// @param	name
//
func (this *Apollo) configType(name string) ConfigType {
	if strings.HasSuffix(name, ".xml") {
		return XML_Type
	}
	if strings.HasSuffix(name, ".json") {
		return JSON_Type
	}
	if strings.HasSuffix(name, ".yml") {
		return YML_Type
	}
	if strings.HasSuffix(name, ".yaml") {
		return YAML_Type
	}
	if strings.HasSuffix(name, ".txt") {
		return TXT_Type
	}

	return DEFAULT_Type
}

// 获取配置信息
//
// @param	name 	空间名
// @param 	args 	配置key
//
func (this *Apollo) Config(name string, args ...string) aReader {
	if 0 == len(args) {
		return aReader{}
	}

	this.rwlock.RLock()
	defer this.rwlock.RUnlock()
	ctype := this.configType(name)
	if ctype != DEFAULT_Type {
		if _, ok := this.conf[name]; !ok {
			return aReader{}
		}
	}

	switch ctype {
	case XML_Type:
		xml := newXml(this.conf[name])

		return xml.Get(args...)
	case JSON_Type:
		json := newJson(this.conf[name])

		return json.Get(args...)
	case YML_Type:
		yaml := newYaml(this.conf[name])

		return yaml.Get(args...)
	case YAML_Type:
		yaml := newYaml(this.conf[name])

		return yaml.Get(args...)
	case TXT_Type:
		txt := newTxt(this.conf[name])

		return txt.Get(args...)
	}

	def := newDefault(this.apollo)
	return def.Get(args...)
}

// 监听apollo配置是否发生变更
//
func (this *Apollo) watchApo() {
	errorCh := this.apollo.Start()
	watchCh := this.apollo.Watch()
	for {
		select {
		case err := <-errorCh:
			fmt.Println("watch err: ", err)

			break
		case resp := <-watchCh:
			this.refresh(resp.Namespace, resp.NewValue["content"])

			reload := false
			var oldi interface{}
			var newi interface{}
			oldi = resp.OldValue["content"]
			newi = resp.NewValue["content"]
			oldValue := oldi.(string)
			newValue := newi.(string)

			if oldValue != newValue {
				err := this.refresh(resp.Namespace, resp.NewValue["content"])
				if nil != err {
					continue
				}

				reload = true
			}

			if reload && nil != this.opts.reloadCall {
				this.opts.reloadCall()
			}
			break
		}
	}
}

// 初始化apollo配置信息
//
// @param 	opts
//
func (this *Apollo) Init(opts ...Option) {
	for _, o := range opts {
		o(&this.opts)
	}

	a, err := agollo.New(this.opts.url, this.opts.appid, agollo.AutoFetchOnCacheMiss())
	if err != nil {
		panic(err)
	}
	this.apollo = a

	// 开始监听
	for _, arg := range this.opts.nameSpaces {
		if _, ok := this.conf[arg]; !ok {
			this.refresh(arg, a.GetNameSpace(arg)["content"])
		}
	}

	go this.watchApo()
	return
}
