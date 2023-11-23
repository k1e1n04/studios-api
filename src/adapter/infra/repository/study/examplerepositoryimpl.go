package study

import (
	modelExample "github.com/k1e1n04/studios-api/example/domain/model.example"
	repositoryExample "github.com/k1e1n04/studios-api/example/domain/repository.example"
)

// ExampleRepositoryImpl は Example リポジトリの実体
type ExampleRepositoryImpl struct {
}

// NewExampleRepository は ExampleRepositoryImpl を生成
func NewExampleRepository() repositoryExample.ExampleRepository {
	return &ExampleRepositoryImpl{}
}

// Get は Exampleエンティティを取得
func (er *ExampleRepositoryImpl) Get() (modelExample.ExampleEntity, error) {
	return modelExample.ExampleEntity{
		Message: "Hello",
	}, nil
}
