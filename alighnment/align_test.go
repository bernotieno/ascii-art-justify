package align

import "testing"

func TestAlignJustify(t *testing.T) {
	type args struct {
		text   string
		banner string
		width  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Single word",
			args: args{
				text:   "Hello",
				banner: "standard",
				width:  50,
			},
		},
		{
			name: "Multiple words",
			args: args{
				text:   "Hello World",
				banner: "standard",
				width:  80,
			},
		},
		{
			name: "Long text",
			args: args{
				text:   "The quick brown fox jumps over the lazy dog",
				banner: "standard",
				width:  100,
			},
		},
		{
			name: "Text too long for width",
			args: args{
				text:   "This text is too long to fit in the specified width",
				banner: "standard",
				width:  20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AlignJustify(tt.args.text, tt.args.banner, tt.args.width)
		})
	}
}

func TestAlignRight(t *testing.T) {
	type args struct {
		text   string
		banner string
		width  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Short text",
			args: args{
				text:   "Hello",
				banner: "standard",
				width:  50,
			},
		},
		{
			name: "Medium text",
			args: args{
				text:   "Hello World",
				banner: "standard",
				width:  80,
			},
		},
		{
			name: "Long text",
			args: args{
				text:   "The quick brown fox jumps over the lazy dog",
				banner: "standard",
				width:  100,
			},
		},
		{
			name: "Text too long for width",
			args: args{
				text:   "This text is too long to fit",
				banner: "standard",
				width:  20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AlignRight(tt.args.text, tt.args.banner, tt.args.width)
		})
	}
}

func TestAlignLeft(t *testing.T) {
	type args struct {
		text   string
		banner string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Short text",
			args: args{
				text:   "Hello",
				banner: "standard",
			},
		},
		{
			name: "Medium text",
			args: args{
				text:   "Hello World",
				banner: "standard",
			},
		},
		{
			name: "Long text",
			args: args{
				text:   "The quick brown fox jumps over the lazy dog",
				banner: "standard",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AlignLeft(tt.args.text, tt.args.banner)
		})
	}
}

func TestAlignCenter(t *testing.T) {
	type args struct {
		text   string
		banner string
		width  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Short text",
			args: args{
				text:   "Hello",
				banner: "standard",
				width:  50,
			},
		},
		{
			name: "Medium text",
			args: args{
				text:   "Hello World",
				banner: "standard",
				width:  80,
			},
		},
		{
			name: "Long text",
			args: args{
				text:   "The quick brown fox jumps over the lazy dog",
				banner: "standard",
				width:  100,
			},
		},
		{
			name: "Text too long for width",
			args: args{
				text:   "This text is too long to fit",
				banner: "standard",
				width:  20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AlignCenter(tt.args.text, tt.args.banner, tt.args.width)
		})
	}
}