package main

import (
	"log"
	"net/http"
	"prometheus-webhook/model"
	"time"

	//	"syscall"
	//	"unsafe"

	"github.com/gin-gonic/gin"
)

/*
func AlertSpeak(text string) {
	ttsdll := syscall.NewLazyDLL("tts.dll")
	speak := ttsdll.NewProc("rapidSpeakText")
	speak.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}
*/
func main() {
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, v := range notification.Alerts {
			//	AlertSpeak(v.Annotations["summary"])
			log.Println(v.Annotations["summary"])
			model.AlertSound("./soundfiles/alert.mp3")
			time.Sleep(100 * time.Millisecond)
		}
		c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})

	})
	router.Run(":9999")
}
