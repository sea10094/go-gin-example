package v1

import (
	"net/http"
	"fmt"

	//"github.com/unknwon/com"
	//"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	//"github.com/EDDYCJY/go-gin-example/pkg/export"
	//"github.com/EDDYCJY/go-gin-example/pkg/logging"
	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	//"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/command_service"
)



// @Summary get task attr 
// @Produce  json
// @Param taskid query string false "task id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getTaskAttr [get]
func GetTaskAttr(c *gin.Context) {
     appG := app.Gin{C: c}

     taskid := c.Query("taskid")

     command := command_service.Command{}

     reply, err := command.GetTaskAttr(taskid)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return
     }


     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"taskinfo": reply,
     })

}

// @Summary get task list 
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getTaskList [get]
func GetTaskList(c *gin.Context) {
     appG := app.Gin{C: c}

     command := command_service.Command{}

     reply, err := command.GetTaskList()
     if(err != nil) {
	     fmt.Println(err.Error())
	     return
     }

     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"tasks": reply,
     })
}

// @Summary get agent list 
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getAgentList [get]
func GetAgentList(c *gin.Context) {
     appG := app.Gin{C: c}

     command := command_service.Command{}
     reply, err := command.GetAgentList()
     if err != nil {
	return
     }

     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"resp": reply,
     })
}

// @Summary get asset details of an agent
// @Produce  json
// @Param taskid query string false "task id"
// @Param agentid query string false "agent id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/assetSearch [get]
func AssetSearch(c *gin.Context) {
     appG := app.Gin{C: c}
     taskid := c.Query("taskid")
     agentid := c.Query("agentid")
     command := command_service.Command{}
     reply, err := command.AssetSearch(taskid, agentid, c.Request.URL.Path)
     if err != nil {
	return
     }



     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"resp": reply,
     })

}

// @Summary get progress of a task 
// @Produce  json
// @Param taskid query string false "task id"
// @Param agentid query string false "agent id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getProgress [get]
func GetProgress(c *gin.Context) {
     appG := app.Gin{C: c}
     taskid := c.Query("taskid")
     agentid := c.Query("agentid")

     command := command_service.Command{}
     reply, err := command.GetProgress(taskid, agentid, c.Request.URL.Path)
     if err != nil {
	return
     }

     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"resp": reply,
     })
}

func GetCommands(c *gin.Context) {
     appG := app.Gin{C: c}
     command := command_service.Command{}
     reply, err := command.GetCommands()
     if err != nil {
	return
     }

     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"resp": reply,
     })
}


// @Summary Start scan task to the pointed agent 
// @Produce  json
// @Param agentid query string false "agent id"
// @Param ips query string false "ip segment"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/startScan [get]
func StartScan(c *gin.Context) {
     appG := app.Gin{C: c}
     agentid := c.Query("agentid")
     ips := c.Query("ips")
     command := command_service.Command{}
     reply, err := command.StartScan(agentid, ips, c.Request.URL.Path)
     if err != nil {
	return
     }

     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"agentid": reply,
     })
}

// @Summary reboot the specialized agent 
// @Produce  json
// @Param agentid query string false "agent id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/rebootAgent [get]
func RebootAgent(c *gin.Context) {
     appG := app.Gin{C: c}
     agentid := c.Query("agentid")
     command := command_service.Command{}
     reply, err := command.RebootAgent(agentid)
     if err != nil {
	return
     }

     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"agentid": reply,
     })
}


// @Summary get agent monitor info 
// @Produce  json
// @Param agentid query string false "agent id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getAgentMonitor [get]
func GetAgentMonitor(c *gin.Context) {
     appG := app.Gin{C: c}
     agentid := c.Query("agentid")
     command := command_service.Command{}
     reply, err := command.GetAgentMonitor(agentid)
     if err != nil {
	return
     }

     appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"agentid": reply,
     })
}
