package count

import (
	"github.com/golang/mock/gomock"
	"mock/mock"
	"testing"
)

func TestFoo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBar := mock.NewMockBar(ctl)
	gomock.InOrder(mockBar.EXPECT().EXE(1).Return(1))
	fo := NewFoo(mockBar)
	if fo.bar.EXE(1) != 1 {
		t.Error("failed")
	}
}
