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
	cache := new(mocks.Cache)
	svc := New(logger, repo, cache)

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

func TestService_Find_in_DB(t *testing.T) {
	logger := zap.NewNop()
	repo := new(mocks.Database)
	cache := new(mocks.Cache)
	svc := New(logger, repo, cache)

	fakeLink := "https://fake.com"
	expectedLink := "https://example.com"

	repo.On("Find", mock.Anything, fakeLink).Return(&expectedLink, nil)
	cache.On("Get", mock.Anything, fakeLink).Return("", nil)
	//cache.On("Set", mock.Anything, fakeLink, expectedLink).Return(nil) // из-за горутины в Set -- не работает

	link, err := svc.Find(context.Background(), fakeLink)
	assert.NoError(t, err)
	assert.Equal(t, expectedLink, link)

	cache.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestService_Find_in_Cache(t *testing.T) {
	logger := zap.NewNop()
	repo := new(mocks.Database)
	cache := new(mocks.Cache)
	svc := New(logger, repo, cache)

	fakeLink := "https://fake.com"
	expectedLink := "https://example.com"

	cache.On("Get", mock.Anything, fakeLink).Return(expectedLink, nil)

	repo.AssertNotCalled(t, "Find", mock.Anything, fakeLink)

	link, err := svc.Find(context.Background(), fakeLink)
	assert.NoError(t, err)
	assert.Equal(t, expectedLink, link)

	cache.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestService_Create_LinkAlreadyExists(t *testing.T) {
	logger := zap.NewNop()
	repo := new(mocks.Database)
	cache := new(mocks.Cache)
	svc := New(logger, repo, cache)

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
