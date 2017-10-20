// Package annotate is a simple try to simulate the annotation parser.
// its my try to be quick. need to rewrite this using proper syntax parser
// like text/scanner
// TODO : rewrite it :)
package annotate

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	group = regexp.MustCompile(`(?ms)^\s*(/[*]|[/]{2}|)\s*@(\S+)\s*{([^}]*)}\s*$`)
	line  = regexp.MustCompile(`\s*([/]{2}\s*|)((\S+)\s*=(.*)|.*)`)
)

type Single struct {
	Items map[string]string
	Name  string
}

type Group []Single

func loadFromGroup(g string) (Single, error) {
	res := Single{Items: make(map[string]string)}
	lne := line.FindAllStringSubmatch(g, -1)
	for i := range lne {
		if len(lne[i]) == 5 {
			l := strings.Trim(lne[i][2], " /\n\t")
			k := strings.Trim(lne[i][3], " \n\t")
			v := strings.Trim(lne[i][4], " \n\t")

			if k != "" {
				res.Items[k] = v
			} else {
				if l != "" {
					return Single{}, fmt.Errorf("invalid line '%s'", l)
				}
			}
		}
	}

	return res, nil
}

// LoadFromText is a function to load the annotation from comment
func LoadFromText(c string) (Group, error) {
	// First find groups
	var res Group
	grps := group.FindAllStringSubmatch(c, -1)

	for i := range grps {
		if len(grps[i]) == 4 {
			a, err := loadFromGroup(grps[i][3])
			if err != nil {
				return nil, err
			}
			a.Name = grps[i][2]
			res = append(res, a)
		}
	}

	return res, nil
}
