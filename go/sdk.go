package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var debug = os.Getenv("EASYPAY_DEBUG") != ""

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

type TXInfo struct {
	TgUserID    int64  `json:"tg_user_id"`
	BlockNumber int64  `json:"block_number"`
	TXID        string `json:"txid"`
	Type        string `json:"type"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       int64  `json:"value"`
	CreatedAt   string `json:"created_at"`
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

type PromotionInfo struct {
	PromotionID       string  `json:"promotion_id"`
	InviteCodeID      string  `json:"invite_code_id"`
	PromoType         string  `json:"promo_type"`
	Code              string  `json:"code"`
	PromoDiscountRate int64   `json:"promo_discount_rate"`
	CommissionRate    float64 `json:"commission_rate"`
	DiscountRate      float64 `json:"discount_rate"`
	Link              string  `json:"link"`
}

type UserInfo struct {
	ID        string `json:"id"`
	TgUserID  int64  `json:"tg_user_id"`
	IsBot     bool   `json:"is_bot"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
	Language  string `json:"language"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type FundLog struct {
	ID         string  `json:"id"`
	OrderID    string  `json:"order_id"`
	Type       string  `json:"type"`
	TgUserID   int64   `json:"tg_user_id"`
	Amount     int64   `json:"amount"`
	RealAmount float64 `json:"real_amount"`
	Remark     string  `json:"remark"`
	Related    struct {
		ID        string `json:"id"`
		OrderID   string `json:"order_id"`
		Type      string `json:"type"`
		TgUserID  int64  `json:"tg_user_id"`
		Amount    int64  `json:"amount"`
		Balance   int64  `json:"balance"`
		Remark    string `json:"remark"`
		CreatedAt string `json:"created_at"`
	} `json:"related"`
	CreatedAt string `json:"created_at"`
}

func (c *Client) url(method string, optQueryString ...string) string {
	extra := ""
	if len(optQueryString) > 0 {
		extra = "?" + optQueryString[0]
	}
	if debug {
		println("WARNING: DEBUG SERVER USED!!!")
		return "https://pay.test.tgbot.link/api/merchant/" + c.token + "/" + method + extra
	}
	return "https://pay.easypaybot.com/api/merchant/" + c.token + "/" + method + extra
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "EasyPaySDK/go1.0")
	req.Header.Add("Content-Type", "application/json")
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
		Result struct {
			Items []*PayLinkItem `json:"items"`
		} `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("!200")
	}
	return reply.Result.Items, nil
}

func (c *Client) Trans(ctx context.Context, typ, txid string) (*TXInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url("trans", "type="+typ+"&txid="+txid), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code    int     `json:"code"`
		Status  string  `json:"status"`
		Message string  `json:"message"`
		Result  *TXInfo `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("!200")
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s:%s)", reply.Code, reply.Status, reply.Message)
	}
	return reply.Result, nil
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

func (c *Client) OrderList(ctx context.Context, page, pageSize int) ([]*OrderInfo, error) {
	queryString := "page=" + strconv.Itoa(page) + "&page_size=" + strconv.Itoa(pageSize)
	req, err := http.NewRequestWithContext(ctx, "GET", c.url("orderlist", queryString), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Result struct {
			CurrPage int          `json:"currpage"`
			Items    []*OrderInfo `json:"items"`
			PageSize int          `json:"page_size"`
			Total    int          `json:"total"`
		} `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s)", reply.Code, reply.Status)
	}
	return reply.Result.Items, nil
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

func (c *Client) InviteCode(ctx context.Context, code, tgUserID string) (*PromotionInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url("invitecode", "code="+code+"&tg_user_id="+tgUserID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code    int            `json:"code"`
		Status  string         `json:"status"`
		Message string         `json:"message"`
		Result  *PromotionInfo `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s:%s)", reply.Code, reply.Status, reply.Message)
	}
	return reply.Result, nil
}

func (c *Client) ChangeCommission(ctx context.Context, commission int, inviteCodeID, tgUserID string) error {
	queryString := url.Values{
		"commission":     {strconv.Itoa(commission)},
		"invite_code_id": {inviteCodeID},
		"tg_user_id":     {tgUserID},
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.url("change_commission", queryString.Encode()), nil)
	if err != nil {
		return err
	}
	resp, err := c.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var reply struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return fmt.Errorf("#%d(%s:%s)", reply.Code, reply.Status, reply.Message)
	}
	return nil
}

func (c *Client) UserInfo(ctx context.Context, tgUserID string) (*UserInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url("userinfo", "tg_user_id="+tgUserID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  struct {
			Data *UserInfo `json:"data"`
		} `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s:%s)", reply.Code, reply.Status, reply.Message)
	}
	return reply.Result.Data, nil
}

func (c *Client) FundLogs(ctx context.Context, page, pageSize int) ([]*FundLog, error) {
	queryString := "page=" + strconv.Itoa(page) + "&page_size=" + strconv.Itoa(pageSize)
	req, err := http.NewRequestWithContext(ctx, "GET", c.url("fundlogs", queryString), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var reply struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  struct {
			CurrPage int        `json:"currpage"`
			Items    []*FundLog `json:"items"`
			PageSize int        `json:"page_size"`
			Total    int        `json:"total"`
		} `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 200 || reply.Status != "ok" {
		return nil, fmt.Errorf("#%d(%s:%s)", reply.Code, reply.Status, reply.Message)
	}
	return reply.Result.Items, nil
}
