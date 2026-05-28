package formatter

import (
	"fmt"
	"strings"

	"github.com/somaz94/contributors-action/internal/github"
)

// Format names accepted by Format.
const (
	FormatTable = "table"
	FormatList  = "list"
	FormatImage = "image"
)

// Format generates a markdown string for the given contributors.
func Format(contributors []github.Contributor, format string, columns, avatarSize int) string {
	switch format {
	case FormatList:
		return formatList(contributors, avatarSize)
	case FormatImage:
		return formatImage(contributors, avatarSize)
	default:
		return formatTable(contributors, columns, avatarSize)
	}
}

// avatarImg renders a contributor's avatar <img> tag. extraAttrs is a leading-space
// string of additional attributes (e.g., ` title="alice"`) or "" for none.
func avatarImg(c github.Contributor, size int, extraAttrs string) string {
	return fmt.Sprintf(`<img src="%s" width="%d" alt="%s"%s/>`,
		c.AvatarURL, size, c.Login, extraAttrs)
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
		fmt.Fprintf(&sb, "      <a href=\"%s\">\n", c.HTMLURL)
		fmt.Fprintf(&sb, "        %s\n", avatarImg(c, avatarSize, ""))
		sb.WriteString("        <br />\n")
		fmt.Fprintf(&sb, "        <sub><b>%s</b></sub>\n", c.Login)
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
		fmt.Fprintf(&sb, "- [%s %s](%s) (%d contributions)\n",
			avatarImg(c, avatarSize, " "), c.Login, c.HTMLURL, c.Contributions)
	}

	return sb.String()
}

func formatImage(contributors []github.Contributor, avatarSize int) string {
	var sb strings.Builder

	for _, c := range contributors {
		titleAttr := fmt.Sprintf(` title="%s" `, c.Login)
		fmt.Fprintf(&sb, "[%s](%s)\n", avatarImg(c, avatarSize, titleAttr), c.HTMLURL)
	}

	return sb.String()
}
