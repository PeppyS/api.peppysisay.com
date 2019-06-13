package service

import (
	"net/http"
)

type ActivityService struct {
	httpClient *http.Client
}

func NewActivity(httpClient *http.Client) *ActivityService {
	return &ActivityService{httpClient}
}

func (a *ActivityService) GetRecentCodeCommits() interface {

}
