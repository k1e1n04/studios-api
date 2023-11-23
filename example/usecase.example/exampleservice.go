package usecase_example

import (
	repositoryExample "github.com/togisuma/standard-echo-serverless/example/domain/repository.example"
)

// ExampleService は Exampleサービス
type ExampleService struct {
	exampleRepository repositoryExample.ExampleRepository
}

// NewExampleService は ExampleService を生成
func NewExampleService(repositoryExample repositoryExample.ExampleRepository) ExampleService {
	return ExampleService{
		exampleRepository: repositoryExample,
	}
}

// Hello は Hello を実行
func (es *ExampleService) Hello() (ExampleDTO, error) {
	entity, err := es.exampleRepository.Get()
	if err != nil {
		return ExampleDTO{}, err
	}
	return ExampleDTO{
		Message: entity.Message,
	}, nil
}
