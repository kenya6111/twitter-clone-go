package application

import (
	"context"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"
	"twitter-clone-go/domain/service"
)

type TweetInfo struct {
	UserId    string              `json:"userId"`
	Content   string              `json:"content"`
	ImgFile   []service.FileInput `json:"image"`
	ReplyToId string              `json:"replyToId"`
}

type TweetUsecaseImpl struct {
	tweetRepo         domain.TweetRepository
	tweetImageRepo    domain.TweetImageRepository
	transaction       domain.Transaction
	fileUploadService service.FileUploader
}

type TweetUsecase interface {
	CreateTweet(ctx context.Context, tweetInfo TweetInfo) (*domain.Tweet, error)
}

func NewTweetUsecase(
	tweetRepo domain.TweetRepository,
	tweetImageRepo domain.TweetImageRepository,
	transaction domain.Transaction,
	fileUploadService service.FileUploader,
) *TweetUsecaseImpl {
	return &TweetUsecaseImpl{
		tweetRepo:         tweetRepo,
		tweetImageRepo:    tweetImageRepo,
		transaction:       transaction,
		fileUploadService: fileUploadService,
	}
}

func (t *TweetUsecaseImpl) CreateTweet(ctx context.Context, request TweetInfo) (*domain.Tweet, error) {
	var response *domain.Tweet

	err := t.transaction.Do(ctx, func(ctx context.Context) error {

		tweet, err := domain.NewTweet(request.UserId, request.Content, request.ReplyToId)
		if err != nil {
			return err
		}
		response, err = t.tweetRepo.Insert(ctx, tweet)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert tweet")
		}

		savePathList, err := t.fileUploadService.UploadFile(request.ImgFile)
		if err != nil {
			return apperrors.SaveLocalFileFailed.Wrap(err, "fail to save local file")
		}

		imagesModel, err := domain.NewTweetImage(tweet.ID, savePathList)
		if err != nil {
			return err
		}
		_, err = t.tweetImageRepo.Insert(ctx, imagesModel)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert tweetImage")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
