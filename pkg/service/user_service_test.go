package service

import (
	"github.com/kimxuanhong/go-campaign-no-02/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"testing"
)

func TestUserServiceImpl_FindUserByEmail(t *testing.T) {

	var arr []models.PersonImpl
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345"), 10)
	arr = append(arr, *models.NewUser(1, "Hung Pham", 22, "dark nong", "hungpham", string(hashedPassword), "admin", "customer"))

	type fields struct {
		arr []models.PersonImpl
	}
	type args struct {
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.PersonImpl
	}{
		{
			name: "Success!",
			fields: fields{
				arr: arr,
			},
			args: args{
				username: "hungpham",
			},
			want: &arr[0],
		},
		{
			name: "Not found!",
			fields: fields{
				arr: arr,
			},
			args: args{
				username: "hungfarm",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserServiceImpl{
				arr: tt.fields.arr,
			}
			if got := r.FindUserByEmail(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
