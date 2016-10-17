package main

import (
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

func main() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"min", "50%", "75%", "max", "release"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})

	const layout = "2006-01-02"
	base := time.Date(2016, 11, 1, 0, 0, 0, 0, time.Local)
	for index := 2; index < 8; index++ {
		estimate := base.AddDate(0, index, -1)
		term := estimate.Sub(base)
		maxTerm := base.Add(time.Duration(term.Hours()*1.5) * time.Hour)
		fifty := estimate.Add(time.Duration(maxTerm.Sub(estimate).Hours()*0.4) * time.Hour)
		seventyFive := fifty.Add(time.Duration(maxTerm.Sub(estimate).Hours()*0.12) * time.Hour)
		release := maxTerm.AddDate(0, 5, 0)
		table.Append([]string{
			estimate.Format(layout),
			fifty.Format(layout),
			seventyFive.Format(layout),
			maxTerm.Format(layout),
			release.Format(layout),
		})
	}
	table.Render()
}
