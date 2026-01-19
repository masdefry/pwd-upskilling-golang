Hello, Full Stack Web Development Lecturers✌️!

🧑‍💻 Bagaimana cara membuat project baru di Golang?

        Project di Golang biasa disebut dengan istilah `module`. Untuk membuat module, kita bisa menggunakan
        perintah berikut di directory tempat kita akan membuat module:

            ➡️ go mod init <module-name>

        Untuk mengeksekusi code yang telah kita buat, kita bisa meng-compile terlebih dahulu dengan command sebagai berikut:

            ➡️ go build

            ➡️ intro-golang.exe

        Cara mengeksekusi code diatas cukup rumit, karena setiap terjadi perubahan code__ kita diharuskan untuk meng-compile ulang code kita menjadi binary file. Untungnya__ Golang support untuk mengeksekusi code kita tanpa harus di compile terlebih dahulu:

            ➡️ go run <filename>

🧑‍💻 Core concept di Golang

        📝 Package

            Package adalah sebuah unit dasar yang digunakan untuk mengelompokan code (file, function, struct, variable, dll) yang saling berhubungan___ agar program:
            ▪️Lebih terstruktur
            ▪️Lebih mudah di maintain
            ▪️Reusable
            ▪️Menghindari penamaan yang bentrok

            1️⃣ Konsep dasar package
                Setiap file .go WAJIB punya deklarasi package di baris paling atas.

            2️⃣ Fungsi utama package
                ▪️Pengelompokan code
                ▪️Namespace (hindari konflik nama)

                    mathutils.Add()
                    mathutils.Subtract()

                  Tanpa package, fungsi Add() bisa bentrok dengan fungsi lain.

                ▪️Reusability
                    Package bisa dipakai di project lain, dengan cara:

                        import "github.com/user/mathutils"

                ▪️Access control (public & private)
                    Go mengatur akses berdasarkan huruf kapital:

                        func Hitung() {} // public (exported)
                        func hitung() {} // private (unexported)

            3️⃣ Aturan Penting Package

                ▪️Satu folder = satu package
                    Semua file .go dalam folder yang sama harus punya nama package yang sama.

                ▪️Nama package biasanya = nama folder
                    utils/ → package utils

                ▪️Package bisa di-alias
                    import u "project/utils"

        📝 Multiple Main Function

            ▪️Di Golang, function dalam module/project adalah unik, artinya kita tidak boleh membuat function dengan nama yang sama.
            ▪️Oleh karena itu, apabila kita membuat file baru (semisal, sample.go), lalu membuat nama function yang sama yaitu `funct main`, maka kita tidak dapat melakukan build module. Karena function tersebut dianggap duplicate dengan function lain yang ada di file `hello.world.go`.
