package service

import "github.com/karlozz157/storicard/src/domain/entity"

type INotificationService interface {
	Notify(summary *entity.Summary) error
}
