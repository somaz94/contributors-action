# Inputs & Outputs

Complete reference for all action inputs and outputs.

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
| `max_contributors` | `0` | Max contributors to include (0 = all) |
| `exclude` | | Comma-separated usernames to exclude |
| `include_bots` | `false` | Include bot accounts |
| `avatar_size` | `100` | Avatar image size in pixels |
| `sort_by` | `contributions` | Sort by: `contributions` or `name` |
| `update_section` | `false` | Update section between markers in existing file |
| `section_start` | `<!-- CONTRIBUTORS-START -->` | Start marker for section update |
| `section_end` | `<!-- CONTRIBUTORS-END -->` | End marker for section update |
| `dry_run` | `false` | Preview without writing to file |

<br/>

### Input Details

<br/>

#### token

GitHub token used for API authentication. The default `GITHUB_TOKEN` is sufficient for public repositories. For private repositories, use a PAT with `repo` scope.

<br/>

#### owner / repo

By default, the action uses `GITHUB_REPOSITORY` to determine the owner and repo. Override these to fetch contributors from a different repository.

<br/>

#### format

Three output formats are supported:

- **`table`** (default) - HTML table with avatars arranged in columns
- **`list`** - Markdown list with avatars, names, and contribution counts
- **`image`** - Inline clickable avatar images

See [Output Formats](output-formats.md) for examples.

<br/>

#### exclude

Comma-separated list of GitHub usernames to exclude. Useful for filtering out bot accounts:

```yaml
exclude: 'dependabot[bot],renovate[bot],github-actions[bot]'
```

<br/>

#### update_section

When `true`, the action updates content between `section_start` and `section_end` markers in the existing `output_file`, preserving all content before and after the markers.

<br/>

## Outputs

| Output | Description |
|---|---|
| `contributors_count` | Number of contributors found |
| `output_file` | Path to the generated/updated file |
| `top_contributor` | Username of the top contributor |
