package main

import (
	"AV II GO/command/handler"
	"AV II GO/internal/appointment"
	"AV II GO/internal/dentist"
	"AV II GO/internal/patient"
	"AV II GO/db"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	sql := store.NewSQL()

	sqlPatient := store.SQLPatient(sql)
	patientRepository := patient.Repository(sqlPatient)
	patientService := patient.Service(patientRepository)
	patientHandler := handler.PatientHandler(patientService )

	patients := r.Group("/patients")
	{
		patients.GET("/id/:id", patientHandler.GetById())
		patients.POST("", patientHandler.Create())
		patients.PUT(":id", patientHandler.Update())
		patients.DELETE(":id", patientHandler.Delete())
	}

	sqlDentist := store.SQLDentist(sql)
	dentistRepository := dentist.Repository(sqlDentist)
	dentistService := dentist.Service(dentistRepository)
	dentistHandler := handler.DentistHandler(dentistService)

	dentists := r.Group("/dentists")
	{
		dentists.GET("/id/:id", dentistHandler.GetById())
		dentists.POST("", dentistHandler.Create())
		dentists.PUT(":id", dentistHandler.Update())
		dentists.DELETE(":id", dentistHandler.Delete())
	}

	sqlAppointment := store.SQLAppointment(sql)
	appointmentRepository := appointment.Repository(sqlAppointment)
	appointmentService := appointment.Service(appointmentRepository)
	appointmentHandler := handler.AppointmentHandler(appointmentService)

	appointments := r.Group("/appointments")
	{
		appointments.GET("/rg/:rg", appointmentHandler.GetByRg())
		appointments.POST("/id/:id-patient/:id-dentist", appointmentHandler.CreateById())
		appointments.PUT(":id", appointmentHandler.Update())
		appointments.PATCH(":id", appointmentHandler.Patch())
		appointments.DELETE(":id", appointmentHandler.Delete())
	}

	r.Run(":8080")
}
