package controllers

import (
	"../structs"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	"toko-ijah/helpers"
	"github.com/dustin/go-humanize"
)

func (idb *InDB) GetBarang(c *gin.Context) {
	var (
		listbarang[] structs.Barang

		result structs.Result
	)

	idb.DB.Find(&listbarang)

	if len(listbarang) <= 0 {
		result.Message = "failed"
		result.Data = nil
	} else {
		result.Message = "success"
		result.Data = listbarang
	}


	c.JSON(http.StatusOK, result)
}


func (idb *InDB) BarangMasuk(c *gin.Context) {
	var (
		Barang structs.Barang
		//New_barang structs.Barang
		datapembelian structs.Pembelian
		newstok structs.Barang
		result gin.H
	)

	//id := c.Query("id")
	c.BindJSON(&datapembelian)

	//get Request json
	sku := datapembelian.Sku
	nama_barang := datapembelian.Nama_barang
	qty_diterima := datapembelian.Qty_diterima
	harga_beli := datapembelian.Harga_beli
	qty := datapembelian.Qty
	created_date := time.Now()
	created_at := created_date.Format("2006-01-02 15:04:05")
	note := datapembelian.Note
	invoice := datapembelian.Invoice_pembelian

	err := idb.DB.Where("sku = ?", sku).First(&Barang).Error
	if err != nil {
		//insert data ke tbl barang
		Barang.Sku = sku
		Barang.Nama_barang = nama_barang
		Barang.Stok = qty_diterima
		Barang.Created_date = created_at

		idb.DB.Create(&Barang)

		//insert data ke table pembelian

		datapembelian.Sku = sku
		datapembelian.Nama_barang = nama_barang
		datapembelian.Qty = qty
		datapembelian.Qty_diterima = qty_diterima
		datapembelian.Harga_beli = harga_beli
		datapembelian.Total = qty * harga_beli
		datapembelian.Created_date = created_at
		datapembelian.Note = note
		datapembelian.Invoice_pembelian = invoice

		idb.DB.Create(&datapembelian)
		result = gin.H{
			"result": "Pembelian Barang Berhasil",
		}
	}else{

		//insert data ke table pembelian

		datapembelian.Sku = sku
		datapembelian.Nama_barang = nama_barang
		datapembelian.Qty = qty
		datapembelian.Qty_diterima = qty_diterima
		datapembelian.Harga_beli = harga_beli
		datapembelian.Total = qty * harga_beli
		datapembelian.Created_date = created_at
		datapembelian.Note = note
		datapembelian.Invoice_pembelian = invoice

		idb.DB.Create(&datapembelian)

		//insert data ke tbl barang
		newstok.Stok = Barang.Stok + qty_diterima
		newstok.Updated_date = created_at

		err = idb.DB.Model(&Barang).Updates(newstok).Error
		if err != nil {
			result = gin.H{
				"result": "update failed",
			}
		} else {
			result = gin.H{
				"result": "Pembelian Dan Penambahan Barang Berhasil",
			}
		}
	}


	c.JSON(http.StatusOK, result)
}

func (idb *InDB) BarangKeluar(c *gin.Context) {
	var (
		Barang structs.Barang
		Penjualan structs.Penjualan
		Request structs.Penjualan
		newstok structs.Barang
		result gin.H
	)

	//id := c.Query("id")
	c.BindJSON(&Request)

	//get Request json
	sku := Request.Sku

	qty := Request.Qty
	harga_jual := Request.Harga_jual
	note := Request.Note
	invoice := Request.Invoice_penjualan
	created_date := time.Now()
	created_at := created_date.Format("2006-01-02 15:04:05")

	err := idb.DB.Where("sku = ?", sku).First(&Barang).Error
	if err != nil {

	}else{

		Penjualan.Sku = Barang.Sku
		Penjualan.Nama_barang = Barang.Nama_barang
		Penjualan.Qty = qty
		Penjualan.Harga_jual = harga_jual
		Penjualan.Total = qty * harga_jual
		Penjualan.Created_at = created_at
		Penjualan.Note = note
		Penjualan.Invoice_penjualan = invoice
		idb.DB.Create(&Penjualan)
		result = gin.H{
			"result": "Penjualan Barang Berhasil",
		}

		newstok.Stok = Barang.Stok - qty
		newstok.Updated_date = created_at

		err = idb.DB.Model(&Barang).Updates(newstok).Error
		if err != nil {
			result = gin.H{
				"result": "update failed",
			}
		} else {
			result = gin.H{
				"result": "Penjualan Dan Pengurangan stok Barang Berhasil",
			}
		}

	}


	c.JSON(http.StatusOK, result)
}

func (idb *InDB) NilaiBarang(c *gin.Context) {
	//file, _ := os.Create("path/file/sales.csv")
	var searchkeytemp structs.Searchkeytemp
	var omzet,laba string
	var start,end string
	//var Rekap structs.Rekapitulasi_barang
	c.BindJSON(&searchkeytemp)
	created_date := time.Now()
	created_at := created_date.Format("2006-01-02 15:04:05")



	var data = [][]string{{"LAPORAN NILAI BARANG"},{""},
		{"Tanggal Cetak",created_at},
		{"Jumlah SKU",start+ " sd "+ end },
		{"Jumlah Total Barang", omzet},
		{"Total Nilai ", laba},
		{""},{"SKU", "Nama Barang","Jumlah","Rata Rata Harga Beli"}}
	var nilai_barang structs.NilaiBarang
	file, err := os.Create("NilaiBarang.csv")
	helper.CheckError("Cannot create file", err)
	defer file.Close()

	println("SELECT bs.sku,bs.nama_barang,bs.stok,(sum(pb.harga_beli) / stok ) as rata, (sum(pb.harga_beli) / stok ) * bs.stok as total " +
		"FROM barangs bs inner join pembelians pb on bs.sku = bs.sku")

	rows, err := idb.DB.Raw("SELECT bs.sku,bs.nama_barang,bs.stok,(sum(pb.harga_beli) / stok ) as rata, (sum(pb.harga_beli) / stok ) * bs.stok as total " +
		"FROM barangs bs inner join pembelians pb on bs.sku = bs.sku").Rows() // (*sql.Rows, error)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nilai_barang.Sku,&nilai_barang.Nama_barang,&nilai_barang.Stok,&nilai_barang.Rata,&nilai_barang.Total)
		data = append(data, []string{
			nilai_barang.Sku,
			nilai_barang.Nama_barang,
			nilai_barang.Stok,
			humanize.Comma(nilai_barang.Rata),
			humanize.Comma(nilai_barang.Total),
		})
	}


	writer := csv.NewWriter(file)
	writer.WriteAll(data)
	defer writer.Flush()


}

func (idb *InDB) LaporanPenjualan(c *gin.Context) {

	var searchkeytemp structs.Searchkeytemp
	var omzet,laba string
	var start string
	var Rekap structs.Rekapitulasi_penjualan

	c.BindJSON(&searchkeytemp)
	created_date := time.Now()
	created_at := created_date.Format("2006-01-02 15:04:05")

	start = searchkeytemp.Start_date
	end := searchkeytemp.End_date

	if(len(start) == 0){
		start = created_date.Format("2006-01-02")+ " 00:00:00"
	}

	if(len(end) == 0){
		end = created_date.Format("2006-01-02")+ " 23:59:59"
	}


	rows1, err1 := idb.DB.Raw("SELECT sum((pj.harga_jual - pb.harga_beli)* pj.qty) as laba,sum(pj.total)as omzet  FROM penjualans pj inner join pembelians pb on pj.sku = pb.sku where pj.created_at >= '2019-08-10 00:00:00' and created_at <= '2019-08-10 23:59:59' GROUP by invoice_penjualan").Rows() // (*sql.Rows, error)
	defer rows1.Close()
	for rows1.Next() {
		if err1 == nil {
			rows1.Scan(&Rekap.Laba_kotor, &Rekap.Omzet)
			omzet = humanize.Comma(Rekap.Omzet)
			laba = humanize.Comma(Rekap.Laba_kotor)
		}
	}


	var data = [][]string{{"LAPORAN PENJUALAN"},{""},
		{"Tanggal Cetak",created_at},
		{"Tanggal",start+ " sd "+ end },
		{"Total Omzet", omzet},
		{"Total Laba Kotor", laba},
		{""},
		{"ID Pesanan", "Waktu","Sku","Nama Barang","Jumlah", "Harga Jual","Total","Harga Beli","Laba"}}
	var laporan structs.LaporanPenjualan
	file, err := os.Create("LaporanPenjualan.csv")
	helper.CheckError("Cannot create file", err)
	defer file.Close()

	println("SELECT invoice_penjualan,pj.created_at,pj.sku,pj.nama_barang,pj.qty,pj.harga_jual,pj.total,pb.harga_beli,(pj.harga_jual - pb.harga_beli) as laba" +
		" FROM penjualans pj inner join pembelians pb on pj.sku = pb.sku where pj.created_at >= '"+start+"' and created_at <= '"+end+"' GROUP by invoice_penjualan")

	rows, err := idb.DB.Raw("SELECT invoice_penjualan,pj.created_at,pj.sku,pj.nama_barang,pj.qty,pj.harga_jual,pj.total,pb.harga_beli,(pj.harga_jual - pb.harga_beli) as laba" +
		" FROM penjualans pj inner join pembelians pb on pj.sku = pb.sku where pj.created_at >= '"+start+"' and created_at <= '"+end+"' GROUP by invoice_penjualan").Rows() // (*sql.Rows, error)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&laporan.Invoice_penjualan,&laporan.Created_at,&laporan.Sku,&laporan.Nama_barang,&laporan.Qty,&laporan.Harga_jual,&laporan.Total,&laporan.Harga_beli,&laporan.Laba)
		data = append(data, []string{
			laporan.Invoice_penjualan,
			laporan.Created_at,
			laporan.Sku,
			laporan.Nama_barang,
			laporan.Qty,
			humanize.Comma(laporan.Harga_jual),
			humanize.Comma(laporan.Total),
			humanize.Comma(laporan.Harga_beli),
			humanize.Comma(laporan.Laba),

		})

	}

	writer := csv.NewWriter(file)
	writer.WriteAll(data)
	defer writer.Flush()


}