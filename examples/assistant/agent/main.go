package main

import (
	"context"
	"fmt"

	"github.com/henomis/lingoose/assistant"
	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/thread"

	pythontool "github.com/henomis/lingoose/tools/python"
	serpapitool "github.com/henomis/lingoose/tools/serpapi"
)

func main() {

	auto := "auto"
	a := assistant.New(
		openai.New().WithModel(openai.GPT4o).WithToolChoice(&auto).WithTools(
			pythontool.New(),
			serpapitool.New(),
		),
	).WithParameters(
		assistant.Parameters{
			AssistantName:      "AI Assistant",
			AssistantIdentity:  "an helpful assistant",
			AssistantScope:     "with their questions.",
			CompanyName:        "",
			CompanyDescription: "",
		},
	).WithThread(
		thread.New().AddMessages(
			thread.NewUserMessage().AddContent(
				thread.NewTextContent("calculate the average temperature in celsius degrees of New York, Rome, and Tokyo."),
			),
		),
	).WithMaxIterations(10)

	err := a.Run(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("----")
	fmt.Println(a.Thread())
	fmt.Println("----")
}
