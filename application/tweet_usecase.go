package application

import (
	"context"
	"mime/multipart"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"
	"twitter-clone-go/domain/service"
)

type TweetInfo struct {
	UserId    string `json:"userId"`
	Content   string `json:"content"`
	Img       string `json:"image"`
	ReplyToId string `json:"replyToId"`
}

type TweetUsecaseImpl struct {
	tweetRepo         domain.TweetRepository
	transaction       domain.Transaction
	fileUploadService service.FileUploader
}

type TweetUsecase interface {
	CreateTweet(ctx context.Context, tweetInfo TweetInfo, form *multipart.Form) (*domain.Tweet, error)
}

func NewTweetUsecase(
	tweetRepo domain.TweetRepository,
	transaction domain.Transaction,
	fileUploadService service.FileUploader,
) *TweetUsecaseImpl {
	return &TweetUsecaseImpl{
		tweetRepo:         tweetRepo,
		transaction:       transaction,
		fileUploadService: fileUploadService,
	}
}

func (t *TweetUsecaseImpl) CreateTweet(ctx context.Context, request TweetInfo, form *multipart.Form) (*domain.Tweet, error) {
	var response *domain.Tweet

	err := t.transaction.Do(ctx, func(ctx context.Context) error {

		savePath, err := t.fileUploadService.UploadFile(form)
		if err != nil {
			return apperrors.SaveLocalFileFailed.Wrap(err, "fail to save file")
		}

		tweet, err := domain.NewTweet(request.UserId, request.Content, savePath, request.ReplyToId)
		if err != nil {
			return err
		}

		response, err = t.tweetRepo.InsertTweet(ctx, tweet)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert tweet")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
