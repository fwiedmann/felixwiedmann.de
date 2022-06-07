package application

//go:generate mockgen -destination mocks/mock.go . Repository,PolicyEnforcementPoint,IdService,TimeService

import (
	"context"
	"errors"
	"time"
)

type Service interface {
	CreateOpinionCommand(ctx context.Context, user AuthenticatedUser, opinion OpinionCreateDTO) (Opinion, error)
	DeleteOpinionCommand(ctx context.Context, user AuthenticatedUser, id OpinionId) error
	HandleUserDeletionEvent(ctx context.Context, event any) error

	CreateVoteCommand(ctx context.Context, user AuthenticatedUser, vote VoteCreateAndUpdateDTO) (Vote, error)
	UpdateVoteCommand(ctx context.Context, user AuthenticatedUser, vote VoteCreateAndUpdateDTO) (Vote, error)
	DeleteVoteCommand(ctx context.Context, user AuthenticatedUser, id OpinionId) (Vote, error)
}

type Repository interface {
	CreateOpinion(ctx context.Context, opinion Opinion) error
	DeleteOpinion(ctx context.Context, id OpinionId) error

	CreateVote(ctx context.Context, vote Vote) error
	UpdateVote(ctx context.Context, vote Vote) error
	DeleteVote(ctx context.Context, id OpinionId) error
}

type PolicyEnforcementPoint interface {
	RequestAccessForUser(userId string, action string) error
}

type IdService interface {
	GenerateId() string
}

type TimeService interface {
	CurrentTime() time.Time
}

func NewOpinionService(point PolicyEnforcementPoint, repository Repository, idService IdService, timeService TimeService) Service {
	return &service{
		pep:         point,
		repo:        repository,
		idService:   idService,
		timeService: timeService,
	}
}

const (
	// ActionCreateOpinion will be used for the user policy enforcement
	ActionCreateOpinion = "CreateOpinion"
)

var (
	// EmptyOpinionStatementError can be returned during OpinionCreateDTO validation
	EmptyOpinionStatementError = errors.New("opinion statement is empty")
)

type service struct {
	pep         PolicyEnforcementPoint
	repo        Repository
	idService   IdService
	timeService TimeService
}

// CreateOpinionCommand handles the create command for the frontend
func (s *service) CreateOpinionCommand(ctx context.Context, user AuthenticatedUser, opinion OpinionCreateDTO) (Opinion, error) {
	if err := s.pep.RequestAccessForUser(string(user.Id), ActionCreateOpinion); err != nil {
		return Opinion{}, err
	}

	if opinion.Statement == "" {
		return Opinion{}, EmptyOpinionStatementError
	}

	o := Opinion{
		ID:        OpinionId(s.idService.GenerateId()),
		Owner:     user.Id,
		CreatedAt: s.timeService.CurrentTime(),
		Statement: opinion.Statement,
	}

	if err := s.repo.CreateOpinion(ctx, o); err != nil {
		return Opinion{}, err
	}
	return o, nil
}

func (s *service) DeleteOpinionCommand(ctx context.Context, user AuthenticatedUser, id OpinionId) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) HandleUserDeletionEvent(ctx context.Context, event any) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) CreateVoteCommand(ctx context.Context, user AuthenticatedUser, vote VoteCreateAndUpdateDTO) (Vote, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateVoteCommand(ctx context.Context, user AuthenticatedUser, vote VoteCreateAndUpdateDTO) (Vote, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteVoteCommand(ctx context.Context, user AuthenticatedUser, id OpinionId) (Vote, error) {
	//TODO implement me
	panic("implement me")
}
