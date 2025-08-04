package queue

import (
	"amocrm2.0/internal/config"
	"github.com/beanstalkd/go-beanstalk"
)

type Beanstalk struct {
	Conn *beanstalk.Conn
}

func InitBeanstalk(cfg *config.Beanstalk) (*Beanstalk, error) {
	addr := cfg.Host + cfg.Port
	conn, err := beanstalk.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Beanstalk{Conn: conn}, nil
}
