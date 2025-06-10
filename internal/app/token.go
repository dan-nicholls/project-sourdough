package app

import (
	"fmt"
	"time"
)

type Token struct {
	ID string
	IssuedAt time.Time
	ExpiresAt time.Time
	Used bool
}

type TokenStore interface {
	SaveToken(token Token) error
	GetToken(id string) (Token, error)
	MarkUsed(id string) error
}

type DemoTokenStore struct {
	Store []Token
}

func NewDemoTokenStore() DemoTokenStore {
	return DemoTokenStore{
		Store: make([]Token,0),
	}
}

func (ts *DemoTokenStore) SaveToken(token Token) error {
	ts.Store = append(ts.Store, token)
	fmt.Printf("Token Added to Store: %v", token) 
	return nil
} 

func (ts *DemoTokenStore) GetToken(id string) (Token, error) {
	for _,t := range ts.Store {
		if id == t.ID {
			fmt.Printf("Token Found in Store: %v", t) 
			return t, nil
		}
	}		
	return Token{}, fmt.Errorf("No token found with ID: %v", id) 
} 
func (ts *DemoTokenStore) MarkUsed(id string) error {
	for i,t := range ts.Store {
		if id == t.ID {
			fmt.Printf("Token marked as used: %v", t) 
			token := &ts.Store[i]
			token.Used = true
			return nil
		}
	}		
	return fmt.Errorf("No token found with ID: %v", id) 
} 
