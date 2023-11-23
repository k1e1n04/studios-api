package repository_example

import modelExample "github.com/togisuma/standard-echo-serverless/example/domain/model.example"

// ExampleRepository は ExampleRepository のインターフェース
type ExampleRepository interface {
	// Get は 挨拶を取得
	Get() (modelExample.ExampleEntity, error)
}
