package repository

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"errors"
	"regexp"
	"strconv"
	"strings"

	eg "api.brutecore/internal/engine"
	"api.brutecore/internal/utility"

	"api.brutecore/libs/lib_db"
	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
)

type SESSRepositoryLayer struct {
	db *lib_db.DB
	p  *eg.Pulling
}

type AlterSess struct {
	ID       *int64      `json:"id"`
	Field    string      `json:"field"`
	NewValue interface{} `json:"new_value"`
}

type MUploadProxy struct {
	SessionID int64  `json:"session_id"`
	ProxyType int    `json:"proxy_type"`
	TimeOut   int    `json:"timeout"`
	Data      string `json:"data"`
}

const (
	GoodStatus = "Хорошие"
	LogStatus  = "Хорошие с логом"
	Line       = "\n-------------\n"
)

func NewSESSRepositoryLayer(d *lib_db.DB, pl *eg.Pulling) *SESSRepositoryLayer {
	return &SESSRepositoryLayer{
		db: d,
		p:  pl,
	}
}

func (r *SESSRepositoryLayer) GetSessions() (*lib_db.DBResult, error) {
	return r.db.QueryRow(lib_db.TxRead, QGetSessions)
}

func (r *SESSRepositoryLayer) CreateSession() (*lib_db.DBResult, error) {
	uuid := uuid.New()
	_, err := r.db.Exec(lib_db.TxWrite, QInsertSession, uuid)
	if err != nil {
		return nil, err
	}

	return r.db.QueryRow(lib_db.TxRead, "SELECT ID, CREATE_TIME, NAME FROM SESSION WHERE NAME = $1", uuid)
}

func (r *SESSRepositoryLayer) DeleteSession(sID string) error {
	id, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(lib_db.TxWrite, QDeleteSession, id)
	return err
}

func (r *SESSRepositoryLayer) GetInfoSession(sID string) (*lib_db.DBResult, error) {
	id, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.db.QueryRow(lib_db.TxRead, QSessionInfo, id)
}

func (r *SESSRepositoryLayer) UploadProxy(resp *MUploadProxy) error {
	if resp.TimeOut == 0 {
		return errors.New("Укажите таймаут")
	}

	if resp.ProxyType < 18 || resp.ProxyType > 20 {
		return errors.New("Тип проси указан некорректно")
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(resp.Data)
	if err != nil {
		return errors.New("Ошибка декодирования данных")
	}
	decodedString := string(decodedBytes)
	lines := strings.Split(decodedString, "\n")
	tx, err := r.db.StartTx(lib_db.TxWrite)
	if err != nil {
		return err
	}
	trueTx := tx.(*sql.Tx)
	trueTx.Exec("DELETE FROM SESSION_PROXY WHERE SESSION_ID = $1", resp.SessionID)
	trueTx.Exec(QInsertSessionDTLTimeOut, resp.SessionID, "P1T", resp.TimeOut, "Таймаут прокси")
	for _, line := range lines {
		var port int
		dline := strings.Split(line, ":")
		if dline[0] == "" {
			continue
		}

		if port, err = strconv.Atoi(strings.ReplaceAll(dline[1], "\r", "")); err != nil {
			continue
		}
		trueTx.Exec(QInsertSessionProxy, resp.SessionID, dline[0], port, resp.ProxyType)
	}

	trueTx.Exec("UPDATE SESSION SET PROXY_ID = -2 WHERE ID = $1", resp.SessionID)

	err = trueTx.Commit()
	if err != nil {
		trueTx.Rollback()
		return err
	}

	return nil
}

func (r *SESSRepositoryLayer) AlterSession(s *AlterSess) error {
	if s.ID == nil {
		return errors.New("Укажите ID сессии")
	}

	if s.Field == "" {
		return errors.New("Укажите поле для изменения")
	}

	if s.NewValue == nil {
		return errors.New("Укажите новое значение")
	}

	switch strings.ToUpper(s.Field) {
	case "DATABASE_ID":
		res, err := r.db.QueryRow(lib_db.TxRead, QGetDatabaseType, s.NewValue)
		if err != nil {
			return err
		}

		if (*res)[0]["type"].(int64) != 3 {
			return errors.New("Можно выбирать только комболист")
		}

	case "PROXY_ID":
	case "MODULE_ID":
	case "STATUS":
		if s.NewValue == 22 {
			if err := r.StartSession(*s.ID); err != nil {
				return err
			}
		}

	case "NAME":
		if len(s.NewValue.(string)) > 60 {
			return errors.New("Длина поля превышает максимальную длину")
		}

	case "WORKER_COUNT":
		if s.NewValue.(float64) > 2000 || s.NewValue.(float64) < 1 {
			return errors.New("Недопустимое значение")
		}
	default:
		return errors.New("Неопознанное поле для изменения")
	}

	_, err := r.db.Exec(lib_db.TxWrite, `UPDATE "session" SET `+s.Field+` = $1 WHERE ID = $2`, s.NewValue, *s.ID)
	if strings.ToUpper(s.Field) == "STATUS" {
		r.p.Iteration()
	}
	return err
}

func (r *SESSRepositoryLayer) StartSession(sessionId int64) error {
	res, err := r.db.QueryRow(lib_db.TxRead, QGetSessionStartInfo, sessionId)
	if err != nil {
		return err
	}

	if (*res)[0]["database_id"] == nil {
		return errors.New("Укажите Комболист для перебора")
	}

	if (*res)[0]["proxy_id"] != nil && (*res)[0]["proxy_id"].(int64) == -2 &&
		(*res)[0]["PROX_COUNT"].(int64) < 5 {
		return errors.New("Прокси выбраны из файла но не загружены")
	}

	if (*res)[0]["module_id"] == nil {
		return errors.New("Укажите модуль для перебора")
	}

	return nil
}

func (r *SESSRepositoryLayer) GetSessionStatistic(sID string) (*lib_db.DBResult, *lib_db.DBResult, error) {
	id, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		return nil, nil, err
	}

	res1, err := r.db.QueryRow(lib_db.TxRead, QGetSessionStatistic, id, id)
	if err != nil {
		return nil, nil, err
	}

	res2, err := r.db.QueryRow(lib_db.TxRead, QGetPercent, id)
	if err != nil {
		return nil, nil, err
	}
	(*res2)[0]["proxycount"] = r.p.GetProxyCount(id)

	return res1, res2, nil
}

func (r *SESSRepositoryLayer) GetResults(sId string, pId string, eCount string, params string) (error, int, *lib_db.DBResult) {
	session_id, err := strconv.ParseInt(sId, 10, 64)
	if err != nil {
		return err, 0, nil
	}
	page_id, err := strconv.ParseInt(pId, 10, 64)
	if err != nil {
		return err, 0, nil
	}
	entries_count, err := strconv.ParseInt(eCount, 10, 64)

	match, err := regexp.MatchString("^\\d+(,\\d+)*$", params)
	if err != nil {
		return err, 0, nil
	}
	if !match {
		return errors.New("Перепроверьте запрашиваемые типы"), 0, nil
	}

	match, err = regexp.MatchString(`\b999\b`, params)
	if err != nil {
		return err, 0, nil
	}
	nine := 0
	if match {
		nine = 999
	}

	res, err := r.db.QueryRow(lib_db.TxRead, strings.Replace(QGetSessionPagesCount, "-_-", params, -1), session_id, nine, session_id)
	if err != nil {
		return err, 0, nil
	}

	dataCount := (*res)[0]["C1"].(int64)
	pagesCount := int(dataCount / entries_count)
	if dataCount%entries_count > 0 {
		pagesCount++
	}
	if pagesCount == 0 {
		pagesCount = 1
	}

	res, err = r.db.QueryRow(lib_db.TxRead, strings.Replace(QGetSessionResults, "-_-", params, -1), session_id, nine, session_id, entries_count, (page_id-1)*entries_count)
	if err != nil {
		return err, 0, nil
	}

	return nil, pagesCount, res
}

func (r *SESSRepositoryLayer) DownloadUniversal(sID string, sFormat string, params string, sType bool) ([]byte, error) {
	var (
		nine  int64
		err   error
		res   *lib_db.DBResult
		files []utility.File
	)
	session_id, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		return nil, err
	}

	format := strings.ToUpper(sFormat)
	if format != "TXT" && format != "XLSX" {
		return nil, errors.New("Неизвестный формат")
	}

	if sType {
		match, err := regexp.MatchString("^\\d+(,\\d+)*$", params)
		if err != nil {
			return nil, err
		}
		if !match {
			return nil, errors.New("Перепроверьте запрашиваемые типы")
		}

		match, err = regexp.MatchString(`\b999\b`, params)
		if err != nil {
			return nil, err
		}
		nine = 0
		if match {
			nine = 999
		}
	}

	if !sType {
		res, err = r.db.QueryRow(lib_db.TxRead, QueryALL, session_id, session_id)
	} else {
		res, err = r.db.QueryRow(lib_db.TxRead, strings.Replace(QuerySelected, "-_-", params, -1), session_id, nine, session_id)
	}
	if err != nil {
		return nil, err
	}

	switch format {
	case "TXT":
		buf := map[string]string{}
		for _, item := range *res {
			status := item["st"].(string)
			data := item["da"].(string)
			log := item["lo"].(string)

			buf[status] = buf[status] + data + "\n"
			if status == GoodStatus && log != "" {
				buf[LogStatus] = buf[LogStatus] + data + "\n" + log + Line
			}
		}
		for key, val := range buf {
			files = append(files, utility.File{
				Name: key + ".txt",
				Body: []byte(val),
			})
		}
	case "XLSX":
		excel := xlsx.NewFile()
		sheet, err := excel.AddSheet("Results")
		if err != nil {
			return nil, err
		}

		headRow := sheet.AddRow()
		for _, value := range []string{"Статус", "Данные", "Лог"} {
			headRow.AddCell().Value = value
		}

		for _, item := range *res {
			rowObj := sheet.AddRow()
			rowObj.AddCell().Value = item["st"].(string)
			rowObj.AddCell().Value = item["da"].(string)
			rowObj.AddCell().Value = item["lo"].(string)
		}

		var buffer bytes.Buffer
		err = excel.Write(&buffer)
		files = append(files, utility.File{
			Name: "Результаты.xlsx",
			Body: buffer.Bytes(),
		})
	}

	return utility.MakeZip(&files)
}
