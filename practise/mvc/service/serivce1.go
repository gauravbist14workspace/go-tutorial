package service

import (
	repo "go_tutorial/practise/mvc/repository"
)

type IService interface {
	GetDataFromRepo(input any) any
}

type Service1 struct {
	dao repo.IRepository1
}

func NewService1(repo1 repo.IRepository1) *Service1 {
	return &Service1{
		dao: repo1,
	}
}

func (service Service1) GetDataFromRepo(input any) any {
	return nil
}
