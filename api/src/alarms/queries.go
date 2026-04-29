// src/alarms/handlers.go
package alarms

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

var globalAlarmContext *AlarmContext

// Initialize - καλέστε αυτή τη function από το main.go
func Init(ctx *AlarmContext) {
    globalAlarmContext = ctx
}

type GetAlarmsRequest struct {
    TransmitterSerial string `json:"transmitter_serial"`
    Device           string `json:"device"`
    AlarmGroup       string `json:"alarm_group"`
}

func HandleGetAlarms(c *gin.Context) {
    var req GetAlarmsRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    alarms, err := GetAlarms(globalAlarmContext, req.TransmitterSerial, req.Device, req.AlarmGroup)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": alarms})
}

type SetAlarmRequest struct {
    TransmitterSerial string `json:"transmitter_serial"`
    Device           string `json:"device"`
    AlarmGroup       string `json:"alarm_group"`
    Alarm            string `json:"alarm"`
}

func HandleSetAlarm(c *gin.Context) {
    var req SetAlarmRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := SetAlarm(globalAlarmContext, req.TransmitterSerial, req.Device, req.AlarmGroup, req.Alarm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}