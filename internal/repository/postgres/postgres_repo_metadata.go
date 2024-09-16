package postgres

import (
	"time"

	"github.com/kenmobility/git-api-service/internal/domain"
)

// Repository represents the Postgres model for the repositories table.
type Repository struct {
	ID                uint   `gorm:"primarykey"`
	PublicID          string `gorm:"type:varchar;uniqueIndex"`
	Name              string `gorm:"type:varchar;unique"`
	Description       string `gorm:"type:text"`
	URL               string `gorm:"type:varchar"`
	Language          string `gorm:"type:varchar"`
	ForksCount        int
	StarsCount        int
	OpenIssuesCount   int
	WatchersCount     int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastFetchedCommit string `gorm:"type:varchar"`
	IsFetching        bool
	LastFetchedPage   int32 `gorm:"default:1"`
}

// ToDomain converts a Postgres Repository object to domain entity RepoMetadata.
func (pr *Repository) ToDomain() *domain.RepoMetadata {
	return &domain.RepoMetadata{
		PublicID:          pr.PublicID,
		Name:              pr.Name,
		Description:       pr.Description,
		URL:               pr.URL,
		Language:          pr.Language,
		ForksCount:        pr.ForksCount,
		StarsCount:        pr.StarsCount,
		OpenIssuesCount:   pr.OpenIssuesCount,
		WatchersCount:     pr.WatchersCount,
		CreatedAt:         pr.CreatedAt,
		UpdatedAt:         pr.UpdatedAt,
		LastFetchedCommit: pr.LastFetchedCommit,
		IsFetching:        pr.IsFetching,
		LastFetchedPage:   pr.LastFetchedPage,
	}
}

// FromDomainRepo returns a Postgres Repository object from domain entity RepoMetadata.
func FromDomainRepo(r *domain.RepoMetadata) *Repository {
	return &Repository{
		PublicID:          r.PublicID,
		Name:              r.Name,
		Description:       r.Description,
		URL:               r.URL,
		Language:          r.Language,
		ForksCount:        r.ForksCount,
		StarsCount:        r.StarsCount,
		OpenIssuesCount:   r.OpenIssuesCount,
		WatchersCount:     r.WatchersCount,
		CreatedAt:         r.CreatedAt,
		UpdatedAt:         r.UpdatedAt,
		LastFetchedCommit: r.LastFetchedCommit,
		IsFetching:        r.IsFetching,
		LastFetchedPage:   r.LastFetchedPage,
	}
}
