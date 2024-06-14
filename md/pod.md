## POD 相关

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

### 从 pod 内拷贝文件至本地

```golang
import (
	"archive/tar"
	"fmt"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
)

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("", "./mytest/config")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	req := clientset.CoreV1().RESTClient().Get().
		Resource("pods").
		Namespace("default").
		Name("testpod").
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Command: []string{"tar", "cf", "-", "test.txt"},
			Stdout:  true,
			Stdin:   true,
			Stderr:  true,
			TTY:     false,
		}, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		log.Fatalln(err)
	}
	pipReader, pipWriter := io.Pipe()

	go func() {
		defer pipWriter.Close()
		exec.Stream(remotecommand.StreamOptions{
			Stdout: pipWriter,
			Stderr: os.Stderr,
			Stdin:  os.Stdin,
			Tty:    false,
		})
	}()
	reader := tar.NewReader(pipReader)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		fmt.Println("读到文件名：" + header.FileInfo().Name())
		file, err := os.Create(header.FileInfo().Name())
		if err != nil {
			fmt.Println(err)
			break
		}
		io.Copy(file, reader)
	}
}
```
