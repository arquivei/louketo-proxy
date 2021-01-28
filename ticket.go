/*
Copyright 2020 OAuth-Proxy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ticket is a structure representing the ticket used in server based
// session storage. It provides a unique per session decryption secret giving
// more security than the shared CookieSecret.
type ticket struct {
	id     string
	secret []byte
}

// newTicket creates a new ticket. The ID & secret will be randomly created
// with 16 byte sizes. The ID will be prefixed & hex encoded.
func newTicket() (*ticket, error) {
	id := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, id); err != nil {
		return nil, fmt.Errorf("failed to create new ticket ID: %v", err)
	}

	secret := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, secret); err != nil {
		return nil, fmt.Errorf("failed to create encryption secret: %v", err)
	}

	return &ticket{
		id:     hex.EncodeToString(id),
		secret: secret,
	}, nil
}

// encodeTicket encodes the Ticket to a string for usage in cookies
func (t *ticket) encodeTicket() string {
	return fmt.Sprintf("%s.%s", t.id, base64.RawURLEncoding.EncodeToString(t.secret))
}

// decodeTicket decodes an encoded ticket string
func decodeTicket(encTicket string) (*ticket, error) {
	ticketParts := strings.Split(encTicket, ".")
	if len(ticketParts) != 2 {
		return nil, errors.New("failed to decode ticket")
	}
	ticketID, secretBase64 := ticketParts[0], ticketParts[1]

	secret, err := base64.RawURLEncoding.DecodeString(secretBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode encryption secret: %v", err)
	}

	return &ticket{
		id:     ticketID,
		secret: secret,
	}, nil
}

// decodeTicketFromRequest retrieves a potential ticket cookie from a request
// and decodes it to a ticket.
func decodeTicketFromRequest(req *http.Request, config *Config) (*ticket, error) {
	cookie, err := req.Cookie(config.CookieTicketName)
	if err != nil {
		// Don't wrap this error to allow `err == http.ErrNoCookie` checks
		return nil, err
	}

	// An existing cookie exists, try to retrieve the ticket

	// Valid cookie, decode the ticket
	return decodeTicket(cookie.Value)
}
