package utils

const (
	MSG_FAILED_JWKS          = "failed to get JWKS: %v"
	MSG_TOKEN_ERROR          = "token parse error: %v"
	MSG_INVALID_TOKEN        = "invalid token"
	MSG_INVALID_TOKEN_CLAIMS = "invalid token claims"
	MSG_INVALID_ISSUER       = "invalid issuer"
	MSG_INVALID_AUDIENCE     = "invalid audience"
	MSG_HEADER_MISSING       = "authorization header missing"
	MSG_JWT_ERROR_VALIDATE   = "JWT validation error: %v"
	MSG_STATUS_TOKEN         = "Token is valid? : %t"

	MSG_FAILED_LOAD_AWS          = "failed to load AWS config: %v"
	MSG_MARSHAL_PAYLOAD          = "failed to marshal payload: %v"
	MSG_FAILED_UNMARSHAL_PAYLOAD = "failed unmarshal lambda response: %v"
	MSG_FAILED_INVOKE_LAMBDA     = "failed to invoke Lambda: %v"
)
