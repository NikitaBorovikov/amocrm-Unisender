package queue

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/beanstalkd/go-beanstalk"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	conn     *beanstalk.Conn
	tubeName string
}

func NewProducer(conn *beanstalk.Conn, tubeName string) *Producer {
	return &Producer{
		conn:     conn,
		tubeName: tubeName,
	}
}

func (p *Producer) AddSyncContactsTask(task SyncContactsTask) (uint64, error) {
	body, err := json.Marshal(task)
	if err != nil {
		return 0, err
	}
	tube := beanstalk.Tube{Conn: p.conn, Name: p.tubeName}
	id, err := tube.Put(body, 1024, 0, 10*time.Minute)
	logrus.Infof("task id = %d in tube", id)
	return id, err
}

func (p *Producer) FetchSyncContactsTask() (*SyncContactsTask, error) {
	tubeSet := beanstalk.NewTubeSet(p.conn, p.tubeName)
	id, body, err := tubeSet.Reserve(30 * time.Second)
	if err != nil {
		if errors.Is(err, beanstalk.ErrTimeout) {
			return nil, nil // задачи нет
		}
		return nil, err
	}
	defer tubeSet.Conn.Delete(id)

	var task SyncContactsTask
	if err := json.Unmarshal(body, &task); err != nil {
		return nil, err
	}
	return &task, nil
}
