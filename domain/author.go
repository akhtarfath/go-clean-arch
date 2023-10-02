package domain

import "context"

// Author representing the Author data struct
type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthorUseCase represent the author's use-cases
type AuthorUseCase interface {
	GetByID(ctx context.Context, id int64) (Author, error)
}

// AuthorRepository represent the author's repository contract
type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Author, error)
}
