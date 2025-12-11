CREATE TABLE Type_Activities(
   Id_Type_Activities CHAR(50),
   Name VARCHAR(50) NOT NULL,
   Del BOOLEAN NOT NULL,
   PRIMARY KEY(Id_Type_Activities)
);

CREATE TABLE Permissions(
   Id_Permission VARCHAR(35),
   Hierarchical_Number BYTE NOT NULL,
   PRIMARY KEY(Id_Permission)
);

CREATE TABLE Activities(
   Id_Activities CHAR(50),
   Location_GPS TEXT NOT NULL,
   Name CHAR(150) NOT NULL,
   Type VARCHAR(50),
   Google_Rate DECIMAL(3,2),
   In_App_Rate DECIMAL(3,2),
   Del BOOLEAN NOT NULL,
   PRIMARY KEY(Id_Activities),
   UNIQUE(Location_GPS)
);

CREATE TABLE Users(
   Id_Users CHAR(50),
   Email CHAR(150) NOT NULL,
   Username CHAR(100) NOT NULL,
   Password VARCHAR(50) NOT NULL,
   Name CHAR(50) NOT NULL,
   Surname CHAR(100) NOT NULL,
   Birth_Date DATE,
   Nationality VARCHAR(100),
   Phone INT NOT NULL,
   Phone_Country_Code SMALLINT NOT NULL,
   Session_duration SMALLINT NOT NULL,
   Currency CHAR(5),
   Languages VARCHAR(50),
   Units VARCHAR(10),
   Theme VARCHAR(25),
   App_notification BOOLEAN NOT NULL,
   Email_notification BOOLEAN NOT NULL,
   Del BOOLEAN NOT NULL,
   Id_Permission VARCHAR(35) NOT NULL,
   PRIMARY KEY(Id_Users),
   UNIQUE(Email),
   UNIQUE(Username),
   FOREIGN KEY(Id_Permission) REFERENCES Permissions(Id_Permission)
);

CREATE TABLE Trips(
   Id_Trips CHAR(100),
   City CHAR(200) NOT NULL,
   Has_car BOOLEAN NOT NULL,
   Walking_distance BYTE NOT NULL,
   Is_Valid BOOLEAN NOT NULL,
   Start_Date DATE NOT NULL,
   End_Date DATE NOT NULL,
   Activities VARCHAR(300),
   Trip_Way TEXT NOT NULL,
   Del BOOLEAN NOT NULL,
   Id_Users CHAR(50) NOT NULL,
   PRIMARY KEY(Id_Trips),
   FOREIGN KEY(Id_Users) REFERENCES Users(Id_Users)
);

CREATE TABLE Trip_Has_Type_Activities(
   Id_Trips CHAR(100),
   Id_Type_Activities CHAR(50),
   PRIMARY KEY(Id_Trips, Id_Type_Activities),
   FOREIGN KEY(Id_Trips) REFERENCES Trips(Id_Trips),
   FOREIGN KEY(Id_Type_Activities) REFERENCES Type_Activities(Id_Type_Activities)
);

CREATE TABLE Have_Favorite_Activities(
   Id_Users CHAR(50),
   Id_Type_Activities CHAR(50),
   PRIMARY KEY(Id_Users, Id_Type_Activities),
   FOREIGN KEY(Id_Users) REFERENCES Users(Id_Users),
   FOREIGN KEY(Id_Type_Activities) REFERENCES Type_Activities(Id_Type_Activities)
);

CREATE TABLE Loved_Activities(
   Id_Users CHAR(50),
   Id_Activities CHAR(50),
   PRIMARY KEY(Id_Users, Id_Activities),
   FOREIGN KEY(Id_Users) REFERENCES Users(Id_Users),
   FOREIGN KEY(Id_Activities) REFERENCES Activities(Id_Activities)
);

CREATE TABLE Trips_Has_Activities(
   Id_Trips CHAR(100),
   Id_Activities CHAR(50),
   PRIMARY KEY(Id_Trips, Id_Activities),
   FOREIGN KEY(Id_Trips) REFERENCES Trips(Id_Trips),
   FOREIGN KEY(Id_Activities) REFERENCES Activities(Id_Activities)
);
