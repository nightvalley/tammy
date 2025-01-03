package vars

import "os"

func ExportVariables() {
	// Defaults
	if os.Getenv("DEFAULT_FORM") == "" {
		os.Setenv("DEFAULT_FORM", "table")
	}
	if os.Getenv("ALLWAYS_DISPLAY_SIZE") == "" {
		os.Setenv("ALLWAYS_DISPLAY_SIZE", "false")
	}
	if os.Getenv("ALLWAYS_SHOW_HIDDEN_FILES") == "" {
		os.Setenv("ALLWAYS_SHOW_HIDDEN_FILES", "false")
	}

	// Form: List
	if os.Getenv("LIST_ENUMERATOR") == "" {
		os.Setenv("LIST_ENUMERATOR", "list.Roman")
	}

	// Form: Tree
	if os.Getenv("TREE_ENUMERATOR") == "" {
		os.Setenv("TREE_ENUMERATOR", "tree.RoundedEnumerator")
	}
}

func GetEnv() map[string]string {
	variables := make(map[string]string)

	variables["defaultForm"] = os.Getenv("DEFAULT_FORM")
	variables["allwaysDisplaySize"] = os.Getenv("ALLWAYS_DISPLAY_SIZE")
	variables["allwaysShowHiddenFiles"] = os.Getenv("ALLWAYS_SHOW_HIDDEN_FILES")
	variables["listEnumerator"] = os.Getenv("LIST_ENUMERATOR")
	variables["treeEnumerator"] = os.Getenv("TREE_ENUMERATOR")

	return variables
}
