# heegapo
apollo配置使用工具,支持yaml,yml,json,xml,txt等配置访问


# 使用
```
apollo := heegapo.NewApollo()
apollo.Init(heegapo.Url("localhost:8080"), heegapo.Appid("12345"),
    heegapo.Namespace([]string{"application", "test.yaml", "test.yml", "test.json", "test.xml"}),
    heegapo.ReloadCall(func() {
        fmt.Println("Apollo config refresh call")

        return
    }),
)

...
apollo.Config("application", "teset").String("default value")
```