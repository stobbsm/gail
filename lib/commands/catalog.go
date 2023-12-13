package commands

import (
	"github.com/stobbm/gail/lib/response"
)

// Catalog of SMTP commands

// Command represents an SMTP command, where args are the expected arguements
type Command interface {
	//Handle accepts arguments (0->many) as strings and returns a CommandResponse
	Handle(args ...string) response.CommandResponse
}

type Ehlo struct{}
type Helo struct{}
type Mail struct{}
type Rcpt struct{}
type Data struct{}
type Rset struct{}
type Vrfy struct{}
type Expn struct{}
type Help struct{}
type Noop struct{}
type Quit struct{}
