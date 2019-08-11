package structs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Barang struct {
	gorm.Model
	Id int `json:"id"`
	Sku string `json:"sku"`
	Nama_barang string `json:"nama_barang"`
	Stok int `json:"stok"`
	Created_date string `json:"created_date"`
	Updated_date string `json:"updated_date"`
}

type Pembelian struct {
	gorm.Model
	Id int `json:"id"`
	Sku string `json:"sku"`
	Nama_barang string `json:"nama_barang"`
	Qty int `json:"qty"`
	Qty_diterima int `json:"qty_diterima"`
	Harga_beli int `json:"harga_beli"`
	Total int `json:"total"`
	Note string `json:"Note"`
	Invoice_pembelian string `json:"invoice_pembelian"`
	Created_date string `json:"created_date"`
}

type Penjualan struct {
	gorm.Model
	Id int `json:"id"`
	Sku string `json:"sku"`
	Nama_barang string `json:"nama_barang"`
	Qty int `json:"qty"`
	Harga_jual int `json:"harga_jual"`
	Total int `json:"total"`
	Note string `json:"Note"`
	Invoice_penjualan string `json:"invoice_penjualan"`
	Created_at string `json:"created_date"`
}

type NilaiBarang struct {
	gorm.Model
	Sku string `json:"sku"`
	Nama_barang string `json:"nama_barang"`
	Stok string `json:"stok"`
	Rata int64 `json:"rata"`
	Total int64 `json:"total"`
}

type LaporanPenjualan struct {
	gorm.Model
	Invoice_penjualan string `json:"id_pesanan"`
	Created_at string `json:"created_at"`
	Sku string `json:"sku"`
	Nama_barang string `json:"nama_barang"`
	Qty string `json:"jumlah"`
	Harga_jual int64 `json:"harga_jual"`
	Harga_beli int64 `json:"harga_beli"`
	Total int64 `json:"total"`
	Laba int64 `json:"laba"`
}

type Result struct {
	gorm.Model
	Message interface{} `json:"message"`
	Data interface{} `json:"data"`
}

type Searchkeytemp struct {
	Start_date string
	End_date string
}

type Rekapitulasi_penjualan struct {
	gorm.Model
	Omzet int64
	Laba_kotor int64
}

type Rekapitulasi_barang struct {
	gorm.Model
	Jml_sku string
	Jml_barang string
	Total_nilai int64

}


