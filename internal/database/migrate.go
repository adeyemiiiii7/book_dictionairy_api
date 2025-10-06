package database

import (
	"example/go_api_tutorial/internal/models"
	"log"
)

// Migrate runs database migrations
func Migrate() error {
	log.Println("Running database migrations...")
	
	// Auto-migrate the schema
	err := DB.AutoMigrate(
		&models.User{},
		&models.Book{},
	)
	
	if err != nil {
		return err
	}
	
	log.Println("Database migrations completed successfully!")
	return nil
}

// SeedData inserts initial data (for development)
func SeedData() error {
	log.Println("Seeding initial data...")
	
	// Check if books already exist
	var count int64
	DB.Model(&models.Book{}).Count(&count)
	
	if count > 0 {
		log.Printf("Books already exist (%d books), skipping seed data", count)
		return nil
	}
	
	// Create comprehensive book collection
	books := []models.Book{
		// Classic Literature
		{Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 3},
		{Title: "Ulysses", Author: "James Joyce", Quantity: 5},
		{Title: "Don Quixote", Author: "Miguel de Cervantes", Quantity: 2},
		{Title: "One Hundred Years of Solitude", Author: "Gabriel García Márquez", Quantity: 4},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 8},
		{Title: "Moby Dick", Author: "Herman Melville", Quantity: 3},
		{Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 2},
		{Title: "Hamlet", Author: "William Shakespeare", Quantity: 6},
		{Title: "The Odyssey", Author: "Homer", Quantity: 4},
		{Title: "Madame Bovary", Author: "Gustave Flaubert", Quantity: 3},
		{Title: "The Divine Comedy", Author: "Dante Alighieri", Quantity: 2},
		{Title: "The Brothers Karamazov", Author: "Fyodor Dostoevsky", Quantity: 3},
		{Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Quantity: 5},
		{Title: "Wuthering Heights", Author: "Emily Brontë", Quantity: 4},
		{Title: "Pride and Prejudice", Author: "Jane Austen", Quantity: 7},
		
		// Modern Classics
		{Title: "1984", Author: "George Orwell", Quantity: 10},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 8},
		{Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 6},
		{Title: "Animal Farm", Author: "George Orwell", Quantity: 9},
		{Title: "Brave New World", Author: "Aldous Huxley", Quantity: 5},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 12},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 15},
		{Title: "Fahrenheit 451", Author: "Ray Bradbury", Quantity: 6},
		{Title: "The Grapes of Wrath", Author: "John Steinbeck", Quantity: 4},
		{Title: "Of Mice and Men", Author: "John Steinbeck", Quantity: 7},
		
		// Contemporary Fiction
		{Title: "The Road", Author: "Cormac McCarthy", Quantity: 5},
		{Title: "Life of Pi", Author: "Yann Martel", Quantity: 6},
		{Title: "The Kite Runner", Author: "Khaled Hosseini", Quantity: 8},
		{Title: "The Book Thief", Author: "Markus Zusak", Quantity: 7},
		{Title: "The Handmaid's Tale", Author: "Margaret Atwood", Quantity: 9},
		{Title: "Gone Girl", Author: "Gillian Flynn", Quantity: 6},
		{Title: "The Girl with the Dragon Tattoo", Author: "Stieg Larsson", Quantity: 5},
		{Title: "The Alchemist", Author: "Paulo Coelho", Quantity: 10},
		{Title: "The Da Vinci Code", Author: "Dan Brown", Quantity: 8},
		{Title: "Angels and Demons", Author: "Dan Brown", Quantity: 6},
		
		// Fantasy & Science Fiction
		{Title: "Harry Potter and the Philosopher's Stone", Author: "J.K. Rowling", Quantity: 20},
		{Title: "Harry Potter and the Chamber of Secrets", Author: "J.K. Rowling", Quantity: 18},
		{Title: "Harry Potter and the Prisoner of Azkaban", Author: "J.K. Rowling", Quantity: 16},
		{Title: "Harry Potter and the Goblet of Fire", Author: "J.K. Rowling", Quantity: 15},
		{Title: "Harry Potter and the Order of the Phoenix", Author: "J.K. Rowling", Quantity: 14},
		{Title: "A Game of Thrones", Author: "George R.R. Martin", Quantity: 12},
		{Title: "A Clash of Kings", Author: "George R.R. Martin", Quantity: 10},
		{Title: "Dune", Author: "Frank Herbert", Quantity: 8},
		{Title: "Foundation", Author: "Isaac Asimov", Quantity: 6},
		{Title: "Ender's Game", Author: "Orson Scott Card", Quantity: 7},
		{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 9},
		{Title: "Neuromancer", Author: "William Gibson", Quantity: 4},
		{Title: "Snow Crash", Author: "Neal Stephenson", Quantity: 5},
		{Title: "The Name of the Wind", Author: "Patrick Rothfuss", Quantity: 8},
		{Title: "American Gods", Author: "Neil Gaiman", Quantity: 7},
		
		// Mystery & Thriller
		{Title: "The Silence of the Lambs", Author: "Thomas Harris", Quantity: 6},
		{Title: "Rebecca", Author: "Daphne du Maurier", Quantity: 4},
		{Title: "And Then There Were None", Author: "Agatha Christie", Quantity: 8},
		{Title: "Murder on the Orient Express", Author: "Agatha Christie", Quantity: 7},
		{Title: "The Hound of the Baskervilles", Author: "Arthur Conan Doyle", Quantity: 5},
		{Title: "The Big Sleep", Author: "Raymond Chandler", Quantity: 3},
		{Title: "The Maltese Falcon", Author: "Dashiell Hammett", Quantity: 3},
		{Title: "In Cold Blood", Author: "Truman Capote", Quantity: 4},
		{Title: "The Shining", Author: "Stephen King", Quantity: 9},
		{Title: "It", Author: "Stephen King", Quantity: 8},
		
		// Literary Fiction
		{Title: "Beloved", Author: "Toni Morrison", Quantity: 5},
		{Title: "The Color Purple", Author: "Alice Walker", Quantity: 6},
		{Title: "Midnight's Children", Author: "Salman Rushdie", Quantity: 3},
		{Title: "The Remains of the Day", Author: "Kazuo Ishiguro", Quantity: 4},
		{Title: "Never Let Me Go", Author: "Kazuo Ishiguro", Quantity: 5},
		{Title: "Atonement", Author: "Ian McEwan", Quantity: 6},
		{Title: "The English Patient", Author: "Michael Ondaatje", Quantity: 3},
		{Title: "The God of Small Things", Author: "Arundhati Roy", Quantity: 4},
		{Title: "White Teeth", Author: "Zadie Smith", Quantity: 5},
		{Title: "The Corrections", Author: "Jonathan Franzen", Quantity: 4},
		
		// Historical Fiction
		{Title: "All the Light We Cannot See", Author: "Anthony Doerr", Quantity: 8},
		{Title: "The Pillars of the Earth", Author: "Ken Follett", Quantity: 7},
		{Title: "Wolf Hall", Author: "Hilary Mantel", Quantity: 4},
		{Title: "The Nightingale", Author: "Kristin Hannah", Quantity: 9},
		{Title: "Memoirs of a Geisha", Author: "Arthur Golden", Quantity: 6},
		{Title: "The Help", Author: "Kathryn Stockett", Quantity: 8},
		{Title: "The Boy in the Striped Pyjamas", Author: "John Boyne", Quantity: 7},
		{Title: "A Thousand Splendid Suns", Author: "Khaled Hosseini", Quantity: 6},
		{Title: "The Shadow of the Wind", Author: "Carlos Ruiz Zafón", Quantity: 5},
		{Title: "The Tattooist of Auschwitz", Author: "Heather Morris", Quantity: 7},
		
		// Non-Fiction
		{Title: "Sapiens", Author: "Yuval Noah Harari", Quantity: 12},
		{Title: "Educated", Author: "Tara Westover", Quantity: 10},
		{Title: "Becoming", Author: "Michelle Obama", Quantity: 15},
		{Title: "The Immortal Life of Henrietta Lacks", Author: "Rebecca Skloot", Quantity: 6},
		{Title: "When Breath Becomes Air", Author: "Paul Kalanithi", Quantity: 8},
		{Title: "The Glass Castle", Author: "Jeannette Walls", Quantity: 7},
		{Title: "Into the Wild", Author: "Jon Krakauer", Quantity: 6},
		{Title: "Unbroken", Author: "Laura Hillenbrand", Quantity: 7},
		{Title: "The Wright Brothers", Author: "David McCullough", Quantity: 4},
		{Title: "Steve Jobs", Author: "Walter Isaacson", Quantity: 8},
		
		// Self-Help & Business
		{Title: "Atomic Habits", Author: "James Clear", Quantity: 15},
		{Title: "The 7 Habits of Highly Effective People", Author: "Stephen Covey", Quantity: 10},
		{Title: "Think and Grow Rich", Author: "Napoleon Hill", Quantity: 8},
		{Title: "How to Win Friends and Influence People", Author: "Dale Carnegie", Quantity: 9},
		{Title: "The Lean Startup", Author: "Eric Ries", Quantity: 7},
		{Title: "Zero to One", Author: "Peter Thiel", Quantity: 6},
		{Title: "Good to Great", Author: "Jim Collins", Quantity: 5},
		{Title: "The 4-Hour Work Week", Author: "Tim Ferriss", Quantity: 8},
		{Title: "Thinking, Fast and Slow", Author: "Daniel Kahneman", Quantity: 6},
		{Title: "Predictably Irrational", Author: "Dan Ariely", Quantity: 5},
	}
	
	// Insert books in batches for better performance
	batchSize := 20
	for i := 0; i < len(books); i += batchSize {
		end := i + batchSize
		if end > len(books) {
			end = len(books)
		}
		
		batch := books[i:end]
		if err := DB.Create(&batch).Error; err != nil {
			return err
		}
		log.Printf("Seeded books %d-%d", i+1, end)
	}
	
	log.Printf("Successfully seeded %d books!", len(books))
	return nil
}

// ResetAndSeedData clears existing data and reseeds (for development)
func ResetAndSeedData() error {
	log.Println("Resetting and reseeding database...")
	
	// Clear existing books
	if err := DB.Exec("DELETE FROM books").Error; err != nil {
		return err
	}
	log.Println("Cleared existing books")
	
	// Reset auto-increment counter
	if err := DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1").Error; err != nil {
		return err
	}
	
	// Seed new data
	return SeedData()
}
