package response

type Code uint

const (
	// Positive status codes
	STATUS     Code = 211
	HELP            = 214
	READY           = 220
	CLOSING         = 221
	COMPLETE        = 250
	FORWARD         = 251
	NOVRFYSEND      = 252

	// Informational status codes
	MAILSTART = 354

	// Server error codes
	SRVUNAVAIL   = 421
	MBOXUNAVAIL  = 450
	ABORT        = 451
	OUTOFSTORAGE = 452
	BADPARAMS    = 455

	// Client error codes
	CMDSYNTAX      = 500
	PARAMSYNTAX    = 501
	CMDNOTIMPL     = 502
	BADSEQ         = 503
	PARAMNOTIMPL   = 504
	DENIED         = 550
	USRNOTLOCAL    = 551
	EXCEEDALLOC    = 552
	NAMENOTALLOWED = 553
	FAILED         = 554
	MAILTOERR      = 555
	RCPTTOERR      = 555
)
