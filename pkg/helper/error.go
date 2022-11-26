package helper

import "github.com/gofiber/fiber/v2"

// fail, goFiber PanicHandler function not available
func PanicIfError(err *error) {
	panic("error")
}

// fail, should directly return error in the main function handler
func BadRequestIfError(err *error) *fiber.Error {
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	return nil
}
