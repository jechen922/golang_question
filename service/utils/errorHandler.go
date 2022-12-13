package utils

import (
	"demo/apierror"
	"errors"
	"log"

	"gorm.io/gorm"
)

type ErrorHandler struct{}

func (h ErrorHandler) NotFoundOrInternalError(errorCode int, err error) *apierror.Error {
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return apierror.New(errorCode)
	}
	log.Println("record unexpected error here")
	return apierror.New(apierror.InternalServiceError, err)
}

func (h ErrorHandler) ErrorOrInternalError(errorCode int, err error) *apierror.Error {
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return apierror.New(errorCode)
	}
	log.Println("record unexpected error here")
	return apierror.New(apierror.InternalServiceError, err)
}
