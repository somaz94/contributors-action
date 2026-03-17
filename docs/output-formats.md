# Output Formats

Contributors Action supports three output formats. Each example below shows the generated Markdown/HTML and how it renders.

<br/>

## Table Format (default)

The `table` format generates an HTML table with avatars and profile links, arranged in the specified number of columns.

**Configuration:**

```yaml
format: table
columns: 4
avatar_size: 100
```

**Generated HTML:**

```html
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
    <td align="center">
      <a href="https://github.com/dave">
        <img src="https://avatars.githubusercontent.com/u/4?v=4" width="100" alt="dave"/>
        <br />
        <sub><b>dave</b></sub>
      </a>
    </td>
  </tr>
  <tr>
    <td align="center">
      <a href="https://github.com/eve">
        <img src="https://avatars.githubusercontent.com/u/5?v=4" width="100" alt="eve"/>
        <br />
        <sub><b>eve</b></sub>
      </a>
    </td>
  </tr>
</table>
```

**Rendered Output:**

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
    <td align="center">
      <a href="https://github.com/dave">
        <img src="https://avatars.githubusercontent.com/u/4?v=4" width="100" alt="dave"/>
        <br />
        <sub><b>dave</b></sub>
      </a>
    </td>
  </tr>
  <tr>
    <td align="center">
      <a href="https://github.com/eve">
        <img src="https://avatars.githubusercontent.com/u/5?v=4" width="100" alt="eve"/>
        <br />
        <sub><b>eve</b></sub>
      </a>
    </td>
  </tr>
</table>

<br/>

## List Format

The `list` format generates a Markdown list with avatar images, usernames, profile links, and contribution counts.

**Configuration:**

```yaml
format: list
avatar_size: 50
```

**Generated Markdown:**

```markdown
- [<img src="https://avatars.githubusercontent.com/u/1?v=4" width="50" alt="alice" /> alice](https://github.com/alice) (150 contributions)
- [<img src="https://avatars.githubusercontent.com/u/2?v=4" width="50" alt="bob" /> bob](https://github.com/bob) (80 contributions)
- [<img src="https://avatars.githubusercontent.com/u/3?v=4" width="50" alt="charlie" /> charlie](https://github.com/charlie) (45 contributions)
```

**Rendered Output:**

- [<img src="https://avatars.githubusercontent.com/u/1?v=4" width="50" alt="alice" /> alice](https://github.com/alice) (150 contributions)
- [<img src="https://avatars.githubusercontent.com/u/2?v=4" width="50" alt="bob" /> bob](https://github.com/bob) (80 contributions)
- [<img src="https://avatars.githubusercontent.com/u/3?v=4" width="50" alt="charlie" /> charlie](https://github.com/charlie) (45 contributions)

<br/>

## Image Format

The `image` format generates a grid of clickable avatar images with hover tooltips.

**Configuration:**

```yaml
format: image
avatar_size: 80
```

**Generated Markdown:**

```markdown
[<img src="https://avatars.githubusercontent.com/u/1?v=4" width="80" alt="alice" title="alice" />](https://github.com/alice)
[<img src="https://avatars.githubusercontent.com/u/2?v=4" width="80" alt="bob" title="bob" />](https://github.com/bob)
[<img src="https://avatars.githubusercontent.com/u/3?v=4" width="80" alt="charlie" title="charlie" />](https://github.com/charlie)
```

**Rendered Output:**

[<img src="https://avatars.githubusercontent.com/u/1?v=4" width="80" alt="alice" title="alice" />](https://github.com/alice)
[<img src="https://avatars.githubusercontent.com/u/2?v=4" width="80" alt="bob" title="bob" />](https://github.com/bob)
[<img src="https://avatars.githubusercontent.com/u/3?v=4" width="80" alt="charlie" title="charlie" />](https://github.com/charlie)

<br/>

## Section Update Example

When using `update_section: true`, the action replaces content between markers in an existing file.

**Before:**

```markdown
# My Project

## Contributors

<!-- CONTRIBUTORS-START -->
<!-- CONTRIBUTORS-END -->

## License
MIT
```

**After:**

```markdown
# My Project

## Contributors

<!-- CONTRIBUTORS-START -->
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
  </tr>
</table>
<!-- CONTRIBUTORS-END -->

## License
MIT
```

Content before and after the markers is preserved.
