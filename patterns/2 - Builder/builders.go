package main

import "strconv"

type MGUparserBuilder struct {
	dbconn          string
	configureParser []string
	logFile         string
}

func (b *MGUparserBuilder) setDBConnection(conn string) {
	b.dbconn = conn
}

func (b *MGUparserBuilder) setcConfigure(cfgs []string) {
	b.configureParser = cfgs
}

func (b *MGUparserBuilder) setLogFile(file string) {
	b.logFile = file
}

func (b *MGUparserBuilder) Build() (parser APIparser) {
	// To work to set connection
	parser.dbConnection = b.dbconn

	// for _, cfg := range cfgs {
	// appply configure to parser
	// }
	// b.configureParser = "MGU Parser"
	parser.parser = "MGU Parser"

	// Preapre log file
	parser.logFile = b.logFile

	return parser
}

type MISISparserBuilder struct {
	dbconn          string
	configureParser []string
	logFile         string
	misisCode       int
}

func (b *MISISparserBuilder) setDBConnection(conn string) {
	b.dbconn = conn
}

func (b *MISISparserBuilder) setcConfigure(cfgs []string) {
	b.configureParser = cfgs
}

func (b *MISISparserBuilder) setLogFile(file string) {
	b.logFile = file
}

func (b *MISISparserBuilder) setMISIScode(code int) {
	b.misisCode = code
}

func (b *MISISparserBuilder) Build() (parser APIparser) {
	// To work to set connection
	parser.dbConnection = b.dbconn

	// Check and apply misis code to do reqs
	// for _, cfg := range cfgs {
	// appply configure to parser
	// }
	// b.configureParser = "MGU Parser"
	parser.parser = "MISIS Parser with " + strconv.Itoa(b.misisCode)

	// Preapre log file
	parser.logFile = b.logFile

	return parser
}
