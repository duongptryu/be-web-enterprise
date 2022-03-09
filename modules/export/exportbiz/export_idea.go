package exportbiz

import (
	"bytes"
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
	"web/modules/export/exportmodel"
	"web/modules/idea/ideastore"
)

type exportIdeaBiz struct {
	ideaStore    ideastore.IdeaStore
	acaYearStore acayearstore.AcademicYearStore
}

func NewExportIdeaBiz(ideaStore ideastore.IdeaStore, acaYearStore acayearstore.AcademicYearStore) *exportIdeaBiz {
	return &exportIdeaBiz{
		ideaStore:    ideaStore,
		acaYearStore: acaYearStore,
	}
}

func (biz *exportIdeaBiz) ExportIdeaBiz(ctx context.Context, data *exportmodel.Export) (*bytes.Buffer, error) {
	acaYearExist, err := biz.acaYearStore.FindAcaYear(ctx, map[string]interface{}{"id": data.AcaYearId})
	if err != nil {
		return nil, err
	}
	if acaYearExist.Id == 0 {
		return nil, common.ErrDataNotFound(acayearmodel.EntityName)
	}

	data.NameAcaYear = acaYearExist.Title

	ideas, err := biz.ideaStore.ListALlIdea(ctx, map[string]interface{}{"aca_year_id": acaYearExist.Id}, "User", "Category", "Department")
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	// Create a new sheet.
	nameSheet := "ideas"
	index := f.NewSheet(nameSheet)

	//set title
	f.SetCellValue(nameSheet, "A1", "Id")
	f.SetCellValue(nameSheet, "B1", "Title")
	f.SetCellValue(nameSheet, "C1", "Content")
	f.SetCellValue(nameSheet, "D1", "Owner Id")
	f.SetCellValue(nameSheet, "E1", "Owner Name")
	f.SetCellValue(nameSheet, "F1", "Owner Email")
	f.SetCellValue(nameSheet, "G1", "Category")
	f.SetCellValue(nameSheet, "H1", "Likes count")
	f.SetCellValue(nameSheet, "I1", "Dislikes count")
	f.SetCellValue(nameSheet, "J1", "Voting rate")
	f.SetCellValue(nameSheet, "K1", "Views count")
	f.SetCellValue(nameSheet, "L1", "Comments count")
	f.SetCellValue(nameSheet, "M1", "Status")
	f.SetCellValue(nameSheet, "N1", "Is anonymous")
	f.SetCellValue(nameSheet, "O1", "Is expire")
	f.SetCellValue(nameSheet, "P1", "Files")
	f.SetCellValue(nameSheet, "Q1", "Created At")
	f.SetCellValue(nameSheet, "R1", "Link")
	// Set value of a cell.

	for i, v := range ideas {
		var files []string
		for _, v := range *v.Files {
			files = append(files, v.Url)
		}
		f.SetCellValue(nameSheet, fmt.Sprintf("A%v", i+2), v.Id)
		f.SetCellValue(nameSheet, fmt.Sprintf("B%v", i+2), v.Title)
		f.SetCellValue(nameSheet, fmt.Sprintf("C%v", i+2), v.Content)
		f.SetCellValue(nameSheet, fmt.Sprintf("D%v", i+2), v.UserId)
		f.SetCellValue(nameSheet, fmt.Sprintf("E%v", i+2), v.User.FullName)
		f.SetCellValue(nameSheet, fmt.Sprintf("F%v", i+2), v.User.Email)
		f.SetCellValue(nameSheet, fmt.Sprintf("G%v", i+2), v.Category.Name)
		f.SetCellValue(nameSheet, fmt.Sprintf("H%v", i+2), v.LikesCount)
		f.SetCellValue(nameSheet, fmt.Sprintf("I%v", i+2), v.DislikesCount)
		f.SetCellValue(nameSheet, fmt.Sprintf("J%v", i+2), v.LikesCount-v.DislikesCount)
		f.SetCellValue(nameSheet, fmt.Sprintf("K%v", i+2), v.ViewsCount)
		f.SetCellValue(nameSheet, fmt.Sprintf("L%v", i+2), v.CommentsCount)
		f.SetCellValue(nameSheet, fmt.Sprintf("M%v", i+2), v.Status)
		f.SetCellValue(nameSheet, fmt.Sprintf("N%v", i+2), v.IsAnonymous)
		f.SetCellValue(nameSheet, fmt.Sprintf("O%v", i+2), v.IsExpire)
		f.SetCellValue(nameSheet, fmt.Sprintf("P%v", i+2), files)
		f.SetCellValue(nameSheet, fmt.Sprintf("Q%v", i+2), v.CreatedAt.Format("2006-01-02T15:04:05"))
		f.SetCellValue(nameSheet, fmt.Sprintf("R%v", i+2), fmt.Sprintf("https://groupbar.me/ideas/%v", v.Id))
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.\
	result, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return result, nil
}
