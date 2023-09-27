package todos

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetTodos(c *fiber.Ctx) error {
	getTodo := GetTodoQuery()
	if len(getTodo) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Data not found",
		})
	}
	return c.Status(200).JSON(getTodo)
}

func AddTodo(c *fiber.Ctx) error {
	type RequestBody struct {
		Title     string `json:"title"`
		Completed string `json:"completed"`
	}

	requestBody := new(RequestBody)

	// Parse the JSON payload from the request body
	err := c.BodyParser(requestBody)
	if err != nil {
		return err
	}

	title := requestBody.Title
	completed := requestBody.Completed

	if title == "" || completed == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Data not complete",
		})
	}

	completedInt, _ := strconv.Atoi(completed)
	data := TodoModel{
		Title:     title,
		Completed: completedInt,
	}

	transactionInsertTodo := TransactionInsertTodoQuery(data)
	if transactionInsertTodo != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": transactionInsertTodo.Error(),
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  true,
			"message": "Success",
		})
	}
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	type RequestBody struct {
		Completed string `json:"completed"`
	}

	requestBody := new(RequestBody)

	// Parse the JSON payload from the request body
	err := c.BodyParser(requestBody)
	if err != nil {
		return err
	}

	completed := requestBody.Completed

	if completed == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Data not complete",
		})
	}

	checkExistTodo := CheckExistTodoQuery(id)
	if checkExistTodo == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "No todo list found",
		})
	}

	completedInt, _ := strconv.Atoi(completed)
	data := TodoModel{
		Completed: completedInt,
	}

	transactionUpdateTodo := TransactionUpdateTodoQuery(data, id)
	if transactionUpdateTodo != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": transactionUpdateTodo.Error(),
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  true,
			"message": "Success",
		})
	}
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	transactionDeleteTodo := TransactionDeleteTodoQuery(id)
	if transactionDeleteTodo != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": transactionDeleteTodo.Error(),
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  true,
			"message": "Success",
		})
	}
}
