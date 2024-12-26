package controllers

import (
	"context"
	"net/http"
	"rest_api_nosql_db/configs"
	"rest_api_nosql_db/models"
	"rest_api_nosql_db/responses"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var supplierCollection *mongo.Collection = configs.GetCollection(configs.DB, "suppliers")
var validate = validator.New()

func CreateSupplier(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var supplier models.Supplier
	defer cancel()

	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SupplierResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&supplier); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SupplierResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newSupplier := models.Supplier{
		Id:      primitive.NewObjectID(),
		Name:    supplier.Name,
		Address: supplier.Address,
		Mobile:  supplier.Mobile,
	}

	result, err := supplierCollection.InsertOne(ctx, newSupplier)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SupplierResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.SupplierResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetSupplier(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	supplierId := c.Params("supplierId")
	var supplier models.Supplier
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(supplierId)

	err := supplierCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&supplier)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SupplierResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.SupplierResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": supplier}})
}

func EditSupplier(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	supplierId := c.Params("supplierId")
	var supplier models.Supplier
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(supplierId)

	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SupplierResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&supplier); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SupplierResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": supplier.Name, "address": supplier.Address, "mobile": supplier.Mobile}

	result, err := supplierCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SupplierResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	var updatedSupplier models.Supplier
	if result.MatchedCount == 1 {
		err := supplierCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedSupplier)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.SupplierResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.SupplierResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedSupplier}})
}

func DeleteSupplier(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	supplierId := c.Params("supplierId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(supplierId)

	result, err := supplierCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SupplierResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.SupplierResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "This SupplierID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.SupplierResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Supplier  deleted!"}},
	)
}

func GetAllSuppliers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var suppliers []models.Supplier
	defer cancel()

	results, err := supplierCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SupplierResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleSupplier models.Supplier
		if err = results.Decode(&singleSupplier); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.SupplierResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		suppliers = append(suppliers, singleSupplier)
	}

	return c.Status(http.StatusOK).JSON(
		responses.SupplierResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": suppliers}},
	)
}
