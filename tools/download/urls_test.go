package download

import "testing"

func TestTrim(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				urlStr: `https://cdn.nlark.com/yuque/0/2022/png/2397010/1668767432343-635ad8ea-c4ac-42bd-ae37-91d6bf00d541.png#averageHue=%2370673e&clientId=ue647dcef-b8ea-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=1040&id=u0b01f546&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1040&originWidth=1920&originalType=binary&ratio=1&rotation=0&showTitle=false&size=134394&status=done&style=none&taskId=u71a2fad7-b512-44b5-85a2-3f682520dca&title=&width=1920`,
			},
			want: "/80/v2-aaa43d382d9acc01cfe060b0228f58c4_720w.webp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimUrl(tt.args.urlStr); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
