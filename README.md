#ini adalah Final Project 2 Golang Studi Independen, yaitu membuat rest API bernama MyGram, yang dimana pada aplikasi ini kita dapat menyimpan foto maupun membuat comment untuk foto orang lain. project ini mengunakan gorm, gin, dan juga database mysql.

##Sebelum menjalankan rest API-nya, file env.example harus direname dulu menjadi .env

###Untuk menggunakannya, bisa dengan aplikasi postman atau insomnia. Lalu menulisakan alamat atau url http://localhost:5000/

#Sebelum melakukan perintah CRUD, diharuskan untuk melakukan POST register terlebih dahulu pada url http://localhost:5000/users/register
-Saat melakukan register, diperlukan username, email, password, dan age yang harus diisi dengan form,
selain dengan form register juga bisa diisi menggunakan syntax json pada bagian body. Misalnya : 
{ 
    "username" : "Fajrian Nugraha", 
    "email" : "nugraha34@gmail.com", 
    "password" : "antetokounmpo" 
    "age" : 14
}

#setelah melakukan register, user harus melakukan login pada url atau http://localhost:5000/users/login di postman atau insomnia 

Untuk melakukan login cukup dengan mengisikan email dan password. Lalu token pada bagian response nya harus di-copy untuk melakukan perintah CRUD
#Untuk melakukan perintah pada user, seperti Update dan Delete user bisa dengan melakukan Request PUT pada url http://localhost:5000/update-user/:penggunaId
dan Request DELETE pada url http://localhost:5000/delete-user/:penggunaId

#Perintah yang bisa dilakukan, yaitu POST, PUT, GET, dan DELETE. Tetapi sebelum itu, pada bagian Header wajib ditambahkan "Authorization" dan memasukkan token yang telah didapat dari login tadi agar perintahnya berjalan.

// endpoint Photos

untuk melakukan POST atau menginsert Photo, bisa dengan url http://localhost:5000/photos/create-photo lalu memasukkan data seperti "title", "caption", "photo_url" dengan form, 
dan juga bisa input menggunakan syntax json pada bagian body, misalnya : 
{
    "title":"Photo Dufan", 
    "caption":"indah sekali", 
    "photo_url":"www.photo.com" 
}

untuk melakukan GET atau melihat photo, bisa dengan url http://localhost:5000/photos/

untuk melakukan PUT atau mengupdate photo, bisa dengan menambahkan photoId dibelakang url, misalnya untuk mengubah photo dengan id 3, maka urlnya http://localhost:5000/photos/update-photo/3

untuk melakukan DELETE atau menghapus photo bisa dengan menambahkan photoId dibelakang url, misalnya jika ingin menghapus photo dengan id 4, maka urlnya http://localhost:5000/photos/delete-photo/4


// endpoint Comment

untuk melakukan POST atau menginsert comment, bisa dengan url http://localhost:5000/comments/create-comment lalu memasukkan data seperti "message", "photo_id" dengan form, 
dan juga bisa input menggunakan syntax json pada bagian body, misalnya : 
{
    "message":"Gaya euy", 
    "photo_id": 1, 
}

untuk melakukan GET atau melihat comment, bisa dengan url http://localhost:5000/comments/

untuk melakukan PUT atau mengupdate comment, bisa dengan menambahkan commentId dibelakang url, misalnya untuk mengubah comment dengan id 3, maka urlnya http://localhost:5000/comments/update-comment/3

untuk melakukan DELETE atau menghapus comment bisa dengan menambahkan commentId dibelakang url, misalnya jika ingin menghapus comment dengan id 4, maka urlnya http://localhost:5000/comments/delete-comment/4


// endpoint socialmedia

untuk melakukan POST atau menginsert  socialmedia, bisa dengan url http://localhost:5000/socialmedias/create-sosmed lalu memasukkan data seperti "name", "social_media_url" dengan form, 
dan juga bisa input menggunakan syntax json pada bagian body, misalnya : 
{
    "name":"@fajrian_nu1", 
    "social_media_url": "www. instagram.com", 
}

untuk melakukan GET atau melihat socialmedia, bisa dengan url http://localhost:5000/socialmedias/

untuk melakukan PUT atau mengupdate socialmedias, bisa dengan menambahkan socialMediaId dibelakang url, misalnya untuk mengubah socialmedia dengan id 3, maka urlnya http://localhost:5000/socialmedias/update-sosmed/3

untuk melakukan DELETE atau menghapus socialmedias bisa dengan menambahkan socialMediaId dibelakang url, misalnya jika ingin menghapus socialmedia dengan id 4, maka urlnya http://localhost:5000/socialmedias/delete-sosmed/4




