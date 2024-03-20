package models

type Errormessage struct {
	Type       string
	Msg        string
	StatusCode int
	Display    bool
}

const (
	BRstatus = 404
	BRtype   = "Bad Request"
	//----------------------
	ISEstatus = 500
	ISEtype   = "Internal Servor Error"
	ISEmsg    = "Oops ! server didn't react as expected"
)
