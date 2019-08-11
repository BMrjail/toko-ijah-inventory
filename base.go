package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"


)



func main() {

	db := config.DBInit()

	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/barang", inDB.GetBarang)
	//router.POST("/pembelian_barang", inDB.InputBarang)
	router.POST("/barang_masuk", inDB.BarangMasuk)
	router.POST("/barang_keluar", inDB.BarangKeluar)
	router.GET("/nilai_barang", inDB.NilaiBarang)
	router.POST("/laporan_penjualan", inDB.LaporanPenjualan)



	router.Run(":7777")

}
