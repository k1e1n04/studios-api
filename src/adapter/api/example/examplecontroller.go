package example

import (
	"github.com/labstack/echo/v4"
	usecase_example "github.com/togisuma/studios-api/example/usecase.example"
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
