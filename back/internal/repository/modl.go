package repository

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"api.brutecore/libs/lib_db"
	"github.com/gofiber/fiber/v2"
)

type MODLRepositoryLayer struct {
	db *lib_db.DB
}

const (
	dirName = "modules"
)

func NewMODLRepositoryLayer(d *lib_db.DB) *MODLRepositoryLayer {
	return &MODLRepositoryLayer{
		db: d,
	}
}

func (r *MODLRepositoryLayer) GetModules() (*lib_db.DBResult, error) {
	return r.db.QueryRow(lib_db.TxRead, QGetMdules)
}

func (r *MODLRepositoryLayer) DeleteModule(sID string) error {
	id, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		return err
	}

	res, err := r.db.QueryRow(lib_db.TxRead, "SELECT * FROM MODULE WHERE ID = $1", id)
	if err != nil {
		return err
	}

	if err := os.Remove((*res)[0]["path"].(string)); err != nil {
		return err
	}

	_, err = r.db.Exec(lib_db.TxWrite, "DELETE FROM MODULE WHERE ID = $1", id)
	return err
}

func (r *MODLRepositoryLayer) UploadModule(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return errors.New("Файл модуля не загружен")
	}

	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return errors.New("Не удалось создать директорию")
	}

	fileName := filepath.Join(dirName, file.Filename)
	if _, err := os.Stat(fileName); err == nil {
		return errors.New("Модуль с таким именем уже существует")
	}

	if err := c.SaveFile(file, fileName); err != nil {
		return errors.New("Не удалось сохранить файл")
	}

	if strings.ToUpper(runtime.GOOS) == "LINUX" {
		_, _ = exec.Command("sudo", "chmod", "+x", fileName).Output()
	}
	cmd, err := exec.Command(fileName, "-getinfo").Output()
	if err != nil {
		return err
	}

	mapa := make(map[string]interface{})
	mapa["name"] = "none"
	mapa["version"] = "none"
	mapa["author"] = "none"
	mapa["data_type"] = 0

	if err := json.Unmarshal(cmd, &mapa); err != nil {
		return err
	}

	_, err = r.db.Exec(lib_db.TxWrite, QInsertModule, mapa["name"].(string), mapa["version"].(string), mapa["author"].(string), mapa["data_type"].(float64), fileName)
	return err
}
