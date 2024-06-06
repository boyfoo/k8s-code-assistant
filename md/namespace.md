### 根据资源对象获取 namespace 唯一 key 值

```golang
import "k8s.io/client-go/tools/cache"

key, err := cache.MetaNamespaceKeyFunc(k8sObj)
# 例：key = default/my-service
```

### 根据 namespace 唯一 key 值获取对象名

```golang
import "k8s.io/client-go/tools/cache"

namespace, name, _ := cache.SplitMetaNamespaceKey(key)
# key =  default/my-service
# namespace = default; name = my-service
```