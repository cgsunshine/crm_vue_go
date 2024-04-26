package utils

import "testing"

func TestFileExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{`D:\gomod\gin_vue\gin-vue-admin-main\server\api\v1\crm\crm_test_.go`},
			false,
		},
		{
			"",
			args{`D:\gomod\gin_vue\gin-vue-admin-main\server\api\v1\crm\enter.go`},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExist(tt.args.path); got != tt.want {
				t.Errorf("FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
