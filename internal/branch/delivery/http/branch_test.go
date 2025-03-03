package http

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/require"
// 	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
// 	"gitlab.com/gma-vietnam/tanca-event/internal/models"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/jwt"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/util"
// )

// func TestDeliveryHTTPBranch_Create(t *testing.T) {
// 	jwtPayload := jwt.Payload{
// 		UserID: "test",
// 		ShopID: "62442f27c83a052e8f39e111",
// 	}

// 	scope := jwt.NewScope(jwtPayload)
// 	id := mongo.ObjectIDFromHexOrNil("62442f27c83a052e8f39e111")
// 	time, _ := util.StrToDateTime("2024-09-26 00:00:00")

// 	type mockUsecase struct {
// 		isCalled bool
// 		input    branch.CreateInput
// 		output   models.Branch
// 		err      error
// 	}

// 	tcs := map[string]struct {
// 		req         string
// 		isUnauth    bool
// 		mockUsecase mockUsecase
// 		wantBody    string
// 		wantCode    int
// 	}{
// 		"success": {
// 			req: `{
// 				"name": "test"
// 			}`,
// 			mockUsecase: mockUsecase{
// 				isCalled: true,
// 				input: branch.CreateInput{
// 					Name: "test",
// 				},
// 				output: models.Branch{
// 					ID:        id,
// 					Name:      "test",
// 					ShopID:    id,
// 					Alias:     "test",
// 					Code:      "test",
// 					CreatedAt: time,
// 				},
// 			},
// 			wantCode: http.StatusOK,
// 			wantBody: `{
// 				"error_code": 0,
// 				"message": "Success",
// 				"data": {
// 					"id": "62442f27c83a052e8f39e111",
// 					"alias": "test",
// 					"sort_index": 0,
// 					"shop_id": "62442f27c83a052e8f39e111",
// 					"code": "test",
// 					"created_at": "2024-09-26 07:00:00"
// 				}
// 			}`,
// 		},
// 		"Invalid body": {
// 			req: `{
// 			}`,
// 			wantCode: http.StatusBadRequest,
// 			wantBody: `{
// 				"error_code": 10000,
// 				"message": "Wrong body"
// 			}`,
// 		},
// 		"Failed validation": {
// 			req: `{
// 				"name": ""
// 			}`,
// 			wantBody: `{
// 				"error_code": 10000,
// 				"message": "Wrong body"
// 			}`,
// 			wantCode: http.StatusBadRequest,
// 		},
// 	}

// 	for name, tc := range tcs {
// 		t.Run(name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			c, engine := gin.CreateTestContext(w)
// 			h, deps := initHandler(t)

// 			engine.POST("/branches", h.create)
// 			c.Request = httptest.NewRequest(http.MethodPost, "/branches", bytes.NewBufferString(tc.req))

// 			if !tc.isUnauth {
// 				c.Request = c.Request.WithContext(
// 					jwt.SetPayloadToContext(c.Request.Context(), jwtPayload),
// 				)
// 			}

// 			if tc.mockUsecase.isCalled {
// 				deps.uc.EXPECT().Create(c.Request.Context(), scope, tc.mockUsecase.input).
// 					Return(
// 						tc.mockUsecase.output,
// 						tc.mockUsecase.err,
// 					)
// 			}

// 			// THEN
// 			engine.ServeHTTP(w, c.Request)
// 			require.Equal(t, tc.wantCode, w.Code)
// 			require.JSONEq(t, tc.wantBody, w.Body.String())
// 		})
// 	}
// }
