package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"

	"strings"

	"github.com/urfave/cli"

	"./client"
	"./shared"
)

func main() {
	gitApp := cli.NewApp()
	gitApp.Name = "git"
	gitApp.Usage = "the stupid content tracker"
	gitApp.Version = "1.0.0"
	gitApp.Authors = []cli.Author{
		cli.Author{
			Name:  "Alexey Borodenko",
			Email: "kosunix@gmail.com",
		},
	}
	gitApp.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			fmt.Printf(strings.Join(c.Args(), " "))
		} else {
			fmt.Println("usage: git [--version] [--help] [-C <path>]")
			fmt.Println("           [--exec-path[=<path>]] [--html]-path]")
			fmt.Println("           [-p|--paginate|--nopager]")
			fmt.Println("           [--git-dir=<path>]")
			fmt.Println("           <command> [<args>]")
			fmt.Println()
			fmt.Println("The most commonly used git commands are:")
			fmt.Println("   add		Add file contents to the index")
			fmt.Println("   bisect	Find by binary search the change that introdued a bug")
		}
		return nil
	}
	gitApp.Commands = []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Get version of git",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "remote, r",
					Usage: "Get remote version of git",
				},
				cli.StringFlag{
					Name:  "local, l",
					Usage: "Get local version of git",
				},
			},
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)
				if name == "remote" || name == "r" {
					conn, err := net.Dial("tcp", "localhost:1234")
					if err != nil {
						log.Fatal("Connectiong:", err)
					}

					git := &client.GitClient{Client: rpc.NewClient(conn)}

					versions := make(chan shared.Version)

					go func() {
						for x := 0; x < 3; x++ {
							versions <- git.GetVersion()
							time.Sleep(1000000000)
						}
						close(versions)
					}()

					fmt.Println("Make RPC calls")

					fmt.Println("Remote versions of git is:")
					for x := range versions {
						fmt.Printf("%s\n", x.Name)
					}
				} else if name == "local" || name == "l" {
					fmt.Println("Make local call")
				} else {
					fmt.Printf("Unknown option: %s\n", name)
				}
				return nil
			},
		},
	}

	gitApp.Run(os.Args)
}
