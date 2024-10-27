package utils

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

const expire = 2      //有效期默认 2 秒
const maxUsages = 200 //ticket 最大使用次数，默认200次

type Ticket struct {
	ID        string
	ExpireAt  time.Time
	MaxUsages int
	Usages    int
	mu        sync.Mutex
}

var CurrentTicket *Ticket

func GenerateTicket() {
	for {
		b := make([]byte, 16)
		rand.Read(b)
		CurrentTicket = &Ticket{
			ID:        hex.EncodeToString(b),                //ticket凭证
			ExpireAt:  time.Now().Add(expire * time.Second), //有效期
			MaxUsages: maxUsages,                            // 每个 ticket 最多使用次数
			mu:        sync.Mutex{},                         //锁，主要保证并发安全
		}
		time.Sleep(expire * time.Second)
	}
}

func ValidateTicket(ticketID string) bool {
	if CurrentTicket == nil {
		return false
	}

	CurrentTicket.mu.Lock()
	defer CurrentTicket.mu.Unlock()

	if CurrentTicket.ID != ticketID {
		return false
	}

	if time.Now().After(CurrentTicket.ExpireAt) || CurrentTicket.Usages >= CurrentTicket.MaxUsages {
		return false
	}

	CurrentTicket.Usages++
	return true
}
