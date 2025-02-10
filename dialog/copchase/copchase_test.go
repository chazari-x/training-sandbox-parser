package copchase

import (
	"encoding/json"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr bool
	}{
		{
			name: "1",
			text: `{FFFFFF}#Лобби  {FFFFFF}Статус  {FFFFFF}Рейтинг: 29831  {FFFFFF}Онлайн
#1  В игре {c08003}8:54{FFFFFF}  0 - 300  {a39700}4 / 8{FFFFFF}
#2     0 - 300  {FFFFFF}0 / 8{FFFFFF}
#3     0 - 300  {FFFFFF}0 / 8{FFFFFF}
#4     20 - 500  {FFFFFF}0 / 8{FFFFFF}
#5     300 - 1500  {FFFFFF}0 / 8{FFFFFF}
#6 В игре {A99306}6:24{FFFFFF}  500 - 3500  {ff0000}8 / 8{FFFFFF}
#7     1000 - 5000  {FFFFFF}0 / 8{FFFFFF}
#8     2000  {FFFFFF}0 / 8{FFFFFF}
#9     3000  {FFFFFF}0 / 8{FFFFFF}
#10     3000  {FFFFFF}0 / 8{FFFFFF}

Гараж`,
			wantErr: false,
		},
		{
			name: "2",
			text: `{FFFFFF}#Лобби  {FFFFFF}Статус  {FFFFFF}Рейтинг: 0  {FFFFFF}Онлайн
#1  В игре {9c9c16}5:04{FFFFFF}  0 - 300  {d46f00}6 / 8{FFFFFF}
#2     0 - 300  {FFFFFF}0 / 8{FFFFFF}
#3     0 - 300  {FFFFFF}0 / 8{FFFFFF}
#4     {FF0000}20{FFFFFF}  {FFFFFF}0 / 8{FFFFFF}
#5     {FF0000}300{FFFFFF}  {FFFFFF}0 / 8{FFFFFF}
#6  В игре {9c9c16}5:04{FFFFFF}  {FF0000}500{FFFFFF}  {ff0000}8 / 8{FFFFFF}
#7     {FF0000}1000{FFFFFF}  {FFFFFF}0 / 8{FFFFFF}
#8     {FF0000}2000{FFFFFF}  {FFFFFF}0 / 8{FFFFFF}
#9     {FF0000}3000{FFFFFF}  {FFFFFF}0 / 8{FFFFFF}
#10     {FF0000}3000{FFFFFF}  {FFFFFF}0 / 8{FFFFFF}

Гараж`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New()
			got, err := p.Parse(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, err := json.Marshal(got)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(string(marshal))
		})
	}
}
