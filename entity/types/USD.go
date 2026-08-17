package usd

import (
	"io"
	"fmt"
	"strconv"
	"github.com/shopspring/decimal"
)

// custom monetary type for handling USD currency.
type USD string

func (u *USD) Unmarshal(v interface{}) error {
	conversion, ok := v.(string)
	if !ok{
		return fmt.Errorf("Error unmarshaling bytes-like object")
	}
	s_conversion, err := strconv.Atoi(conversion)
	if err != nil{
		return fmt.Errorf("Error occurred during conversion of string to int")
	}
	f_conversion   := float64(s_conversion) / 100
	d  := decimal.NewFromFloat(f_conversion).RoundBank(2)
	*u  = USD(d.String())
	return nil
}

func (u USD) Marshal(w io.Writer) {
	d_conversion, err := strconv.Atoi(string(u))
	if err != nil{
		fmt.Errorf("Error converting string to int")
	}
	f_conversion := float64(d_conversion) * 100
	s_conversion := strconv.Itoa(int(f_conversion))
	w.Write([]byte(s_conversion))
}
