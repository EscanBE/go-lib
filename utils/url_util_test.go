package utils

import "testing"

func TestExtractHostAndPort(t *testing.T) {
	//goland:noinspection HttpUrlsUsage
	tests := []struct {
		url     string
		want    string
		wantErr bool
	}{
		{
			url:     "",
			want:    "",
			wantErr: true,
		},
		{
			url:     "&&&tcp://&&&&",
			want:    "",
			wantErr: true,
		},
		{
			url:     "g:80",
			want:    "",
			wantErr: true,
		},
		{
			url:     "localhost",
			want:    "",
			wantErr: true,
		},
		{
			url:     "google.com",
			want:    "",
			wantErr: true,
		},
		{
			url:     "http://google.com",
			want:    "google.com",
			wantErr: false,
		},
		{
			url:     "https://google.com",
			want:    "google.com",
			wantErr: false,
		},
		{
			url:     "https://google.com:8080",
			want:    "google.com:8080",
			wantErr: false,
		},
		{
			url:     "https://google.com/search?q=google&rlz=AAA&oq=google&aqs=chrome..123.123&sourceid=chrome&ie=UTF-8",
			want:    "google.com",
			wantErr: false,
		},
		{
			url:     "https://google.com:8080?q=google&rlz=AAA&oq=google&aqs=chrome..123.123&sourceid=chrome&ie=UTF-8",
			want:    "google.com:8080",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		name, _ := FirstNonEmptyString(tt.want, tt.url)
		t.Run(name, func(t *testing.T) {
			got, err := ExtractHostAndPort(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractHostAndPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractHostAndPort() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractHostAndPortOrKeep(t *testing.T) {
	//goland:noinspection HttpUrlsUsage
	tests := []struct {
		url  string
		want string
	}{
		{
			url:  "",
			want: "",
		},
		{
			url:  "&&&tcp://&&&&",
			want: "&&&tcp://&&&&",
		},
		{
			url:  "g:80",
			want: "g:80",
		},
		{
			url:  "localhost",
			want: "localhost",
		},
		{
			url:  "google.com",
			want: "google.com",
		},
		{
			url:  "http://google.com",
			want: "google.com",
		},
		{
			url:  "https://google.com",
			want: "google.com",
		},
		{
			url:  "https://google.com:8080",
			want: "google.com:8080",
		},
		{
			url:  "https://google.com/search?q=google&rlz=AAA&oq=google&aqs=chrome..123.123&sourceid=chrome&ie=UTF-8",
			want: "google.com",
		},
		{
			url:  "https://google.com:8080?q=google&rlz=AAA&oq=google&aqs=chrome..123.123&sourceid=chrome&ie=UTF-8",
			want: "google.com:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := ExtractHostAndPortOrKeep(tt.url); got != tt.want {
				t.Errorf("ExtractHostAndPortOrKeep() = %v, want %v", got, tt.want)
			}
		})
	}
}
