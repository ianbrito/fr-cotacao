// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Carrier struct {
	Reference        int64
	Name             string
	RegisteredNumber string
	StateInscription string
	LogoUrl          string
}

type Dispatcher struct {
	ID                         string
	RequestID                  string
	RegisteredNumberShipper    string
	RegisteredNumberDispatcher string
	ZipcodeOrigin              int32
}

type Offer struct {
	ID                           string
	DispatcherID                 sql.NullString
	Offer                        int32
	SimulationType               int32
	CarrierID                    int64
	Service                      string
	ServiceCode                  sql.NullString
	ServiceDescription           sql.NullString
	DeliveryTime                 json.RawMessage
	OriginalDeliveryTime         json.RawMessage
	Identifier                   sql.NullString
	DeliveryNote                 sql.NullString
	HomeDelivery                 sql.NullBool
	CarrierNeedsToReturnToSender sql.NullBool
	Expiration                   time.Time
	CostPrice                    float64
	FinalPrice                   float64
	Weights                      json.RawMessage
	Composition                  json.RawMessage
	Esg                          json.RawMessage
	Modal                        sql.NullString
}
