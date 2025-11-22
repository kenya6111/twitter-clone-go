package http

import "twitter-clone-go/domain"

// UserResponse はユーザー情報のレスポンス用DTO
type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"isActive"`
}

// ToUserResponse はdomain.UserからUserResponseに変換する
func ToUserResponse(user *domain.User) *UserResponse {
	if user == nil {
		return nil
	}
	return &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
	}
}

// ToUserListResponse はdomain.UserのスライスからUserResponseのスライスに変換する
func ToUserListResponse(users []domain.User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, user := range users {
		responses[i] = UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			IsActive: user.IsActive,
		}
	}
	return responses
}

// CreateTweetResponse はツイート作成APIのレスポンス用DTO
type CreateTweetResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Content   string `json:"content"`
	ImageUrl  string `json:"imageUrl"`
	ReplyToID string `json:"replyToId,omitempty"`
}

// ToCreateTweetResponse はdomain.TweetからCreateTweetResponseに変換する
func ToCreateTweetResponse(tweet *domain.Tweet) *CreateTweetResponse {
	if tweet == nil {
		return nil
	}
	return &CreateTweetResponse{
		// ID:        tweet.ID,
		UserID:  tweet.UserID,
		Content: tweet.Content,
		// ImageUrl:  tweet.ImageUrl,
		// ReplyToID: tweet.ReplyToID,
	}
}
