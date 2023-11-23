package repository_example

import modelExample "github.com/k1e1n04/studios-api/example/domain/model.example"

// ExampleRepository は ExampleRepository のインターフェース
type ExampleRepository interface {
	// Get は 挨拶を取得
	Get() (modelExample.ExampleEntity, error)
}
