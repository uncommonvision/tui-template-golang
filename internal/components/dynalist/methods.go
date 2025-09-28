package dynalist

import (
// "os/exec"
// @TODO: RESEARCH USING BROWSER PACKAGE
// "github.com/pkg/browser"
)

const (
	URL MetaInfo = iota
)

func (m *Model) SelectedIndex() int {
	return m.selectedIndex
}

func (m *Model) SelectedRow() Row {
	return m.rows[m.selectedIndex]
}

// func (r Row) OpenURL() {
// 	if r.meta[URL] != "" {
// 		_ = exec.Command("xdg-open", r.meta[URL]).Start()
// 	}
// }

func (r Row) GetURL() string {
	return r.meta[URL]
}
