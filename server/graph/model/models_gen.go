// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Email struct {
	Email string `json:"email" bson:"email"`
}

type ID struct {
	ID string `json:"id" bson:"_id"`
}

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type LocationInput struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type LocationSearch struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Radius      int       `json:"radius" bson:"radius"`
}

type Money struct {
	Currency string  `json:"currency" bson:"currency"`
	Value    float64 `json:"value" bson:"value"`
}

type MoneyInput struct {
	Currency string  `json:"currency" bson:"currency"`
	Value    float64 `json:"value" bson:"value"`
}

type Product struct {
	ID          *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Price       *Money    `json:"price" bson:"price"`
	Description string    `json:"description" bson:"description"`
	Image       string    `json:"image" bson:"image"`
	Author      string    `json:"author" bson:"author"`
	StoreName   string    `json:"storeName" bson:"storeName"`
	Location    *Location `json:"location" bson:"location"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

type ProductInput struct {
	Name        string         `json:"name" bson:"name"`
	Price       *MoneyInput    `json:"price" bson:"price"`
	Description string         `json:"description" bson:"description"`
	Image       string         `json:"image" bson:"image"`
	Author      string         `json:"author" bson:"author"`
	StoreName   string         `json:"storeName" bson:"storeName"`
	Location    *LocationInput `json:"location" bson:"location"`
}

type SessionID struct {
	Sessionid string `json:"sessionid" bson:"sessionid"`
}

type User struct {
	ID        *string `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string  `json:"email" bson:"email"`
	Password  *string `json:"password,omitempty" bson:"password,omitempty"`
	Sessionid *string `json:"sessionid,omitempty" bson:"sessionid,omitempty"`
	Points    int     `json:"points" bson:"points"`
}

type UserInput struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}