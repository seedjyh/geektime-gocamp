package students

import (
	"context"
	"fmt"
)

func New(repo Repo) *Service {
	return &Service{repo: repo}
}

type Service struct {
	repo Repo
}

func (s *Service) Register(ctx context.Context, do *StudentDO) error {
	if !do.UID.Valid() {
		return fmt.Errorf("invalid uid")
	}
	if !do.RealName.Valid() {
		return fmt.Errorf("invalid real name")
	}
	return s.repo.Add(ctx, do)
}

func (s *Service) Unregister(ctx context.Context, uid UID) error {
	if !uid.Valid() {
		return fmt.Errorf("invalid uid")
	}
	return s.repo.DeleteByUID(ctx, uid)
}
