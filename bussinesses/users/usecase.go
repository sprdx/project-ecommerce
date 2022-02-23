package users

import (
	context "context"
	"time"
)

type Usecase struct {
	Repo           DomainRepository
	contextTimeout time.Duration
}

func NewUsecase(repo DomainRepository, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (u *Usecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	message := domain.ValidateNewUserData()
	if message != "VALID" {
		var err error
		return Domain{}, err
	}

	return u.Repo.Create(ctx, domain)
}
