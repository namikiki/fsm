package domain

import (
	"context"

	"fsm/pkg/ent"
)

type DirRepository interface {
	Create(ctx context.Context, f ent.Dir) error
	Delete(ctx context.Context, f ent.Dir) error
	Rename(ctx context.Context, f ent.Dir, newName string) error
	ReadDir(ctx context.Context, f ent.Dir) ([]ent.Dir, error)
	WalkDirByPath(ctx context.Context, f ent.Dir) ([]ent.Dir, error)
	WalkDirBySyncID(ctx context.Context, userID, syncID string) ([]ent.Dir, error)
}
