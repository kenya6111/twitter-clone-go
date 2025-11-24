package postgres

import (
	"context"
	"log"

	"twitter-clone-go/domain"
	db "twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TweetRepository struct {
	client *client
}

func NewTweetRepository(pool *pgxpool.Pool) *TweetRepository {
	return &TweetRepository{&client{pool: pool}}
}

func (tr *TweetRepository) Insert(ctx context.Context, model *domain.Tweet) (*domain.Tweet, error) {
	q := tr.client.Querier(ctx)
	tweetInfo := db.CreateTweetParams{
		UserID:    model.UserID,
		Content:   model.Content,
		ReplyToID: toPgtypeInt4(model.ReplyToID),
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
	var replyID *int
	if in.ReplyToID.Valid {
		val := int(in.ReplyToID.Int32)
		replyID = &val
	}
	return domain.Tweet{
		ID:        int(in.ID),
		UserID:    in.UserID,
		Content:   in.Content,
		ReplyToID: replyID,
	}
}

func toPgtypeText(val string) pgtype.Text {
	return pgtype.Text{
		String: val,
		Valid:  val != "",
	}
}
func toPgtypeInt4(val *int) pgtype.Int4 {
	if val == nil {
		return pgtype.Int4{}
	}
	return pgtype.Int4{
		Int32: int32(*val),
		Valid: true,
	}
}
