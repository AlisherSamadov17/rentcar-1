package service

import (
	"rent-car/pkg/logger"
	"rent-car/storage"
)



type IServiceManager interface {
	Car() carService
    Customer() customerService
	Order() orderService
}

type Service struct {
	carService carService
	customerService customerService
	orderService orderService
	logger logger.ILogger
}

func New(storage storage.IStorage,log logger.ILogger) Service  {
	services := Service{}
	services.carService = NewCarService(storage,log)
	services.customerService = NewCustomerService(storage,log)
	services.orderService = NewOrderService(storage,log)
	services.logger=log

	return services
}

func (s Service) Car() carService {
	return s.carService
}

func (s Service) Customer() customerService {
	return s.customerService
}

func (s Service) Order() orderService {
	return s.orderService
}