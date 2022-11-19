package download

import (
	"os"
	"testing"
)

func TestDownLoadImpl_DownLoad(t *testing.T) {
	var dir string = "tmp/"

	// 删除临时文件
	os.RemoveAll(dir)

	d := NewDownLoader(dir)

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
			want:    dir + "v2-aaa43d382d9acc01cfe060b0228f58c4_720w.webp",
			wantErr: false,
		},
		{
			args: args{
				url: `https://cdn.nlark.com/yuque/0/2022/png/2397010/1668767432343-635ad8ea-c4ac-42bd-ae37-91d6bf00d541.png`,
			},
			want:    dir + "1668767432343-635ad8ea-c4ac-42bd-ae37-91d6bf00d541.png",
			wantErr: false,
		},
		{
			args: args{
				url: `https://cdn.nlark.com/yuque/__puml/2aa76dcc256f08335083bd039b5b4e49.svg#lake_card_v2=eyJ0eXBlIjoicHVtbCIsImNvZGUiOiJAc3RhcnR1bWxcblxufEEgU2VjdGlvbnxcbnN0YXJ0XG46c3RlcDE7XG58I0FudGlxdWVXaGl0ZXxCIFNlY3Rpb258XG46c3RlcDI7XG46c3RlcDM7XG58QSBTZWN0aW9ufFxuOnN0ZXA0O1xufEIgU2VjdGlvbnxcbjpzdGVwNTtcbnN0b3BcblxuQGVuZHVtbCIsInVybCI6Imh0dHBzOi8vY2RuLm5sYXJrLmNvbS95dXF1ZS9fX3B1bWwvMmFhNzZkY2MyNTZmMDgzMzUwODNiZDAzOWI1YjRlNDkuc3ZnIiwiaWQiOiJtanN4OCIsIm1hcmdpbiI6eyJ0b3AiOnRydWUsImJvdHRvbSI6dHJ1ZX0sImNhcmQiOiJkaWFncmFtIn0=`,
			},
			want:    dir + "2aa76dcc256f08335083bd039b5b4e49.svg",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
