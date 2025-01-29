package utils

type todoErr int

const (
	NoErr todoErr = iota
	NotValidCommand
	NotValidArgs
)
