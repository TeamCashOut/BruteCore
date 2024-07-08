package handler

import (
	"encoding/json"
	"strings"
	"time"

	eg "api.brutecore/internal/engine"

	"api.brutecore/internal/repository"
	"api.brutecore/libs/lib_db"
	"github.com/gofiber/fiber/v2"
)

type SESSLayer struct {
	repo repository.SESSRepositoryLayer
}

func NewSESSLayer(d *lib_db.DB, pl *eg.Pulling) *SESSLayer {
	return &SESSLayer{
		repo: *repository.NewSESSRepositoryLayer(d, pl),
	}
}

func (sl *SESSLayer) GetSessions(c *fiber.Ctx) error {
	sessions, err := sl.repo.GetSessions()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "sessions": sessions,
		})
	}
}

func (sl *SESSLayer) CreateSession(c *fiber.Ctx) error {
	res, err := sl.repo.CreateSession()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "info": (*res)[0],
		})
	}
}

func (sl *SESSLayer) DeleteSession(c *fiber.Ctx) error {
	err := sl.repo.DeleteSession(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "msg_txt": "Элемент удален",
		})
	}
}

func (sl *SESSLayer) GetInfoSession(c *fiber.Ctx) error {
	session, err := sl.repo.GetInfoSession(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "session": (*session)[0],
		})
	}
}

func (sl *SESSLayer) AlterSession(c *fiber.Ctx) error {
	data := new(repository.AlterSess)
	if err := json.Unmarshal(c.Body(), data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": "Invalid JSON data",
		})
	}

	err := sl.repo.AlterSession(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "msg_txt": "Field was been changed",
		})
	}
}

func (sl *SESSLayer) UploadProxy(c *fiber.Ctx) error {
	data := new(repository.MUploadProxy)
	if err := json.Unmarshal(c.Body(), data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": "Invalid JSON data",
		})
	}

	err := sl.repo.UploadProxy(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "msg_txt": "Прокси были загружены",
		})
	}
}

func (sl *SESSLayer) GetStatistic(c *fiber.Ctx) error {
	result, statistic, err := sl.repo.GetSessionStatistic(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		status := (*statistic)[0]["STATUS_NAME"].(string)
		delete((*statistic)[0], "STATUS_NAME")

		finish_time := "Не закончено"
		if  (*statistic)[0]["FINISH_TIME"] != nil {
			finish_time = (*statistic)[0]["FINISH_TIME"].(time.Time).Format("2006-01-02 15:04:05")
		}
		delete((*statistic)[0], "FINISH_TIME")

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "result": result, "statistic": (*statistic)[0],
			"status": status, "finish_time": finish_time,
		})
	}
}

func (sl *SESSLayer) GetResults(c *fiber.Ctx) error {
	err, pagescount, data := sl.repo.GetResults(c.Query("session_id"), c.Query("page"), c.Query("entriescount"), c.Query("params"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "pagescount": pagescount, "data": data,
		})
	}
}

func (sl *SESSLayer) DownloadUniversal(c *fiber.Ctx) error {
	selected := strings.Contains(strings.ToUpper(c.Path()), "DOWNLOADSELECTED")

	archive, err := sl.repo.DownloadUniversal(c.Query("session_id"), c.Query("format"), c.Query("params"), selected)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false, "msg_txt": err.Error(),
		})
	} else {
		c.Set("Content-Disposition", "attachment; filename=results.zip")
		c.Set("Content-Type", "application/zip")
		return c.Send(archive)
	}
}
