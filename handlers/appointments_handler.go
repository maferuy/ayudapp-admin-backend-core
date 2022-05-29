package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maferuy/ayudapp-admin-backend-core/services"
	"github.com/maferuy/ayudapp-admin-backend-core/utils"
)

// @Summary	Obtiene todos las citas
// @ID 		get-appointments
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} services.GetAppointmentsResponse
// @Failure 400 {object} services.GetAppointmentsResponse
// @Router 	/admin/appointments [get]
func handleGetAppointments(service services.IAppointmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appointments, err := service.GetAppointments()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(appointments))
	}
}

// @Summary Crea una cita
// @ID 		create-appointment
// @Accept 	json
// @Produce json
// @Security ApiKeyAuth
// @Param   CreateAppointmentRequest body services.CreateAppointmentRequest true "Datos de la cita"
// @Success 200 {object} services.CreateAppointmentResponse
// @Failure 400 {object} services.CreateAppointmentResponse
// @Router 	/admin/appointments [post]
func handleCreateAppointment(service services.IAppointmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req services.CreateAppointmentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		appointmentID, err := service.CreateAppointment(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(appointmentID))
	}
}

// @Summary Obtiene una cita
// @ID 		get-appointment
// @Produce json
// @Security ApiKeyAuth
// @Param 	id path int true "ID de la cita"
// @Success 200 {object} services.GetAppointmentResponse
// @Failure 400 {object} services.GetAppointmentResponse
// @Router 	/admin/appointments/{id} [get]
func handleGetAppointment(service services.IAppointmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errors.New("el id es requerido")))
			return
		}

		appointment, err := service.GetAppointment(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(appointment))
	}
}

// @Summary Actualiza una cita
// @ID 		update-appointment
// @Accept 	json
// @Produce json
// @Security ApiKeyAuth
// @Param   id 					path int 						true "ID de la cita"
// @Param 	UpdateAppointmentRequest 	body services.UpdateAppointmentRequest true "Datos de la cita"
// @Success 200 {object} services.UpdateAppointmentResponse
// @Failure 400 {object} services.UpdateAppointmentResponse
// @Router 	/admin/appointments/{id} [put]
func handleUpdateAppointment(service services.IAppointmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req services.UpdateAppointmentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}

		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errors.New("el id es requerido")))
			return
		}

		appointment, err := service.UpdateAppointment(id, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(appointment))
	}
}

// @Summary Elimina una cita
// @ID 		delete-appointment
// @Produce json
// @Security ApiKeyAuth
// @Param 	id path int true "ID de la cita"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router 	/admin/appointments/{id} [delete]
func handleDeleteAppointment(service services.IAppointmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errors.New("el id es requerido")))
			return
		}

		err := service.DeleteAppointment(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(nil))
	}
}

/** Crea un nuevo grupo de endpoints
 *
 * @param group *gin.IRoutes "El grupo de endpoints padre"
 * @param service services.IAppointmentService "El servicio de citas"
 * @return *gin.IRoutes "El grupo de endpoints creado"
 */
func newAppointmentHandler(group gin.IRoutes, service services.IAppointmentService) *gin.IRoutes {
	group.GET("/", handleGetAppointments(service))
	group.POST("/", handleCreateAppointment(service))

	group.GET("/:id", handleGetAppointment(service))
	group.PUT("/:id", handleUpdateAppointment(service))
	group.DELETE("/:id", handleDeleteAppointment(service))

	return &group
}
