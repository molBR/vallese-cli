package commands

import (
	"fmt"
	"time"
	"vallese-cli/out"
	"vallese-cli/utils"

	"github.com/eiannone/keyboard"
	"github.com/urfave/cli"
)

var watchdog = false

func print(key *keyboard.Key, watchdog *bool) {
	up := time.Now()
	fmt.Println("RECORIDNG... press CTRL + A again to stop")
	go utils.Record(watchdog)
	for *watchdog {

		down := time.Now()
		diff := down.Local().Sub(up)
		fmt.Printf("\033[2K\r%g", diff.Seconds())
		time.Sleep(300 * time.Nanosecond)
	}
	fmt.Printf("\n")
	fmt.Println("\n \nFinished recording")
}

func Listen(c *cli.Context) {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press Ctrl + Q to quit, Ctrl + A to record")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyCtrlQ {
			break
		}
		if key == keyboard.KeyCtrlA {
			watchdog = !watchdog
			if watchdog {
				go print(&key, &watchdog)
			} else {
				fmt.Println("\nSending file...")
				jsonObject := out.SendRequestFileUpload()
				srt := out.SendRequestTranscribe(jsonObject.Url)
				//time.Sleep(30 * time.Second)
				rtr := out.SendRequestTranscribeId(srt.Id)
				for rtr.Status != "completed" {
					time.Sleep(1 * time.Second)
					rtr = out.SendRequestTranscribeId(srt.Id)
				}
				fmt.Println("TEXT: ", rtr.Text)
				out.SendRequestToOpenAI(rtr.Text)

			}
		}
	}
}
