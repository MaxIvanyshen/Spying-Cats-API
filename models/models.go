package models

type Cat struct {
    Id int `gorm:"primaryKey"`
    Name string `json:"name"`
    YearsOfExperience int `json:"yearsOfExperience"`
    Breed string `json:"breed"`
    Salary int `json:"salary"`
}

type Mission struct {
    Id int `gorm:"primaryKey"`
    Cat *Cat `gorm:"foreignKey:Id" json:"cat"`
    Targets []*Target `gorm:"foreignKey:MissionID" json:"targets"`
    Complete bool `json:"complete"`
}

type Target struct {
    Id int `gorm:"primaryKey"`
    MissionID int
    Name string `json:"name"`
    Country string `json:"country"` 
    Notes []*Note `gorm:"foreignKey:TargetID" json:"notes"`
    Complete bool `json:"complete"`
}

type Note struct {
    Id     int    `gorm:"primaryKey"`
    TargetID int  
    Content string `json:"content"`
}


