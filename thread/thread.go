package thread

type Thread struct {
	Messages []*Message
}

type ContentType string

const (
	ContentTypeText         ContentType = "text"
	ContentTypeImage        ContentType = "image"
	ContentTypeToolCall     ContentType = "tool_call"
	ContentTypeToolResponse ContentType = "tool_response"
)

type Content struct {
	Type ContentType
	Data any
}

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

type Message struct {
	Role     Role
	Contents []*Content
}

type ToolResponseData struct {
	ID     string
	Name   string
	Result string
}

type ToolCallData struct {
	ID       string
	Function ToolCallFunction
}

type ToolCallFunction struct {
	Name      string
	Arguments string
}

type MediaData struct {
	Raw any
	URL *string
}

func NewTextContent(text string) *Content {
	return &Content{
		Type: ContentTypeText,
		Data: text,
	}
}

func NewImageContent(mediaData *MediaData) *Content {
	return &Content{
		Type: ContentTypeImage,
		Data: mediaData,
	}
}

func NewToolResponseContent(toolResponseData *ToolResponseData) *Content {
	return &Content{
		Type: ContentTypeToolResponse,
		Data: toolResponseData,
	}
}

func NewToolCallContent(data []*ToolCallData) *Content {
	return &Content{
		Type: ContentTypeToolCall,
		Data: data,
	}
}

func (m *Message) AddContent(content *Content) *Message {
	m.Contents = append(m.Contents, content)
	return m
}

func NewUserMessage() *Message {
	return &Message{
		Role: RoleUser,
	}
}

func NewAssistantMessage() *Message {
	return &Message{
		Role: RoleAssistant,
	}
}

func NewToolMessage() *Message {
	return &Message{
		Role: RoleTool,
	}
}

func (t *Thread) AddMessage(message *Message) *Thread {
	t.Messages = append(t.Messages, message)
	return t
}

func (t *Thread) AddMessages(messages []*Message) *Thread {
	t.Messages = append(t.Messages, messages...)
	return t
}

func (t *Thread) CountMessages() int {
	return len(t.Messages)
}

func NewThread() *Thread {
	return &Thread{}
}

func (t *Thread) String() string {
	str := "Thread:\n"
	for _, message := range t.Messages {
		str += string(message.Role) + ":\n"
		for _, content := range message.Contents {
			str += "\tType: " + string(content.Type) + "\n"
			switch content.Type {
			case ContentTypeText:
				str += "\tText: " + content.Data.(string) + "\n"
			case ContentTypeImage:
				str += "\tImage URL: " + *content.Data.(*MediaData).URL + "\n"
			case ContentTypeToolCall:
				for _, toolCallData := range content.Data.([]*ToolCallData) {
					str += "\tTool Call ID: " + toolCallData.ID + "\n"
					str += "\tTool Call Function Name: " + toolCallData.Function.Name + "\n"
					str += "\tTool Call Function Arguments: " + toolCallData.Function.Arguments + "\n"
				}
			case ContentTypeToolResponse:
				str += "\tTool ID: " + content.Data.(*ToolResponseData).ID + "\n"
				str += "\tTool Name: " + content.Data.(*ToolResponseData).Name + "\n"
				str += "\tTool Result: " + content.Data.(*ToolResponseData).Result + "\n"
			}
		}
	}
	return str
}
