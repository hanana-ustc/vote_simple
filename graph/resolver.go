package graph

import (
	"Tiny_Vote/db"
	"Tiny_Vote/utils"
	"context"
	"encoding/binary"
	"log"
	"sync"

	"github.com/dgraph-io/badger/v4"
)

type Resolver struct {
	mutex sync.Mutex
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Vote(ctx context.Context, usernames []string, ticket string) (bool, error) {
	if !utils.ValidateTicket(ticket) {
		return false, nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 事务+重试 or 锁 最终选加锁 事务+重试失败率还是太高了
	err := db.DB.Update(func(txn *badger.Txn) error {
		for _, username := range usernames {
			item, err := txn.Get([]byte(username))
			var count int32
			if err == badger.ErrKeyNotFound {
				count = 0 // 用户不存在，初始化计数为 0
			} else if err != nil {
				return err
			} else {
				err = item.Value(func(val []byte) error {
					count = int32(binary.BigEndian.Uint32(val))
					return nil
				})
				if err != nil {
					return err
				}
			}
			// 增加计数并保存
			count++
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, uint32(count))
			err = txn.Set([]byte(username), buf)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetTicket(ctx context.Context) (string, error) {
	return utils.CurrentTicket.ID, nil
}

func (r *queryResolver) QueryVotes(ctx context.Context, username string) (int, error) {
	var count int32
	err := db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(username))
		if err == badger.ErrKeyNotFound {
			count = 0 //用户不存在返回初始化计数
		} else if err != nil {
			return err
		} else {
			err = item.Value(func(val []byte) error {
				count = int32(binary.BigEndian.Uint32(val))
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	return int(count), nil
}
