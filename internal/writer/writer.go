package writer

import (
	"fmt"
	"os"
	"strings"

	"github.com/somaz94/contributors-action/internal/config"
)

// Write writes the content to the output file, optionally updating a section.
func Write(cfg *config.Config, content string) error {
	if cfg.UpdateSection {
		return updateSection(cfg.OutputFile, content, cfg.SectionStart, cfg.SectionEnd)
	}
	return os.WriteFile(cfg.OutputFile, []byte(content), 0o644)
}

func updateSection(filePath, content, startMarker, endMarker string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("reading file %s: %w", filePath, err)
	}

	original := string(data)
	startIdx := strings.Index(original, startMarker)
	endIdx := strings.Index(original, endMarker)

	if startIdx == -1 || endIdx == -1 {
		return fmt.Errorf("section markers not found in %s (start: %q, end: %q)", filePath, startMarker, endMarker)
	}

	if endIdx <= startIdx {
		return fmt.Errorf("end marker appears before start marker in %s", filePath)
	}

	var sb strings.Builder
	sb.WriteString(original[:startIdx+len(startMarker)])
	sb.WriteString("\n")
	sb.WriteString(content)
	sb.WriteString(original[endIdx:])

	return os.WriteFile(filePath, []byte(sb.String()), 0o644)
}
