package formatter

import (
	"fmt"
	"strings"

	"github.com/somaz94/contributors-action/internal/github"
)

// Format generates a markdown string for the given contributors.
func Format(contributors []github.Contributor, format string, columns, avatarSize int) string {
	switch format {
	case "list":
		return formatList(contributors, avatarSize)
	case "image":
		return formatImage(contributors, avatarSize)
	default:
		return formatTable(contributors, columns, avatarSize)
	}
}

func formatTable(contributors []github.Contributor, columns, avatarSize int) string {
	if len(contributors) == 0 {
		return ""
	}

	var sb strings.Builder

	sb.WriteString("<table>\n")

	for i, c := range contributors {
		if i%columns == 0 {
			if i > 0 {
				sb.WriteString("  </tr>\n")
			}
			sb.WriteString("  <tr>\n")
		}
		sb.WriteString("    <td align=\"center\">\n")
		sb.WriteString(fmt.Sprintf("      <a href=\"%s\">\n", c.HTMLURL))
		sb.WriteString(fmt.Sprintf("        <img src=\"%s\" width=\"%d\" alt=\"%s\"/>\n", c.AvatarURL, avatarSize, c.Login))
		sb.WriteString("        <br />\n")
		sb.WriteString(fmt.Sprintf("        <sub><b>%s</b></sub>\n", c.Login))
		sb.WriteString("      </a>\n")
		sb.WriteString("    </td>\n")
	}

	sb.WriteString("  </tr>\n")
	sb.WriteString("</table>\n")

	return sb.String()
}

func formatList(contributors []github.Contributor, avatarSize int) string {
	var sb strings.Builder

	for _, c := range contributors {
		sb.WriteString(fmt.Sprintf("- [<img src=\"%s\" width=\"%d\" alt=\"%s\" /> %s](%s) (%d contributions)\n",
			c.AvatarURL, avatarSize, c.Login, c.Login, c.HTMLURL, c.Contributions))
	}

	return sb.String()
}

func formatImage(contributors []github.Contributor, avatarSize int) string {
	var sb strings.Builder

	for _, c := range contributors {
		sb.WriteString(fmt.Sprintf("[<img src=\"%s\" width=\"%d\" alt=\"%s\" title=\"%s\" />](%s)\n",
			c.AvatarURL, avatarSize, c.Login, c.Login, c.HTMLURL))
	}

	return sb.String()
}
