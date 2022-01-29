package entities

import (
	"reflect"
	"sync"
	"testing"

	"github.com/emirpasic/gods/trees/binaryheap"
)

func TestAuctionBids_ValidateBid(t *testing.T) {
	type fields struct {
		mu        sync.Mutex
		auctionId int
		bids      *binaryheap.Heap
	}
	type args struct {
		bid *Bid
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aucbids := &AuctionBids{
				mu:        tt.fields.mu,
				auctionId: tt.fields.auctionId,
				bids:      tt.fields.bids,
			}
			if err := aucbids.ValidateBid(tt.args.bid); (err != nil) != tt.wantErr {
				t.Errorf("AuctionBids.ValidateBid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuctionBids_GetHighestBid(t *testing.T) {
	type fields struct {
		mu        sync.Mutex
		auctionId int
		bids      *binaryheap.Heap
	}
	tests := []struct {
		name   string
		fields fields
		want   *Bid
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aucbids := &AuctionBids{
				mu:        tt.fields.mu,
				auctionId: tt.fields.auctionId,
				bids:      tt.fields.bids,
			}
			if got := aucbids.GetHighestBid(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuctionBids.GetHighestBid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bidComp(t *testing.T) {
	type args struct {
		b1 interface{}
		b2 interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bidComp(tt.args.b1, tt.args.b2); got != tt.want {
				t.Errorf("bidComp() = %v, want %v", got, tt.want)
			}
		})
	}
}
