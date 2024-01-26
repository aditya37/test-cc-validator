package repository

import (
	"errors"
	"sync"
	"testing"

	"github.com/aditya37/test-cc-validator/model"
	"github.com/stretchr/testify/assert"
)

func Test_cardStore_InsertCardNumber(t *testing.T) {
	dumyData := map[string]model.MSTCard{}
	dumyData["111 11 11"] = model.MSTCard{
		Number: "111 11 11",
	}

	// dumyDataPositive
	dumyDataPositive := map[string]model.MSTCard{}
	dumyDataPositive["2023 2022 2021"] = model.MSTCard{
		Number: "2023 2022 2021",
	}
	type fields struct {
		db      map[string]model.MSTCard
		RWMutex sync.RWMutex
	}
	type args struct {
		value   string
		network string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Negative: duplicate card number",
			fields: fields{
				db:      dumyData,
				RWMutex: sync.RWMutex{},
			},
			args: args{
				value: "111 11 11",
			},
			wantErr: errors.New("duplicate number"),
		},
		{
			name: "Positive: ",
			fields: fields{
				db:      dumyDataPositive,
				RWMutex: sync.RWMutex{},
			},
			args:    args{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &cardStore{
				db:      tt.fields.db,
				RWMutex: tt.fields.RWMutex,
			}
			err := cs.InsertCardNumber(tt.args.value, tt.args.network)
			if err != nil {
				assert.Error(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
