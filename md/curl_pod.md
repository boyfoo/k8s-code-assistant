## 访问 pod 方式

> 以下都默认`pod`的端口是80，如果不是80，需要在访问时指定端口

### 容器外

1. IP访问

```shell
# 使用 kubectl get pod -o wide 查看pod的ip
$ curl {pod-ip}
```

2. ClusterIp访问，给`pod`添加`Service`资源

```shell
# 使用 kubectl get svc 查看的cluster-ip
$ curl {cluster-ip}
```

3. 使用externalIPs访问，给`Server`添加`externalIPs`字段

```yaml
apiVersion: v1
kind: Service
metadata:
  name: testsvc001-svc
spec:
  type: ClusterIP
  externalIPs:
    - 192.168.56.57 # 如宿主机的ip
  ports:
    - port: 80
      targetPort: 80
  selector:
    app: testpod001-label
```

请求访问： `curl {externalIPs}`

使用`iptables -t nat -nvL | grep  192.168.56.57`可以查看到添加了路由规则

4. IP拼接特定后缀访问

```shell
# 容器ip拼接特定后缀,使用 kubectl get pod -o wide 查看pod的ip
curl 10-244-1-19.default.pod.cluster.local
```

### 容器内

上面容器外的方式加上都可以在容器内访问

1. 使用完全限定名访问pod**只能自己pod访问自己**，需要设置`subdomain`值

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: testpod001
spec:
  nodeName: node001
  subdomain: testsub
```

```shell
# 完全限定名 
curl testpod001.testsub.default.svc.cluster.local
curl testpod002.testsub002.default.svc.cluster.local
```