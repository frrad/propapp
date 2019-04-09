package counties

import (
	"bufio"
	"bytes"
	"html/template"
)

type County struct {
	StateName     string `toml:"state_name"`
	StateAbbrev   string `toml:"state_abbreviation"`
	CountyName    string `toml:"county_name"`
	AppraiserSite string `toml:"appraiser_site"`
}

type Counties struct {
	Counties []County `toml:"county"`
}

func (c Counties) AsHTML() (string, error) {
	t := template.Must(template.New("").Parse(`
<table>
<tr>
<th>State Name</th>
<th>State Abbrev</th>
<th>County Name</th>
<th>Appraiser Site</th>
</tr>
{{range .Counties}}
<tr>
<td>{{.StateName}}</td>
<td>{{.StateAbbrev}}</td>
<td>{{.CountyName}}</td>
<td><a href="{{.AppraiserSite}}">link</a></td>
</tr>
{{end}}
</table>
`))

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	if err := t.Execute(writer, c); err != nil {
		return "", err
	}

	writer.Flush()
	return b.String(), nil
}
