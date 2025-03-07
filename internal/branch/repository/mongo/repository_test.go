package mongo

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// 	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
// 	"gitlab.com/gma-vietnam/tanca-event/internal/models"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/util"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type mockDeps struct {
// 	db  *mongo.MockDatabase
// 	col *mongo.MockCollection
// }

// func initRepo(t *testing.T, mockTime time.Time) (branch.Repository, mockDeps) {
// 	l := log.InitializeTestZapLogger()

// 	db := mongo.NewMockDatabase(t)
// 	col := mongo.NewMockCollection(t)

// 	db.EXPECT().Collection("branches").
// 		Return(col)

// 	repo := &implRepository{
// 		l:     l,
// 		db:    db,
// 		clock: func() time.Time { return mockTime },
// 	}
// 	return repo, mockDeps{
// 		db:  db,
// 		col: col,
// 	}
// }

// func TestBranchRepository_Create(t *testing.T) {
// 	scope := models.Scope{
// 		UserID: "user-id",
// 	}

// 	mockTime, _ := util.StrToDateTime("2024-01-01 00:00:00")
// 	mockID, _ := primitive.ObjectIDFromHex("667788990011223344556677")

// 	type mockCreateInput struct {
// 		branch models.Branch
// 	}

// 	type mockInsertOne struct {
// 		isCalled bool
// 		input    mockCreateInput
// 		output   interface{}
// 		err      error
// 	}

// 	type mockNewObjectID struct {
// 		isCalled bool
// 		output   primitive.ObjectID
// 		err      error
// 	}

// 	type mockColl struct {
// 		insertOne   mockInsertOne
// 		newObjectID mockNewObjectID
// 	}

// 	tcs := map[string]struct {
// 		input    branch.CreateOptions
// 		mockColl mockColl
// 		wantRes  models.Branch
// 		wantErr  error
// 	}{
// 		"success": {
// 			input: branch.CreateOptions{
// 				Name:  "Branch Name",
// 				Alias: "branch-name",
// 				Code:  "BRANCH-NAME",
// 			},
// 			mockColl: mockColl{
// 				insertOne: mockInsertOne{
// 					isCalled: true,
// 					input: mockCreateInput{
// 						branch: models.Branch{
// 							ID:        mockID,
// 							Name:      "Branch Name",
// 							Alias:     "branch-name",
// 							Code:      "BRANCH-NAME",
// 							CreatedAt: mockTime,
// 							UpdatedAt: mockTime,
// 						},
// 					},
// 					output: interface{}(nil),
// 					err:    nil,
// 				},
// 				newObjectID: mockNewObjectID{
// 					isCalled: true,
// 					output:   mockID,
// 					err:      nil,
// 				},
// 			},
// 			wantRes: models.Branch{
// 				ID:        mockID,
// 				Name:      "Branch Name",
// 				Alias:     "branch-name",
// 				Code:      "BRANCH-NAME",
// 				CreatedAt: mockTime,
// 				UpdatedAt: mockTime,
// 			},
// 			wantErr: nil,
// 		},
// 	}

// 	for name, tc := range tcs {
// 		t.Run(name, func(t *testing.T) {
// 			ctx := context.Background()

// 			repo, deps := initRepo(t, mockTime)

// 			if tc.mockColl.insertOne.isCalled {
// 				deps.col.EXPECT().InsertOne(ctx, tc.mockColl.insertOne.input.branch).
// 					Return(tc.mockColl.insertOne.output, tc.mockColl.insertOne.err)
// 			}

// 			if tc.mockColl.newObjectID.isCalled {
// 				deps.db.EXPECT().NewObjectID().
// 					Return(tc.mockColl.newObjectID.output)
// 			}

// 			gotBranch, gotErr := repo.Create(ctx, scope, tc.input)

// 			require.Equal(t, tc.wantRes, gotBranch)
// 			require.Equal(t, tc.wantErr, gotErr)
// 		})
// 	}
// }
