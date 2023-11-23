package example

import (
	usecase_example "github.com/k1e1n04/studios-api/example/usecase.example"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ExampleController は Exampleコントローラ
type ExampleController struct {
	exampleService usecase_example.ExampleService
}

// NewExampleController は ExampleController を生成
func NewExampleController(exampleService usecase_example.ExampleService) ExampleController {
	return ExampleController{
		exampleService: exampleService,
	}
}

// Hello は Hello を実行
func (ec *ExampleController) Hello(c echo.Context) error {
	dto, err := ec.exampleService.Hello()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, FromDTO(dto))
}
