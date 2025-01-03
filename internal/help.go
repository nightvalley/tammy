package help

import (
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/charmbracelet/glamour"
)

func ShowHelpMessage() {
	helpmessaage := `
# Usage
Display information about files in the current directory:
tammy

- tammy -f:
  + Change output format. Available forms: table, list, total, tree (default - table).
- tammy -h:
  + Show hidden files.
- tammy -s:
  + Show file size.
- tammy -p:
  + Specify the path to the directory in which to count lines. It is not necessary to specify the path. The path can also be specified at the very end: tammy -f list -s -h ~/Documents.
- tammy -ft:
  + Count lines only in files with a certain extension. Example: tammy -ft md, or tammy -ft .md.
- tammy -t:
  + Show execution time.
- tammy -version
  + Check for updates.

# Configuring
The utility is configured using environment variables. Available variables:
- DEFAULT_FORM
  + Available values:
    + export DEFAULT_FORM="table" - default
    + export DEFAULT_FORM="list"
    + export DEFAULT_FORM="tree"
    + export DEFAULT_FORM="total"
- ALLWAYS_DISPLAY_SIZE
  + Available values:
    + export ALLWAYS_DISPLAY_SIZE="false" - default
    + export ALLWAYS_DISPLAY_SIZE="true"
- ALLWAYS_SHOW_HIDDEN_FILES
  + Available values:
    + export ALLWAYS_SHOW_HIDDEN_FILES="false" - default
    + export ALLWAYS_SHOW_HIDDEN_FILES="true"
- LIST_ENUMERATOR
  + Available values:
    + export LIST_ENUMERATOR="roman" - default
    + export LIST_ENUMERATOR="arabic"
    + export LIST_ENUMERATOR="dash"
    + export LIST_ENUMERATOR="alphabet"
    + export LIST_ENUMERATOR="bullet"
    + export LIST_ENUMERATOR="asterisk"
- TREE_ENUMERATOR
  + Available values:
    + export TREE_ENUMERATOR="rounded" - default
    + export TREE_ENUMERATOR="default_enumerator"
    + export TREE_ENUMERATOR="default_indenter"`

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)

	out, err := r.Render(helpmessaage)
	if err != nil {
		log.Error(err)
	}

	fmt.Print(out)
}
