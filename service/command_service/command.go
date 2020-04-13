package command_service

import (
	"encoding/json"
	"bytes"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"io/ioutil"
	"strings"
	//"strconv"
	//"time"
	"fmt"
	"net/http"
	"net/url"

	//"github.com/360EntSecGroup-Skylar/excelize"
	//"github.com/tealeg/xlsx"

        "github.com/EDDYCJY/go-gin-example/models"
	"github.com/tidwall/gjson"
	//"github.com/EDDYCJY/go-gin-example/pkg/export"
	//"github.com/EDDYCJY/go-gin-example/pkg/file"
	//"github.com/EDDYCJY/go-gin-example/pkg/logging"
)


type Resp struct {
     StatusCode int `json:"statusCode"`
     Messages   string `json:"messages"`
     Data       *Data  `json:"data"`
}

type Data struct {
     TaskId     string `json:"taskID"`
}

type Command struct {
}

func (t *Command) GetProgress(taskid string, agentid string, path string) (string, error) {
     obj := make(map[string]interface{})
     obj["task"] = taskid

     data, err := json.Marshal(obj)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     dest_url := "http://"
     dest_url += agentid
     dest_url += path

     reader := bytes.NewReader(data)
     request, err := http.NewRequest("POST", dest_url, reader)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }
     request.Header.Set("Content-Type", "application/json;charset=UTF-8")

     proxy, _ := url.Parse("http://127.0.0.1:8080")
     transport := &http.Transport{
           Proxy: http.ProxyURL(proxy),
     }
     client := &http.Client{
          Transport: transport,
     }


     client.Transport = transport

     resp, err := client.Do(request)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     respBytes, err := ioutil.ReadAll(resp.Body)
     if err != nil {
        fmt.Println(err.Error())
        return "",err
     }

     return string(respBytes),nil
}

func (t *Command) AssetSearch(taskid string, agentid string, path string) (string, error) {
     str := "taskid=" + taskid
     obj := make(map[string]interface{})
     obj["query"] = str

     data, err := json.Marshal(obj)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     dest_url := "http://"
     dest_url += agentid
     dest_url += path

     reader := bytes.NewReader(data)
     request, err := http.NewRequest("POST", dest_url, reader)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }
     request.Header.Set("Content-Type", "application/json;charset=UTF-8")

     proxy, _ := url.Parse("http://127.0.0.1:8080")
     transport := &http.Transport{
           Proxy: http.ProxyURL(proxy),
     }
     client := &http.Client{
          Transport: transport,
     }


     client.Transport = transport


     resp, err := client.Do(request)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     respBytes, err := ioutil.ReadAll(resp.Body)
     if err != nil {
        fmt.Println(err.Error())
        return "",err
     }

     return string(respBytes),nil
}

func (t *Command) GetAgentList() (interface{}, error) {
     url := "http://127.0.0.1:7500/api/proxy/http"

     resp, err := http.Get(url)
     if err != nil {
	return "", err
     }
     defer resp.Body.Close()
     s,err:=ioutil.ReadAll(resp.Body)

     fmt.Println(string(s))
     value := gjson.GetBytes(s, "proxies.#.name")
     var str []string
     for _,name := range value.Array() {
	 fmt.Println(name.String())
	 str = append(str, name.String())
     }

     //data, err := json.Marshal(str)
     return str,nil
}


func (t *Command) GetTaskAttr(taskid string) (string, error) {
     reply, err := gredis.Get(taskid)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return reply, err
     }

     return reply, nil
}

func (t *Command) GetTaskList() ([]string, error) {
     reply, err := gredis.SMembers("TaskSet")
     if(err != nil) {
	     fmt.Println(err.Error())
	     return reply, err
     }

     return reply, nil
}

func(t *Command) GetAgentMonitor(agentid string) (interface{}, error) {
     dest_url := "http://"
     dest_url += agentid
     dest_url += ".control"
     dest_url += "/api/agent_info"

     request, err := http.NewRequest("GET", dest_url, nil)

     proxy, _ := url.Parse("http://127.0.0.1:8080")
     transport := &http.Transport{
           Proxy: http.ProxyURL(proxy),
     }
     client := &http.Client{
          Transport: transport,
     }


     client.Transport = transport

     resp, err := client.Do(request)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     respBytes, err := ioutil.ReadAll(resp.Body)
     if err != nil {
        fmt.Println(err.Error())
        return "",err
     }

     return respBytes,nil
}

func(t *Command) RebootAgent(agentid string) (interface{}, error) {
     dest_url := "http://"
     dest_url += agentid
     dest_url += ".control"
     dest_url += "/api/reboot"

     request, err := http.NewRequest("GET", dest_url, nil)

     proxy, _ := url.Parse("http://127.0.0.1:8080")
     transport := &http.Transport{
           Proxy: http.ProxyURL(proxy),
     }
     client := &http.Client{
          Transport: transport,
     }


     client.Transport = transport

     resp, err := client.Do(request)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     respBytes, err := ioutil.ReadAll(resp.Body)
     if err != nil {
        fmt.Println(err.Error())
        return "",err
     }

     return respBytes,nil
}

func(t *Command) StartScan(agentid string, ips string, path string) (string, error) {
     obj := make(map[string]interface{})
     asset := make(map[string]interface{})

     obj["asset"] = asset
     asset["ips"] = strings.Split(ips, ",")
     asset["ports"] = "21,22"

     data, err := json.Marshal(obj)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     reader := bytes.NewReader(data)

     dest_url := "http://"
     dest_url += agentid
     dest_url += path 


     request, err := http.NewRequest("POST", dest_url, reader)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     request.Header.Set("Content-Type", "application/json;charset=UTF-8")

     proxy, _ := url.Parse("http://127.0.0.1:8080")
     transport := &http.Transport{
           Proxy: http.ProxyURL(proxy),
     }
     client := &http.Client{
          Transport: transport,
     }


     client.Transport = transport

     resp, err := client.Do(request)
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }

     respBytes, err := ioutil.ReadAll(resp.Body)
     if err != nil {
        fmt.Println(err.Error())
        return "",err
     }

     respData := Resp{}
     err = json.Unmarshal(respBytes, &respData)
     if err != nil {
        fmt.Println(err.Error())
        return "",err
     }

     fmt.Println(respData.Data.TaskId)
     gredis.SAdd("TaskSet", respData.Data.TaskId)

     obj["AgentId"] = agentid

     fmt.Println(string(data))
     gredis.Set(respData.Data.TaskId, obj, 24*3600)


     return respData.Data.TaskId, nil
}

func(t *Command) GetCommands() (string, error) {
     commands, err := models.GetCommands()
     if(err != nil) {
	     fmt.Println(err.Error())
	     return "",err
     }
     fmt.Println(commands)

     for _, command := range commands{
         fmt.Println(command.AgentId)
     }


     return "",nil
}
