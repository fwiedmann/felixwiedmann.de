package application_test

import (
	"context"
	"errors"
	"github.com/fwiedmann/site/backend/internal/opinions/application"
	mock_application "github.com/fwiedmann/site/backend/internal/opinions/application/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func Test_service_CreateOpinionCommand(t *testing.T) {
	const testUserId application.UserId = "1"
	const testDefaultId = "187"
	const testStatement = "copy and pasta is good!"

	repoError := errors.New("repo error")
	pepErrpr := errors.New("pep error")
	testDate := time.Now()

	type fields struct {
		repoError error
		pepError  error
	}
	type args struct {
		ctx     context.Context
		user    application.AuthenticatedUser
		opinion application.OpinionCreateDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    application.Opinion
		wantErr error
	}{
		{
			name: "Should throw error because empty opinion",
			fields: fields{
				repoError: nil,
			},
			args: args{
				ctx: context.Background(),
				user: application.AuthenticatedUser{
					Id: testUserId,
				},
				opinion: application.OpinionCreateDTO{
					Statement: "",
				},
			},
			want:    application.Opinion{},
			wantErr: application.EmptyOpinionStatementError,
		},
		{
			name: "Should throw error because repo error",
			fields: fields{
				repoError: repoError,
			},
			args: args{
				ctx: context.Background(),
				user: application.AuthenticatedUser{
					Id: testUserId,
				},
				opinion: application.OpinionCreateDTO{
					Statement: testStatement,
				},
			},
			want:    application.Opinion{},
			wantErr: repoError,
		},
		{
			name: "Should throw error because pep error",
			fields: fields{
				pepError: pepErrpr,
			},
			args: args{
				ctx: context.Background(),
				user: application.AuthenticatedUser{
					Id: testUserId,
				},
				opinion: application.OpinionCreateDTO{
					Statement: "copy pasta is good!",
				},
			},
			want:    application.Opinion{},
			wantErr: pepErrpr,
		},
		{
			name:   "Should successfully create an opinion",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				user: application.AuthenticatedUser{
					Id: testUserId,
				},
				opinion: application.OpinionCreateDTO{
					Statement: testStatement,
				},
			},
			want: application.Opinion{
				ID:        testDefaultId,
				Owner:     testUserId,
				CreatedAt: testDate,
				Statement: testStatement,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			idService := mock_application.NewMockIdService(ctrl)
			idService.EXPECT().GenerateId().Return(testDefaultId).MaxTimes(1)

			timeService := mock_application.NewMockTimeService(ctrl)
			timeService.EXPECT().CurrentTime().Return(testDate).MaxTimes(1)

			pep := mock_application.NewMockPolicyEnforcementPoint(ctrl)
			pep.EXPECT().RequestAccessForUser(gomock.Any(), gomock.Any()).Return(tt.fields.pepError)

			repo := mock_application.NewMockRepository(ctrl)
			repo.EXPECT().CreateOpinion(gomock.Any(), gomock.Any()).Return(tt.fields.repoError).MaxTimes(1)

			s := application.NewOpinionService(pep, repo, idService, timeService)
			got, err := s.CreateOpinionCommand(tt.args.ctx, tt.args.user, tt.args.opinion)

			if (err != nil) && tt.wantErr == nil {
				t.Errorf("CreateOpinionCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateOpinionCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOpinionCommand() got = %v, want %v", got, tt.want)
			}
		})
	}
}
