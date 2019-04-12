package counties

import (
	"bufio"
	"bytes"
	"html/template"
)

// materialized
type CountyTable []CountyRow

type CountyRow struct {
	StateName     string
	StateAbbrev   string
	CountyName    string
	AppraiserSite string
}

// at rest
type UserData map[string]State

type State struct {
	FullName string   `toml:"full_name"`
	Counties []County `toml:"counties"`
}

type County struct {
	Name          string `toml:"name"`
	AppraiserSite string `toml:"appraiser_site"`
}

func (d UserData) AsTable() CountyTable {
	ans := []CountyRow{}

	for stateAbbrev, stateData := range d {
		for _, county := range stateData.Counties {
			ans = append(ans, CountyRow{
				StateName:     stateData.FullName,
				StateAbbrev:   stateAbbrev,
				CountyName:    county.Name,
				AppraiserSite: county.AppraiserSite,
			})
		}
	}

	return ans
}

func (c CountyTable) AsHTML() (string, error) {
	t := template.Must(template.New("").Parse(`
<table>
<tr>
<th>State Name</th>
<th>State Abbrev</th>
<th>County Name</th>
<th>Appraiser Site</th>
</tr>
{{range .}}
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
