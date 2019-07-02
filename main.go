package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	//đặt tên cho app
	app.Name = "My Command App"
	//đặt phiên bản cho app
	app.Version = "1.0.0"
	//đặt mục đích sử dụng app
	app.Usage = "My Commandline App"

	//tạo những command con cho app
	//khi call sẽ có dạng mycommand add
	app.Commands = []cli.Command{
		{
			//khởi tạo subcommand
			Name: "add",
			//tạo alias cho subcommand
			//khi call sẽ có thể gọi mycommand add hoặc mycommand a
			Aliases: []string{"a"},
			//định nghĩa action cho command
			//có sử dụng thêm Flag
			Action: func(c *cli.Context) error {
				if c.String("lang") == "english" {
					fmt.Println("Call add command")
				} else {
					fmt.Println("Gọi command add")
				}

				return nil
			},
			//định nghĩa Category cho subcommand
			Category: "Templates Action",
			//tạo các Flag để sử dụng cho command
			Flags: []cli.Flag{
				cli.StringFlag{
					//định nghĩa flag
					//khi call sẽ có dạng mycommand a --lang vietnam hoặc mycommand a --l vietnam
					Name: "lang, l",
					//định nghĩa giá trị mặc định
					Value: "english",
					//hiển thị mục đích của flag
					Usage: "Show content by language specific",
				},
			},
		},
		{
			Name:    "del",
			Aliases: []string{"d"},

			Action: func(c *cli.Context) error {
				fmt.Println("Call del command")
				return nil
			},
			Category: "Templates Action",
		},
	}

	//thực thi app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
