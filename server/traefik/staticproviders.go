package traefik

import (
	"fmt"
	"strings"
)

func GetStaticProviders(tabStart uint8) string {
	initTabs := strings.Repeat("\t", int(tabStart))

	return fmt.Sprintf("%smanager-https-redir:\n%s\tredirectScheme:\n%s\t\tscheme: https\n%s\t\tpermanent: false\n\n",
		initTabs, initTabs, initTabs, initTabs)
}
