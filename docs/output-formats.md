# Output Formats

Contributors Action supports three output formats. Each example below shows the generated Markdown/HTML.

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
      <a href="https://github.com/user-a">
        <img src="https://avatars.githubusercontent.com/u/...?v=4" width="100" alt="user-a"/>
        <br />
        <sub><b>user-a</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/user-b">
        <img src="https://avatars.githubusercontent.com/u/...?v=4" width="100" alt="user-b"/>
        <br />
        <sub><b>user-b</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/user-c">
        <img src="https://avatars.githubusercontent.com/u/...?v=4" width="100" alt="user-c"/>
        <br />
        <sub><b>user-c</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/user-d">
        <img src="https://avatars.githubusercontent.com/u/...?v=4" width="100" alt="user-d"/>
        <br />
        <sub><b>user-d</b></sub>
      </a>
    </td>
  </tr>
  <tr>
    <td align="center">
      <a href="https://github.com/user-e">
        <img src="https://avatars.githubusercontent.com/u/...?v=4" width="100" alt="user-e"/>
        <br />
        <sub><b>user-e</b></sub>
      </a>
    </td>
  </tr>
</table>
```

Each cell contains the contributor's avatar image, a link to their GitHub profile, and their username.

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
- [<img src="https://avatars.githubusercontent.com/u/...?v=4" width="50" alt="user-a" /> user-a](https://github.com/user-a) (150 contributions)
- [<img src="https://avatars.githubusercontent.com/u/...?v=4" width="50" alt="user-b" /> user-b](https://github.com/user-b) (80 contributions)
- [<img src="https://avatars.githubusercontent.com/u/...?v=4" width="50" alt="user-c" /> user-c](https://github.com/user-c) (45 contributions)
```

Each item shows the avatar, username with profile link, and total contribution count.

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
[<img src="https://avatars.githubusercontent.com/u/...?v=4" width="80" alt="user-a" title="user-a" />](https://github.com/user-a)
[<img src="https://avatars.githubusercontent.com/u/...?v=4" width="80" alt="user-b" title="user-b" />](https://github.com/user-b)
[<img src="https://avatars.githubusercontent.com/u/...?v=4" width="80" alt="user-c" title="user-c" />](https://github.com/user-c)
```

Avatars are displayed inline. Hovering shows the username as a tooltip, and clicking links to the profile.

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
      <a href="https://github.com/user-a">
        <img src="https://avatars.githubusercontent.com/u/...?v=4" width="100" alt="user-a"/>
        <br />
        <sub><b>user-a</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/user-b">
        <img src="https://avatars.githubusercontent.com/u/...?v=4" width="100" alt="user-b"/>
        <br />
        <sub><b>user-b</b></sub>
      </a>
    </td>
  </tr>
</table>
<!-- CONTRIBUTORS-END -->

## License
MIT
```

Content before and after the markers is preserved.
