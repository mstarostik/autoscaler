// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3outposts

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Access was denied for this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// There was a conflict with this action, and it could not be completed.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// There was an exception with the internal server.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeOutpostOfflineException for service response error code
	// "OutpostOfflineException".
	//
	// The service link connection to your Outposts home Region is down. Check your
	// connection and try again.
	ErrCodeOutpostOfflineException = "OutpostOfflineException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// There was an exception validating this data.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"ConflictException":         newErrorConflictException,
	"InternalServerException":   newErrorInternalServerException,
	"OutpostOfflineException":   newErrorOutpostOfflineException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ThrottlingException":       newErrorThrottlingException,
	"ValidationException":       newErrorValidationException,
}
