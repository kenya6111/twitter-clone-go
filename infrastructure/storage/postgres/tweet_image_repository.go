package postgres

import (
	"context"
	"log"

	"twitter-clone-go/domain"
	db "twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TweetImageRepository struct {
	client *client
}

func NewTweetImageRepository(pool *pgxpool.Pool) *TweetImageRepository {
	return &TweetImageRepository{&client{pool: pool}}
}

func (tr *TweetImageRepository) Insert(ctx context.Context, models []domain.TweetImage) ([]domain.TweetImage, error) {
	var resultSet []domain.TweetImage
	q := tr.client.Querier(ctx)
	for _, model := range models {
		tweetInfo := db.CreateTweetImageParams{
			TweetID:  int32(model.TweetID),
			ImageUrl: model.ImageUrl,
		}
		image, err := q.CreateTweetImage(ctx, tweetInfo)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result := toTweetImageDomain(&image)
		resultSet = append(resultSet, result)
	}
	return resultSet, nil
}

func toTweetImageDomain(in *db.TweetImage) domain.TweetImage {
	return domain.TweetImage{
		ID:       int(in.ID),
		TweetID:  int(in.TweetID),
		ImageUrl: in.ImageUrl,
	}
}
