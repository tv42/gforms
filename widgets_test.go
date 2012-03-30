package gforms

import (
	"html/template"
	"testing"
)

func TestWidgets(t *testing.T) {
	textWidget := NewTextWidget()
	checkboxWidget := NewCheckboxWidget()

	var widgetTests = []struct {
		given, expected template.HTML
	}{
		{
			textWidget.Render(nil, ""),
			template.HTML(`<input type="text" value="" />`),
		},
		{
			textWidget.Render([]string{"name", "foo"}, ""),
			template.HTML(`<input type="text" name="foo" value="" />`),
		},
		{
			checkboxWidget.Render(nil, ""),
			template.HTML(`<input type="checkbox" value="" />`),
		},
	}

	for _, tt := range widgetTests {
		if tt.given != tt.expected {
			t.Errorf("got %v, but expected %v", tt.given, tt.expected)
		}
	}
}