package history

import (
	"reflect"
	"testing"
)

func TestBrowserHistory_Back(t *testing.T) {
	type fields struct {
		history      []string
		currentIndex int
	}
	type args struct {
		steps int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"", fields{[]string{"aaa"}, 0}, args{1}, "aaa"},
		{"", fields{[]string{"aaa", "bbb"}, 1}, args{1}, "aaa"},
		{"", fields{[]string{"aaa", "bbb", "ccc"}, 1}, args{1}, "aaa"},
		{"", fields{[]string{"aaa", "bbb", "ccc", "ddd"}, 3}, args{2}, "bbb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser := &BrowserHistory{
				history:      tt.fields.history,
				currentIndex: tt.fields.currentIndex,
			}
			if got := browser.Back(tt.args.steps); got != tt.want {
				t.Errorf("Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowserHistory_Forward(t *testing.T) {
	type fields struct {
		history      []string
		currentIndex int
	}
	type args struct {
		steps int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"", fields{[]string{"aaa"}, 0}, args{1}, "aaa"},
		{"", fields{[]string{"aaa", "bbb"}, 1}, args{1}, "bbb"},
		{"", fields{[]string{"aaa", "bbb", "ccc"}, 1}, args{1}, "ccc"},
		{"", fields{[]string{"aaa", "bbb", "ccc", "ddd"}, 0}, args{2}, "ccc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser := &BrowserHistory{
				history:      tt.fields.history,
				currentIndex: tt.fields.currentIndex,
			}
			if got := browser.Forward(tt.args.steps); got != tt.want {
				t.Errorf("Forward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowserHistory_Visit(t *testing.T) {
	type fields struct {
		history      []string
		currentIndex int
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{"", fields{[]string{"aaa"}, 0}, args{"bbb"}, fields{[]string{"aaa", "bbb"}, 1}},
		{"", fields{[]string{"aaa", "bbb"}, 1}, args{"ccc"}, fields{[]string{"aaa", "bbb", "ccc"}, 2}},
		{"", fields{[]string{"aaa", "bbb", "ccc"}, 1}, args{"ddd"}, fields{[]string{"aaa", "bbb", "ddd"}, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser := &BrowserHistory{
				history:      tt.fields.history,
				currentIndex: tt.fields.currentIndex,
			}
			browser.Visit(tt.args.url)
			want := &BrowserHistory{tt.want.history, tt.want.currentIndex}
			if !reflect.DeepEqual(*browser, *want) {
				t.Errorf("Forward() = %v, want %v", browser, want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	history := Constructor("aaa")
	if history.currentIndex != 0 || reflect.DeepEqual(history.history, [1]string{"aaa"}) {
		t.Errorf("Forward() = %v, want %v", history, "[aaa] 0")

	}
}
