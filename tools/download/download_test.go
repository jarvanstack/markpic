package download

import "testing"

func TestDownLoadImpl_DownLoad(t *testing.T) {
	type fields struct {
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				url: `https://pic1.zhimg.com/80/v2-aaa43d382d9acc01cfe060b0228f58c4_720w.webp`,
			},
			want:    "v2-aaa43d382d9acc01cfe060b0228f58c4_720w.webp",
			wantErr: false,
		},
		{
			args: args{
				url: `https://cdn.nlark.com/yuque/0/2022/png/2397010/1668767432343-635ad8ea-c4ac-42bd-ae37-91d6bf00d541.png`,
			},
			want:    "1668767432343-635ad8ea-c4ac-42bd-ae37-91d6bf00d541.png",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DownLoadImpl{
				"./tmp",
			}
			got, err := d.DownLoad(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownLoadImpl.DownLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownLoadImpl.DownLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
