package models

import "database/sql/driver"

const AllAircrafts = 0
const LastAircrafts = 1

// jsonResponse is a type for post data handlers response
type JSONResponse struct {
	OK          bool   `json:"ok"`
	Message     string `json:"message,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
}

// FlightRecord is a type for logbook flight records
type FlightRecord struct {
	UUID      string `json:"uuid"`
	Date      string `json:"date"`
	MDate     string `json:"m_date"`
	Departure struct {
		Place string `json:"place"`
		Time  string `json:"time"`
	} `json:"departure"`
	Arrival struct {
		Place string `json:"place"`
		Time  string `json:"time"`
	} `json:"arrival"`
	Aircraft struct {
		Model string `json:"model"`
		Reg   string `json:"reg_name"`
	} `json:"aircraft"`
	Time struct {
		SE         string `json:"se_time"`
		ME         string `json:"me_time"`
		MCC        string `json:"mcc_time"`
		Total      string `json:"total_time"`
		Night      string `json:"night_time"`
		IFR        string `json:"ifr_time"`
		PIC        string `json:"pic_time"`
		CoPilot    string `json:"co_pilot_time"`
		Dual       string `json:"dual_time"`
		Instructor string `json:"instructor_time"`
	} `json:"time"`
	Landings struct {
		Day   int `json:"day"`
		Night int `json:"night"`
	} `json:"landings"`
	SIM struct {
		Type string `json:"type"`
		Time string `json:"time"`
	} `json:"sim"`
	PIC     string `json:"pic_name"`
	Remarks string `json:"remarks"`

	Distance int
}

// Airpot is a structure for airport record
type Airport struct {
	ICAO      string  `json:"icao"`
	IATA      string  `json:"iata"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Elevation int     `json:"elevation"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
}

type ColumnsWidth struct {
	Col1  float64 `json:"col1"`
	Col2  float64 `json:"col2"`
	Col3  float64 `json:"col3"`
	Col4  float64 `json:"col4"`
	Col5  float64 `json:"col5"`
	Col6  float64 `json:"col6"`
	Col7  float64 `json:"col7"`
	Col8  float64 `json:"col8"`
	Col9  float64 `json:"col9"`
	Col10 float64 `json:"col10"`
	Col11 float64 `json:"col11"`
	Col12 float64 `json:"col12"`
	Col13 float64 `json:"col13"`
	Col14 float64 `json:"col14"`
	Col15 float64 `json:"col15"`
	Col16 float64 `json:"col16"`
	Col17 float64 `json:"col17"`
	Col18 float64 `json:"col18"`
	Col19 float64 `json:"col19"`
	Col20 float64 `json:"col20"`
	Col21 float64 `json:"col21"`
	Col22 float64 `json:"col22"`
	Col23 float64 `json:"col23"`
}

type ExportPDF struct {
	LogbookRows int          `json:"logbook_rows"`
	Fill        int          `json:"fill"`
	LeftMargin  float64      `json:"left_margin"`
	LeftMarginA float64      `json:"left_margin_a"`
	LeftMarginB float64      `json:"left_margin_b"`
	TopMargin   float64      `json:"top_margin"`
	BodyRow     float64      `json:"body_row_height"`
	FooterRow   float64      `json:"footer_row_height"`
	PageBreaks  string       `json:"page_breaks"`
	Columns     ColumnsWidth `json:"columns"`
}

type ExportXLS struct {
	ConvertTime bool `json:"convert_time"`
}

type ExportCSV struct {
	Delimeter string `json:"delimeter"`
	CRLF      bool   `json:"crlf"`
}

// Settings is a type for settings
type Settings struct {
	OwnerName              string            `json:"owner_name"`
	SignatureText          string            `json:"signature_text"`
	AircraftClasses        map[string]string `json:"aircraft_classes"`
	AuthEnabled            bool              `json:"auth_enabled"`
	Login                  string            `json:"login"`
	Password               string            `json:"password"`
	Hash                   string            `json:"hash"`
	EnableFlightRecordHelp bool              `json:"enable_flightrecord_help"`

	ExportA4  ExportPDF `json:"export_a4"`
	ExportA5  ExportPDF `json:"export_a5"`
	ExportXLS ExportXLS `json:"export_xls"`
	ExportCSV ExportCSV `json:"export_csv"`
}

// License is a type for licesing
type License struct {
	UUID         string `json:"uuid"`
	Category     string `json:"category"`
	Name         string `json:"name"`
	Number       string `json:"number"`
	Issued       string `json:"issued"`
	ValidFrom    string `json:"valid_from"`
	ValidUntil   string `json:"valid_until"`
	Remarks      string `json:"remarks"`
	DocumentName string `json:"document_name"`
	Document     []byte
}

// Attachment is a type for attachments
type Attachment struct {
	UUID         string `json:"uuid"`
	RecordID     string `json:"record_id"`
	DocumentName string `json:"document_name"`
	Document     []byte
}

// TableData is a type for Datatables
type TableData struct {
	Data [][]string `json:"data"`
}

// Mock is a type for mocking sql requests
type Mock struct {
	Query  string
	Rows   []string
	Values []driver.Value
	Args   []driver.Value
}
