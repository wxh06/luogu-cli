package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/urfave/cli/v2"
	"github.com/wxh06/luogu-cli"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "user",
				Aliases: []string{"u"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "style",
						Value: "notty",
					},
				},
				ArgsUsage: "uid",
				Action: func(cCtx *cli.Context) (err error) {
					data, err := luogu.Request[luogu.UserData]("GET", "https://www.luogu.com.cn/user/"+cCtx.Args().First(), nil)
					if err != nil {
						return
					}

					fmt.Println(data.CurrentData.User.Name)
					introduction, err := glamour.Render(data.CurrentData.User.Introduction, cCtx.String("style"))
					if err != nil {
						return
					}
					fmt.Println(introduction)
					return
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
