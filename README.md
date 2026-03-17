# Contributors Action

[![Continuous Integration](https://github.com/somaz94/contributors-action/actions/workflows/ci.yml/badge.svg)](https://github.com/somaz94/contributors-action/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Latest Tag](https://img.shields.io/github/v/tag/somaz94/contributors-action)](https://github.com/somaz94/contributors-action/tags)
[![Top Language](https://img.shields.io/github/languages/top/somaz94/contributors-action)](https://github.com/somaz94/contributors-action)
[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-Contributors%20Action-blue?logo=github)](https://github.com/marketplace/actions/contributors-action)

A GitHub Action that generates and updates a contributors list from GitHub repository data.

<br/>

## Features

- **Multiple Formats**: Table (HTML), list (Markdown), or image grid
- **Section Update**: Update a specific section in an existing file (e.g., README.md) using markers
- **Filtering**: Exclude specific users, filter out bots
- **Sorting**: Sort by contribution count or username
- **Customizable**: Avatar size, column count, max contributors
- **Dry Run**: Preview output without writing files

<br/>

## Usage

<br/>

### Basic

```yaml
- name: Update Contributors
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    output_file: CONTRIBUTORS.md
```

<br/>

### Update Section in README

```yaml
- name: Update Contributors Section
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    output_file: README.md
    update_section: true
    section_start: '<!-- CONTRIBUTORS-START -->'
    section_end: '<!-- CONTRIBUTORS-END -->'
    columns: 6
```

<br/>

### Custom Options

```yaml
- name: Generate Contributors List
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    output_file: CONTRIBUTORS.md
    format: list
    max_contributors: 20
    exclude: 'dependabot[bot],renovate[bot]'
    sort_by: contributions
    avatar_size: 50
```

<br/>

## Inputs

| Input | Default | Description |
|---|---|---|
| `token` | `${{ github.token }}` | GitHub token for API access |
| `owner` | (current repo owner) | Repository owner |
| `repo` | (current repo name) | Repository name |
| `output_file` | `CONTRIBUTORS.md` | Output file path |
| `format` | `table` | Output format: `table`, `list`, or `image` |
| `columns` | `6` | Number of columns for table format |
| `max_contributors` | `0` | Max contributors (0 = all) |
| `exclude` | | Comma-separated usernames to exclude |
| `include_bots` | `false` | Include bot accounts |
| `avatar_size` | `100` | Avatar image size in pixels |
| `sort_by` | `contributions` | Sort by: `contributions` or `name` |
| `update_section` | `false` | Update section between markers in existing file |
| `section_start` | `<!-- CONTRIBUTORS-START -->` | Start marker for section update |
| `section_end` | `<!-- CONTRIBUTORS-END -->` | End marker for section update |
| `dry_run` | `false` | Preview without writing to file |

<br/>

## Outputs

| Output | Description |
|---|---|
| `contributors_count` | Number of contributors found |
| `output_file` | Path to the generated/updated file |
| `top_contributor` | Username of the top contributor |

<br/>

## Output Formats

<br/>

### Table (default)

Generates an HTML table with avatars and profile links, arranged in the specified number of columns.

<br/>

### List

Generates a Markdown list with avatar images, usernames, profile links, and contribution counts.

<br/>

### Image

Generates a grid of clickable avatar images.

<br/>

## Local Development

```bash
# Build
make build

# Test
make test

# Coverage
make cover

# Format
make fmt
```

<br/>

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
