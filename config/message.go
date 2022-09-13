package config

import "errors"

const (
	MsgCityAdded   = "City has been added"
	MsgCityDeleted = "City has been successfully deleted"
	MsgCityUpdated = "City has been updated"
	MsgFavAdded    = "City has been added to favourites"
	MsgFavRemoved  = "City has been removed from favourites"
)

var (
	ErrMissingBasicAuth            = errors.New("Authorization must be required in header")
	ErrWrongPayload                = errors.New("Wrong payload, please try again")
	ErrRecordNotFound              = errors.New("Record not found")
	ErrParameterMissing            = errors.New("Parameter missing")
	ErrTokenMissing                = errors.New("Error token missing")
	ErrInvalidHashKey              = errors.New("Invalid hash key")
	ErrInvalidHttpMethod           = errors.New("Invalid http method")
	ErrHttpCallBadRequest          = errors.New("Bad request")
	ErrHttpCallUnauthorized        = errors.New("Unauthorized")
	ErrHttpCallNotFound            = errors.New("Call not found")
	ErrHttpCallInternalServerError = errors.New("Internal server error")
	ErrWentWrong                   = errors.New("Something went wrong")
)
