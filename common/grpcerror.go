package common

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrorOK               = status.Error(codes.OK, "OK")
	ErrorUnknown          = status.Error(codes.Unknown, "Unknown")
	ErrorInvalidArgument  = status.Error(codes.InvalidArgument, "InvalidArgument")
	ErrorNotFound         = status.Error(codes.NotFound, "NotFound")
	ErrorAlreadyExists    = status.Error(codes.AlreadyExists, "AlreadyExists")
	ErrorPermissionDenied = status.Error(codes.PermissionDenied, "PermissionDenied")
	ErrorInternal         = status.Error(codes.Internal, "Internal")
)

// GrpcErrorOK OK
func GrpcErrorOK(err error) error {
	return status.Error(codes.OK, err.Error())
}

// GrpcErrorUnknown Unknown
func GrpcErrorUnknown(err error) error {
	return status.Error(codes.Unknown, err.Error())
}

// GrpcErrorInvalidArgument InvalidArgument
func GrpcErrorInvalidArgument(err error) error {
	return status.Error(codes.InvalidArgument, err.Error())
}

// GrpcErrorNotFound NotFound
func GrpcErrorNotFound(err error) error {
	return status.Error(codes.NotFound, err.Error())
}

// GrpcErrorAlreadyExists AlreadyExists
func GrpcErrorAlreadyExists(err error) error {
	return status.Error(codes.AlreadyExists, err.Error())
}

// GrpcErrorPermissionDenied PermissionDenied
func GrpcErrorPermissionDenied(err error) error {
	return status.Error(codes.PermissionDenied, err.Error())
}

// GrpcErrorInternal Internal
func GrpcErrorInternal(err error) error {
	return status.Error(codes.Internal, err.Error())
}
