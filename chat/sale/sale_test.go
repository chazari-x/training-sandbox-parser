package sale

import (
	"github.com/chazari-x/training-sandbox-parser/model"
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	p := New()

	tests := []struct {
		name    string
		text    string
		want    *model.MessageSale
		wantErr bool
	}{
		{
			name: "",
			text: "[SALE]: {FFFFFF}\"Подписка PREMIUM 30 дней\": {ffde00}162р.{FFFFFF} | {ffde00}-46%{FFFFFF} | Прошлая цена: {ffde00}300р.",
			want: &model.MessageSale{
				Type:      model.ChatMessageTypeSale,
				Timestamp: 0, // Timestamp will be set to current time in Parse method
				Message:   "Подписка PREMIUM 30 дней",
				Price:     "162р.",
				Discount:  "-46%",
				OldPrice:  "300р.",
			},
			wantErr: false,
		},
		{
			name: "",
			text: "[SALE]: {FFFFFF}\"Приписка перед ником\": {FFDE00}Скидка: 10%",
			want: &model.MessageSale{
				Type:      model.ChatMessageTypeSale,
				Timestamp: 0, // Timestamp will be set to current time in Parse method
				Message:   "Приписка перед ником",
				Price:     "",
				Discount:  "10%",
				OldPrice:  "",
			},
		},
		{
			name: "",
			text: "[SALE]: {FFFFFF} Поменяй свой РП ник! *Бесплатно!",
			want: &model.MessageSale{
				Type:      model.ChatMessageTypeSale,
				Timestamp: 0, // Timestamp will be set to current time in Parse method
				Message:   "Поменяй свой РП ник! *Бесплатно!",
				Price:     "",
				Discount:  "",
				OldPrice:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("got:", got)
			got.Timestamp = 0 // Reset timestamp to compare only the message content
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
