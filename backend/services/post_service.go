package services

import (
	"github.com/sasanzare/go-cms/models"
	"github.com/sasanzare/go-cms/utils"
	"errors"
	"time"
	"strings"

	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

// CreatePost creates a new post with validation
func (s *PostService) CreatePost(post *models.Post) error {
	// Validation
	if post.Title == "" {
		return errors.New(utils.ValidationFailedMsg + ": title is required")
	}
	if len(post.Title) > 255 {
		return errors.New(utils.ValidationFailedMsg + ": title exceeds 255 characters")
	}
	if post.Content == "" {
		return errors.New(utils.ValidationFailedMsg + ": content is required")
	}
	if post.AuthorID == 0 {
		return errors.New(utils.ValidationFailedMsg + ": author ID is required")
	}

	// Set defaults
	if post.Status == "" {
		post.Status = models.PostStatusDraft
	}
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	// Generate slug if empty
	if post.Slug == "" {
		post.Slug = generateSlug(post.Title)
	}

	return s.db.Create(post).Error
}

// GetPostByID retrieves a post by ID with relationships
func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := s.db.
		Preload("Author").
		Preload("Category").
		Preload("Tags").
		First(&post, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("post not found")
	}
	return &post, err
}

// UpdatePost updates an existing post
func (s *PostService) UpdatePost(id uint, updates map[string]interface{}) (*models.Post, error) {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		return nil, errors.New("post not found")
	}

	// Validate status transition
	if status, ok := updates["status"].(string); ok {
		if !isValidStatusTransition(post.Status, status) {
			return nil, errors.New("invalid status transition")
		}
		if status == models.PostStatusPublished && post.PublishedAt == nil {
			now := time.Now()
			updates["published_at"] = &now
		}
	}

	updates["updated_at"] = time.Now()

	if err := s.db.Model(&post).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// DeletePost soft-deletes a post
func (s *PostService) DeletePost(id uint) error {
	result := s.db.Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}

// ListPosts retrieves posts with filters
func (s *PostService) ListPosts(filter PostFilter) ([]models.Post, error) {
	var posts []models.Post
	query := s.db.Preload("Author").Preload("Category")

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.CategoryID != 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.AuthorID != 0 {
		query = query.Where("author_id = ?", filter.AuthorID)
	}

	err := query.Order("created_at DESC").Find(&posts).Error
	return posts, err
}

// PublishPost changes post status to published
func (s *PostService) PublishPost(id uint) (*models.Post, error) {
	now := time.Now()
	return s.UpdatePost(id, map[string]interface{}{
		"status":       models.PostStatusPublished,
		"published_at": &now,
	})
}

// IncrementViewCount increments the view count
func (s *PostService) IncrementViewCount(id uint) error {
	return s.db.Model(&models.Post{}).
		Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).
		Error
}

// Helper functions
func generateSlug(title string) string {
	// Implement your slug generation logic
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func isValidStatusTransition(oldStatus, newStatus string) bool {
	// Implement your status transition rules
	return true
}

// PostFilter defines filtering options for listing posts
type PostFilter struct {
	Status     string
	CategoryID uint
	AuthorID   uint
}