Project name : DaCol (Data Collection)

DB:

-  Product :
   -  name
   -  qty
   -  price
-  User :
   -  email
   -  password
-  Selling :
   -  id_product
   -  qty
   -  date

Usange:

-  Backend : Go
-  Frontend : React

Route :

-  Auth :

-  / (main dashboard) (GET)
   Data need : Product[]
-  /detail/:id (GET) :
   Product and Selling

-  /add (add to database) (POST)
-  /delete/:id (delete by id) (DELETE)
-  /logout (DELETE)

Not Auth :

-  /login (login page) (POST)
-  /register (register page) (POST)
