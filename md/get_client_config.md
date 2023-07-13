## 获取 client-go config

### pod 内的默认配置

```
import "k8s.io/client-go/rest"

config := rest.InClusterConfig()
```

### 指定路径文件获取

```
import "k8s.io/client-go/tools/clientcmd"

// 默认使用常量所表示的 ~/.kube/config 路径 如果都为空则使用rest.InClusterConfig()
config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
```

### 根据运行环境自动加载最适合的配置

```
import "k8s.io/cli-runtime/pkg/genericclioptions"

flags := genericclioptions.NewConfigFlags(true)
config := flags.ToRawKubeConfigLoader()
```

### 从 bytes 中加载

```
import (
  "fmt"
  "io/ioutil"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
)

func main() {
  file, _ := ioutil.ReadFile("./config")
  clientConfig, _ := clientcmd.NewClientConfigFromBytes(file)
  config, _ := clientConfig.ClientConfig()
  clientSet, _ := kubernetes.NewForConfig(config)
  serverVersion, _ := clientSet.ServerVersion()
  fmt.Println(serverVersion)
}
```