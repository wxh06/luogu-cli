/*
Copyright © 2023 汪心禾

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package usercmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wxh06/luogu-cli/pkg/luogu"
	"github.com/wxh06/luogu-cli/pkg/markdown"
)

type uidKey struct{}

// UserCmd represents the user command
var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		uid, err := cmd.Flags().GetUint("uid")
		if err != nil {
			panic(err)
		}
		username, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}

		if username != "" && uid == 0 {
			data, err := luogu.SearchUser(username)
			if err != nil {
				return err
			}
			if strings.EqualFold(data.Users[0].Name, username) {
				uid = uint(data.Users[0].Uid)
			} else {
				return errors.New("用户未找到")
			}
		}
		cmd.SetContext(context.WithValue(cmd.Context(), uidKey{}, uid))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		uid := cmd.Context().Value(uidKey{}).(uint)
		style, err := cmd.Flags().GetString("style")
		if err != nil {
			panic(err)
		}
		infoFlag, err := cmd.Flags().GetStringSlice("info")
		if err != nil {
			panic(err)
		}

		data, err := luogu.Request[luogu.DataResponse[luogu.UserData]]("GET", fmt.Sprintf("https://www.luogu.com.cn/user/%d", uid), nil)
		if err != nil {
			return
		}

		for _, info := range infoFlag {
			switch info {
			case "uid":
				fmt.Println(data.CurrentData.User.Uid)
			case "name":
				fmt.Println(data.CurrentData.User.Name)
			case "introduction":
				intro, err := markdown.Render(data.CurrentData.User.Introduction, style)
				if err != nil {
					return err
				}
				fmt.Println(intro)
			default:
				return errors.New("未知用户信息：" + info)
			}
		}
		return
	},
}

func init() {
	UserCmd.PersistentFlags().Uint("uid", 0, "User ID")
	UserCmd.PersistentFlags().String("name", "", "Username")
	// 虽然洛谷只返回完全匹配的用户名，但还是写一下补全好了
	if err := UserCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var names []string
		data, err := luogu.SearchUser(toComplete)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}
		for _, user := range data.Users {
			if user.Name != "" {
				names = append(names, user.Name)
			}
		}
		return names, cobra.ShellCompDirectiveDefault
	}); err != nil {
		panic(err)
	}
	UserCmd.MarkFlagsMutuallyExclusive("uid", "name")
	UserCmd.MarkFlagsOneRequired("uid", "name")

	UserCmd.Flags().StringSlice("info", []string{"name"}, "")
}
