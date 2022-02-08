package grpc

import "errors"

var (
	errInputData     = errors.New("invalid input data")
	errAlreadyExists = errors.New("user already exists")
	errInvalidEmail  = errors.New("invalid email")
	errNotExists     = errors.New("user doesn't exist")
)

func (s *UserManagementServer) getMessageError(err error) string {
	msg := s.messages.Errors.Default
	switch err {
	case errInputData:
		msg = s.messages.Errors.InvalidInputData
	case errAlreadyExists:
		msg = s.messages.Errors.AlreadyExists
	case errInvalidEmail:
		msg = s.messages.Errors.InvalidEmail
	case errNotExists:
		msg = s.messages.Errors.NotExists
	}

	return msg
}
