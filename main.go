package main

import (
	"fmt"
	"net/http"
	"runtime"
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
		for i := 0; i < 10; i++ {
			kb.Press()
		}
		time.Sleep(3 * time.Second)
		kb.Release()
		kb.Clear()

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
