package utils

import (
	"artisancraft-platform/models"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	
	// Get database configuration from environment
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}
	
	// Force use the new database name
	dbName := "artisancraft.db"
	
	// Initialize database based on type
	if dbType == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	} else {
		// For future PostgreSQL support
		log.Fatal("Only SQLite is currently supported")
	}
	
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	err = DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Seed initial data
	seedData()
}

func seedData() {
	// Check if users already exist
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		// Create demo user
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("artisan123"), bcrypt.DefaultCost)
		demoUser := models.User{
			Username: "admin",
			Password: string(hashedPassword),
		}
		DB.Create(&demoUser)
		log.Println("Demo user created: admin / artisan123")
	}

	// Always recreate items for fresh data
	DB.Exec("DELETE FROM items")
	
	// Seed items
	items := []models.Item{
		{
			Name:        "Vintage Leather Backpack",
			Description: "Handcrafted full-grain leather backpack with brass hardware and adjustable straps for everyday adventures",
			Price:       189.99,
			Category:    "Fashion & Accessories",
			Rating:      4.8,
			Reviews:     1247,
			Image:       "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?w=400&h=300&fit=crop",
			InStock:     true,
		},
		{
			Name:        "Organic Herbal Tea Collection",
			Description: "Premium loose-leaf tea blend with chamomile, lavender, and mint for relaxation and wellness",
			Price:       24.99,
			Category:    "Food & Beverages",
			Rating:      4.9,
			Reviews:     2156,
			Image:       "https://images.unsplash.com/photo-1544787219-7f47ccb76574?w=400&h=300&fit=crop",
			InStock:     true,
		},
		{
			Name:        "Handmade Ceramic Plant Pot",
			Description: "Artisan-crafted ceramic pot with drainage hole, perfect for succulents and small houseplants",
			Price:       39.99,
			Category:    "Home & Garden",
			Rating:      4.7,
			Reviews:     892,
			Image:       "https://images.unsplash.com/photo-1485955900006-10f4d324d411?w=400&h=300&fit=crop",
			InStock:     true,
		},
		{
			Name:        "Yoga Mat Premium",
			Description: "Non-slip eco-friendly yoga mat with alignment lines and carrying strap for your practice",
			Price:       79.99,
			Category:    "Sports & Fitness",
			Rating:      4.8,
			Reviews:     3421,
			Image:       "https://images.unsplash.com/photo-1544367567-0f2fcb009e0b?w=400&h=300&fit=crop",
			InStock:     true,
		},
		{
			Name:        "Wireless Bluetooth Speaker",
			Description: "Portable waterproof speaker with 360-degree sound and 20-hour battery life for outdoor adventures",
			Price:       129.99,
			Category:    "Electronics",
			Rating:      4.6,
			Reviews:     678,
			Image:       "https://images.unsplash.com/photo-1608043152269-423dbba4e7e1?w=400&h=300&fit=crop",
			InStock:     true,
		},
		{
			Name:        "Handwoven Cotton Throw Blanket",
			Description: "Soft, lightweight throw blanket made from organic cotton, perfect for cozy evenings",
			Price:       89.99,
			Category:    "Home & Garden",
			Rating:      4.9,
			Reviews:     945,
			Image:       "https://images.unsplash.com/photo-1522771739844-6a9f6d5f14af?w=400&h=300&fit=crop",
			InStock:     true,
		},
		{
			Name:        "Artisan Soap Collection",
			Description: "Handmade natural soaps with essential oils and organic ingredients for gentle skin care",
			Price:       34.99,
			Category:    "Beauty & Wellness",
			Rating:      4.7,
			Reviews:     1567,
			Image:       "https://images.unsplash.com/photo-1556228720-195a672e8a03?w=400&h=300&fit=crop",
			InStock:     true,
		},
		{
			Name:        "Bamboo Cutting Board Set",
			Description: "Sustainable bamboo cutting boards in three sizes with juice groove and non-slip feet",
			Price:       59.99,
			Category:    "Home & Garden",
			Rating:      4.8,
			Reviews:     1123,
			Image:       "https://images.unsplash.com/photo-1582735689369-4fe89db7110c?w=400&h=300&fit=crop",
			InStock:     true,
		},
	}

	for _, item := range items {
		DB.Create(&item)
	}

	log.Println("ArtisanCraft database seeded successfully")
} 