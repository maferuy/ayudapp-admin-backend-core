package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maferuy/ayudapp-admin-backend-core/services"
	"github.com/maferuy/ayudapp-admin-backend-core/utils"
)

// @Summary	Crea una categoría
// @ID 		create-category
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} services.CreateCategoryResponse
// @Failure 400 {object} services.CreateCategoryResponse
// @Router 	/admin/categories [post]
func handleCreateCategory(service services.ICategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req services.CreateCategoryRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		categoryId, err := service.CreateCategory(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(categoryId))
	}
}

// @Summary	Obtiene las categorías
// @ID 		get-categories
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} services.GetCategoriesResponse
// @Failure 400 {object} services.GetCategoriesResponse
// @Router 	/admin/categories [get]
func handleGetCategories(service services.ICategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories, err := service.GetCategories()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(categories))
	}
}

// @Summary	Obtiene una categoría por su ID
// @ID 		get-category
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} services.GetCategoryResponse
// @Failure 400 {object} services.GetCategoryResponse
// @Router 	/admin/categories/{id} [get]
func handleGetCategory(service services.ICategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			err := errors.New("el id es requerido")
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		category, err := service.GetCategory(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(category))
	}
}

// @Summary	Actualiza una categoría
// @ID 		update-category
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} services.UpdateCategoryResponse
// @Failure 400 {object} services.UpdateCategoryResponse
// @Router 	/admin/categories/{id} [put]
func handleUpdateCategory(service services.ICategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req services.UpdateCategoryRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		id := ctx.Param("id")
		if id == "" {
			err := errors.New("el id es requerido")
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		category, err := service.UpdateCategory(id, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(category))
	}
}

// @Summary	Elimina una categoría
// @ID 		delete-category
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router 	/admin/categories/{id} [delete]
func handleDeleteCategory(service services.ICategoryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			err := errors.New("el id es requerido")
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		err := service.DeleteCategory(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(nil))
	}
}

/**
 * @param group *gin.IRoutes "El grupo de endpoints padre"
 * @param service services.ICategoryService "El servicio de categorias"
 * @return *gin.Routes "El grupo de endpoints creado"
 */
func newCategoryHandler(group gin.IRoutes, service services.ICategoryService) *gin.IRoutes {
	group.POST("/", handleCreateCategory(service))
	group.GET("/", handleGetCategories(service))
	group.GET("/:id", handleGetCategory(service))
	group.PUT("/:id", handleUpdateCategory(service))
	group.DELETE("/:id", handleDeleteCategory(service))

	return &group
}
