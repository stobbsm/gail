package response

// 421
func SrvNotAvailable(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: SRVUNAVAIL,
		msg:  []string{"service not available"},
	}
	return s.append(msg...)
}

// 450
func ActionNotTaken(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: MBOXUNAVAIL,
		msg:  []string{"mailbox not available"},
	}
	return s.append(msg...)
}

// 451
func Aborted(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: ABORT,
		msg:  []string{"action aborted: server error"},
	}
	return s.append(msg...)
}

// 452
func NoStorage(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: OUTOFSTORAGE,
		msg:  []string{"aborted: storage full"},
	}
	return s.append(msg...)
}

// 455
func ParamError(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: BADPARAMS,
		msg:  []string{"server error: illegal parameters"},
	}
	return s.append(msg...)
}

// 500
func CmdSyntax(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: CMDSYNTAX,
		msg:  []string{"command syntax error"},
	}
	return s.append(msg...)
}

// 501
func ParamSyntax(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: PARAMSYNTAX,
		msg:  []string{"parameter syntax error"},
	}
	return s.append(msg...)
}

// 502
func CmdNotImpl(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: CMDNOTIMPL,
		msg:  []string{"command not implemented"},
	}
	return s.append(msg...)
}

// 503
func ParamNotImpl(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: BADSEQ,
		msg:  []string{"bad command sequence"},
	}
	return s.append(msg...)
}

// 550
func MailboxUnavil(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: DENIED,
		msg:  []string{"delivery failed: mailbox unavailable"},
	}
	return s.append(msg...)
}

// 551
func NotLocalUser(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: USRNOTLOCAL,
		msg:  []string{"delivery failed: user not local"},
	}
	return s.append(msg...)
}

// 552
func StorageExceeded(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: EXCEEDALLOC,
		msg:  []string{"action aborted: exceeded storage"},
	}
	return s.append(msg...)
}

// 553
func MailboxSyntax(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: NAMENOTALLOWED,
		msg:  []string{"Mailbox name syntax error"},
	}
	return s.append(msg...)
}

// 554
func TransactionFailed(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: FAILED,
		msg:  []string{"Failed to commit transaction"},
	}
	return s.append(msg...)
}

// 555
func MailFromErr(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: MAILTOERR,
		msg:  []string{"mail from not recognized"},
	}
	return s.append(msg...)
}

// 555
func RcptToErr(msg ...string) *SMTPResponse {
	s := &SMTPResponse{
		code: MAILTOERR,
		msg:  []string{"rcpt to not recognized"},
	}
	return s.append(msg...)
}
