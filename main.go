package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"

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

func main() {

	CreatePaymentSchedule()

}

func GetHeader() []string {

	return []string{"Loan Number", "Loan Booked Date",
		"Loan Type", "Principal Amount(Rs.)", "Interest Rate(%)", "Tenure(months)", "Outstanding Principal(Rs.)"}
}

func CreatePaymentSchedule() {
	marginX := 10.0
	marginY := 20.0
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginX, marginY, marginX)
	pdf.AddPage()
	pdf.SetAuthor("HDFC Ccredit Card Division", true)
	hdfcLogoPath := "D:\\Work\\GoWorkspace\\DB4\\HDFCLogo.png"
	pdf.ImageOptions(hdfcLogoPath, 10, 10, 40, 10, false, gofpdf.ImageOptions{}, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(50, 50, "Loan EMI Table")

	pdf.Ln(-1)
	pdf.SetFont("Arial", "B", 5)
	for _, val := range GetHeader() {
		pdf.CellFormat(25, 8, val, "1", 0, "CM", false, 0, "")

	}
	pdf.Ln(-1)
	//add data
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
	//err := pdf.OutputFileAndClose("Payment.pdf")
	EncodePaymentSchedule(pdf)
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
