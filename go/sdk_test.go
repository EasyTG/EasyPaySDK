package sdk

import (
	"context"
	"testing"
	"time"
)

var (
	c = New("626f63fd5d6816e8087e69d2", "tIjkzXPv3z4sj1VM")
)

func TestMe(t *testing.T) {
	ctx := context.Background()
	info, err := c.Me(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("info: %+v", info)
}

func TestTransfer(t *testing.T) {
	ctx := context.Background()
	result, err := c.Transfer(ctx, &TransferParams{
		OrderID:  time.Now().String(),
		Name:     "转账",
		Amount:   100,
		ToUserID: 5045334922,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", result)
}

func TestDeduct(t *testing.T) {
	ctx := context.Background()
	result, err := c.Deduct(ctx, &DeductParams{
		OrderID:  time.Now().String(),
		UniqueID: time.Now().String(),
		Name:     "转账",
		Amount:   100,
		UsePromo: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", result)
}

func TestPayLink(t *testing.T) {
	ctx := context.Background()
	items, err := c.PayLink(ctx, &PayLinkParams{
		Items: []*PayLinkItem{
			{
				UniqueID:  "1",
				Name:      "item1",
				Amount:    100,
				AutoRenew: true,
			},
			{
				UniqueID:  "2",
				Name:      "item2",
				Amount:    200,
				AutoRenew: false,
			},
			{
				UniqueID:  "3",
				Name:      "item3",
				Amount:    500,
				AutoRenew: false,
			},
		},
		Params:    "a=b&c=d",
		ReturnURL: "https://t.me/RBQ4Bot?start=xxxxxxx",
		ExpiredAt: time.Now().AddDate(1, 0, 0).Unix(),
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range items {
		t.Logf("item: %+v", v)
	}
}

func TestTrans(t *testing.T) {
	ctx := context.Background()
	result, err := c.Trans(ctx, "tron", "12a9100a92a4a4efca1a6f173eb77898a36d15a8df7093e696d175b7dff522a9")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tx: %+v", result)
}

func TestOrderList(t *testing.T) {
	ctx := context.Background()
	result, err := c.OrderList(ctx, 1, 1000)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result {
		t.Logf("order: %+v", v)
	}
}

func TestInviteCode(t *testing.T) {
	ctx := context.Background()
	result, err := c.InviteCode(ctx, "ejbu6nfn", "812342452")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", result)
}

func TestChangeCommission(t *testing.T) {
	ctx := context.Background()
	err := c.ChangeCommission(ctx, 10, "626f881cf0b6d7625ac9c23b", "812342452")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserInfo(t *testing.T) {
	ctx := context.Background()
	result, err := c.UserInfo(ctx, "5045334922")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", result)
}

func TestFundLogs(t *testing.T) {
	ctx := context.Background()
	items, err := c.FundLogs(ctx, 1, 1000)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range items {
		t.Logf("log: %+v", v)
	}
}
