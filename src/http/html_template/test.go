package main

import "bytes"
import "fmt"
import "html/template"
import "strings"

var input = `
  {{ define "LAYOUT" }}
    <html>
      <body>
        {{ template "CONTENT" . }}
      </body>
    </html>
  {{ end }}

  {{ define "CONTENT" }}
    Unsafe content: {{ .Unsafe }}
    Newlines converted to <br/> follow:
    {{ .Normal }}
  {{ end }}

  {{ template "LAYOUT" . }}
`

var other = `
  Hello
  World
  Again
`

var other2 = `
  <script>alert("Owned!");</script>
`

func main() {

    var t, err = template.New("sample").Parse(input)
    if err != nil {
        panic(err)
    }

    var fixed = strings.Replace(other, "\n", "\n<br/>", -1)
    var model = map[string]interface{}{
        "Normal": template.HTML(fixed),
        "Unsafe": other2,
    }

    var out bytes.Buffer
    t.Execute(&out, model)

    var raw = out.String()
    fmt.Printf("%s", raw)
}
