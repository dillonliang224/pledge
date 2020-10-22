package service

import (
	"context"
	"errors"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"

	pb "git.dillonliang.cn/micro-svc/pledge/src/base/book/api"
)

func (s *Service) FindById(ctx context.Context, id string) (resp *pb.BookResp, err error) {
	bookId, ok := mongodb.ConvertToObjectID(id)
	if !ok {
		return nil, errors.New("bookId error")
	}

	book, err := s.dao.FindBookInfo(ctx, bookId)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, errors.New("NO_RECORD")
	}

	return &pb.BookResp{
		Book: &pb.Book{
			Id:            book.ID.String(),
			Title:         book.Title,
			Author:        book.Author,
			Cover:         book.Cover,
			ShortIntro:    book.ShortIntro,
			LongIntro:     book.LongIntro,
			LastChapter:   *book.LastChapter,
			ChaptersCount: book.ChaptersCount,
			IsSerial:      book.IsSerial,
			Updated:       0,
		},
	}, nil
}
