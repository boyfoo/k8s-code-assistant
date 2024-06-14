```golang
import (
  "errors"
  "k8s.io/client-go/util/retry"
)
func main() {
	n := 0
	retry.OnError(retry.DefaultRetry, func(err error) bool {
		return true
	}, func() error {
		fmt.Println("重试")
		if n < 100 {
			n++
			return errors.New("错误")
		}
		return nil
	})
}
```
