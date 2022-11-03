package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/go-vgo/robotgo"
	"github.com/liuzl/gocc"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() {
	onExit := func() {
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}
	go func() {
		mainthread.Init(fn)
	}()
	systray.Run(onReady, onExit)

}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)

		mQuit := systray.AddMenuItem("退出", "Quit the whole app")

		// Sets the icon of a menu item. Only available on Mac.
		mQuit.SetIcon(icon.Data)

		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit now...")
				return
			}
		}
	}()
}
func fn() {
	err := listenHotkey(hotkey.KeyS, hotkey.ModCtrl, hotkey.ModShift)
	if err != nil {
		log.Println(err)
	} else {
		fn()
	}
}

func listenHotkey(key hotkey.Key, mods ...hotkey.Modifier) (err error) {
	ms := []hotkey.Modifier{}
	ms = append(ms, mods...)
	hk := hotkey.New(ms, key)

	err = hk.Register()
	if err != nil {
		return
	}

	// Blocks until the hokey is triggered.
	<-hk.Keydown()
	robotgo.KeyTap("a", "cmd")
	robotgo.MilliSleep(100)
	robotgo.KeyTap("c", "cmd")
	text, err := robotgo.ReadAll()
	mtext := string(text[:])
	language, err := gocc.New("hk2s")
	if err != nil {
		fmt.Println("robotgo.ReadAll err is: ", err)
	}
	out, err := language.Convert(mtext)
	if err != nil {
		log.Fatal(err)
	}
	e := robotgo.WriteAll(out)
	if e != nil {
		fmt.Println("robotgo.WriteAll err is: ", e)
	}

	robotgo.KeyTap("v", "cmd")

	log.Printf("hotkey: %v is down\n", hk)
	//<-hk.Keyup()
	//log.Printf("hotkey: %v is up\n", hk)
	hk.Unregister()
	return
}
