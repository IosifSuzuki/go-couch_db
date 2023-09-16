package client

import "testing"

func TestPath(t *testing.T) {
	t.Run("build simple url with a query parameter", func(t *testing.T) {
		path := NewPath(HttpProtocol, "https://google.com").
			AddQuery("query_text", "news").
			Build()
		if path != "https://google.com?query_text=news" {
			t.Errorf("PathBuilder produce unexpected url: %v", path)
		}
	})
	t.Run("build simple url with path segmets", func(t *testing.T) {
		path := NewPath(HttpProtocol, "https://google.com").
			AddSegment("my").
			AddSegment("fresh").
			AddSegment("news").
			Build()
		if path != "https://google.com/my/fresh/news" {
			t.Errorf("PathBuilder produce unexpected url: %v", path)
		}
	})
	t.Run("build simple url with path segmets", func(t *testing.T) {
		path := NewPath(HttpProtocol, "https://google.com").
			AddQuery("query_text", "news").
			AddSegment("my").
			AddSegment("fresh").
			AddSegment("news").
			Build()
		if path != "https://google.com/my/fresh/news?query_text=news" {
			t.Errorf("PathBuilder produce unexpected url: %v", path)
		}
	})
}
