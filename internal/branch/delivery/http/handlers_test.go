package http

// import (
// 	"testing"

// 	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
// 	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
// )

// type mockDeps struct {
// 	uc *branch.MockUsecase
// }

// func initHandler(t *testing.T) (Handler, mockDeps) {
// 	l := log.InitializeTestZapLogger()
// 	uc := branch.NewMockUsecase(t)

// 	return New(l, uc), mockDeps{
// 		uc: uc,
// 	}
// }
