package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type LoanInfo struct {
	LoanNumber           string
	LoanBookedDate       string
	LoanType             string
	PrincipalAmount      float64
	InterestRate         float64
	Tenure               int
	OutStandingPrincipal float64
}
type EMIInfo struct {
	PrincipalAmount float64
	InterestAmount  float64
	EMIDate         string
}

func main() {

}

func CreatePaymentSchedule(fn string, numEMI int) {
	var pdf *gofpdf.Fpdf
	pdf = PageSetup(pdf)
	pdf.SetFont("Arial", "U", 10)
	pdf.Cell(250, 50, "Loan EMI Table")
	pdf.Ln(-1)
	AddLoanTableToPage(pdf)
	AddEMIDataToPage(pdf, numEMI)
	err := pdf.OutputFileAndClose(fn)
	if err != nil {
		log.Fatal(err)
	}
}
func AddEMIDataToPage(pdf *gofpdf.Fpdf, numEMI int) {
	//set header
	pdf.SetFont("Arial", "", 5)
	for _, val := range GetEMIHeader() {
		pdf.CellFormat(60, 5, val, "1", 0, "CM", false, 0, "")
	}
	pdf.Ln(-1)
	for _, val := range MockEMIData(numEMI) {
		amount, interest := strconv.FormatFloat(val.PrincipalAmount, 'f', 2, 64), strconv.FormatFloat(val.InterestAmount, 'f', 2, 64)

		pdf.CellFormat(60, 5, amount, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(60, 5, interest, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(60, 5, val.EMIDate, "1", 0, "CM", false, 0, "")
		pdf.Ln(-1)

	}

}
func AddLoanTableToPage(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 5)
	for _, val := range GetHeader() {
		pdf.CellFormat(25, 8, val, "1", 0, "CM", false, 0, "")
	}
	pdf.Ln(-1)

	loanInfo := LoanInfo{LoanNumber: "002001001002", LoanBookedDate: "31 Aug 2020",
		LoanType: "JUMBOLOAN", PrincipalAmount: 4626.3,
		InterestRate:         17.3,
		Tenure:               18,
		OutStandingPrincipal: 3882.93,
	}
	pdf.CellFormat(25, 8, loanInfo.LoanNumber, "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, loanInfo.LoanBookedDate, "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, loanInfo.LoanType, "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.PrincipalAmount), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.InterestRate), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.Tenure), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.OutStandingPrincipal), "1", 0, "CM", false, 0, "")
	//fill EMI data
	pdf.Ln(-1)
	pdf.Ln(-1)

}
func PageSetup(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	marginX := 10.0
	marginY := 20.0
	PaperSize := "A4"
	pdf = gofpdf.New("P", "mm", PaperSize, "")

	pdf.SetMargins(marginX, marginY, marginX)
	pdf.AddPage()
	pdf.SetAuthor("HDFC Credit Card Division", true)
	hdfcLogoPath := "D:\\Work\\GoWorkspace\\DB4\\HDFCLogo.png"
	pdf.ImageOptions(hdfcLogoPath, 10, 10, 40, 10, false, gofpdf.ImageOptions{}, 0, "")
	return pdf
}

func GetHeader() []string {

	return []string{"Loan Number", "Loan Booked Date",
		"Loan Type", "Principal Amount(Rs.)", "Interest Rate(%)", "Tenure(months)", "Outstanding Principal(Rs.)"}
}
func GetEMIHeader() []string {

	return []string{"Principal Amount (Rs)", "Interest Amount (Rs)", "EMI Date"}
}

func MockEMIData(numEMI int) []EMIInfo {
	//generate random float amounts.
	rand.Seed(time.Now().UnixNano())

	info := make([]EMIInfo, 0)
	for t := 0; t < numEMI; t++ {
		principal := fmt.Sprintf("%0.2f", rand.Float64()*100)
		principalf, _ := strconv.ParseFloat(principal, 64)
		interest := fmt.Sprintf("%0.2f", rand.Float64()*100)
		interestf, _ := strconv.ParseFloat(interest, 64)
		info = append(info, EMIInfo{PrincipalAmount: principalf, InterestAmount: interestf,
			EMIDate: GetRandomDate()})
	}
	return info

}
func GetRandomDate() string {
	months := []string{"Jan", "Feb", "Mar", "Apr",
		"May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	dates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28}
	years := []int{2019, 2020, 2021, 2022, 2023}
	rdate := dates[rand.Intn(len(dates)-1)]
	rmonth := months[rand.Intn(len(months)-1)]
	ryear := years[rand.Intn(len(years)-1)]
	val := fmt.Sprintf("%d %s %d", rdate, rmonth, ryear)
	return val
}

//test

func CreatePaymentSchedule1() {
	marginX := 10.0
	marginY := 20.0
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetMargins(marginX, marginY, marginX)
	pdf.AddPage()
	pdf.SetAuthor("HDFC Ccredit Card Division", true)
	hdfcLogoPath := "D:\\Work\\GoWorkspace\\DB4\\HDFCLogo.png"
	pdf.ImageOptions(hdfcLogoPath, 10, 10, 40, 10, false, gofpdf.ImageOptions{}, 0, "")
	pdf.SetFont("Arial", "U", 10)
	pdf.Cell(200, 50, "Loan EMI Table")

	pdf.Ln(-1)
	pdf.SetFont("Arial", "B", 5)
	for _, val := range GetHeader() {
		pdf.CellFormat(25, 8, val, "1", 0, "CM", false, 0, "")

	}
	pdf.Ln(-1)
	loanInfo := LoanInfo{LoanNumber: "002001001002", LoanBookedDate: "31 Aug 2020",
		LoanType: "JUMBOLOAN", PrincipalAmount: 4626.3,
		InterestRate:         17.3,
		Tenure:               18,
		OutStandingPrincipal: 3882.93,
	}
	pdf.CellFormat(25, 8, loanInfo.LoanNumber, "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, loanInfo.LoanBookedDate, "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, loanInfo.LoanType, "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.PrincipalAmount), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.InterestRate), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.Tenure), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(25, 8, fmt.Sprintf("%v", loanInfo.OutStandingPrincipal), "1", 0, "CM", false, 0, "")
	//fill EMI data
	pdf.Ln(-1)
	pdf.Ln(-1)
	for _, val := range GetEMIHeader() {
		pdf.CellFormat(60, 8, val, "1", 0, "CM", false, 0, "")
	}
	pdf.Ln(-1)
	err := pdf.OutputFileAndClose("Payment.pdf")
	if err != nil {
		log.Fatal(err)
	}
	//EncodePaymentSchedule(pdf)
}
func EncodePaymentSchedule(payment gofpdf.Pdf) {
	var buf bytes.Buffer
	err := payment.Output(&buf)

	if err != nil {
		log.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println(str)
	fn, err := os.Create("sm.pdf")
	dec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	num, err2 := fn.Write(dec)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(num)
	defer fn.Close()

}

/*func LoadCCData() {
	blockCodes := []rune{'H', 'P', 'J', 'Q', 'B', 'K', 'E'}
	bins := []string{"461786", "405081", "492518", "482217", "563719", "471247", "504018"}
	logos := []string{"371", "415", "304", "561", "461", "983", "861", "739", "688"}
	idApps := []string{"RS", "ZP"}
	idTXN := []string{"PCT", "URX", "PQA", "VPD", "VBC", "UBN", "PZT", "ZQG"}
	dm := []string{"04-07-10", "20-05-10", "26-03-08", "05-4-14", "26-05-07", "08-05-12",
		"07-02-16", "04-06-11", "06-02-13", "29-04-11"}
	makrs := []string{"swaminathan", "ramnarayan", "laxminarayan", "srinivas", "AUTO", "JK92", "LX71"}
	ckd := []string{"subbrao", "raghavendra", "vasantha", "P271", "AUTO", "k0281"}
	flgMnt := []rune{'A', 'Q'}
	mntAction := []rune{'A', 'P'}
}*/

/*
	func DecodePdf(arr []byte) {
		fn, err := os.Create("sm.pdf")
		if err != nil {
			log.Fatal(err)
		}
		defer fn.Close()

		buf := make([]byte, len(arr))
		fmt.Println("BUF Length :: ", len(buf))
		_, err1 := base64.StdEncoding.Decode(buf, arr)
		if err1 != nil {
			log.Fatal(err1)
		}
		fn.WriteString(string(buf))

}
*/
func UsingGoPdf() {
	//creating sample pdf file
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(40, 20, "Hello World")
	err := pdf.OutputFileAndClose("Hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

/*func Sample1() {
	//msg := "TestResourceDictInheritanceDemoPDF"
	xRefTable, err := pdf.CreateResourceDictInheritanceDemoXRef()
	if err != nil {
		log.Fatal(err)
	}

	api.CreatePDFFile(xRefTable, "sample.pdf", nil)
}*/

/*func Sample() {
	msg := "TestCreateDemoPDF"
	//mediaBox := types.NewStringSet([]string{"PRASSANNA","DIWADKAR"})
	p1 := model.Page{Buf: bytes.NewBufferString("HELLO")}

	pdf.CreateTestPageContent(p1)
	xRefTable, err := pdf.CreateDemoXRef()
	xRefTable.Author = "PRASANNA"
	if err != nil {
		log.Fatalf("%s: %v\n", msg, err)
	}
	rootDict, err := xRefTable.Catalog()

	if err != nil {
		log.Fatalf("%s: %v\n", msg, err)
	}
	if err = pdf.AddPageTreeWithSamplePage(xRefTable, rootDict, p1); err != nil {
		log.Fatalf("%s: %v\n", msg, err)
	}

	err2 := api.CreatePDFFile(xRefTable, "Sw.pdf", &model.Configuration{WriteObjectStream: true})

	if err2 != nil {
		log.Fatal(err2)
	}
}*/
