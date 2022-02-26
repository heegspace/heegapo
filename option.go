package heegapo

type ReloadFunc func()
type Options struct {
	url        string
	appid      string
	nameSpaces []string

	reloadCall ReloadFunc
}

type Option func(*Options)

// apollo连接地址
// 用于获取和监听apollo配置的地址信息
//
// @param 	url	地址
//
func Url(url string) Option {
	return func(o *Options) {
		o.url = url
	}
}

// apollo appid,主要是要使用那个配置
//
// @param 	appid
//
func Appid(appid string) Option {
	return func(o *Options) {
		o.appid = appid
	}
}

// 需要直接加载哪些空间的配置
//
// @param	nameSpace 	空间列表
//
func Namespace(nameSpaces []string) Option {
	return func(o *Options) {
		o.nameSpaces = nameSpaces
	}
}

// 配置变更回调函数
//
// @param	call
//
func ReloadCall(call ReloadFunc) Option {
	return func(o *Options) {
		o.reloadCall = call
	}
}
