package example

import usecaseExample "github.com/togisuma/studios-api/example/usecase.example"

// ExampleResponse は Exampleレスポンスの構造体
type ExampleResponse struct {
	// Message は メッセージ
	Message string `json:"message"`
}

// FromDTO は ExampleDto から ExampleResponse を生成
func FromDTO(dto usecaseExample.ExampleDTO) ExampleResponse {
	return ExampleResponse{
		Message: dto.Message,
	}

}
