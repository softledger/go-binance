package binance

import (
	"context"
	"encoding/json"
)

// ListWithdrawsService list withdraws
type ListDustConversionService struct {
	c         *Client
	asset     *string
	status    *int
	startTime *int64
	endTime   *int64
}

// Do send request
func (s *ListDustConversionService) Do(ctx context.Context) ([]*DustConversion, error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/userAssetDribbletLog.html",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(DustConversionResponse)
	fmt.Println(res)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res.Result.DustConversions, nil
}

// WithdrawHistoryResponse define withdraw history response
type DustConversionResponse struct {
	Result  *DustConversionResult `json:"results"`
	Success bool                  `json:"success"`
}

// WithdrawHistoryResponse define withdraw history response
type DustConversionResult struct {
	DustConversions []*DustConversion `json:"rows"`
	Total           int64             `json:"total"`
}

// Withdraw define withdraw info
type DustConversion struct {
	TransferedTotal    string     `json:"transfered_total"`
	ServiceChargeTotal string     `json:"service_charge_total"`
	TranID             int64      `json:"tran_id"`
	OperateTime        string     `json:"operate_time"`
	Logs               []*DustLog `json:"logs"`
}

type DustLog struct {
	TranID              int64  `json:"tranId"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	UID                 string `json:"uid"`
	Amount              string `json:"amount"`
	OperateTime         string `json:"operateTime"`
	TransferedAmount    string `json:"transferedAmount"`
	FromAsset           string `json:"fromAsset"`
}
