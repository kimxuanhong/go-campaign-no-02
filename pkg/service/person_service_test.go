package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/service/mocks"
)

func TestPersonService_GenPerson(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := mocks.NewMockSumService(ctrl)

	type fields struct {
		sumService SumService
	}
	type args struct {
		age int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		dutuoi bool
	}{
		{
			name: "Test 01",
			fields: fields{
				sumService: mock,
			},
			args: args{
				age: 18,
			},
			want:   "Hong nho tuoi",
			dutuoi: false,
		},

		{
			name: "Test 02",
			fields: fields{
				sumService: mock,
			},
			args: args{
				age: 18,
			},
			want:   "Hong da du 18 tuoi",
			dutuoi: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.dutuoi {
				mock.EXPECT().Valid(gomock.Any()).Return(true)
			} else {
				mock.EXPECT().Valid(gomock.Any()).Return(false)
			}

			service := &PersonServiceImpl{
				sumService: tt.fields.sumService,
			}

			if got := service.GenPerson(tt.args.age); got.GetName() != tt.want {
				t.Errorf("PersonService.GenPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
