package base

import "mandiri_payment_gateway_service/internal/response"

type ServiceInterface interface {
	Do() response.ServiceResponse
}
