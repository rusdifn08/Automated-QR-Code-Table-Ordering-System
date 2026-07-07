package http

import (
	"strconv"
	"myapp/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpHandler struct {
	MenuUsecase  domain.MenuUsecase
	OrderUsecase domain.OrderUsecase
	TableUsecase domain.TableUsecase
}

func NewHttpHandler(app *fiber.App, mu domain.MenuUsecase, ou domain.OrderUsecase, tu domain.TableUsecase) {
	handler := &HttpHandler{
		MenuUsecase:  mu,
		OrderUsecase: ou,
		TableUsecase: tu,
	}

	api := app.Group("/api")

	api.Get("/menus", handler.GetMenus)
	api.Get("/tables/:number", handler.GetTable)
	
	api.Post("/orders", handler.CreateOrder)
	api.Get("/orders/:id", handler.GetOrder)
	api.Patch("/orders/:id/status", handler.UpdateOrderStatus)
}

func (h *HttpHandler) GetMenus(c *fiber.Ctx) error {
	menus, err := h.MenuUsecase.FetchAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(menus)
}

func (h *HttpHandler) GetTable(c *fiber.Ctx) error {
	numberStr := c.Params("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid table number"})
	}

	table, err := h.TableUsecase.GetTableByNumber(number)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Table not found"})
	}

	return c.JSON(table)
}

type createOrderRequest struct {
	TableID uuid.UUID          `json:"table_id"`
	Items   []domain.OrderItem `json:"items"`
}

func (h *HttpHandler) CreateOrder(c *fiber.Ctx) error {
	var req createOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	order, err := h.OrderUsecase.CreateOrder(req.TableID, req.Items)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(order)
}

func (h *HttpHandler) GetOrder(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	order, err := h.OrderUsecase.GetOrder(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}

	return c.JSON(order)
}

type updateStatusRequest struct {
	Status string `json:"status"`
}

func (h *HttpHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	var req updateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err = h.OrderUsecase.UpdateOrderStatus(id, req.Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Status updated successfully"})
}
