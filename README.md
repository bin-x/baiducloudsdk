# baiducloudsdk
a golang sdk for baidu cloud, only image search now.

### 使用
```
go get github.com/bin-x/baiducloudsdk
```

##### 使用示例
```
package main

import (
	"fmt"
	baidu "github.com/bin-x/baiducloudsdk"
)

func main() {
	client := baidu.NewHttpClient("YOU API KEY", "YOU API SECRET")
	response,_ := client.SimilarAddUrl(
		"https://dss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=3892521478,1695688217&fm=26&gp=0.jpg",
		"{xx}",
		map[string]string{"tags": "11,12"},
	)
	fmt.Print(string(response))
}
```
