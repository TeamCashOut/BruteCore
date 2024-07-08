package engine

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"

	"api.brutecore/internal/utility"
	"api.brutecore/libs/lib_db"
)

type Proxy struct {
	id         int64
	db         *lib_db.DB
	buffer     []ProxyRecord
	interval   int64
	timeout    int
	update     bool
	lastGetter int
	links      []links
	Worker     bool
	mup        sync.Mutex
}

type ProxyRecord struct {
	host  string
	port  int
	type_ int
}

type links struct {
	id         int64
	link       string
	link_type  int64
	proxy_type int
}

func NewProxy(d *lib_db.DB, proxy_id int64, session_id int64) (*Proxy, error) {
	prox := Proxy{
		id:         proxy_id,
		db:         d,
		lastGetter: -1,
		Worker:     true,
	}

	var (
		res *lib_db.DBResult
		err error
	)

	switch prox.id {
	case -2:
		if res, err = d.QueryRow(lib_db.TxRead, QGetSessionProxyDTL, session_id); err != nil {
			return nil, err
		}
		if res.Count() != 1 {
			return nil, errors.New("Таймаут прокси не был указан вручную")
		}
		prox.timeout, _ = strconv.Atoi((*res)[0]["value"].(string))

		if res, err = d.QueryRow(lib_db.TxRead, QGetSessionProxys, session_id); err != nil {
			return nil, err
		}
		if res.Count() < 5 {
			return nil, errors.New("Прокси не загружены в БД")
		}
		prox.fullfillBuff(&prox.buffer, res)

	case -1:
		return &prox, nil

	default:
		if res, err = d.QueryRow(lib_db.TxRead, QGetProxyInfo, prox.id); err != nil {
			return nil, err
		}
		if res.Count() != 1 {
			prox.id = -1
			return &prox, nil
		}

		prox.interval = (*res)[0]["interval"].(int64)
		prox.timeout = int((*res)[0]["timeout"].(int64))
		prox.update = ((*res)[0]["use_update"].(int64) == 1)
		for _, i := range *res {
			prox.links = append(prox.links, links{
				id:         i["id"].(int64),
				link:       i["link"].(string),
				link_type:  i["link_type"].(int64),
				proxy_type: int(i["proxy_type"].(int64)),
			})
		}
		if err = prox.UpdateProxyData(); err != nil {
			return nil, err
		}
	}

	if len(prox.buffer) < 5 {
		return nil, errors.New("Прокси не были загружены, перепроверьте пресет")
	}
	return &prox, nil
}

func (p *Proxy) fullfillBuff(b *[]ProxyRecord, r *lib_db.DBResult) {
	for _, i := range *r {
		*b = append(*b, ProxyRecord{
			host:  i["host"].(string),
			port:  int(i["port"].(int64)),
			type_: int(i["type"].(int64)),
		})
	}
}

func (p *Proxy) UpdateProxyData() error {
	var tempbuff []ProxyRecord
	for _, i := range p.links {
		var (
			body []byte
			err  error
		)

		if i.link_type == 15 || (i.link_type == 16 && !p.update) {
			res, err := p.db.QueryRow(lib_db.TxRead, QGetProxys, p.id, i.id)
			if err != nil {
				continue
			}
			p.fullfillBuff(&tempbuff, res)
		}

		if i.link_type == 16 && p.update {
			body, err = ioutil.ReadFile(i.link)
			if err != nil {
				continue
			}
		}

		if i.link_type == 17 {
			response, err := http.Get(i.link)
			if err != nil {
				continue
			}
			defer response.Body.Close()

			if body, err = ioutil.ReadAll(response.Body); err != nil {
				continue
			}
		}

		re := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}`)
		matches := re.FindAllString(string(body), -1)
		for _, match := range matches {
			if host, port, err := utility.SplitProxy(match); err == nil {
				tempbuff = append(tempbuff, ProxyRecord{
					host:  host,
					port:  port,
					type_: i.proxy_type,
				})
			}
		}
	}

	if len(tempbuff) != 0 {
		p.mup.Lock()
		defer p.mup.Unlock()
		p.buffer = tempbuff
	}
	return nil
}

func (p *Proxy) ProxyWorker() {
	for p.update && p.Worker {
		p.UpdateProxyData()
		runtime.Gosched()
		time.Sleep(time.Second * time.Duration(p.interval))
	}
}

func (p *Proxy) GiveProxy() *ProxyRecord {
	p.mup.Lock()
	defer p.mup.Unlock()

	if p.lastGetter == len(p.buffer)-1 {
		p.lastGetter = -1
	}
	p.lastGetter++
	return &p.buffer[p.lastGetter]
}