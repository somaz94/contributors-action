package writer

import (
	"os"
	"strings"
	"testing"

	"github.com/somaz94/contributors-action/internal/config"
)

func TestWrite(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "contributors-*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	cfg := &config.Config{
		OutputFile:    tmpFile.Name(),
		UpdateSection: false,
	}

	content := "# Contributors\n\n- alice\n- bob\n"
	if err := Write(cfg, content); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != content {
		t.Errorf("expected %q, got %q", content, string(data))
	}
}

func TestWriteUpdateSection(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "readme-*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	original := `# My Project

## Contributors

<!-- CONTRIBUTORS-START -->
old content here
<!-- CONTRIBUTORS-END -->

## License
MIT
`
	if _, err := tmpFile.WriteString(original); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	cfg := &config.Config{
		OutputFile:    tmpFile.Name(),
		UpdateSection: true,
		SectionStart:  "<!-- CONTRIBUTORS-START -->",
		SectionEnd:    "<!-- CONTRIBUTORS-END -->",
	}

	newContent := "<table><tr><td>alice</td></tr></table>\n"
	if err := Write(cfg, newContent); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	result := string(data)
	if !strings.Contains(result, "# My Project") {
		t.Error("expected header preserved")
	}
	if !strings.Contains(result, "<table><tr><td>alice</td></tr></table>") {
		t.Error("expected new content inserted")
	}
	if !strings.Contains(result, "## License") {
		t.Error("expected footer preserved")
	}
	if strings.Contains(result, "old content here") {
		t.Error("expected old content removed")
	}
}

func TestWriteUpdateSectionReversedMarkers(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "readme-*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// End marker appears before start marker
	original := "<!-- END -->\nsome content\n<!-- START -->\n"
	if _, err := tmpFile.WriteString(original); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	cfg := &config.Config{
		OutputFile:    tmpFile.Name(),
		UpdateSection: true,
		SectionStart:  "<!-- START -->",
		SectionEnd:    "<!-- END -->",
	}

	err = Write(cfg, "content")
	if err == nil {
		t.Fatal("expected error for reversed markers")
	}
	if !strings.Contains(err.Error(), "end marker appears before start marker") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWriteUpdateSectionFileNotFound(t *testing.T) {
	cfg := &config.Config{
		OutputFile:    "/nonexistent/path/file.md",
		UpdateSection: true,
		SectionStart:  "<!-- START -->",
		SectionEnd:    "<!-- END -->",
	}

	err := Write(cfg, "content")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestWriteUpdateSectionMissingMarkers(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "readme-*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString("no markers here"); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	cfg := &config.Config{
		OutputFile:    tmpFile.Name(),
		UpdateSection: true,
		SectionStart:  "<!-- START -->",
		SectionEnd:    "<!-- END -->",
	}

	err = Write(cfg, "content")
	if err == nil {
		t.Fatal("expected error for missing markers")
	}
}
