### 新疆UserAccount用户

生成证书：

```
$ openssl genrsa -out client.key 2048
# 用户名称zx
$ openssl req -new -key client.key -out client.csr -subj "/CN=zx"
# 使用k8s的ca签发 时间365天
$ openssl x509 -req -in client.csr -CA /etc/kubernetes/pki/ca.crt -CAkey=/etc/kubernetes/pki/ca.key -CAcreateserial -out client.crt -days 365
```

创建成功后，请求api测试是否成功`curl --cert ./client.crt --key ./client.key --cacert /etc/kubernetes/pki/ca.crt -s https://{k8s-endpoint}/api`

其中`k8s-endpoint`可以通过 `kubectl cluster-info`或`kubectl get endpoints`获取

通过证书查看用户`openssl x509 -noout -subject -in client.crt`

将用户设置到`config`中：

`kubectl config set-credentials zx --client-certificate=client.crt --client-key=client.key --embed-certs=true`

设置上下文：`kubectl config set-context zx-context --cluster=kubernetes --user=zx`

切换上下文：`kubectl config use-context zx-context`

