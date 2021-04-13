Viper是一个方便Go语言应用程序处理配置信息的库。它可以处理多种格式的配置。它支持的特性：

- 设置默认值
- 从JSON、TOML、YAML、HCL和Java properties文件中读取配置数据
- 可以监视配置文件的变动、重新读取配置文件
- 从环境变量中读取配置数据
- 从远端配置系统中读取数据，并监视它们（比如etcd、Consul）
- 从命令参数中读物配置
- 从buffer中读取
- 调用函数设置配置信息


### 设置默认值

```
viper.SetDefault("time", "2019-7-14")
viper.SetDefault("notifyList", []string{"maple","ffm"})
```

### 监视配置文件, 重新读取配置数据

```
package main

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)
viper:=viper.New()
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  fmt.Println("Config file changed:", e.Name)
})
```
