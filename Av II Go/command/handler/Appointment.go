package handler

import (
	"errors"
	"net/http"

	"strconv"

	"AV II GO/internal/domain"
	"AV II GO/internal/appointment"
	"AV II GO/http/"
	

	"github.com/gin-gonic/gin"
)

type appointment struct {
	s appointment.Service
}

func Appointment(s appointment.Service) *appointment {
	return &appointment{
		s: s,
	}
}


func (h *appointment) GetByRg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rg := ctx.Param("rg")
		appointments, err := h.s.ReadByRg(rg)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, appointments)
	}
}

func (h *appointment) CreateById() gin.HandlerFunc {
	type Request struct {
		Date        string `json:"date" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var request Request
		idPatient, err := strconv.Atoi(ctx.Param("id-patient"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid patient id"))
			return
		}
		idDentist, err := strconv.Atoi(ctx.Param("id-dentist"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid dentist id"))
			return
		}

		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		appointment := domain.Appointment{
			Date:        request.Date,
			Description: request.Description,
		}

		createdAppointment, err := h.s.CreateById(appointment, idPatient, idDentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdAppointment)
	}
}

func (h *appointment) Update() gin.HandlerFunc {
	type Request struct {
		PatientId   int `json:"patient_id"`
		DentistId   int `json:"dentist_id"`
		Date        string `json:"date"`
		Description string `json:"description"`
	}
	return func(ctx *gin.Context) {
		var request Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		if request.PatientId == 0 || request.DentistId == 0 || request.Date == "" || request.Description == "" {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		updateRequestAppointment := domain.Appointment{
			Patient: domain.Patient{
				Id: request.PatientId,
			},
			Dentist: domain.Dentist{
				Id: request.DentistId,
			},
			Date:        request.Date,
			Description: request.Description,
		}
		updatedAppointment, err := h.s.Update(id, updateRequestAppointment)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, updatedAppointment)
	}
}


func (h *appointment) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}