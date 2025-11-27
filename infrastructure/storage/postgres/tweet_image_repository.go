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

func (tr *TweetImageRepository) Insert(ctx context.Context, models []domain.TweetImage) (int64, error) {
	var imageList []db.BulkInsertTweetImageParams
	q := tr.client.Querier(ctx)

	for _, m := range models {
		p := toTweetImageParam(m)
		imageList = append(imageList, *p)
	}
	count, err := q.BulkInsertTweetImage(ctx, imageList)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

func toTweetImageDomain(in *db.TweetImage) domain.TweetImage {
	return domain.TweetImage{
		ID:       int(in.ID),
		TweetID:  int(in.TweetID),
		ImageUrl: in.ImageUrl,
	}
}

func toTweetImageParam(model domain.TweetImage) *db.BulkInsertTweetImageParams {
	return &db.BulkInsertTweetImageParams{
		TweetID:  int32(model.TweetID),
		ImageUrl: model.ImageUrl,
	}
}
