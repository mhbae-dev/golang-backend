package main

import (
	"database/sql"
	"errors"
)

type ticket struct {
  SerialNumber int `json:"serial_number"`
	DeviceType string	`json:"device_type"`
	DeviceInfo string `json:""device_info`
}

// func (p *ticket) getTicket(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

// func (p *ticket) updateTickets(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

// func (p *ticket) deleteTickets(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

// func (p *ticket) createTickets(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

func getTickets(db *sql.DB, start, count int) ([]ticket, error) {
  return nil, errors.New("Not implemented")
}