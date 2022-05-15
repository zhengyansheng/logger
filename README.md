# Logger

## Download

```
go get github.com/zhengyansheng/logger
```

## Use
```go
package main

import (
	"fmt"
	"log"
	
	"github.com/zhengyansheng/logger"
)

func main()  {
	zapLogger := logger.NewZapLogger(
		fmt.Sprintf("var/log/xxx.log"),
		"debug",  // 日志级别
		100,      // 单个文件的最大量
		3,        // 最多的备份数量
		1,        // 每天切割1次
	)
	err := zapLogger.InitLog()
	if err != nil {
		log.Fatal(err)
    }
}
```
