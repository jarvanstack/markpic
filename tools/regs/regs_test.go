package regs

import (
	"reflect"
	"testing"
)

func TestGetMarkdownImageUrls(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				text: `- 阅读: review 自己的文章, 就像代码一样, 你没有 review 的习惯, 但是这个习惯是至关重要的, 对比, 优化, 反馈才能进步
				- 分享: 关于一些知识, 比如说 no tech 思考方法, 可以分享给别人
				- 知识网络: 类似 notion 的知识网络, 也许有用的
				
				**主要是阅读**
				![](https://cdn.nlark.com/yuque/__puml/2aa76dcc256f08335083bd039b5b4e49.svg#lake_card_v2=eyJ0eXBlIjoicHVtbCIsImNvZGUiOiJAc3RhcnR1bWxcblxufEEgU2VjdGlvbnxcbnN0YXJ0XG46c3RlcDE7XG58I0FudGlxdWVXaGl0ZXxCIFNlY3Rpb258XG46c3RlcDI7XG46c3RlcDM7XG58QSBTZWN0aW9ufFxuOnN0ZXA0O1xufEIgU2VjdGlvbnxcbjpzdGVwNTtcbnN0b3BcblxuQGVuZHVtbCIsInVybCI6Imh0dHBzOi8vY2RuLm5sYXJrLmNvbS95dXF1ZS9fX3B1bWwvMmFhNzZkY2MyNTZmMDgzMzUwODNiZDAzOWI1YjRlNDkuc3ZnIiwiaWQiOiJtanN4OCIsIm1hcmdpbiI6eyJ0b3AiOnRydWUsImJvdHRvbSI6dHJ1ZX0sImNhcmQiOiJkaWFncmFtIn0=)`,
			},
			want: []string{
				`https://cdn.nlark.com/yuque/__puml/2aa76dcc256f08335083bd039b5b4e49.svg#lake_card_v2=eyJ0eXBlIjoicHVtbCIsImNvZGUiOiJAc3RhcnR1bWxcblxufEEgU2VjdGlvbnxcbnN0YXJ0XG46c3RlcDE7XG58I0FudGlxdWVXaGl0ZXxCIFNlY3Rpb258XG46c3RlcDI7XG46c3RlcDM7XG58QSBTZWN0aW9ufFxuOnN0ZXA0O1xufEIgU2VjdGlvbnxcbjpzdGVwNTtcbnN0b3BcblxuQGVuZHVtbCIsInVybCI6Imh0dHBzOi8vY2RuLm5sYXJrLmNvbS95dXF1ZS9fX3B1bWwvMmFhNzZkY2MyNTZmMDgzMzUwODNiZDAzOWI1YjRlNDkuc3ZnIiwiaWQiOiJtanN4OCIsIm1hcmdpbiI6eyJ0b3AiOnRydWUsImJvdHRvbSI6dHJ1ZX0sImNhcmQiOiJkaWFncmFtIn0=`,
			},
		},
		{
			args: args{
				text: `
				![](http://1.png)
				发顺丰的
				![](http://2.png)
				`,
			},
			want: []string{
				"http://1.png",
				"http://2.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRemoteImg(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarkdownImageUrls() = %v, want %v", got, tt.want)
			}
		})
	}
}
