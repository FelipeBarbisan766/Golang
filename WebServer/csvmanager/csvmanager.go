package csvmanager

import (
	"encoding/csv"
	"errors"
	"os"
)

type Org struct {
	Index           string
	OrganizationId  string
	Name            string
	Website         string
	Country         string
	Description     string
	Founded         string
	Industry        string
	NumberEmployees string
}
type CSVManager struct {
	filePath string 
}
func NewCSVManager() *CSVManager {
	return &CSVManager{
		filePath: "utils/organizations-100000.csv",
	}
}

func (c *CSVManager) ReadAll() ([]Org, error) {
	file, err := os.Open(c.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var orgs []Org
	for i, cols := range lines {
		if i == 0 { // Ignorar cabeçalho
			continue
		}
		orgs = append(orgs, Org{
			Index:           cols[0],
			OrganizationId:  cols[1],
			Name:            cols[2],
			Website:         cols[3],
			Country:         cols[4],
			Description:     cols[5],
			Founded:         cols[6],
			Industry:        cols[7],
			NumberEmployees: cols[8],
		})
	}
	return orgs, nil
}
func (c *CSVManager) ReadByLine(process func(Org)) error {
	file, err := os.Open(c.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Ignora cabeçalho
	_, _ = reader.Read()

	for {
		cols, err := reader.Read()
		if err != nil {
			break
		}
		org := Org{
			Index:           cols[0],
			OrganizationId:  cols[1],
			Name:            cols[2],
			Website:         cols[3],
			Country:         cols[4],
			Description:     cols[5],
			Founded:         cols[6],
			Industry:        cols[7],
			NumberEmployees: cols[8],
		}
		process(org)
	}
	return nil
}
func (c *CSVManager) Write(org Org) error {
	file, err := os.OpenFile(c.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		org.Index,
		org.OrganizationId,
		org.Name,
		org.Website,
		org.Country,
		org.Description,
		org.Founded,
		org.Industry,
		org.NumberEmployees,
	}

	return writer.Write(record)
}
func (c *CSVManager) SearchByField(field, value string) ([]Org, error) {
	orgs, err := c.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []Org
	for _, o := range orgs {
		switch field {
		case "Index":
			if o.Index == value {
				result = append(result, o)
			}
		case "Name":
			if o.Name == value {
				result = append(result, o)
			}
		case "Country":
			if o.Country == value {
				result = append(result, o)
			}
		default:
			return nil, errors.New("campo inválido")
		}
	}
	return result, nil
}
