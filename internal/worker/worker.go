package worker

import (
	"context"
	"time"

	"amocrm2.0/internal/core/amocrm"
	"amocrm2.0/internal/infrastructure/queue"
	"amocrm2.0/internal/infrastructure/transport/http/dto"
	"amocrm2.0/internal/infrastructure/transport/http/handlers"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	ID       int
	Handlers *handlers.Handlers
	Queue    *queue.Producer
}

func NewWorker(id int, handlers *handlers.Handlers, queue *queue.Producer) *Worker {
	return &Worker{
		ID:       id,
		Handlers: handlers,
		Queue:    queue,
	}
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			w.processTask()
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker) processTask() {
	task, err := w.Queue.FetchSyncContactsTask()
	if err != nil {
		logrus.Errorf("failed to get task from queue: %v", err)
		return
	}

	if task == nil {
		return
	}

	logrus.Infof("Worker %d is processing task %s", w.ID, task.TaskType)
	w.syncContacts(task)
}

func (w *Worker) syncContacts(task *queue.SyncContactsTask) {
	switch task.TaskType {
	case "first_sync":
		w.handleFirstSync(task)
	case "webhook_sync":
		w.handleWebhookSync(task)
	default:
		logrus.Error("invalid task type")
		return
	}
}

func (w *Worker) handleFirstSync(taskInfo *queue.SyncContactsTask) {
	contacts, err := w.Handlers.GetContacts(taskInfo.AccountID)
	if err != nil {
		logrus.Error(err)
		return
	}

	importData := prepareUnisenderImportData(taskInfo.UnisenderKey, contacts)
	if err := w.Handlers.SendContactsToUnisender(importData); err != nil {
		logrus.Error(err)
		return
	}
	// TODO: сохранение контактов в БД в зависимости от статуса ответа Unisender.
	logrus.Infof("successfull sync: accountId = %d, taskType = %s", taskInfo.AccountID, taskInfo.TaskType)
}

func (w *Worker) handleWebhookSync(taskInfo *queue.SyncContactsTask) {
	apiKey, err := w.Handlers.UseCases.AccountUC.GetUnisenderKey(taskInfo.AccountID)
	if err != nil {
		logrus.Errorf("failed to get Unisender API key: %v", err)
		return
	}

	importData := prepareUnisenderImportData(apiKey, taskInfo.Contacts)
	if err := w.Handlers.SendContactsToUnisender(importData); err != nil {
		logrus.Error(err)
		return
	}
	// TODO: сохранение контактов в БД в зависимости от статуса ответа Unisender.
	logrus.Infof("successfull sync: accountId = %d, taskType = %s", taskInfo.AccountID, taskInfo.TaskType)
}

func prepareUnisenderImportData(apiKey string, contacts []amocrm.Contact) *dto.UnisenderImportRequest {
	fieldsName := []string{"name", "email"}
	data := make([][]string, len(contacts))
	for idx, contact := range contacts {
		data[idx] = []string{contact.Name, contact.Email}
	}

	return &dto.UnisenderImportRequest{
		APIKey:     apiKey,
		FieldNames: fieldsName,
		Data:       data,
	}
}
