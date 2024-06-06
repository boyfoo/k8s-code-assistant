### 获取对象的 OwnerReference

```golang
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

ownerReference := metav1.GetControllerOf(obj)
```

### 设置一个对象的 OwnerReference

```golang
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

obj.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
  // 此处的 Service 类似只是示例，根据实际情况修改
  // SchemeGroupVersion 要使用 kind 类型对应的包， 如 Service 是 codev1 包的，否则自动删除效果会失效 
  *metav1.NewControllerRef(service, corev1.SchemeGroupVersion.WithKind("Service")),
}
```
