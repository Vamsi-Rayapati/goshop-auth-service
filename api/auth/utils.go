package auth

import "github.com/FusionAuth/go-client/pkg/fusionauth"

func ParseFusionAuthError(errors *fusionauth.Errors) *fusionauth.Error {
	if errors != nil && errors.FieldErrors != nil {
		for _, v := range errors.FieldErrors {
			return &v[0]
		}
	}

	return nil
}
