package application

import (
	"context"
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
	CreateOpinion(ctx context.Context, user AuthorizedUser, opinion Opinion) error
	DeleteOpinion(ctx context.Context, user AuthorizedUser, id OpinionId) error

	CreateVote(ctx context.Context, user AuthenticatedUser, vote Vote) error
	UpdateVote(ctx context.Context, user AuthenticatedUser, vote Vote) error
	DeleteVote(ctx context.Context, user AuthenticatedUser, id OpinionId) error
}
