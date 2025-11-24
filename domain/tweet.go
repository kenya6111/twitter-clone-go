package domain

import (
	"context"
	"fmt"
	"strconv"
	"twitter-clone-go/apperrors"
	"unicode/utf8"
)

type Tweet struct {
	ID        int
	UserID    string
	Content   string
	Images    []TweetImage
	ReplyToID *int
}

type TweetRepository interface {
	Insert(ctx context.Context, model *Tweet) (*Tweet, error)
}

func NewTweet(userId string, content string, ReplyToID string) (*Tweet, error) {
	// content のドメインルール
	if utf8.RuneCountInString(content) < contentLengthMin || utf8.RuneCountInString(content) > contentLengthMax {
		return nil, apperrors.ReqBadParam.Wrap(apperrors.ErrMismatchData, fmt.Sprintf("content must be between %s and %s characters", contentLengthMin, contentLengthMax))
	}
	// ReplyToID
	replyId, err := strconv.Atoi(ReplyToID)
	if err != nil {
		return nil, apperrors.ReqBadParam.Wrap(apperrors.ErrMismatchData, fmt.Sprintf("invalid reply_to_id: %v", ReplyToID))
	}
	return &Tweet{
		UserID:    userId,
		Content:   content,
		ReplyToID: &replyId,
	}, nil
}

const (
	contentLengthMax  = 140
	contentLengthMin  = 1
	imageUrlLengthMax = 999
	imageUrlLengthMin = 0
)
