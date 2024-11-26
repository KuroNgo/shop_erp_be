package handler

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoleData struct {
	Role   string   `json:"role" example:"admin"`
	API    []string `json:"api" example:"['GET', 'POST', 'PUT']"`
	Method []string `json:"method" example:"['GET', 'DELETE']"`
}

type UserRole struct {
	UserID []primitive.ObjectID `json:"user_id"`
	Role   string               `json:"role"`
}

type APIData struct {
	API    string   `json:"api"`
	Role   []string `json:"role"`
	Method []string `json:"method"`
}

type Role struct {
	Role string `json:"role"`
}

type APIRole struct {
	API  string   `json:"api"`
	Role []string `json:"role"`
}

type RoleAPI struct {
	Role string   `json:"role"`
	API  []string `json:"api"`
}
