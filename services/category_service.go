package services

import (
	"errors"

	"github.com/maferuy/ayudapp-admin-backend-core/models"
	"github.com/maferuy/ayudapp-admin-backend-core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryResponse struct {
	CategoryID string `json:"category_id"`
}

type GetCategoriesResponse struct {
	Categories []models.Category `json:"categories"`
}

type GetCategoryResponse struct {
	Category models.Category `json:"category"`
}

type UpdateCategoryResponse struct {
	Category models.Category `json:"category"`
}

type ICategoryService interface {
	CreateCategory(req CreateCategoryRequest) (response CreateCategoryResponse, err error)
	GetCategories() (response GetCategoriesResponse, err error)
	GetCategory(id string) (response GetCategoryResponse, err error)
	UpdateCategory(id string, req UpdateCategoryRequest) (response UpdateCategoryResponse, err error)
	DeleteCategory(id string) (err error)
}

type CategoryService struct {
	db *mongo.Database
}

/** Crea una categoría
 *
 * @param req CreateCategoryRequest "Los datos de la categoría"
 * @return response CreateCategoryResponse "El id de la categoría"
 * @return err error "El error de la operación"
 */
func (service CategoryService) CreateCategory(req CreateCategoryRequest) (response CreateCategoryResponse, err error) {
	collection := service.db.Collection("categories")

	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		return
	}

	response.CategoryID = result.InsertedID.(primitive.ObjectID).Hex()
	return
}

/** Obtiene todas las categorías
 *
 * @return response GetCategoriesResponse "Las categorías"
 * @return err error "El error de la operación"
 */
func (service CategoryService) GetCategories() (response GetCategoriesResponse, err error) {
	var categories []models.Category
	collection := service.db.Collection("categories")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &categories); err != nil {
		return
	}

	response.Categories = categories
	return
}

/** Obtiene una categoría
 *
 * @param categoryId string "El id de la categoría"
 * @return response GetCategoryResponse "La categoría"
 */
func (service CategoryService) GetCategory(categoryId string) (response GetCategoryResponse, err error) {
	var category models.Category
	collection := service.db.Collection("categories")

	id, err := primitive.ObjectIDFromHex(categoryId)
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return
	} else if count == 0 {
		err = errors.New("categoría no encontrada")
		return
	}

	result := collection.FindOne(ctx, filter)
	if err = result.Err(); err != nil {
		return
	}

	err = result.Decode(&category)
	if err != nil {
		return
	}

	response.Category = category
	return
}

/** Actualiza una categoryia
 *
 * @param categoryId string "El id de la categoría"
 * @param req UpdateCategoryRequest "Los datos de la categoría"
 * @return response UpdateCategoryResponse "La categoría actualizada"
 * @return err error "El error de la operación"
 */
func (service CategoryService) UpdateCategory(categoryId string, req UpdateCategoryRequest) (response UpdateCategoryResponse, err error) {
	var category models.Category
	collection := service.db.Collection("categories")

	id, err := primitive.ObjectIDFromHex(categoryId)
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return
	} else if count == 0 {
		err = errors.New("categoría no encontrada")
		return
	}

	result := collection.FindOne(ctx, filter)
	if err = result.Err(); err != nil {
		return
	}

	if err = result.Decode(&category); err != nil {
		return
	}

	updatedValues, err := utils.ToDoc(req)
	if err != nil {
		return
	}

	update := bson.D{{"$set", updatedValues}}
	if _, err = collection.UpdateOne(ctx, filter, update); err != nil {
		return
	}

	response.Category = category
	return
}

/** Elimina una categoría
 *
 * @param categoryId string "El id de la categoría"
 * @return err error "El error de la operación"
 */
func (service CategoryService) DeleteCategory(categoryId string) (err error) {
	collection := service.db.Collection("categories")

	id, err := primitive.ObjectIDFromHex(categoryId)
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return
	} else if count == 0 {
		err = errors.New("categoría no encontrada")
		return
	}

	result := collection.FindOneAndDelete(ctx, filter)
	if err = result.Err(); err != nil {
		return
	}

	return

}

func NewCategoryService(db *mongo.Database) ICategoryService {
	return &CategoryService{
		db: db,
	}
}
