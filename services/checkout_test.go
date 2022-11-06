package services

import (
	"testing"

	"github.com/adiwenak/seek-checkout-system/entity"
	"github.com/shopspring/decimal"
)

// Ideally we need to mock repository to avoid call to data store.
// but since this is a simple project repository already mock with static data @ repository/stub.go
func TestCalculatePrice(t *testing.T) {
	type args struct {
		customerid entity.CustomerId
		items      []CheckoutItem
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No items",
			args: args{
				customerid: "1",
				items:      []CheckoutItem{},
			},
			want: decimal.NewFromFloat(0.0).String(),
		},
		{
			name: "Item does not exist",
			args: args{
				customerid: "1",
				items: []CheckoutItem{
					{
						Id:       99,
						Quantity: 1,
					},
				},
			},
			want: decimal.NewFromFloat(0.0).String(),
		},
		{
			name: "No special rule",
			args: args{
				customerid: "1",
				items: []CheckoutItem{
					{
						Id:       1,
						Quantity: 1,
					},
					{
						Id:       2,
						Quantity: 1,
					},
					{
						Id:       3,
						Quantity: 1,
					},
				},
			},
			want: decimal.NewFromFloat(987.97).String(),
		},
		{
			name: "Discounted price",
			args: args{
				customerid: "3",
				items: []CheckoutItem{
					{
						Id:       2,
						Quantity: 1,
					},
				},
			},
			want: decimal.NewFromFloat(299.99).String(),
		},
		{
			name: "Combination of special package price and regular",
			args: args{
				customerid: "2",
				items: []CheckoutItem{
					{
						Id:       1,
						Quantity: 5,
					},
				},
			},
			want: decimal.NewFromFloat(1079.96).String(),
		},
		{
			name: "Combination of special package price and discounted price",
			args: args{
				customerid: "4",
				items: []CheckoutItem{
					{
						Id:       2,
						Quantity: 5,
					},
					{
						Id:       3,
						Quantity: 1,
					},
				},
			},
			want: decimal.NewFromFloat(1681.95).String(),
		},
	}

	service := NewCheckoutService()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.CalculateTotalPrice(tt.args.customerid, tt.args.items); got != tt.want {
				t.Errorf("CalculatePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
