package rest

import (
	"assignment2/item"
	"assignment2/order"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Order struct {
	orderService order.Service
	itemService  item.Service
}

func NewOrder(orderService order.Service, itemService item.Service) *Order {
	return &Order{orderService, itemService}
}

func (o *Order) GetOrders(c *gin.Context) {
	orders, err := o.orderService.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}

	var orderResponse []order.OrderResponse

	for _, ordr := range orders {
		items, err := o.itemService.GetByFK(ordr.OrderID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": err,
			})
			return
		}
		var itemResponse []item.ItemResponse

		for _, i := range items {
			res := item.ItemResponse{
				ItemID:      i.ItemID,
				ItemCode:    i.ItemCode,
				Description: i.Description,
				Quantity:    i.Quantity,
			}
			itemResponse = append(itemResponse, res)
		}

		res := order.OrderResponse{
			OrderID:      ordr.OrderID,
			CustomerName: ordr.CustomerName,
			OrderedAt:    ordr.OrderedAt,
			Items:        itemResponse,
		}

		orderResponse = append(orderResponse, res)
	}

	if orderResponse == nil {
		c.JSON(
			http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "order kosong",
			})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  "success",
			"message": "order berhasil didapatkan",
			"data":    orderResponse,
		})
}

func (o *Order) CreateOrder(c *gin.Context) {
	var orderRequest order.OrderRequest

	err := c.ShouldBindJSON(&orderRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  errorMessages,
		})
		return
	}

	order, err := o.orderService.Create(orderRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err,
		})
		return
	}

	items := orderRequest.Items

	for _, i := range items {

		item := item.ItemRequest{
			ItemCode:    i.ItemCode,
			Description: i.Description,
			Quantity:    i.Quantity,
			OrderID:     order.OrderID,
		}
		_, err := o.itemService.Create(item)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  err,
			})
			return
		}
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  "success",
			"message": "order berhasil dimasukkan",
		})
}

func (o *Order) UpdateOrder(c *gin.Context) {
	var orderRequest order.OrderRequest

	err := c.ShouldBindJSON(&orderRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  errorMessages,
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("orderId"))

	order, err := o.orderService.Update(id, orderRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err,
		})
		return
	}

	for _, item := range orderRequest.Items {
		item, err := o.itemService.Update(item.ItemID, item)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  err,
			})
			return
		}
		_ = item
	}

	_ = order

	c.JSON(
		http.StatusOK, gin.H{
			"status":  "success",
			"message": "order berhasil diubah",
		})
}

func (o *Order) DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("orderId"))

	order, err := o.orderService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err,
		})
		return
	}

	_ = order

	c.JSON(
		http.StatusOK, gin.H{
			"status":  "success",
			"message": "order berhasil dihapus",
		})
}
