package responses

import "net/http"

func BadRequestResponse(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": message,
	}
	return result
}

func SuccessResponseNonData(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": message,
	}
	return result
}

func SuccessResponseData(message string, data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": message,
		"Data":    data,
	}
	return result
}
