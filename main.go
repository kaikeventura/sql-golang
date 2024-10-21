package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var fake = faker.New()

// BaseEntity adiciona os timestamps
type BaseEntity struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// FiscalReceipt representa o recibo fiscal
type FiscalReceipt struct {
	BaseEntity
	ID     uuid.UUID `gorm:"type:char(36);primaryKey"`
	State  string    `gorm:"default:OPEN"`
	Source string    `gorm:"default:ANYWHERE"`
	Target string    `gorm:"default:SOMETHING"`
	Amount int64
	Items  []Item `gorm:"foreignKey:FiscalReceiptID"`
}

// Função para criar um FiscalReceipt aleatório
func aFiscalReceipt() *FiscalReceipt {
	return &FiscalReceipt{
		ID:     uuid.New(),
		State:  "OPEN",
		Amount: rand.Int63n(999999) + 1, // Gera um valor entre 1 e 999999
	}
}

// Item representa um item associado a um recibo
type Item struct {
	BaseEntity
	ID               uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price            int64
	Amount           int
	FiscalReceiptID  uuid.UUID       `gorm:"type:char(36)"`
	BookItemID       *uuid.UUID      `gorm:"type:char(36)"`
	BookItem         *BookItem       `gorm:"foreignKey:BookItemID"`
	ChemicalItemID   *uuid.UUID      `gorm:"type:char(36)"`
	ChemicalItem     *ChemicalItem   `gorm:"foreignKey:ChemicalItemID"`
	ClothingItemID   *uuid.UUID      `gorm:"type:char(36)"`
	ClothingItem     *ClothingItem   `gorm:"foreignKey:ClothingItemID"`
	DrinkItemID      *uuid.UUID      `gorm:"type:char(36)"`
	DrinkItem        *DrinkItem      `gorm:"foreignKey:DrinkItemID"`
	ElectronicItemID *uuid.UUID      `gorm:"type:char(36)"`
	ElectronicItem   *ElectronicItem `gorm:"foreignKey:ElectronicItemID"`
	FilmItemID       *uuid.UUID      `gorm:"type:char(36)"`
	FilmItem         *FilmItem       `gorm:"foreignKey:FilmItemID"`
	FitnessItemID    *uuid.UUID      `gorm:"type:char(36)"`
	FitnessItem      *FitnessItem    `gorm:"foreignKey:FitnessItemID"`
	FoodItemID       *uuid.UUID      `gorm:"type:char(36)"`
	FoodItem         *FoodItem       `gorm:"foreignKey:FoodItemID"`
	GameItemID       *uuid.UUID      `gorm:"type:char(36)"`
	GameItem         *GameItem       `gorm:"foreignKey:GameItemID"`
	HardwareItemID   *uuid.UUID      `gorm:"type:char(36)"`
	HardwareItem     *HardwareItem   `gorm:"foreignKey:HardwareItemID"`
	HygieneItemID    *uuid.UUID      `gorm:"type:char(36)"`
	HygieneItem      *HygieneItem    `gorm:"foreignKey:HygieneItemID"`
	MusicItemID      *uuid.UUID      `gorm:"type:char(36)"`
	MusicItem        *MusicItem      `gorm:"foreignKey:MusicItemID"`
	SoftwareItemID   *uuid.UUID      `gorm:"type:char(36)"`
	SoftwareItem     *SoftwareItem   `gorm:"foreignKey:SoftwareItemID"`
	ToolItemID       *uuid.UUID      `gorm:"type:char(36)"`
	ToolItem         *ToolItem       `gorm:"foreignKey:ToolItemID"`
	VehicleItemID    *uuid.UUID      `gorm:"type:char(36)"`
	VehicleItem      *VehicleItem    `gorm:"foreignKey:VehicleItemID"`
	ToyItemID        *uuid.UUID      `gorm:"type:char(36)"`
	ToyItem          *ToyItem        `gorm:"foreignKey:ToyItemID"`
}

func NewItem(price int64, amount int, fiscalReceipt *FiscalReceipt, any interface{}) *Item {
	item := Item{
		ID:              uuid.New(),
		Price:           price,
		Amount:          amount,
		FiscalReceiptID: fiscalReceipt.ID,
	}

	switch v := any.(type) {
	case BookItem:
		item.BookItemID = &v.ID
	case ChemicalItem:
		item.ChemicalItemID = &v.ID
	case ClothingItem:
		item.ClothingItemID = &v.ID
	case DrinkItem:
		item.DrinkItemID = &v.ID
	case ElectronicItem:
		item.ElectronicItemID = &v.ID
	case FilmItem:
		item.FilmItemID = &v.ID
	case FitnessItem:
		item.FitnessItemID = &v.ID
	case FoodItem:
		item.FoodItemID = &v.ID
	case GameItem:
		item.GameItemID = &v.ID
	case HardwareItem:
		item.HardwareItemID = &v.ID
	case HygieneItem:
		item.HygieneItemID = &v.ID
	case MusicItem:
		item.MusicItemID = &v.ID
	case SoftwareItem:
		item.SoftwareItemID = &v.ID
	case ToolItem:
		item.ToolItemID = &v.ID
	case VehicleItem:
		item.VehicleItemID = &v.ID
	case ToyItem:
		item.ToyItemID = &v.ID
	default:
		fmt.Println("Tipo não reconhecido")
	}

	return &item
}

type BookItem struct {
	BaseEntity
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price     int64
	Name      string
	Genre     string
	Author    string
	Publisher string
}

func NewBookItem(price int64) *BookItem {
	return &BookItem{
		ID:        uuid.New(),
		Price:     price,
		Name:      fake.Lorem().Word(),
		Genre:     fake.Lorem().Word(),
		Author:    fake.Lorem().Word(),
		Publisher: fake.Lorem().Word(),
	}
}

type ChemicalItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
}

func NewChemicalItem(price int64) *ChemicalItem {
	return &ChemicalItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.App().Name(),
	}
}

type ClothingItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
}

func NewClothingItem(price int64) *ClothingItem {
	return &ClothingItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.Person().LastName(),
	}
}

type DrinkItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
	Style string
}

func NewDrinkItem(price int64) *DrinkItem {
	return &DrinkItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.Beer().Name(),
		Style: fake.Beer().Malt(),
	}
}

type ElectronicItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
	Type  string
}

func NewElectronicItem(price int64) *ElectronicItem {
	return &ElectronicItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.Lorem().Word(),
		Type:  fake.Lorem().Word(),
	}
}

type FilmItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
	Time  string
}

func NewFilmItem(price int64) *FilmItem {
	return &FilmItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.Lorem().Word(),
		Time:  fake.App().Version(),
	}
}

type FitnessItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
	Team  string
}

func NewFitnessItem(price int64) *FitnessItem {
	return &FitnessItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.Lorem().Word(),
		Team:  fake.Lorem().Word(),
	}
}

type FoodItem struct {
	BaseEntity
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price       int64
	Name        string
	Ingredient  string
	Measurement string
}

func NewFoodItem(price int64) *FoodItem {
	return &FoodItem{
		ID:          uuid.New(),
		Price:       price,
		Name:        fake.Food().Fruit(),
		Ingredient:  fake.Food().Vegetable(),
		Measurement: fake.Food().Fruit(),
	}
}

type GameItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
}

func NewGameItem(price int64) *GameItem {
	return &GameItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.Currency().Currency(),
	}
}

type HardwareItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
	Name  string
	Type  string
}

func NewHardwareItem(price int64) *HardwareItem {
	return &HardwareItem{
		ID:    uuid.New(),
		Price: price,
		Name:  fake.Lorem().Text(15),
		Type:  fake.Lorem().Text(10),
	}
}

type HygieneItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
}

func NewHygieneItem(price int64) *HygieneItem {
	return &HygieneItem{
		ID:    uuid.New(),
		Price: price,
	}
}

type MusicItem struct {
	BaseEntity
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price      int64
	Genre      string
	Instrument string
}

func NewMusicItem(price int64) *MusicItem {
	return &MusicItem{
		ID:         uuid.New(),
		Price:      price,
		Genre:      fake.Music().Genre(),
		Instrument: fake.Music().Name(),
	}
}

type ToolItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
}

func NewToolItem(price int64) *ToolItem {
	return &ToolItem{
		ID:    uuid.New(),
		Price: price,
	}
}

type SoftwareItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
}

func NewSoftwareItem(price int64) *SoftwareItem {
	return &SoftwareItem{
		ID:    uuid.New(),
		Price: price,
	}
}

type ToyItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
}

func NewToyItem(price int64) *ToyItem {
	return &ToyItem{
		ID:    uuid.New(),
		Price: price,
	}
}

type VehicleItem struct {
	BaseEntity
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Price int64
}

func NewVehicleItem(price int64) *VehicleItem {
	return &VehicleItem{
		ID:    uuid.New(),
		Price: price,
	}
}

func createRandomList(size int) []string {
	itemTypes := []string{
		"BookItem",
		"ChemicalItem",
		"ClothingItem",
		"DrinkItem",
		"ElectronicItem",
		"FilmItem",
		"FitnessItem",
		"FoodItem",
		"GameItem",
		"HardwareItem",
		"HygieneItem",
		"MusicItem",
		"SoftwareItem",
		"ToolItem",
		"ToyItem",
		"VehicleItem",
	}

	rand.Seed(time.Now().UnixNano())

	randomItems := make([]string, size)

	for i := 0; i < size; i++ {
		randomItems[i] = itemTypes[rand.Intn(len(itemTypes))]
	}

	return randomItems
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to database:", err)
		return
	}

	db.AutoMigrate(&FiscalReceipt{}, &Item{}, &BookItem{}, &ChemicalItem{}, &ClothingItem{}, &DrinkItem{}, &ElectronicItem{}, &FilmItem{}, &FitnessItem{}, &FoodItem{}, &GameItem{}, &HardwareItem{}, &HygieneItem{}, &MusicItem{}, &SoftwareItem{}, &ToolItem{}, &ToyItem{}, &VehicleItem{})

	Courotine(db)
}

func Courotine(db *gorm.DB) {
	semaphore := make(chan struct{}, 1)
	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			for i := 0; i < 2; i++ {
				CreateFiscalReceipt(db)
			}
		}()
	}

	wg.Wait()
}

func CreateFiscalReceipt(db *gorm.DB) {
	fiscalReceipt := aFiscalReceipt()
	db.Create(&fiscalReceipt)

	itemTypes := createRandomList(rand.Intn(250) + 1)

	semaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for i := 0; i < len(itemTypes); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			for _, itemType := range itemTypes {
				price := rand.Int63n(999999) + 1
				amount := rand.Intn(50) + 1

				switch itemType {
				case "BookItem":
					book := NewBookItem(price)
					db.Create(&book)
					db.Create(NewItem(price, amount, fiscalReceipt, *book))

				case "ChemicalItem":
					chemical := NewChemicalItem(price)
					db.Create(&chemical)
					db.Create(NewItem(price, amount, fiscalReceipt, *chemical))

				case "ClothingItem":
					clothing := NewClothingItem(price)
					db.Create(&clothing)
					db.Create(NewItem(price, amount, fiscalReceipt, *clothing))

				case "DrinkItem":
					drink := NewDrinkItem(price)
					db.Create(&drink)
					db.Create(NewItem(price, amount, fiscalReceipt, *drink))

				case "ElectronicItem":
					electronic := NewElectronicItem(price)
					db.Create(&electronic)
					db.Create(NewItem(price, amount, fiscalReceipt, *electronic))

				case "FilmItem":
					film := NewFilmItem(price)
					db.Create(&film)
					db.Create(NewItem(price, amount, fiscalReceipt, *film))

				case "FitnessItem":
					fitness := NewFitnessItem(price)
					db.Create(&fitness)
					db.Create(NewItem(price, amount, fiscalReceipt, *fitness))

				case "FoodItem":
					food := NewFoodItem(price)
					db.Create(&food)
					db.Create(NewItem(price, amount, fiscalReceipt, *food))

				case "GameItem":
					game := NewGameItem(price)
					db.Create(&game)
					db.Create(NewItem(price, amount, fiscalReceipt, *game))

				case "HardwareItem":
					hardware := NewHardwareItem(price)
					db.Create(&hardware)
					db.Create(NewItem(price, amount, fiscalReceipt, *hardware))

				case "HygieneItem":
					hygiene := NewHygieneItem(price)
					db.Create(&hygiene)
					db.Create(NewItem(price, amount, fiscalReceipt, *hygiene))

				case "MusicItem":
					music := NewMusicItem(price)
					db.Create(&music)
					db.Create(NewItem(price, amount, fiscalReceipt, *music))

				case "SoftwareItem":
					software := NewSoftwareItem(price)
					db.Create(&software)
					db.Create(NewItem(price, amount, fiscalReceipt, *software))

				case "ToolItem":
					tool := NewToolItem(price)
					db.Create(&tool)
					db.Create(NewItem(price, amount, fiscalReceipt, *tool))

				case "ToyItem":
					toy := NewToyItem(price)
					db.Create(&toy)
					db.Create(NewItem(price, amount, fiscalReceipt, *toy))

				case "VehicleItem":
					vehicle := NewVehicleItem(price)
					db.Create(&vehicle)
					db.Create(NewItem(price, amount, fiscalReceipt, *vehicle))
				}
			}
		}()
	}

	wg.Wait()
}
