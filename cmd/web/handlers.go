package main

import "github.com/abassGarane/muscles/domain"

type handler struct {
	service domain.Service
}

func NewHandler(s domain.Service) *handler {
	return &handler{s}
}
