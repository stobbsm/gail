package response

// 421
type SrvNotAvailable struct{}

// 450
type ActionNotTaken struct{}

// 451
type Aborted struct{}

// 452
type NoStorage struct{}

// 455
type ParamError struct{}

// 500
type CmdSyntax struct{}

// 501
type ParamSyntax struct{}

// 502
type CmdNotImpl struct{}

// 503
type ParamNotImpl struct{}

// 550
type MailboxUnavil struct{}

// 551
type NotLocalUser struct{}

// 552
type StorageExceeded struct{}

// 553
type MailboxSyntax struct{}

// 554
func TransactionFailed(msg ...string) {
  s := &SMTPResponse{
	  code: 554,
  	msg:  []string{"Failed to commit transaction"},
  }
}

// 555
type MailFromOrRcpt struct{}
