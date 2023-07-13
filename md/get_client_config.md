## 获取 client-go config

### pod内获取默认config

```go
import "k8s.io/client-go/rest"

config := rest.InClusterConfig()
```