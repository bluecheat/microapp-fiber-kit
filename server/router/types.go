package router

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-querystring/query"
)

// fiber에서 bind tag로 설정된 값
//const (
//	queryTag     = "query"
//	reqHeaderTag = "reqHeader"
//	bodyTag      = "form"
//	paramsTag    = "params"
//)

type serviceFunc[I interface{}, O interface{}] func(req *I) (*O, error)

// 생성한 구조체에 대해 Validate 함수 생성
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// requestValidator 함수는 req를 파싱하고 검증합니다.
// T 형식의 req를 인자로 받고 검증 에러가 없다면 nil을 반환합니다.
func bodyValidator[T any](c *fiber.Ctx, req *T) error {
	if err := c.BodyParser(req); err != nil {
		//return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate struct
	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}

func queryValidator[T any](c *fiber.Ctx, req *T) error {
	if err := c.QueryParser(req); err != nil {
		//return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate struct
	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}

func pathValidator[T any](c *fiber.Ctx, req *T) error {
	if err := c.ParamsParser(req); err != nil {
		//return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate struct
	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}

// 공통 응답
func response[O interface{}](c *fiber.Ctx, response O, err error) error {
	if err != nil {
		return err
	}
	return c.JSON(response)
}

func responseRedirect(c *fiber.Ctx, url string, response interface{}, err error) error {
	if err != nil {
		return err
	}
	v, _ := query.Values(response)
	fmt.Print(v.Encode())
	return c.Redirect(url+"?"+v.Encode(), 301)
}

func wrapHandler[I interface{}, O interface{}](srv serviceFunc[I, O]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req = new(I)

		if err := pathValidator(ctx, req); err != nil {
			return err
		}
		if err := queryValidator(ctx, req); err != nil {
			return err
		}
		if err := bodyValidator(ctx, req); err != nil {
			return err
		}

		result, err := srv(req)
		return response(ctx, result, err)
	}
}

func wrapQueryHandler[I interface{}, O interface{}](input I, srv serviceFunc[I, O]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req = &input
		if err := queryValidator(ctx, req); err != nil {
			return err
		}
		result, err := srv(req)
		return response(ctx, result, err)
	}
}
func wrapBodyHandler[I interface{}, O interface{}](input I, srv serviceFunc[I, O]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req = &input
		if err := bodyValidator(ctx, req); err != nil {
			return err
		}
		result, err := srv(req)
		return response(ctx, result, err)
	}
}

func wrapParamHandler[I interface{}, O interface{}](input I, srv serviceFunc[I, O]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req = &input
		if err := pathValidator(ctx, req); err != nil {
			return err
		}
		result, err := srv(req)
		return response(ctx, result, err)
	}
}
