package services

import (
	"recipe/models"
	"strconv"
)

type Material struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Process struct {
	ProcessName string     `json:"processName"`
	DisplayName string     `json:"displayName"`
	Materials   []Material `json:"materials"`
	Description string     `json:"description"`
}

func GetProcess(processid string) (Process, error) {
	process,err := models.GetProcessById(processid)

	// エラー処理
	if err != nil {
		return Process{}, err
	}

	materials := []Material{}

	for _, val := range process.Material {
		material := Material{
			Name:     val.Name,
			Quantity:  Float32ToString(val.Count) + " " + val.Unit,
		}
		materials = append(materials, material)
	}

	return Process{
		ProcessName: process.Name,
		DisplayName: process.Name,
		Materials:   materials,
		Description: process.Description,
	},nil
}

func Float32ToString(num float32) string {
	return strconv.FormatFloat(float64(num), 'f', -1, 32)
}