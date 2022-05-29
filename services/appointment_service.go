package services

import (
	"errors"
	"time"

	"github.com/maferuy/ayudapp-admin-backend-core/models"
	"github.com/maferuy/ayudapp-admin-backend-core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateAppointmentRequest struct {
	Date      time.Time     `json:"date"`
	Duration  time.Duration `json:"duration"`
	Address   string        `json:"address"`
	Status    string        `json:"status"`
	Helper    string        `json:"helper"`
	CreatedBy string        `json:"created_by"`
}

type UpdateAppointmentRequest struct {
	Date      time.Time     `json:"date"`
	Duration  time.Duration `json:"duration"`
	Address   string        `json:"address"`
	Status    string        `json:"status"`
	Helper    string        `json:"helper"`
	CreatedBy string        `json:"created_by"`
}

type GetAppointmentsResponse struct {
	Appointments []models.Appointment `json:"appointments"`
}

type CreateAppointmentResponse struct {
	AppointmentID string `json:"appointment_id"`
}

type GetAppointmentResponse struct {
	Appointment models.Appointment `json:"appointment"`
}

type UpdateAppointmentResponse struct {
	Appointment models.Appointment `json:"appointment"`
}

type IAppointmentService interface {
	GetAppointments() (response GetAppointmentsResponse, err error)
	CreateAppointment(req CreateAppointmentRequest) (response CreateAppointmentResponse, err error)

	GetAppointment(id string) (response GetAppointmentResponse, err error)
	UpdateAppointment(id string, req UpdateAppointmentRequest) (response UpdateAppointmentResponse, err error)
	DeleteAppointment(id string) (err error)
}

type AppointmentService struct {
	db *mongo.Database
}

/** Obtiene todos las citas
 *
 * @return GetAppointmentsResponse "Las citas"
 * @return err error "El error de la operación"
 */
func (service *AppointmentService) GetAppointments() (response GetAppointmentsResponse, err error) {
	var appointments []models.Appointment
	collection := service.db.Collection("appointments")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	if err = cursor.All(ctx, &appointments); err != nil {
		return
	}

	response.Appointments = appointments
	return
}

/** Crea una cita
 *
 * @param req CreateAppointmentRequest "Los valores de la cita a crear"
 * @return CreateAppointmentResponse "El id de la cita creado"
 * @return err error "El error de la operación"
 */
func (service *AppointmentService) CreateAppointment(req CreateAppointmentRequest) (response CreateAppointmentResponse, err error) {
	collection := service.db.Collection("appointments")

	createdBy, err := primitive.ObjectIDFromHex(req.CreatedBy)
	if err != nil {
		err = errors.New("el ayudante es inválido")
		return
	}

	helper, err := primitive.ObjectIDFromHex(req.Helper)
	if err != nil {
		err = errors.New("el ayudante es inválido")
		return
	}

	appointment := models.Appointment{
		Date:      req.Date,
		Duration:  req.Duration,
		Address:   req.Address,
		Status:    req.Status,
		CreatedBy: createdBy,
		Helper:    helper,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := collection.InsertOne(ctx, appointment)
	if err != nil {
		return
	}

	response.AppointmentID = result.InsertedID.(primitive.ObjectID).Hex()
	return
}

/** Obtiene una cita
 *
 * @param id string "El ID de la cita"
 * @return GetAppointmentResponse "La cita"
 * @return err error "El error de la operación"
 */
func (service *AppointmentService) GetAppointment(appointmentId string) (response GetAppointmentResponse, err error) {
	collection := service.db.Collection("appointments")
	var appointment models.Appointment

	id, err := primitive.ObjectIDFromHex(appointmentId)
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return
	}
	if count == 0 {
		err = errors.New("no se encontró e la cita")
		return
	}

	result := collection.FindOne(ctx, filter)
	if err = result.Err(); err != nil {
		return
	}

	err = result.Decode(&appointment)
	if err != nil {
		return
	}

	response.Appointment = appointment
	return

}

/** Actualiza una cita
 *
 * @param req UpdateAppointmentRequest "Los valores de la cita a actualizar"
 * @param id string "El id de la cita"
 * @return UpdateAppointmentResponse "Los datos de la cita actualizado"
 * @return err error "El error de la operación"
 */
func (service *AppointmentService) UpdateAppointment(appointmentId string, req UpdateAppointmentRequest) (response UpdateAppointmentResponse, err error) {
	collection := service.db.Collection("appointments")
	var appointment models.Appointment

	id, err := primitive.ObjectIDFromHex(appointmentId)
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return
	}
	if count == 0 {
		err = errors.New("no se encontró e la cita")
		return
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return
	}

	if err = cursor.Decode(&appointment); err != nil {
		return
	}

	appointment.UpdatedAt = time.Now()

	updatedValues, err := utils.ToDoc(appointment)
	if err != nil {
		return
	}

	update := bson.D{{"$set", updatedValues}}
	if _, err = collection.UpdateOne(ctx, filter, update); err != nil {
		return
	}

	response.Appointment = appointment
	return
}

/** Elimina una cita
 *
 * @param id string "El id de la cita"
 * @return err error "El error de la operación"
 */
func (service *AppointmentService) DeleteAppointment(appointmentId string) (err error) {
	collection := service.db.Collection("appointments")

	id, err := primitive.ObjectIDFromHex(appointmentId)
	if err != nil {
		return
	}

	filter := bson.M{"id": id}
	result := collection.FindOneAndDelete(ctx, filter)
	if err = result.Err(); err != nil {
		return
	}

	return
}

func NewAppointmentService(db *mongo.Database) IAppointmentService {
	return &AppointmentService{db: db}
}
