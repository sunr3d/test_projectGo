package link_service_handler_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"link_service/internal/handlers/link_service"
	"link_service/internal/interfaces/services"
	"link_service/mocks"
	pb "link_service/proto/link_service"
)

func TestLinkService_GetLink(t *testing.T) {
	serviceMock := new(mocks.Service)
	handler := link_service_handler.New(serviceMock)

	fakeLink := "http://fake.com"
	expectedLink := "http://example.com"

	serviceMock.On("Find", mock.Anything, fakeLink).Return(expectedLink, nil)

	req := &pb.GetLinkRequest{Link: fakeLink}
	res, err := handler.GetLink(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expectedLink, res.Link)
	serviceMock.AssertExpectations(t)
}

func TestLinkService_InputLink(t *testing.T) {
	serviceMock := new(mocks.Service)
	handler := link_service_handler.New(serviceMock)

	inputLink := services.InputLink{
		Link:      "http://example.com",
		FakeLink:  "http://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	serviceMock.On("Create", mock.Anything, mock.MatchedBy(func(link services.InputLink) bool {
		return link.Link == inputLink.Link && link.FakeLink == inputLink.FakeLink
	})).Return(1, nil)

	req := &pb.InputLinkRequest{
		Link:      inputLink.Link,
		FakeLink:  inputLink.FakeLink,
		EraseTime: timestamppb.New(inputLink.EraseTime),
	}
	res, err := handler.InputLink(context.Background(), req)

	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "Link successfully added.", res.Message)
	assert.Equal(t, "1", res.Id)
	serviceMock.AssertExpectations(t)
}

func TestLinkService_InputLink_LinkAlreadyExists(t *testing.T) {
	serviceMock := new(mocks.Service)
	handler := link_service_handler.New(serviceMock)

	inputLink := services.InputLink{
		Link:      "http://example.com",
		FakeLink:  "http://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	serviceMock.On("Create", mock.Anything, mock.MatchedBy(func(link services.InputLink) bool {
		return link.Link == inputLink.Link && link.FakeLink == inputLink.FakeLink
	})).Return(0, errors.New("link already exists"))

	req := &pb.InputLinkRequest{
		Link:      inputLink.Link,
		FakeLink:  inputLink.FakeLink,
		EraseTime: timestamppb.New(inputLink.EraseTime),
	}
	res, err := handler.InputLink(context.Background(), req)

	assert.Error(t, err)
	assert.False(t, res.Success)
	assert.Equal(t, "link already exists", res.Message)
	assert.Equal(t, "n/a", res.Id)
	serviceMock.AssertExpectations(t)
}
