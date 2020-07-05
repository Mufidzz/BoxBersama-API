package structs

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title        string
	Description  string `gorm:"type:mediumtext"`
	NeededAmount int64
	//Association
	ArticleImages []ArticleImage `gorm:"foreignkey:ArticleID;association_foreignkey:ID"`
	Donor         []DonorMin     `gorm:"foreignkey:ArticleID;association_foreignkey:ID"`
}

type ArticleImage struct {
	gorm.Model
	Reference string
	//Key
	ArticleID int
}

type Donor struct {
	gorm.Model
	Name           string
	Email          string
	DonationAmount int64
	PhoneNumber    string
	Motivation     string
	UniqueCode     int
	Status         int `gorm:"default:'0'"`
	//Key
	ArticleID     int
	BankAccountID int
	//Association
	Article     Article     `gorm:"foreignkey:ID;association_foreignkey:ArticleID;Preload:false"`
	BankAccount BankAccount `gorm:"foreignkey:ID;association_foreignkey:BankAccountID"`
}

type BankAccount struct {
	gorm.Model
	AccountNumber string
	AccountName   string
	//Key
	BankID int
	//Association
	Bank Bank `gorm:"foreignkey:ID;association_foreignkey:BankID"`
}

type Bank struct {
	gorm.Model
	Name  string
	Image string
}

type Request struct {
	gorm.Model
	Proposer
	Recipient
	//Association
	Images []RequestImage `gorm:"foreignkey:RequestID;association_foreignkey:ID"`
}

type Proposer struct {
	ProposerName        string
	ProposerEmail       string
	ProposerPhoneNumber string
}

type Recipient struct {
	RecipientName        string
	RecipientAddress     string
	RecipientDescription string
}

type RequestImage struct {
	gorm.Model
	Reference string
	//Key
	RequestID int
}

type User struct {
	gorm.Model
	Username string
	Password string
}
