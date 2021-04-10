### log best practices 


zap 有两种格式

- zap.Logger 使用结构体的格式来记录（性能比较高）
- zap.SugaredLogger 使用传统的 printf 的格式来记录（性能相较上一种低）

zap 打印一条日志的大致流程

- 通过 zapcore.Entry 结构体生成一条信息
- 通过 zapcore.Core 的 check 函数检查日志级别，产生一个zapcore.CheckEntry 对象
- 根据日志级别判断是否需要 panic 和 fatal 并退出应用
- 通过zapcore.Option 的设置，zap.AddCaller()，zap.AddStacktrace(1) 或其他自定义的 fields
- 通过 Encoder 为 zapcore.CheckEntry() 提供编码器，通过 EncodeEntry 方法吧 zapcore.Entry 和 zapcore.Fields 转成最后输出格式，json 和 console 转成 buffer.Buffer, memory encode 转成 map
 
zap 自带的日志记录器

- zap.NewProduction()
- zap.NewDevelopment()
- zap.NewExample()

zap 自定义日志记录器

- 通过 build 创建日志记录器
- 通过 zap.New 创建日志记录器

