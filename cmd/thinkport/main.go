package main

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"github.com/rodaine/table"
)

type members struct {
	name string
	position string
	email string
	linkedin string
}


func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Position", "Email", "LinkedIn")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	c := colly.NewCollector()
	memberCount := 0

	m := make(map[string]members)

	// Find all divs with class "lae-team-member " and print them
	c.OnHTML(".lae-team-member", func(e *colly.HTMLElement) {
		memberCount++
		
		if e.ChildText(".lae-title") != "" {
				mail := e.ChildAttr("a[href^=mailto]", "href")
			if len(mail) > 0 {
				mail = mail[7:]
			}
			// Append member to slice
			m[e.ChildText(".lae-title")] = members{
				name: e.ChildText(".lae-title"),
				position: e.ChildText(".lae-team-member-position"),
				email: mail,
				linkedin: e.ChildAttr(".lae-linkedin", "href"),
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://thinkport.digital/thinkport-cloud-experten-uber-uns/")


	// Sort map by first name
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)


	// Add members to table
	for _, k := range keys {
		tbl.AddRow(m[k].name, m[k].position, m[k].email, m[k].linkedin)
	}

	tbl.Print()
	fmt.Println("\nFound", memberCount, "team members")
}
