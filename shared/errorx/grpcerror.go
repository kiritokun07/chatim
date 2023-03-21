package errorx

import (
	"chatim/shared/errorx/grpcerrordetails"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsNotFound CodeError code is 404
// status.Status code is codes.NotFound
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	codeErr, ok := FromError(err)
	if ok {
		return codeErr.Code() == 404
	}
	grpcErr, ok := status.FromError(err)
	if ok {
		return codes.NotFound == grpcErr.Code()
	}
	return false
}

// NewGrpcErrorFromCodeError convert CodeError to grpc status error.
func NewGrpcErrorFromCodeError(err error) error {
	if err == nil {
		return nil
	}
	codeErr, ok := FromError(err)
	if ok {
		c := codes.FailedPrecondition
		if codeErr.Code() == 404 {
			c = codes.NotFound
		}
		st := status.New(c, codeErr.Desc())
		ds, err := st.WithDetails(
			&grpcerrordetails.GrpcErrorDetails{
				Code: int64(codeErr.Code()),
				Desc: codeErr.Desc(),
			},
		)
		if err != nil {
			return err
		}
		return ds.Err()
	}
	return err
}
