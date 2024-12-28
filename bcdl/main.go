package main

import (
	"os"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/ncmdl/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/bcdl"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			ncmdl.AppLogger.Fatalln("程序发生了异常", r)
		}
	}()

	if err := newApp().Run(os.Args); err != nil {
		ncmdl.AppLogger.Fatalln(err)
		return
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:  "BanCamp歌曲下载工具",
		Usage: "支持下载专辑",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "打印下载信息，所有歌曲的相关信息",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "tmp",
				Aliases: []string{"t"},
				Usage:   "下载到当前目录下的tmp目录，而不是下载到当前目录",
				Value:   false,
			},
			&cli.IntFlag{
				Name:    "pool",
				Aliases: []string{"p"},
				Usage:   "使用多任务下载，默认的任务池大小",
				Value:   ncmdl.PoolSize,
			},
		},
		Action: action(),
	}

	return app
}

func action() cli.ActionFunc {
	return func(c *cli.Context) error {
		inputLinks := types.Slice[string](c.Args().Slice())
		if inputLinks.IsEmpty() {
			return bcdl.ErrInputLinksAreEmpty
		}

		// 异步任务数量
		poolSize := c.Int("pool")
		if poolSize > 0 {
			bcdl.DefaultPoolSize = poolSize
		}
		// 仅显示下载信息，不进行下载
		info := c.Bool("info")
		// 是否下载到tmp目录
		tmp := c.Bool("tmp")

		inputLinks.ForEach(func(s string, i int) {
			t, err := bcdl.NewTask(
				s,
				bcdl.OptionDryRun(info),
				bcdl.OptionTmp(tmp),
			)
			if err != nil {
				err = errors.WithMessage(err, "new task")
				ncmdl.AppLogger.Errorln(err)
				return
			}

			err = t.Download()
			if err != nil {
				err = errors.WithMessage(err, "download")
				ncmdl.AppLogger.Errorln(err)
				return
			}
		})

		return nil
	}
}
