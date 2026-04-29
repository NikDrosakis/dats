// src/alarms/handlers.go
package alarms

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

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

    // Get context from somewhere (global or passed)
    alarms, err := GetAlarms(globalCtx, req.TransmitterSerial, req.Device, req.AlarmGroup)
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

    if err := SetAlarm(globalCtx, req.TransmitterSerial, req.Device, req.AlarmGroup, req.Alarm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}