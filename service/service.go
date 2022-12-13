package service

import (
	"demo/apierror"
	"demo/model/bo"
	"demo/repository"
	"demo/service/utils"
)

type MemberService struct {
	memberRepo repository.MemberRepo
	utils.ErrorHandler
}

func (s *MemberService) Member(id int) (*bo.Member, *apierror.Error) {
	member, err := s.memberRepo.GetByID(id)
	if err != nil {
		return nil, s.NotFoundOrInternalError(apierror.RecordNotFound, err)
	}
	return &bo.Member{Name: member.NickName}, nil
}

func (s *MemberService) Create(name string) *apierror.Error {
	exist, err := s.memberRepo.IsNameExist(name)
	if exist || err != nil {
		return s.ErrorOrInternalError(apierror.Conflict, err)
	}

	if err = s.memberRepo.Create(name); err != nil {
		return apierror.New(apierror.InternalServiceError, err)
	}
	return nil
}

type ArticleService struct {
	errorHandler utils.ErrorHandler
}
