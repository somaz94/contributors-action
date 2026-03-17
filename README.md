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

## Quick Start

```yaml
- name: Update Contributors
  uses: somaz94/contributors-action@v1
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
    output_file: CONTRIBUTORS.md
```

<br/>

## Example Output

### Table Format (default)

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/alice">
        <img src="https://avatars.githubusercontent.com/u/1?v=4" width="100" alt="alice"/>
        <br />
        <sub><b>alice</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/bob">
        <img src="https://avatars.githubusercontent.com/u/2?v=4" width="100" alt="bob"/>
        <br />
        <sub><b>bob</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/charlie">
        <img src="https://avatars.githubusercontent.com/u/3?v=4" width="100" alt="charlie"/>
        <br />
        <sub><b>charlie</b></sub>
      </a>
    </td>
  </tr>
</table>

### List Format

- [<img src="https://avatars.githubusercontent.com/u/1?v=4" width="50" alt="alice" /> alice](https://github.com/alice) (150 contributions)
- [<img src="https://avatars.githubusercontent.com/u/2?v=4" width="50" alt="bob" /> bob](https://github.com/bob) (80 contributions)

### Image Format

[<img src="https://avatars.githubusercontent.com/u/1?v=4" width="80" alt="alice" title="alice" />](https://github.com/alice)
[<img src="https://avatars.githubusercontent.com/u/2?v=4" width="80" alt="bob" title="bob" />](https://github.com/bob)
[<img src="https://avatars.githubusercontent.com/u/3?v=4" width="80" alt="charlie" title="charlie" />](https://github.com/charlie)

<br/>

## Documentation

| Guide | Description |
|---|---|
| [Usage Guide](docs/usage-guide.md) | Detailed examples, workflow patterns, section update, cross-repo, dry run |
| [Inputs & Outputs](docs/inputs-outputs.md) | Complete reference for all action inputs and outputs |
| [Output Formats](docs/output-formats.md) | Format examples with generated code and rendered output |

<br/>

## Local Development

```bash
make build    # Build the binary
make test     # Run unit tests
make cover    # Generate coverage report
make fmt      # Format code
```

<br/>

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
