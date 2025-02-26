package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

type templateData struct {
	Data        map[string]interface{}
	API         map[string]string
	AuthEnabled bool
	Version     string
	NewVersion  bool
}

var functions = template.FuncMap{
	"formatRemarks":  formatRemarks,
	"formatLandings": formatLandings,
	"formatNumber":   formatNumber,
}

//go:embed templates
var templateFS embed.FS

// formatRemarks cuts the remark field if it's too long
func formatRemarks(remarks string) string {
	if len(remarks) > 13 {
		return fmt.Sprintf("%.10s...", remarks)
	} else {
		return remarks
	}
}

func formatLandings(landing int) string {
	if landing == 0 {
		return ""
	} else {
		return fmt.Sprintf("%d", landing)
	}
}

func formatNumber(n int) string {
	s := fmt.Sprintf("%d", n)

	if n >= 1000 {
		s = s[:len(s)-3] + " " + s[len(s)-3:]
	}
	if n >= 1000000 {
		s = s[:len(s)-7] + " " + s[len(s)-7:]
	}
	if n >= 1000000000 {
		s = s[:len(s)-11] + " " + s[len(s)-11:]
	}

	return s
}

// addDefaultData adds default values/consts to all templates
func (app *application) addDefaultData(td *templateData, req *http.Request) *templateData {
	td.API = make(map[string]string)
	td.API["URL"] = fmt.Sprintf("http://localhost:%d", app.config.port)
	td.API["Root"] = APIRoot
	td.API["Logbook"] = APILogbook
	td.API["LogbookData"] = APILogbookData
	td.API["LogbookUUID"] = APILogbookUUID
	td.API["LogbookNew"] = APILogbookNew
	td.API["LogbookSave"] = APILogbookSave
	td.API["LogbookDelete"] = APILogbookDelete
	td.API["LogbookNight"] = APILogbookNight
	td.API["Aircrafts"] = APIAircrafts
	td.API["AircraftsFilter"] = APIAircraftsFilter
	td.API["LogbookUUIDAttachments"] = APILogbookUUIDAttachments
	td.API["LogbookAttachmentsUpload"] = APILogbookAttachmentsUpload
	td.API["LogbookAttachmentsDelete"] = APILogbookAttachmentsDelete
	td.API["LogbookAttachmentsDownload"] = APILogbookAttachmentsDownload
	td.API["LogbookAttachmentsDownloadUUID"] = APILogbookAttachmentsDownloadUUID
	td.API["Export"] = APIExport
	td.API["ExportFormat"] = APIExportFormat
	td.API["Airport"] = APIAirport
	td.API["AirportID"] = APIAirportID
	td.API["AirportUpdate"] = APIAirportUpdate
	td.API["Settings"] = APISettings
	td.API["SettingsAircraftClasses"] = APISettingsAircraftClasses
	td.API["Stats"] = APIStats
	td.API["StatsTotals"] = APIStatsTotals
	td.API["StatsTotalsByType"] = APIStatsTotalsByType
	td.API["StatsTotalsByClass"] = APIStatsTotalsByClass
	td.API["StatsLimits"] = APIStatsLimits
	td.API["Map"] = APIMap
	td.API["MapData"] = APIMapData
	td.API["Licensing"] = APILicensing
	td.API["LicensingData"] = APILicensingData
	td.API["LicensingUUID"] = APILicensingUUID
	td.API["LicensingNew"] = APILicensingNew
	td.API["LicensingDownload"] = APILicensingDownload
	td.API["LicensingDownloadUUID"] = APILicensingDownloadUUID
	td.API["LicensingSave"] = APILicensingSave
	td.API["LicensingDelete"] = APILicensingDelete
	td.API["LicensingAttachmentDelete"] = APILicensingAttachmentDelete
	td.API["Login"] = APILogin
	td.API["Logout"] = APILogout

	td.API["ExportFormatA4"] = exportA4
	td.API["ExportFormatA5"] = exportA5
	td.API["ExportFormatCSV"] = exportCSV
	td.API["ExportFormatXLS"] = exportXLS

	td.AuthEnabled = app.isAuthEnabled
	td.Version = app.version
	td.NewVersion = app.isNewVersion

	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, req *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.templateCache[templateToRender]

	if app.config.env == "prod" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(page, templateToRender, partials)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, req)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(page string, templateToRender string, partials []string) (t *template.Template, err error) {

	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partials.gohtml", x)
		}
	}

	if len(partials) > 0 {
		partials = append(partials, "templates/base.gohtml")
		partials = append(partials, templateToRender)
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, partials...)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.gohtml", templateToRender)
	}

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t

	return t, nil
}
