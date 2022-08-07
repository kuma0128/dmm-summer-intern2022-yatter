package object

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_Account_CreateStatusobject(t *testing.T) {
	type (
		args struct {
			content string
			account Account
		}

		want struct {
			status Status
			err    error
		}
	)

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "正しいstatus情報を生成できる",
			args: args{
				content: "content",
				account: Account{
					ID: 10,
				},
			},
			want: want{
				status: Status{
					AccountID: 10,
					Account:   Account{},
					Content:   "content",
				},
			},
		},
		{
			name: "contentが120文字より多いときはエラー",
			args: args{
				content: "usernameusernameusernameusernameusernameusernameuseusernameusernameusernameusernameusernameusernrnameusernameusernameusernameusernameusernameusernameusernameusernameusername",
				account: Account{
					ID: 10,
				},
			},
			want: want{
				err: errors.New("status content is too long"),
			},
		},
		{
			name: "passwordが0文字ときはエラー",
			args: args{
				content: "",
				account: Account{
					ID: 10,
				},
			},
			want: want{
				err: errors.New("need status content"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateStatusobject(tt.args.content, &tt.args.account)

			if tt.want.err == nil {
				// got（受け取った値）が想定どおりかみてる
				if got.Content != tt.want.status.Content {
					t.Error("invalid content")
				}
				// if got.Account != tt.want.status.Account {
				// 	t.Error("invalid account")
				// }
				if err != nil {
					t.Error("error is happened")
				}
			} else {
				if err.Error() != tt.want.err.Error() {
					t.Error("エラーが違う")
				}
			}
		})
	}
}
