package controllers

import (
	"barang-bekas/config"
	"barang-bekas/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

// Tampilkan semua produk
func Index(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.HTML(http.StatusOK, "index.html", gin.H{"products": products})
}

// Tampilkan form tambah
func New(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

// Proses tambah produk
func Create(c *gin.Context) {
	var product models.Product

	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))

	// Upload gambar
	file, err := c.FormFile("image")
	filename := ""
	if err == nil {
		filename = file.Filename
		c.SaveUploadedFile(file, "uploads/"+filename)
	}

	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	product.Price = price
	product.Stock = stock
	product.ImageURL = filename
	product.UserID = 1 // default sementara

	config.DB.Create(&product)
	c.Redirect(http.StatusFound, "/products")
}

// Form edit
func Edit(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	config.DB.First(&product, id)
	c.HTML(http.StatusOK, "form.html", gin.H{"product": product})
}

// Proses update produk
func Update(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	config.DB.First(&product, id)

	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))

	// Jika ada file baru di-upload, simpan
	file, err := c.FormFile("image")
	if err == nil {
		filename := file.Filename
		c.SaveUploadedFile(file, "uploads/"+filename)
		product.ImageURL = filename
	}

	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	product.Price = price
	product.Stock = stock

	config.DB.Save(&product)
	c.Redirect(http.StatusFound, "/products")
}

// Hapus produk
func Delete(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	config.DB.Delete(&product, id)
	c.Redirect(http.StatusFound, "/products")
}

// Export PDF
func ExportPDF(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Daftar Produk")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)

	for _, p := range products {
		pdf.Cell(40, 10, p.Name)
		pdf.Cell(40, 10, fmt.Sprintf("Rp %.0f", p.Price))
		pdf.Ln(10)
	}

	err := pdf.Output(c.Writer)
	if err != nil {
		c.String(500, "Gagal membuat PDF")
	}
}

// Export Excel
func ExportExcel(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	f := excelize.NewFile()
	sheet := "Produk"
	f.SetSheetName("Sheet1", sheet)

	f.SetCellValue(sheet, "A1", "Nama Produk")
	f.SetCellValue(sheet, "B1", "Harga")
	f.SetCellValue(sheet, "C1", "Stok")

	for i, p := range products {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), p.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.Price)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), p.Stock)
	}

	c.Header("Content-Disposition", "attachment; filename=produk.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	_ = f.Write(c.Writer)
}
