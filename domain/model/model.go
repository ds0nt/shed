package model

type Provider string

const (
	Gemini  Provider = "gemini"
	Mistral Provider = "mistral"
	Ollama  Provider = "ollama"
	OpenAI  Provider = "openai"
)

type Model struct {
	Provider Provider `json:"source"`
	Name     string   `json:"name"`
}
