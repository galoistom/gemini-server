package main

import (
	"fmt"
	"net"
)

type geminiResponse struct {
	conn        net.Conn
	status      int
	meta        string
	wroteHeader bool
}

func newResponse(conn net.Conn) *geminiResponse {
	return &geminiResponse{
		conn:        conn,
		status:      20,
		meta:        "text/gemini",
		wroteHeader: false,
	}
}

func (w *geminiResponse) WriteHeader(status int, meta string) error {
	if w.wroteHeader {
		return fmt.Errorf("header already written")
	}
	w.SetStatus(status, meta)
	header := fmt.Sprintf("%d %s\r\n", status, meta)
	_, err := w.conn.Write([]byte(header))
	if err != nil {
		return err
	}
	w.wroteHeader = true
	return nil
}

func (w *geminiResponse) Write(data []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(w.status, w.meta)
	}
	return w.conn.Write(data)
}

func (w *geminiResponse) WriteString(s string) (int, error) {
	return w.conn.Write([]byte(s))
}

func (w *geminiResponse) SetStatus(status int, meta string) {
	w.status = status
	w.meta = meta
}
