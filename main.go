package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micmonay/keybd_event"
)

type WebhookReq struct {
	Version      string `json:"version"`
	CreatedAt    string `json:"created_at"`
	Id           string `json:"id"`
	AmountRaw    int    `json:"amount_raw"`
	Cut          int    `json:"cut"`
	DonatorName  string `json:"donator_name"`
	DonatorEmail string `json:"donator_email"`
	Message      string `json:"message"`
}

var ContainNameAkil = map[string]bool{
	"akil":   true,
	"yummin": true,
	"yummy":  true,
	"yami":   true,
}

var Ambatukam = map[string]bool{
	"ambatukam":  true,
	"ambadeblou": true,
	"ambasing":   true,
	"ambatunat":  true,
}

var EnoBening = map[string]bool{
	"eno":    true,
	"bening": true,
}

var Dhika = map[string]bool{
	"dhika":      true,
	"dika":       true,
	"ceo":        true,
	"kalinsun":   true,
	"cleansound": true,
	"andhikanug": true,
}

func main() {
	r := gin.Default()
	r.POST("/webhook", func(ctx *gin.Context) {
		var payload WebhookReq
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, "GAGAL BROOO")
			return
		}
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}

		// For linux, it is very important to wait 2 seconds
		if runtime.GOOS == "linux" {
			time.Sleep(2 * time.Second)
		}

		var howManyTimesWwillClicked int = 10
		time.Sleep(5 * time.Second)
		fmt.Println("Ini ammoutnya", payload.AmountRaw)
		switch payload.AmountRaw {
		case 10000:
			fmt.Println("Jaln nih")
			kb.SetKeys(keybd_event.VK_SPACE)
		case 15000:
			kb.SetKeys(keybd_event.VK_W)
		case 20000:
			kb.SetKeys(keybd_event.VK_SCROLL)
		case 50000:
			kb.HasALT(true)
			kb.SetKeys(keybd_event.VK_F4)
			howManyTimesWwillClicked = 4
		case 100000:
			kb.SetKeys(139)
		case 1000000:
			kb.SetKeys(keybd_event.VK_R)
		default:
			kb.SetKeys(keybd_event.VK_1)
			kb.SetKeys(keybd_event.VK_G)
		}

		// Press the selected keys
		err = kb.Launching()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("jalan nih")
		}

		// Or you can use Press and Release
		for i := 0; i < howManyTimesWwillClicked; i++ {
			kb.Press()
		}
		time.Sleep(3 * time.Second)
		kb.Release()
		kb.Clear()

		go func() {
			var err error
			var url string
			var contain bool
			messageArray := strings.Split(payload.Message, " ")

			for _, val := range messageArray {
				tmp := strings.ToLower(val)
				if _, contain := ContainNameAkil[tmp]; contain {
					url = "https://www.youtube.com/watch?v=84qzw26k5Oc&t=7s"
					contain = true
					break
				}

				if _, ambatukam := Ambatukam[tmp]; ambatukam {
					url = "https://www.youtube.com/watch?v=rV5ynCW-kVw"
					contain = true
				}

				if _, eno := EnoBening[tmp]; eno {
					url = "https://www.youtube.com/watch?v=4zD-3lxMivo&t=27s"
					contain = true
					break
				}

				if _, dhika := Dhika[tmp]; dhika {
					url = "https://www.youtube.com/watch?v=JiIHhAOjytw"
					contain = true
					break
				}
			}

			if contain {
				err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

			}

			if err != nil {
				log.Fatal(err)
			}
		}()

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
