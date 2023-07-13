## 获取 client-go config

### pod 内的默认配置

```
import "k8s.io/client-go/rest"

config := rest.InClusterConfig()
```

### 指定路径文件获取

```
import "k8s.io/client-go/tools/clientcmd"

config, err := clientcmd.BuildConfigFromFlags("", ".kube/config")
```