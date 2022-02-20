package students

import "context"

type Repo interface {
	Add(ctx context.Context, do *StudentDO) error
	DeleteByUID(ctx context.Context, uid UID) error
}
