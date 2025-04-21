package stats

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
			text: `{80BCFF}Статистика аккаунта: {FFFFFF}Wufus #569679
{dd9a18}[PREMIUM] Подписка активна до 02.05.2025.

{80BCFF}Рейтинг Social Credits:  {FFFFFF}7168.0
{80BCFF}Предупреждения:   {FFFFFF}0
{80BCFF}Убийств/Смертей:   {FFFFFF}453/1563
{80BCFF}Рейтинг CopChase:   {FFFFFF}6825
{80BCFF}Количество BonusPoints:{FFFFFF} 28
{80BCFF}Количество времени за день: {FFFFFF}0ч. 53м

{80BCFF}Подтвержденный аккаунт: {FFFFFF}tiktok.com/@wufus33

{80BCFF}Росписи от модераторов:{FFFFFF}

05/04/2025 {80BCFF}| kentuha: {FFFFFF}айтишник
05/04/2025 {80BCFF}| LINCOLN: {FFFFFF}тикитокер
05/04/2025 {80BCFF}| AMEPUKA: {FFFFFF}Чемпион
06/04/2025 {80BCFF}| IntelCoreBot: {FFFFFF}Тики Ток
06/04/2025 {80BCFF}| orangefest: {FFFFFF}Молодой Тик - Так
06/04/2025 {80BCFF}| Hamster: {FFFFFF}Тиктакер
07/04/2025 {80BCFF}| jonathan_abrams: {FFFFFF}https://www.youtube.com/@jaxveller6277/videos
08/04/2025 {80BCFF}| .LIONHEART.: {FFFFFF}Вуф-Вуфус!
13/04/2025 {80BCFF}| czo.ooo: {FFFFFF}новичек программист

`,
			wantErr: false,
		},
		{
			name: "2",
			text: `{80BCFF}Статистика аккаунта: {FFFFFF}Wufus #569679
{dd9a18}[PREMIUM] Подписка активна до 02.05.2025.

{80BCFF}Рейтинг Social Credits:  {FFFFFF}7168.0
{80BCFF}Предупреждения:   {FFFFFF}0
{80BCFF}Убийств/Смертей:   {FFFFFF}453/1563
{80BCFF}Рейтинг CopChase:   {FFFFFF}6825
{80BCFF}Количество BonusPoints:{FFFFFF} 28
{80BCFF}Количество времени за день: {FFFFFF}0ч. 53м

{80BCFF}Подтвержденный аккаунт: {FFFFFF}tiktok.com/@wufus33

{80BCFF}Росписи от модераторов:{FFFFFF}
`,
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
