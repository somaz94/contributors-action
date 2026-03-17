# Usage Guide

Detailed usage examples and workflow patterns for Contributors Action.

<br/>

## Basic Usage

```yaml
- name: Update Contributors
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    output_file: CONTRIBUTORS.md
```

<br/>

## Update Section in README

Instead of generating a standalone file, you can update a specific section in an existing file using HTML comment markers.

Add markers to your README.md:

```markdown
## Contributors

<!-- CONTRIBUTORS-START -->
<!-- CONTRIBUTORS-END -->
```

Then configure the action:

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

The action replaces everything between the markers while preserving the rest of the file.

<br/>

## Custom Options

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

## Cross-Repository Contributors

Fetch contributors from a different repository:

```yaml
- name: Fetch External Contributors
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    owner: another-org
    repo: another-repo
    output_file: CONTRIBUTORS.md
```

<br/>

## Automated Workflow

Run the action automatically when PRs are merged or on a schedule:

```yaml
name: Update Contributors

on:
  push:
    branches: [main]
  schedule:
    - cron: '0 0 * * 0'  # Weekly on Sunday
  workflow_dispatch:

permissions:
  contents: write

jobs:
  update-contributors:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v6

      - name: Generate Contributors
        uses: somaz94/contributors-action@v1
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          output_file: CONTRIBUTORS.md
          format: table
          columns: 6
          exclude: 'dependabot[bot],renovate[bot]'

      - name: Commit changes
        uses: somaz94/go-git-commit-action@v1
        with:
          commit_message: "docs: update CONTRIBUTORS.md"
          branch: main
          file_pattern: "CONTRIBUTORS.md"
          github_token: ${{ secrets.GITHUB_TOKEN }}
```

<br/>

## Dry Run

Preview the output without writing any files:

```yaml
- name: Preview Contributors
  id: preview
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    dry_run: true

- name: Show Results
  run: |
    echo "Found ${{ steps.preview.outputs.contributors_count }} contributors"
    echo "Top contributor: ${{ steps.preview.outputs.top_contributor }}"
```

<br/>

## Using Outputs

The action provides outputs that can be used in subsequent steps:

```yaml
- name: Generate Contributors
  id: contributors
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    output_file: CONTRIBUTORS.md

- name: Check Results
  run: |
    echo "Contributors: ${{ steps.contributors.outputs.contributors_count }}"
    echo "Output file: ${{ steps.contributors.outputs.output_file }}"
    echo "Top contributor: ${{ steps.contributors.outputs.top_contributor }}"
```
