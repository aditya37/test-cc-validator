package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aditya37/test-cc-validator/mock"
	"github.com/aditya37/test-cc-validator/model"
	"github.com/aditya37/test-cc-validator/repository"
	"github.com/golang/mock/gomock"
)

type fields struct {
	repo *mock.MockCCWriterReader
}

func initTest(t *testing.T) fields {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mdb := mock.NewMockCCWriterReader(ctrl)
	return fields{
		repo: mdb,
	}
}

func TestService_GetCardNumber(t *testing.T) {
	initTest(t)
	type args struct {
		ctx   context.Context
		value string
	}
	tests := []struct {
		name    string
		mock    func(fields)
		args    args
		want    model.ResponseValidateCard
		wantErr bool
	}{
		{
			name: "Negative: without prefix 4",
			mock: func(f fields) {},
			args: args{
				ctx:   context.Background(),
				value: "31234567890123456789",
			},
			want:    model.ResponseValidateCard{},
			wantErr: true,
		},
		{
			name: "Negative: prefix contine 4 not found",
			mock: func(f fields) {
				f.repo.EXPECT().GetCardNumber(gomock.All()).Return(model.MSTCard{}, errors.New("number not found"))
			},
			args: args{
				ctx:   context.Background(),
				value: "41234567890123456789",
			},
			want:    model.ResponseValidateCard{},
			wantErr: true,
		},
		{
			name: "Positive: prefix containe 4",
			mock: func(f fields) {
				f.repo.
					EXPECT().
					GetCardNumber(gomock.All()).
					Return(model.MSTCard{
						CardNetwork: "Visa",
					}, nil)
			},
			args: args{
				ctx:   context.Background(),
				value: "41234567890123456789",
			},
			want: model.ResponseValidateCard{
				CardNetwork: "Visa",
				Number:      "41234567890123456789",
				IsValid:     true,
			},
			wantErr: false,
		},
		{
			name: "Negative: prefix containe 5 error get number",
			mock: func(f fields) {
				f.repo.EXPECT().GetCardNumber(gomock.All()).Return(model.MSTCard{}, errors.New("number not found"))
			},
			args: args{
				ctx:   context.Background(),
				value: "5233111222211112222",
			},
			want:    model.ResponseValidateCard{},
			wantErr: true,
		},
		{
			name: "Positive: prefix containe 5  get number",
			mock: func(f fields) {
				f.repo.EXPECT().GetCardNumber(gomock.All()).Return(model.MSTCard{
					CardNetwork: "MasterCard",
				}, nil)
			},
			args: args{
				ctx:   context.Background(),
				value: "5233111222211112222",
			},
			want: model.ResponseValidateCard{
				CardNetwork: "MasterCard",
				Number:      "5233111222211112222",
				IsValid:     true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := initTest(t)
			if tt.mock != nil {
				tt.mock(f)
			}
			s := NewService(f.repo)
			got, err := s.GetCardNumber(tt.args.ctx, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetCardNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCardNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_InsertCardNumber(t *testing.T) {
	initTest(t)
	type args struct {
		ctx     context.Context
		request model.RequestAddCard
	}
	tests := []struct {
		name    string
		mock    func(fields)
		args    args
		want    model.ResponseAddCard
		wantErr bool
	}{
		{
			name: "Negative: duplicate card number",
			mock: func(f fields) {
				f.repo.EXPECT().InsertCardNumber(gomock.All(), gomock.All()).Return(repository.ErrDuplicate)
			},
			args: args{
				ctx: context.Background(),
				request: model.RequestAddCard{
					Number:      "xxxx",
					CardNetwork: "VISA",
				},
			},
			want:    model.ResponseAddCard{},
			wantErr: true,
		},
		{
			name: "Positive: insert card number",
			mock: func(f fields) {
				f.repo.EXPECT().InsertCardNumber(gomock.All(), gomock.All()).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				request: model.RequestAddCard{
					Number:      "xxxx",
					CardNetwork: "VISA",
				},
			},
			want: model.ResponseAddCard{
				CardNetwork: "VISA",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := initTest(t)
			if tt.mock != nil {
				tt.mock(f)
			}
			s := NewService(f.repo)
			got, err := s.InsertCardNumber(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.InsertCardNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.InsertCardNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
