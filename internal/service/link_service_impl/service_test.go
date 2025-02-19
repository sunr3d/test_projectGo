package link_service_impl

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"link_service/internal/interfaces/services"
	"link_service/mocks"
)

func TestService_Create(t *testing.T) {
	logger := zap.NewNop()
	repo := new(mocks.Database)
	svc := New(logger, repo)

	inputLink := services.InputLink{
		Link:      "https://example.com",
		FakeLink:  "https://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	var emptyLink *string
	repo.On("Find", mock.Anything, inputLink.FakeLink).Return(emptyLink, nil)
	repo.On("Create", mock.Anything, mock.Anything).Return(nil)

	err := svc.Create(context.Background(), inputLink)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestService_Find(t *testing.T) {
	logger := zap.NewNop()
	repo := new(mocks.Database)
	svc := New(logger, repo)

	fakeLink := "https://fake.com"
	expectedLink := "https://example.com"

	repo.On("Find", mock.Anything, fakeLink).Return(&expectedLink, nil)

	link, err := svc.Find(context.Background(), fakeLink)
	assert.NoError(t, err)
	assert.Equal(t, expectedLink, link)

	repo.AssertExpectations(t)
}

func TestService_Create_LinkAlreadyExists(t *testing.T) {
	logger := zap.NewNop()
	repo := new(mocks.Database)
	svc := New(logger, repo)

	inputLink := services.InputLink{
		Link:      "https://example.com",
		FakeLink:  "https://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	repo.On("Find", mock.Anything, inputLink.FakeLink).Return(&inputLink.Link, nil)

	err := svc.Create(context.Background(), inputLink)
	assert.Error(t, err)
	assert.Equal(t, "link already exists", err.Error())

	repo.AssertExpectations(t)
}
