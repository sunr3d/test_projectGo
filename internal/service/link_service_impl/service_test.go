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
		Link:      "http://example.com",
		FakeLink:  "http://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	repo.On("Find", mock.Anything, inputLink.FakeLink).Return("", nil)
	repo.On("Create", mock.Anything, mock.Anything).Return(1, nil)

	id, err := svc.Create(context.Background(), inputLink)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	repo.AssertExpectations(t)
}

func TestService_Find(t *testing.T) {
	logger := zap.NewNop()
	repo := new(mocks.Database)
	svc := New(logger, repo)

	fakeLink := "http://fake.com"
	expectedLink := "http://example.com"

	repo.On("Find", mock.Anything, fakeLink).Return(expectedLink, nil)

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
		Link:      "http://example.com",
		FakeLink:  "http://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	repo.On("Find", mock.Anything, inputLink.FakeLink).Return("http://example.com", nil)

	id, err := svc.Create(context.Background(), inputLink)
	assert.Error(t, err)
	assert.Equal(t, 0, id)
	assert.Equal(t, "link already exists", err.Error())

	repo.AssertExpectations(t)
}
