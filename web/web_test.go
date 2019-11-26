package web

import (
	"bytes"
	"strings"
	"testing"
)

func TestHeaderTemplate(t *testing.T) {
	var (
		b bytes.Buffer
		s = "MyTitle"
	)

	err := header.Execute(&b, s)
	if err != nil {
		t.Errorf("failed to execute template: %v", err)
	}
	if !strings.Contains(b.String(), s) {
		t.Errorf("should contain %q got: %q", s, b.String())
	}
}

func TestTemplateAssembly(t *testing.T) {
	var (
		b  bytes.Buffer
		ti = "MyTitle"
		bd = "MyBody"
	)
	header.Execute(&b, ti)
	body.Execute(&b, bd)
	footer.Execute(&b, nil)
	o := b.String()
	if strings.Contains(o, ti) && strings.Contains(o, bd) {
		t.Errorf("unexpected output: %s", o)
	}
}
