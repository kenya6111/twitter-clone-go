package postgres

import (
	"context"
	"log"

	"twitter-clone-go/domain"
	db "twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TweetRepository struct {
	client *client
}

func NewTweetRepository(pool *pgxpool.Pool) *TweetRepository {
	return &TweetRepository{&client{pool: pool}}
}

func (tr *TweetRepository) InsertTweet(ctx context.Context, model *domain.Tweet) (*domain.Tweet, error) {
	q := tr.client.Querier(ctx)
	tweetInfo := db.CreateTweetParams{
		UserID:    model.UserID,
		Content:   model.Content,
		ImgUrl:    model.ImageUrl,
		ReplyToID: model.ReplyToID,
	}
	tweet, err := q.CreateTweet(ctx, tweetInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultSet := toTweetDomain(&tweet)
	return &resultSet, nil
}

func toTweetDomain(in *db.Tweet) domain.Tweet {
	return domain.Tweet{
		ID:        in.ID,
		UserID:    in.UserID,
		Content:   in.Content.String,
		ImageUrl:  in.ImgUrl.String,
		ReplyToID: in.ReplyToID.String,
	}
}
