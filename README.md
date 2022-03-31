# logger

```go
func initLogger() error {
    zapLogger := logger.NewZapLogger(
    		fmt.Sprintf("%s/%s/alarm.log", rootDir, conf.GetString(logKeyPrefix+"dir")),
    		conf.GetString(logKeyPrefix+"level"),  // 日志级别
    		conf.GetInt(logKeyPrefix+"maxSize"),   // 单个文件的最大量
    		conf.GetInt(logKeyPrefix+"maxBackup"), // 最多的备份数量
    		conf.GetInt(logKeyPrefix+"maxAge"),    // 每天切割1次
    	)
    return zapLogger.InitLog()
}
```
