package usecase

// import (
// 	context "context"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
// 	"gitlab.com/gma-vietnam/tanca-event/internal/models"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
// )

// func TestBranchUsecase_Create(t *testing.T) {
// 	scope := models.Scope{
// 		UserID: "test",
// 	}

// 	type mockRepo struct {
// 		expCall bool
// 		input   branch.CreateOptions
// 		want    models.Branch
// 		err     error
// 	}

// 	tcs := map[string]struct {
// 		input    branch.CreateInput
// 		mockRepo mockRepo
// 		want     models.Branch
// 		wantErr  error
// 	}{
// 		"success": {
// 			input: branch.CreateInput{
// 				Name: "test",
// 			},
// 			mockRepo: mockRepo{
// 				expCall: true,
// 				input: branch.CreateOptions{
// 					Name:  "test",
// 					Alias: "test",
// 					Code:  "TEST",
// 				},
// 			},
// 			want: models.Branch{
// 				Name:   "test",
// 				Alias:  "test",
// 				Code:   "TEST",
// 				ShopID: mongo.ObjectIDFromHexOrNil("62442f27c83a052e8f39e111"),
// 			},
// 		},
// 		"error": {
// 			input: branch.CreateInput{
// 				Name: "test",
// 			},
// 			mockRepo: mockRepo{
// 				expCall: true,
// 				input: branch.CreateOptions{
// 					Name:  "test",
// 					Alias: "test",
// 					Code:  "TEST",
// 				},
// 			},
// 			wantErr: errors.New("Some error"),
// 		},
// 	}

// 	for desc, tc := range tcs {
// 		t.Run(desc, func(t *testing.T) {
// 			ctx := context.Background()

// 			uc, deps := initUseCase(t)

// 			deps.repo.EXPECT().Create(ctx, scope, tc.mockRepo.input).
// 				Return(tc.mockRepo.want, tc.mockRepo.err)

// 			res, err := uc.Create(ctx, scope, tc.input)
// 			if err != nil {
// 				require.EqualError(t, err, tc.wantErr.Error())
// 			} else {
// 				require.NoError(t, err)
// 				require.Equal(t, tc.mockRepo.want, res)
// 			}
// 		})
// 	}
// }
