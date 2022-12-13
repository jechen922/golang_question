package repository

import "demo/model/po"

type MemberRepo interface {
	GetByID(id int) (*po.Member, error)
	IsNameExist(name string) (bool, error)
	Create(name string) error
}
