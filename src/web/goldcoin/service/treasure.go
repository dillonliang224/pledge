package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/model"
)

func (s *Service) GetTreasureRecords(ctx context.Context, activityId string, productId string, period string) ([]*model.TreasureRecord, error) {
	aId, _ := primitive.ObjectIDFromHex(activityId)
	pId, _ := primitive.ObjectIDFromHex(productId)
	records, err := s.dao.FindTreasureRecords(ctx, aId, pId, period)
	if err != nil {
		return nil, err
	}

	return records, nil
}
