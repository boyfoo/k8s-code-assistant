## 获取 client-go config

### pod 内获取默认 config

```go
import "k8s.io/client-go/rest"

config := rest.InClusterConfig()
```