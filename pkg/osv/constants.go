package osv

type Ecosystem string

const (
	Alpine  Ecosystem = "Alpine"
	Android Ecosystem = "Android"
)

/*
"Alpine"
"Android"
"crates.io"
"Debian"
"DWF"
"GitHub Actions"
"Go"
"GSD"
"Hex"
"Linux"
"Maven"
"npm"
"NuGet"
"OSS-Fuzz"
"Packagist"
"Pub"
"PyPI"
"RubyGems"
"UVI"
*/

func (e *Ecosystem) Lookup() string {
	return ""
}
