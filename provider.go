package PackageE

import (
	"os"

	"github.com/urfave/cli"
)

func FuncInPackageE() {
	cli.NewApp().Run(os.Args)
}
