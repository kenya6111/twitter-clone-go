package domain

import (
	"context"
	"fmt"
	"twitter-clone-go/apperrors"
	"unicode/utf8"

	ulid "github.com/oklog/ulid/v2"
)

type Tweet struct {
	ID        string
	UserID    string
	Content   string
	ImageUrl  string
	ReplyToID string
}

type TweetRepository interface {
	InsertTweet(ctx context.Context, model *Tweet) (*Tweet, error)
}

func NewTweet(userId string, content string, ImageUrl string, ReplyToID string) (*Tweet, error) {
	// userIDの存在チェック

	// content のドメインルール
	if utf8.RuneCountInString(content) < contentLengthMin || utf8.RuneCountInString(content) > contentLengthMax {
		return nil, apperrors.ReqBadParam.Wrap(apperrors.ErrMismatchData, fmt.Sprintf("content must be between %s and %s characters", contentLengthMin, contentLengthMax))
	}
	// ImageUrl のドメインルール_
	if utf8.RuneCountInString(ImageUrl) < imageUrlLengthMin || utf8.RuneCountInString(ImageUrl) > imageUrlLengthMax {
		return nil, apperrors.ReqBadParam.Wrap(apperrors.ErrMismatchData, fmt.Sprintf("imageUrl must be between %s and %s characters", imageUrlLengthMin, imageUrlLengthMax))
	}
	// ReplyToIDの存在チェック

	return &Tweet{
		ID:        ulid.Make().String(),
		UserID:    userId,
		Content:   content,
		ImageUrl:  ImageUrl,
		ReplyToID: ReplyToID,
	}, nil
}

func ReconstructTweet(id string, userId string, content string, ImageUrl string, ReplyToID string) (*Tweet, error) {
	// userIDの存在チェック

	// content のドメインルール
	if utf8.RuneCountInString(content) < contentLengthMin || utf8.RuneCountInString(content) > contentLengthMax {
		return nil, apperrors.ReqBadParam.Wrap(apperrors.ErrMismatchData, fmt.Sprintf("content must be between %s and %s characters", contentLengthMin, contentLengthMax))
	}
	// ImageUrl のドメインルール_
	if utf8.RuneCountInString(ImageUrl) < imageUrlLengthMin || utf8.RuneCountInString(ImageUrl) > imageUrlLengthMax {
		return nil, apperrors.ReqBadParam.Wrap(apperrors.ErrMismatchData, fmt.Sprintf("imageUrl must be between %s and %s characters", imageUrlLengthMin, imageUrlLengthMax))
	}
	// ReplyToIDの存在チェック

	return &Tweet{
		ID:        id,
		UserID:    userId,
		Content:   content,
		ImageUrl:  ImageUrl,
		ReplyToID: ReplyToID,
	}, nil
}

const (
	contentLengthMax  = 140
	contentLengthMin  = 1
	imageUrlLengthMax = 999
	imageUrlLengthMin = 0
)
