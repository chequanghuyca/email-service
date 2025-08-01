package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode 	int 	`json:"status_code"`
	RootErr 	error 	`json:"root_err"`
	Message 	string 	`json:"message"`
	Log 		string 	`json:"log"`
	Key 		string 	`json:"key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: 	http.StatusBadRequest,
		RootErr: 		root,
		Message: 		msg,
		Log: 			log,
		Key: 			key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: 	statusCode,
		RootErr: 		root,
		Message: 		msg,
		Log: 			log,
		Key: 			key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: 	http.StatusUnauthorized,
		RootErr: 		root,
		Message: 		msg,
		Key: 			key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(
		err, 
		"something went wrong with DB", 
		err.Error(), 
		"DB_ERROR",
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(
		err,
		"Invalid request",
		err.Error(),
		"ERR_INVALID_REQUEST",
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity), 
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity), 
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity), 
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity), 
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity), 
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s was deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity), 
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already existed", strings.ToLower(entity)),
		fmt.Sprintf("Err%sExisted", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrNoPermission( err error) *AppError { 
	return NewUnauthorized(
		err,
		"You don't have permission to do this action",
		"ErrNoPermission",
	)
}

var RecordNotFound = errors.New("record not found")