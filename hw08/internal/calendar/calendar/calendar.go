package calendar

import (
	"github.com/dark705/otus/hw08/internal/calendar/event"
	"github.com/dark705/otus/hw08/internal/config"
	"github.com/dark705/otus/hw08/internal/storage"
	"github.com/sirupsen/logrus"
)

type Calendar struct {
	Config  config.Config
	Storage storage.Interface
	Logger  logrus.Logger
}

func (c Calendar) AddEvent(e event.Event) error {
	c.Logger.Debug("Try add to storage, Event:", e)
	err := c.Storage.Add(e)
	if err != nil {
		c.Logger.Debug("Fail add Event to storage:", err)
		return err
	}
	c.Logger.Info("Success add to storage, Event:", e)
	return nil
}

func (c Calendar) DelEvent(id int) error {
	c.Logger.Debug("Try del form storage Event, with Id:", id)
	err := c.Storage.Del(id)
	if err != nil {
		c.Logger.Debug("Fail del Event from storage:", err)
		return err
	}
	c.Logger.Info("Success del from storage Event, with Id:", id)
	return nil
}

func (c Calendar) GetEvent(id int) (event.Event, error) {
	c.Logger.Debug("Try get Event form storage, with Id:", id)
	event, err := c.Storage.Get(id)
	if err != nil {
		c.Logger.Debug("Fail get Event from storage:", err)
		return event, err
	}
	c.Logger.Info("Success get from storage Event, with Id:", id)
	return event, nil
}

func (c Calendar) GetAllEvents() ([]event.Event, error) {
	c.Logger.Debug("Try get all Events form storage")
	events, err := c.Storage.GetAll()
	if err != nil {
		c.Logger.Debug("Fail get Events from storage:", err)
		return events, err
	}
	c.Logger.Info("Success get from storage Events", len(events))
	return events, nil
}

func (c Calendar) EditEvent(e event.Event) error {
	c.Logger.Debug("Try edit Event in storage, with Id:", e.Id)
	err := c.Storage.Edit(e)
	if err != nil {
		c.Logger.Debug("Fail edit Event in storage:", err)
		return err
	}
	c.Logger.Info("Success edit Event in storage with Id:", e.Id)
	return nil
}