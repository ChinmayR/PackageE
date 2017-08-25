package PackageE

import "github.com/ChinmayR/PackageC"

func FuncInPackageE() string {
	return "From PackageE: " + PackageC.FuncInPackageC()
}

