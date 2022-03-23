package model

import "github.com/pkg/errors"

var (
	ErrAuthFiled                         = errors.New("auth uid failed")
	ErrFindUid                           = errors.New("Failed to find the uid")
	ErrWrongDataFormat                   = errors.New("The data format is wrong")
	ErrFieldIsNonexistentInCompareValues = errors.New("CompareValues: field does not exist")
	//ErrFieldIsNonexistentInGenerateBodyOfRequest = errors.New("GenerateBodyOfRequest: field does not exist")
	ErrParseConfig                = errors.New("Failed to parse the configuration")
	ErrUnmarshalBodyOfResponse    = errors.New("Failed to unmarshal the body of the response")
	ErrInvalidRequest             = errors.New("The request is invalid")
	ErrCheckDigest                = errors.New("Failed to check the digest")
	ErrQueryKey                   = errors.New("Failed to query the key")
	ErrKeyIsNonexistent           = errors.New("The key does not exist")
	ErrLoadConfig                 = errors.New("Failed to query the configuration")
	ErrConfigurationIsNonexistent = errors.New("The configuration does not exist")
	ErrHTTPCommunication          = errors.New("A HTTP communication error has occurred")
	ErrInvalidToken               = errors.New("The token is Invalid")
)
