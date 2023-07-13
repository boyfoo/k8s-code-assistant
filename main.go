package main

import "fmt"

type Student struct {
	Name string
}

func sa(arr []int) {
	sum := 10
	ret := make([][2]int, 0)
	hash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if v, ok := hash[sum-arr[i]]; ok {
			ret = append(ret, [2]int{arr[v], arr[i]})
		} else {
			hash[arr[i]] = i
		}
	}
	fmt.Println(ret)
}

func main() {
	ss := []int{2, 4, 1, 6, 8, 4}
	sa(ss)
}

//func mainss() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//	s := os.Getenv("OPENAI_KEY")
//
//	client := openai.NewClient(s)
//	resp, err := client.CreateChatCompletion(
//		context.Background(),
//		openai.ChatCompletionRequest{
//			Model: openai.GPT3Dot5Turbo0613,
//			Messages: []openai.ChatCompletionMessage{
//				{
//					Role:    openai.ChatMessageRoleUser,
//					Content: "你好!",
//				},
//			},
//		},
//	)
//
//	if err != nil {
//		fmt.Printf("ChatCompletion error: %v\n", err)
//		return
//	}
//
//	fmt.Println(resp.Choices[0].Message.Content)
//}
