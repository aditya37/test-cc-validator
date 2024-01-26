package repository

import (
	"errors"
	"sync"
	"testing"

	"github.com/aditya37/test-cc-validator/model"
	"github.com/stretchr/testify/assert"
)

func Test_cardStore_GetCardNumber(t *testing.T) {

	cardFound := map[string]model.MSTCard{}
	cardFound["2024 2025 2026"] = model.MSTCard{
		Number: "2024 2025 2026",
	}
	type fields struct {
		db      map[string]model.MSTCard
		RWMutex sync.RWMutex
	}
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.MSTCard
		wantErr error
	}{
		{
			name: "Negative: get card number not found",
			fields: fields{
				db:      make(map[string]model.MSTCard),
				RWMutex: sync.RWMutex{},
			},
			args: args{
				value: "2023 2022 2021",
			},
			want:    model.MSTCard{},
			wantErr: errors.New("number not found"),
		},
		{
			name: "Positive: get card number  found",
			fields: fields{
				db:      cardFound,
				RWMutex: sync.RWMutex{},
			},
			args: args{
				value: "2024 2025 2026",
			},
			want: model.MSTCard{
				Number: "2024 2025 2026",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &cardStore{
				db:      tt.fields.db,
				RWMutex: tt.fields.RWMutex,
			}
			resp, err := cs.GetCardNumber(tt.args.value)
			if err != nil {
				assert.Error(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.want, resp)
		})
	}
}
