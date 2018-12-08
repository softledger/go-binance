package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type dustServiceTestSuite struct {
	baseTestSuite
}

func TestDustConversionService(t *testing.T) {
	suite.Run(t, new(dustServiceTestSuite))
}

func (s *dustServiceTestSuite) TestListDustConversions() {
	data := []byte(`{
    "success": true, 
    "results": {
      "total": 2,
      "rows": [
        {
          "transfered_total": "0.00132256",
          "service_charge_total": "0.00002699",
          "tran_id": 4359321,
          "logs": [
            {
	            "tranId": 4359321,
	            "serviceChargeAmount": "0.000009",
	            "uid": "10000015",
	            "amount": "0.0009",
	            "operateTime": "2018-05-03 17:07:04",
	            "transferedAmount": "0.000441",
	            "fromAsset": "USDT"
            },
            {
              "tranId": 4359321,
              "serviceChargeAmount": "0.00001799",
              "uid": "10000015",
              "amount": "0.0009",
              "operateTime": "2018-05-03 17:07:04",
              "transferedAmount": "0.00088156",
              "fromAsset": "ETH"
            }
          ],
          "operate_time": "2018-05-03 17:07:04"
        }
      ]
    }
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	dustConversions, err := s.client.NewListDustConversionService().Do(newContext())
	r := s.r()
	r.NoError(err)
	s.Len(dustConversions, 1)
	e1 := &DustConversion{
		TransferedTotal:    "0.00132256",
		ServiceChargeTotal: "0.00002699",
		TranID:             4359321,
		OperateTime:        "2018-05-03 17:07:04",
		Logs: []*DustLog{{
			TranID:              4359321,
			ServiceChargeAmount: "0.000009",
			UID:                 "10000015",
			Amount:              "0.0009",
			OperateTime:         "2018-05-03 17:07:04",
			TransferedAmount:    "0.000441",
			FromAsset:           "USDT",
		}, {
			TranID:              4359321,
			ServiceChargeAmount: "0.00001799",
			UID:                 "10000015",
			Amount:              "0.0009",
			OperateTime:         "2018-05-03 17:07:04",
			TransferedAmount:    "0.00088156",
			FromAsset:           "ETH",
		}},
	}
	s.Equal(e1, dustConversions[0])
}
