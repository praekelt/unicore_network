package server

import (
	"github.com/go-martini/martini"
)

func New(identity Ident) *martini.ClassicMartini {
	m := martini.Classic()
	m.Map(identity)
	m.Use(DB())
	m.Get("/identity", GetIdentity)
	return m
}