package domain

import (
	"context"
	"fmt"
	"twitter-clone-go/apperrors"
	"unicode/utf8"
)

type TweetImage struct {
	ID       int
	TweetID  int
	ImageUrl string
}

type TweetImageRepository interface {
	Insert(ctx context.Context, model []TweetImage) (int64, error)
}

func NewTweetImage(tweetId int, imageUrls []string) ([]TweetImage, error) {
	var result []TweetImage
	for _, url := range imageUrls {
		// ImageUrl のドメインルール_
		if utf8.RuneCountInString(url) < imageUrlLengthMin || utf8.RuneCountInString(url) > imageUrlLengthMax {
			return nil, apperrors.ReqBadParam.Wrap(apperrors.ErrMismatchData, fmt.Sprintf("imageUrl must be between %s and %s characters", imageUrlLengthMin, imageUrlLengthMax))
		}
		image := &TweetImage{
			TweetID:  tweetId,
			ImageUrl: url,
		}
		result = append(result, *image)
	}

	return result, nil
}
