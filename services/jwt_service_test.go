package services

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_GenerateToken(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := newMockGenerateToken(ctrl)

	m.
		EXPECT()

}
