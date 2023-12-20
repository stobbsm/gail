package response

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// Default errors for SMTP

// CommandResponse contains the result of a command. Error's should be checked
// after each command using the Error() method
type CommandResponse interface {
	//Error returns either an error or nil. If it contains an error, that will be
	//sent as the message with the code in the SMTP response for the command.
	Error() error
	//Code is the integer result code of the SMTP result, as specificed in
	// https://www.rfc-editor.org/rfc/rfc5321.html#section-4.2
	Code() int
	//Msg returns the reply string of the valid message. Multiline responses are
	//returned this way, prefixing with 'Code()-{Msg Line}CRLF'
	Msg() []string
	//String returns the reply as a string, combining multiline responses as
	//needed
	String() string
}

// SMTPError handles and formats SMTP errors
type SMTPResponse struct {
	code int
	msg  []string
}

// Error returns the SMTPError in the format expected by the SMTP protocol,
// ending in a CRLF
func (s *SMTPResponse) Error() string {
	return s.String()
}

func (s *SMTPResponse) Code() int     { return s.code }
func (s *SMTPResponse) Msg() []string { return s.msg }
func (s *SMTPResponse) String() string {
	var b strings.Builder
	if err := WriteSMTPResponse(&b, s.code, s.msg...); err != nil {
		log.Println(err)
	}
	return b.String()
}

func (s *SMTPResponse) append(msg ...string) *SMTPResponse {
	for _, m := range msg {
		s.msg = append(s.msg, m)
	}
	return s
}

// WriteSMTPResponse uses the code and builds the given lines, as strings,
// into a singular string output to write to the response buffer
func WriteSMTPResponse(b io.Writer, code int, lines ...string) error {
	if len(lines) == 0 {
		_, err := fmt.Fprintf(b, "%d\r\n", code)
		return err
	}
	for n, v := range lines {
		if n == 0 {
			if _, err := fmt.Fprintf(b, "%d %s\r\n", code, v); err != nil {
				return err
			}
			continue
		}
		if _, err := fmt.Fprintf(b, "%d-%s\r\n", code, v); err != nil {
			return err
		}
	}
	return nil
}

// 211
type SysStat struct{}

// 214
type HelpMsg struct{}

// 220
type SrvReady struct{}

// 221
type SrvClose struct{}

// 250
type RqstComplete struct{}

// 251
type FwdTo struct{}

// 252
type NoVrfyWillAttempt struct{}

// 354
type StartMail struct{}
