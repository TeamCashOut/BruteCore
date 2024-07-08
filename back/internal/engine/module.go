package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"api.brutecore/internal/utility"
	"api.brutecore/libs/lib_db"
	"github.com/shirou/gopsutil/process"
)

type Module struct {
	id       int64
	pid      int
	fileName string
	port     string
	cmd      *exec.Cmd
	mup      sync.Mutex
}

type ExecuteInput struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Pin       string `json:"pin,omitempty"`
	Data      string `json:"data,omitempty"`
	TimeOut   int    `json:"timeout,omitempty"`
	ProxyHost string `json:"proxy_host,omitempty"`
	ProxyPort int    `json:"proxy_port,omitempty"`
	ProxyType int    `json:"proxy_type,omitempty"`
}

type ExecuteOutput struct {
	Status int     `json:"status"`
	Log    *string `json:"log,omitempty"`
}

const (
	subp = "http://127.0.0.1:%s/ExecuteModule"
)

func NewModule(d *lib_db.DB, module_id int64) (*Module, error) {
	res, err := d.QueryRow(lib_db.TxRead, "SELECT PATH FROM MODULE WHERE ID = $1", module_id)
	if err != nil {
		return nil, err
	}

	fileName := (*res)[0]["path"].(string)
	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}

	return &Module{
		id:       module_id,
		fileName: fileName,
	}, nil
}

func (m *Module) InitalizeModule() {
	for {
		port, err := utility.FindFreePort()
		if err != nil {
			continue
		}

		m.port = strconv.Itoa(port)
		m.cmd = exec.Command(m.fileName, "-port="+m.port)
		if err := m.cmd.Start(); err != nil {
			continue
		} else {
			m.pid = m.cmd.Process.Pid
			break
		}
	}
	time.Sleep(time.Second * 3)
}

func (m *Module) UnInitalizeModule() error {
	return m.cmd.Process.Kill()
}

func (m *Module) IsRunningProccess() bool {
	_, err := process.NewProcess(int32(m.pid))
	return err == nil
}

func (m *Module) CheckRunningProccess() {
	m.mup.Lock()
	defer m.mup.Unlock()

	if !m.IsRunningProccess() {
		m.InitalizeModule()
	}
}

func (m *Module) ExecuteModule(c *ComboListRecord, p *ProxyRecord, timeout int, data_type int64) (int, *string) {
	strct := ExecuteInput{
		TimeOut:   timeout,
	}
	
	if p != nil {
		strct.ProxyHost = p.host
		strct.ProxyPort = p.port
		strct.ProxyType = p.type_
	}

	switch data_type {
	case 4:
		strct.Email = c.data
	case 5:
		strct.Username = c.data
	case 6:
		strct.Password = c.data
	case 7:
		strct.Pin = c.data
	case 8, 9:
		parts := strings.Split(c.data, ":")
		if len(parts) == 2 {
			strct.Password = parts[1]
			if data_type == 8 {
				strct.Username = parts[0]
			} else {
				strct.Email = parts[0]
			}
		} else {
			return 28, nil
		}
	case 10:
		strct.Data = c.data
	default:
		return 28, nil
	}

	jsonData, err := json.Marshal(strct)
	if err != nil {
		return 28, nil
	}

	resp, err := http.Post(fmt.Sprintf(subp, m.port), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		m.CheckRunningProccess()
		return m.ExecuteModule(c, p, timeout, data_type)
	}
	defer resp.Body.Close()

	var outputData ExecuteOutput
	if err := json.NewDecoder(resp.Body).Decode(&outputData); err != nil {
		return 28, nil
	}

	return outputData.Status, outputData.Log
}
