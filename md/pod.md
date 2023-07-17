## POD相关

### 远程 exec pod

```go
import (
    "fmt"
    v1 "k8s.io/api/core/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/kubernetes/scheme"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/tools/remotecommand"
    "os"
)

func main() {
    config, _ := clientcmd.BuildConfigFromFlags("", "./config")
    client, _ := kubernetes.NewForConfig(config)
    req := client.CoreV1().RESTClient().Post().
        Resource("pods").Namespace("default").Name("nginx-deployment-579fdb4c89-wdp28").
        SubResource("exec").VersionedParams(
        &v1.PodExecOptions{
            Container: "nginx",
            Command:   []string{"sh", "-c", "ls"},
            Stdin:     true,
            Stdout:    true,
            Stderr:    true,
        },
        scheme.ParameterCodec,
    )
    fmt.Println(req.URL())
    exec, _ := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
    exec.Stream(remotecommand.StreamOptions{
        Stdin:  os.Stdin,
        Stdout: os.Stdout,
        Stderr: os.Stderr,
        Tty:    true,
    })
}
```