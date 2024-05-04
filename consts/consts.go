package consts

// When error occurs, the error message will be stored in the gin context with this key. LogMiddleware will catch this error and log it
const GinErrorKey = "api_error"
const GinDetailErrorKey = "api_additional_error"

// When the response is ready, the response data will be stored in the gin context with this key. ResponseMiddleware will catch this data and send it to the client

const GinResponseKey = "api_response"

const GinAuthUid = "uid_auth"
