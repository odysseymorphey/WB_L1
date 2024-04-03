package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Структура данных в формате JSON
type EmployeeJSON struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

// Структура данных в формате XML
type EmployeeXML struct {
	Name     string `xml:"name"`
	Position string `xml:"position"`
}

// Адаптер для преобразования данных из JSON в XML
type JSONToXMLAdapter struct {
	Employee EmployeeJSON
}

// Метод для преобразования данных из JSON в XML
func (a *JSONToXMLAdapter) ConvertToXML() (string, error) {
	// Преобразуем данные из JSON в XML
	employeeXML := EmployeeXML{
		Name:     a.Employee.Name,
		Position: a.Employee.Position,
	}

	// Кодируем данные в формате XML
	xmlData, err := xml.Marshal(employeeXML)
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}

func main() {
	// Данные о работнике в формате JSON
	jsonData := `{"name": "John Doe", "position": "Developer"}`

	// Распаковываем данные JSON
	var employeeJSON EmployeeJSON
	err := json.Unmarshal([]byte(jsonData), &employeeJSON)
	if err != nil {
		fmt.Println("Ошибка при распаковке данных JSON:", err)
		return
	}

	// Создаем адаптер
	adapter := JSONToXMLAdapter{Employee: employeeJSON}

	// Преобразуем данные из JSON в XML
	xmlData, err := adapter.ConvertToXML()
	if err != nil {
		fmt.Println("Ошибка при преобразовании данных из JSON в XML:", err)
		return
	}

	// Выводим данные в формате XML
	fmt.Println("Данные в формате XML:")
	fmt.Println(xmlData)
}
