package application

import "time"

// OpinionId unique identifier for an opinion in the system
type OpinionId string

// UserId unique identifier for an user in the system
type UserId string

// Opinion is submitted by an authenticated user.
// Only the owner or the admin is able to delete the opinion.
// If a User gets deleted in the system, all related opinions should be deleted too.
type Opinion struct {
	ID        OpinionId
	Owner     UserId
	CreatedAt time.Time
	Statement string
}

// OpinionCreateDTO holds required information to perform a create action on a opinion
type OpinionCreateDTO struct {
	Statement string
}

// Vote represents a users agreement or disagreement on the given opinion.
// A Vote can be created, updated or deleted.
type Vote struct {
	Agreement bool
	Opinion   OpinionId
	Voter     UserId
	CreatedAt time.Time
	UpdatedAt time.Time
}

// VoteCreateAndUpdateDTO holds required information to perform a create or update action on a vote
type VoteCreateAndUpdateDTO struct {
	Agreement bool
	Opinion   OpinionId
}

// AuthenticatedUser is capable to perform actions on opinions and votes
type AuthenticatedUser struct {
	Id UserId
}

// AuthorizedUser represents a user identity which is permitted to perform the action on the given resource
type AuthorizedUser struct {
	id       UserId
	action   string
	resource string
}

func (a AuthorizedUser) Id() UserId {
	return a.id
}

func (a AuthorizedUser) Action() string {
	return a.action
}

func (a AuthorizedUser) Resource() string {
	return a.resource
}
