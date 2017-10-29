// Package noopmailer contains an no-op implementation of the mailer interface
package noopmailer

import (
	mailer "github.com/Nivl/go-mailer"
)

// Makes sure Noop implements Mailer
var _ mailer.Mailer = (*Mailer)(nil)

// Mailer is a mailer that does nothing
type Mailer struct {
}

// SendStackTrace emails the current stacktrace to the default FROM
func (m *Mailer) SendStackTrace(trace []byte, message string, context map[string]string) error {
	return nil
}

// Send is used to send an email
func (m *Mailer) Send(msg *mailer.Message) error {
	return nil
}
