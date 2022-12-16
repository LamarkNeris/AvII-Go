package handler

import (
	"errors"
	"net/http"

	"AV II GO/internal/domain"
	"AV II GO/internal/dentist"
	"AV II GO/http"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

func DentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

func (h *dentistHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.ReadById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, dentist)
	}
}

func (h *dentistHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		err := ctx.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}

		createdDentist, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdDentist)
	}
}

func (h *dentistHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		var dentist domain.Dentist
		err = ctx.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		createdDentist, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdDentist)
	}
}

func (h *dentistHandler) Delete() gin.HandlerFunc {
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
