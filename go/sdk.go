package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// https://pay.easypaybot.com/api/merchant/<token>/METHOD_NAME
type Client struct {
	Client *http.Client
	token  string
}

func New(appID, key string) *Client {
	return &Client{token: appID + ":" + key}
}

type MerchantInfo struct {
	AppID     string `json:"app_id"`
	AdminID   int64  `json:"admin_id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Credit    int64  `json:"credit"`
	WebHook   string `json:"webhook"`
	Withdraw  string `json:"withdraw"`
	Status    string `json:"status"`
	IsShow    bool   `json:"is_show"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Balance   int64  `json:"balance"`
	Freeze    int64  `json:"freeze"`
	KYCLevel  int64  `json:"kyc_level"`
}

type PayLinkItem struct {
	UniqueID  string `json:"unique_id"`
	Name      string `json:"name"`
	Amount    int64  `json:"amount"`
	AutoRenew bool   `json:"auto_renew"`

	// 以下字段由服务器端填充，请勿设置。
	Link    string `json:"link"`
	Command string `json:"command"`
	BotLink string `json:"bot_link"`
}

type PayLinkParams struct {
	Items     []*PayLinkItem `json:"items"`
	Params    string         `json:"params"`
	ReturnURL string         `json:"return_url"`
	ExpiredAt int64          `json:"expired_at"`
}

type OrderInfo struct {
	OrderID         string `json:"order_id"`
	AppID           string `json:"app_id"`
	UniqueID        string `json:"unique_id"`
	ParentID        string `json:"parent_id"`
	Name            string `json:"name"`
	Amount          int64  `json:"amount"`
	OriginalAmount  int64  `json:"original_amount"`
	Params          string `json:"params"`
	ReturnURL       string `json:"return_url"`
	PayUserID       int64  `json:"pay_user_id"`
	ToUserID        int64  `json:"to_user_id"`
	AutoRenew       bool   `json:"auto_renew"`
	PromoCodeID     string `json:"promo_code_id"`
	PromoCode       string `json:"promo_code"`
	PromoType       string `json:"promo_type"`
	DiscountRate    int64  `json:"discount_rate"`
	DiscountValue   int64  `json:"discount_value"`
	CommissionRate  int64  `json:"commission_rate"`
	CommissionValue int64  `json:"commission_value"`
	Status          string `json:"status"`
	CreatedAt       string `json:"created_at"`
	ExpiredAt       string `json:"expired_at"`
}

type TransferParams struct {
	OrderID  string `json:"order_id"`
	Name     string `json:"name"`
	Amount   int64  `json:"amount"`
	ToUserID int64  `json:"to_user_id"`
}
type TransferReceipt struct {
	OrderID   string `json:"order_id"`
	AppID     string `json:"app_id"`
	ParentID  string `json:"parent_id"`
	Name      string `json:"name"`
	Amount    int64  `json:"amount"`
	PayUserID int64  `json:"pay_user_id"`
	ToUserID  int64  `json:"to_user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type DeductParams struct {
	OrderID  string `json:"order_id"`
	UniqueID string `json:"unique_id"`
	Name     string `json:"name"`
	Amount   int64  `json:"amount"`
	UsePromo bool   `json:"use_promo"`
}
type DeductReceipt struct {
	OrderID   string `json:"order_id"`
	AppID     string `json:"app_id"`
	ParentID  string `json:"parent_id"`
	Name      string `json:"name"`
	Amount    int64  `json:"amount"`
	Params    string `json:"params"`
	ReturnURL string `json:"return_url"`
	PayUserID int64  `json:"pay_user_id"`
	ToUserID  int64  `json:"to_user_id"`
	AutoRenew bool   `json:"auto_renew"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func (c *Client) url(method string, optQueryString ...string) string {
	extra := ""
	if len(optQueryString) > 0 {
		extra = "?" + optQueryString[0]
	}
	return "https://pay.easypaybot.com/api/merchant/" + c.token + "/" + method + extra
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "EasyPayClient/go1.0")
	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(req)
}

func (c *Client) Me(ctx context.Context) (*MerchantInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url("me"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code   int           `json:"code"`
		Status string        `json:"status"`
		Result *MerchantInfo `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s)", reply.Code, reply.Status)
	}
	return reply.Result, nil
}

func (c *Client) PayLink(ctx context.Context, params *PayLinkParams) ([]*PayLinkItem, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.url("paylink"), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Items []*PayLinkItem `json:"items"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("!200")
	}
	return reply.Items, nil
}

func (c *Client) Order(ctx context.Context, id string) (*OrderInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url("order", "order_id="+id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code   int        `json:"code"`
		Status string     `json:"status"`
		Result *OrderInfo `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s)", reply.Code, reply.Status)
	}
	return reply.Result, nil
}

func (c *Client) Transfer(ctx context.Context, params *TransferParams) (*TransferReceipt, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.url("transfer"), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code   int              `json:"code"`
		Status string           `json:"status"`
		Result *TransferReceipt `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s)", reply.Code, reply.Status)
	}
	return reply.Result, nil
}

func (c *Client) Deduct(ctx context.Context, params *DeductParams) (*DeductReceipt, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.url("deduct"), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code   int            `json:"code"`
		Status string         `json:"status"`
		Result *DeductReceipt `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s)", reply.Code, reply.Status)
	}
	return reply.Result, nil
}
