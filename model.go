package main

import (
	"database/sql"
	"errors"
)

type ticket struct {
  SerialNumber string `json:"serialnumber"`
	DeviceType string	`json:"devicetype"`
	DeviceInfo string `json:""devicetype`
}

func (p *ticket) getTickets(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (p *ticket) updateTickets(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (p *ticket) deleteTickets(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (p *ticket) createTickets(db *sql.DB) error {
  return errors.New("Not implemented")
}

func getTickets(db *sql.DB, start, count int) ([]ticket, error) {
  return nil, errors.New("Not implemented")
}