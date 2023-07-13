## Deployment
### 获取deployment调度的最新pod
```
import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

config, _ := clientcmd.BuildConfigFromFlags("", "./config")
client, _ := kubernetes.NewForConfig(config)
deploy, _ := client.AppsV1().Deployments("default").Get(context.Background(), "nginx-deployment", metav1.GetOptions{})
reList, _ := client.AppsV1().ReplicaSets("default").List(context.Background(), metav1.ListOptions{
    LabelSelector: metav1.FormatLabelSelector(deploy.Spec.Selector),
})
var rs *v1.ReplicaSet
// 获取deployment调度的最新的replicaset
for _, rsItem := range reList.Items {
    if rsItem.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] == deploy.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] {
        for _, reference := range rsItem.OwnerReferences {
            if reference.Kind == "Deployment" && reference.UID == deploy.UID {
                rs = rsItem.DeepCopy()
                break
            }
        }
    }
}
if rs != nil {
    podList, _ := client.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{
        LabelSelector: metav1.FormatLabelSelector(rs.Spec.Selector),
    })
    for _, pod := range podList.Items {
        fmt.Println(pod.Name)
    }
}
```