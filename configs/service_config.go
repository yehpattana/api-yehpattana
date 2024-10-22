package configs

import (
	"fmt"
	"time"
)

type service struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
}

func (s *service) Url() string                 { return fmt.Sprintf("%s:%d", s.host, s.port) } // host:port
func (s *service) Name() string                { return s.name }
func (s *service) Version() string             { return s.version }
func (s *service) ReadTimeout() time.Duration  { return s.readTimeout }
func (s *service) WriteTimeout() time.Duration { return s.writeTimeout }
func (s *service) BodyLimit() int              { return s.bodyLimit }
func (s *service) FileLimit() int              { return s.fileLimit }
func (s *service) Host() string                { return s.host }
func (s *service) Port() int                   { return s.port }
