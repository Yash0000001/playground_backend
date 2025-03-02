package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yash0000001/playgroundbackend/runner"
)

type CodeRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func Run_handler(c *fiber.Ctx) error {
	var req CodeRequest
	if err := c.BodyParser(&req); err != nil {
		// fmt.Println("Error parsing request:", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": "Invalid Request", "details": err.Error()})
	}
	output, stderr, time, err := runner.CompileAndRun(req.Code, req.Language)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Execution failed",
			"stderr":  stderr,
			"details": err.Error()})
	}

	return c.JSON(fiber.Map{"output": output, "stderr": stderr, "time": time})
}
