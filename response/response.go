package response

import (
	"net/http"
)

// function response false param
func FalseParamResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "False Param",
	}
	return result
}

// function response bad request
func BadRequestResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Bad Request",
	}
	return result
}

// function response access forbidden
func AccessForbiddenResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Access Forbidden",
	}
	return result
}

// function response success dengan paramater
func SuccessResponseData(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    data,
	}
	return result
}

// function response success tanpa parameter
func SuccessResponseNonData() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	}
	return result
}

// function response login failure
func LoginFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Email or Password Invalid",
	}
	return result
}

// function response login success
func LoginSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data":    data,
	}
	return result
}

func RentingSuccessResponse(data, payment interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":           http.StatusOK,
		"message":        "Your Resevation Success",
		"reservation_id": data,
		"payment":        payment,
	}
	return result
}

func CheckFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Item not available",
	}
	return result
}

func CheckSuccessResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Item available",
	}
	return result
}

func PasswordCannotLess5() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "password cannot less than 5 character",
	}
	return result
}

func NameCannotEmpty() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "name cannot be empty",
	}
	return result
}

func EmailCannotEmpty() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "email cannot be empty",
	}
	return result
}

func EmailPasswordCannotEmpty() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "email or password cannot be empty",
	}
	return result
}

func PasswordCannotEmpty() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "password cannot be empty",
	}
	return result
}

func PhoneNumberCannotEmpty() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "phone number cannot be empty",
	}
	return result
}

func IsExist() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Email or Phone Number is Exist",
	}
	return result
}

func FormatEmailInvalid() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Email must contain email format",
	}
	return result
}

func ItemsNotFoundResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Items not found",
	}
	return result
}

func DateInvalidResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Renting Date Invalid",
	}
	return result
}

func UploadErrorResponse(err error) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": err.Error(),
		"error":   true,
	}
	return result
}

func ProductsBadGatewayResponse(message string) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
	}
	return result
}

func BookingNotFoundResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Booking not found",
	}
	return result
}

func BookingOwnProductsFailed() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Cannot booking own products",
	}
	return result
}

func CheckoutSuccessResponse(data, payment interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":        http.StatusOK,
		"message":     "Your CheckOut Success",
		"checkout_id": data,
		"payment":     payment,
	}
	return result
}

func CheckoutSuccessResponseOVO(data, payment interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":        http.StatusOK,
		"message":     "Your CheckOut Success",
		"checkout_id": data,
		"payment":     payment,
	}
	return result
}

func CheckoutCODSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":        http.StatusOK,
		"message":     "Your CheckOut Success",
		"checkout_id": data,
	}
	return result
}

func CheckOutMissingResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "You Must Choose Payment Method",
	}
	return result
}
