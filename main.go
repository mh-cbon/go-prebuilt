package main

import (
  "fmt"
  "os"
  "io"
  "path/filepath"
  "os/exec"

  "github.com/urfave/cli"
  "github.com/mitchellh/go-homedir"
  "github.com/badgerodon/penv"
)

var VERSION = "0.0.0"

func main () {

  	app := cli.NewApp()
  	app.Name = "go-prebuilt"
  	app.Version = VERSION
  	app.Usage = "Go-prebuilt easy way to share and install go pre-built packages"
  	app.UsageText = "go-prebuilt <cmd> <options>"
  	app.Commands = []cli.Command{
  		{
  			Name:   "setup",
  			Usage:  "Setup go-prebuilt on your computer",
  			Action: add,
  			Flags: []cli.Flag{
  				cli.StringFlag{
  					Name:  "username, u",
  					Value: "",
  					Usage: "Github username",
  				},
  			},
  		},
  		{
  			Name:   "uninstall",
  			Usage:  "Uninstall go-prebuilt from your computer",
  			Action: uninstall,
  			Flags: []cli.Flag{
  				cli.StringFlag{
  					Name:  "username, u",
  					Value: "",
  					Usage: "Github username",
  				},
  			},
  		},
  		{
  			Name:   "add",
  			Usage:  "Add a pre built binary",
  			Action: add,
  			Flags: []cli.Flag{
  				cli.StringFlag{
  					Name:  "source, s",
  					Value: "",
  					Usage: "Source of the pre built package",
  				},
  			},
  		},
  		{
  			Name:   "check",
  			Usage:  "check availability of a pre built package",
  			Action: check,
  			Flags: []cli.Flag{
  				cli.StringFlag{
  					Name:  "source, s",
  					Value: "",
  					Usage: "Source of the pre built package",
  				},
  			},
  		},
  		{
  			Name:   "remove",
  			Usage:  "Remove a pre built binary",
  			Action: remove,
  			Flags: []cli.Flag{
  				cli.StringFlag{
  					Name:  "name, n",
  					Value: "",
  					Usage: "Name of the pre built package to remove",
  				},
  			},
  		},
  	}

  app.Run(os.Args)
}


func setup (c *cli.Context) error {
  isInstalled := testGoPrebuilt()
  if isInstalled {
    fmt.Println("go-prebuilt is already installed")
  } else {
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println("An error occured while looking for your home directory")
      fmt.Println(err)
      os.Exit(1)
    }
    fmt.Println(home)

    goPrebuiltPath := filepath.Join(home, "go-prebuit")

    err = penv.AppendEnv("PATH", goPrebuiltPath)
    if err != nil {
      fmt.Println("An error occured while adding go-prebuilt to your PATH variable")
      fmt.Println(err)
      os.Exit(1)
    }

    err = os.MkdirAll(goPrebuiltPath, 0644)
    if err != nil {
      fmt.Println("An error occured while creating the path "+goPrebuiltPath)
      fmt.Println(err)
      os.Exit(1)
    }

    currentBinPath := os.Args[0]
    installBinPath := filepath.Join(goPrebuiltPath, "go-prebuilt")
    err = copyFileContents(currentBinPath, installBinPath)
    if err != nil {
      fmt.Println("An error occured while copying go-prebuilt binary to the path "+goPrebuiltPath)
      fmt.Println(err)
      os.Exit(1)
    }
    err = os.Chmod(installBinPath, 0744)
    if err != nil {
      fmt.Println("An error occured while setting permissions on the path "+goPrebuiltPath)
      fmt.Println(err)
      os.Exit(1)
    }
    fmt.Println("All done!")
  }
  return nil
}

func uninstall (c *cli.Context) error {
  fmt.Println("Unhandled for now")
  os.Exit(1)
  return nil
}

func check (c *cli.Context) error {
  source := c.String("source")

  fmt.Println(source)

  return nil

}

func add (c *cli.Context) error {
  fmt.Println("Unhandled for now")
  os.Exit(1)
  return nil
}

func remove (c *cli.Context) error {
  fmt.Println("Unhandled for now")
  os.Exit(1)
  return nil
}

func testGoPrebuilt () bool {
  oCmd := exec.Command("go-prebuilt", []string{"-v"}...)
	err := oCmd.Run()
	if err != nil {
		return false
	}
  return true
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}
